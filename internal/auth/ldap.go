package auth

import (
	"context"
	"fmt"
	"time"

	"github.com/go-ldap/ldap/v3"
)

type Client struct {
	Host       string
	Port       int
	BaseDN     string
	UserFilter string
	GroupDN    string
}

func NewClient(host string, port int, baseDN, userFilter, groupDN string) *Client {
	return &Client{
		Host:       host,
		Port:       port,
		BaseDN:     baseDN,
		UserFilter: userFilter,
		GroupDN:    groupDN,
	}
}

func (c *Client) Authenticate(ctx context.Context, username, password string) (bool, error) {
	// Create a new LDAP connection
	l, err := ldap.DialURL(fmt.Sprintf("ldap://%s:%d", c.Host, c.Port))
	if err != nil {
		return false, fmt.Errorf("failed to connect to LDAP server: %v", err)
	}
	defer l.Close()

	// Construct the user's DN
	userDN := fmt.Sprintf("uid=%s,%s", username, c.UserFilter)

	// Set a timeout for the connection
	ldapTimeout := 10 * time.Second
	l.SetTimeout(ldapTimeout)

	// Attempt to bind as the user
	err = l.Bind(userDN, password)
	if err != nil {
		// If binding fails, the credentials are invalid
		return false, nil
	}

	// Search for the user to check group membership
	searchRequest := ldap.NewSearchRequest(
		c.BaseDN,
		ldap.ScopeWholeSubtree, ldap.NeverDerefAliases, 0, int(ldapTimeout.Seconds()), false,
		fmt.Sprintf("(&(objectClass=user)(uid=%s)(memberOf=%s))", ldap.EscapeFilter(username), c.GroupDN),
		[]string{"dn"},
		nil,
	)

	// Perform the search
	sr, err := l.Search(searchRequest)
	if err != nil {
		return false, fmt.Errorf("LDAP search error: %v", err)
	}

	// Check if the search was interrupted by context cancellation
	select {
	case <-ctx.Done():
		return false, ctx.Err()
	default:
		// If we found an entry, the user is authenticated and in the correct group
		return len(sr.Entries) > 0, nil
	}
}
