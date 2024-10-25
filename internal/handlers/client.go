package handlers

import (
	ldap "github.com/10ego/gthp/internal/auth"
	"github.com/10ego/gthp/internal/config"
	"github.com/10ego/gthp/internal/database"

	"go.uber.org/zap"
)

func New(cfg *config.Config, db *database.DB, ldapClient *ldap.Client, log *zap.SugaredLogger) *Handler {
	return &Handler{config: cfg, db: db, ldapClient: ldapClient, log: log}
}
