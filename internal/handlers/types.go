package handlers

import (
	"html/template"

	ldap "github.com/10ego/gthp/internal/auth"
	"github.com/10ego/gthp/internal/config"
	"github.com/10ego/gthp/internal/database"
	"go.uber.org/zap"
)

type Handler struct {
	config     *config.Config
	db         *database.DB
	ldapClient *ldap.Client
	templates  *template.Template
	log        *zap.SugaredLogger
}
