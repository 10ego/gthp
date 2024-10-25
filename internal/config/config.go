package config

import (
	"fmt"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	ServerAddr     string
	DatabaseURL    string
	LDAPHost       string
	LDAPPort       int
	LDAPBaseDN     string
	LDAPUserFilter string
	LDAPGroupDN    string
	Development    bool
	Title          string
}

func Load() (*Config, error) {
	err := godotenv.Load()
	if err != nil {
		fmt.Errorf("Failed to load environment variables", "error", err)
		return nil, err
	}

	return &Config{
		ServerAddr:     getEnv("SERVER_ADDR", ":8080"),
		DatabaseURL:    getEnv("DATABASE_URL", "postgres://user:password@localhost/gthp?sslmode=disable"),
		LDAPHost:       getEnv("LDAP_HOST", "ldap.my.local"),
		LDAPPort:       getEnvInt("LDAP_PORT", 3389),
		LDAPBaseDN:     getEnv("LDAP_BASE_DN", "dc=my,dc=local"),
		LDAPUserFilter: getEnv("LDAP_USER_FILTER", "ou=ad,dc=my,dc=local"),
		LDAPGroupDN:    getEnv("LDAP_GROUP_DN", "cn=mygroup,ou=groups,dc=my,dc=local"),
		Development:    getEnvBool("DEVELOPMENT", true),
		Title:          getEnv("TITLE", "default site title"),
	}, nil
}

func getEnv(key string, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}
func getEnvInt(key string, fallback int) int {
	if value, ok := os.LookupEnv(key); ok {
		n, err := strconv.Atoi(value)
		if err != nil {
			panic(err)
		}
		return n
	}
	return fallback
}
func getEnvBool(key string, fallback bool) bool {
	if value, ok := os.LookupEnv(key); ok {
		b, err := strconv.ParseBool(value)
		if err == nil {
			return b
		}
	}
	return fallback
}
