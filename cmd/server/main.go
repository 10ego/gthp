package main

import (
	"net/http"
	"os"
	"path/filepath"

	"github.com/10ego/gthp/internal/auth"
	"github.com/10ego/gthp/internal/config"
	"github.com/10ego/gthp/internal/database"
	"github.com/10ego/gthp/internal/handlers"
	"github.com/10ego/gthp/internal/logger"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		os.Exit(1)
	}
	log := logger.New(cfg.Development)
	defer log.Sync()

	db, err := database.Connect(cfg.DatabaseURL, log)
	if err != nil {
		log.Fatalw("Failed to connect to database", "error", err)
	}
	defer db.Close()

	ldapClient := auth.NewClient(
		cfg.LDAPHost,
		cfg.LDAPPort,
		cfg.LDAPBaseDN,
		cfg.LDAPUserFilter,
		cfg.LDAPGroupDN,
	)

	h := handlers.New(cfg, db, ldapClient, log)

	mux := http.NewServeMux()
	mux.HandleFunc("GET /", h.IndexHandler)
	mux.HandleFunc("GET /login", h.LoginHandler)
	mux.HandleFunc("POST /login", h.LoginHandler)
	mux.HandleFunc("POST /logout", h.LogoutHandler)

	fs := http.FileServer(http.Dir(filepath.Join("internal", "static")))
	mux.Handle("GET /static/", http.StripPrefix("/static/", fs))

	log.Infow("Server starting", "address", cfg.ServerAddr)
	log.Fatal(http.ListenAndServe(cfg.ServerAddr, mux))
	//
	// srv := &http.Server{
	// 	Addr:    cfg.ServerAddr,
	// 	Handler: mux,
	// }
	//
	// // Start the server in a goroutine
	// go func() {
	// 	log.Infow("Server starting", "address", cfg.ServerAddr)
	// 	if err := srv.ListenAndServe(); err != http.ErrServerClosed {
	// 		log.Fatalw("ListenAndServe Failed", "error", err)
	// 	}
	// }()
	//
	// // Wait for interrupt signal to gracefully shutdown the server
	// quit := make(chan os.Signal, 1)
	// signal.Notify(quit, os.Interrupt)
	// <-quit
	// log.Info("Server is shutting down...")
	//
	// ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	// defer cancel()
	// if err := srv.Shutdown(ctx); err != nil {
	// 	log.Fatalw("Server forced to shutdown", "error", err)
	// }
	//
	// log.Info("Server exiting")
}
