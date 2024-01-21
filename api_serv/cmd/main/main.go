package main

import (
	"context"
	"fmt"
	"github.com/elecshen/shopping_list/internal/Server"
	"github.com/elecshen/shopping_list/internal/handler"
	"github.com/elecshen/shopping_list/internal/repository"
	"github.com/elecshen/shopping_list/internal/service"
	"github.com/go-oauth2/oauth2/v4/generates"
	"github.com/go-oauth2/oauth2/v4/manage"
	"github.com/go-oauth2/oauth2/v4/models"
	oaSrv "github.com/go-oauth2/oauth2/v4/server"
	"github.com/go-oauth2/oauth2/v4/store"
	"github.com/golang-jwt/jwt"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
	"os"
	"os/signal"
	"syscall"
)

// @title Shopping list API
// @version 1.0
// @description API Server for ShoppingList Application

// @host localhost:80
// @BasePath /

// @securityDefinitions.apikey JWTAuth
// @in header
// @name Authorization

func main() {
	cfg := repository.Config{
		Host:     os.Getenv("AUCTION_DB_HOST"),
		Port:     os.Getenv("AUCTION_DB_PORT"),
		DBName:   os.Getenv("AUCTION_DB_NAME"),
		Username: os.Getenv("AUCTION_DB_USERNAME"),
		Password: os.Getenv("AUCTION_DB_PASSWORD"),
		SSLMode:  os.Getenv("AUCTION_DB_SSLMODE"),
	}

	logrus.Info(fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
		cfg.Host, cfg.Port, cfg.Username, cfg.DBName, cfg.Password, cfg.SSLMode))

	db, err := repository.NewPostgresDB(cfg)
	if err != nil {
		logrus.Fatalf("failed to init db: %s", err.Error())
	}

	driver, err := postgres.WithInstance(db.DB, &postgres.Config{})
	m, err := migrate.NewWithDatabaseInstance(
		"file:///go/src/schema/",
		"postgres", driver)
	if err != nil {
		logrus.Fatalf("failed to connect for migration: %s", err.Error())
	}
	err = m.Up()
	if err != nil && err != migrate.ErrNoChange {
		logrus.Fatalf("failed to migrate db: %s", err.Error())
	}

	var (
		idVar     string = "000000"
		secretVar string = "123456"
		domainVar string = "http://localhost:80/oauth"
	)

	manager := manage.NewDefaultManager()
	manager.SetAuthorizeCodeTokenCfg(manage.DefaultAuthorizeCodeTokenCfg)

	manager.MustTokenStorage(store.NewMemoryTokenStore())
	//manager.MapAccessGenerate(generates.NewAccessGenerate())
	manager.MapAccessGenerate(generates.NewJWTAccessGenerate("", []byte("rgjef#4#%8GHNr43bj#rgek4FRMN"), jwt.SigningMethodHS256))

	clientStore := store.NewClientStore()
	clientStore.Set(idVar, &models.Client{
		ID:     idVar,
		Secret: secretVar,
		Domain: domainVar,
	})
	manager.MapClientStorage(clientStore)

	oauthSrv := oaSrv.NewServer(oaSrv.NewConfig(), manager)

	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	handlers := handler.NewHandler(services, oauthSrv)

	srv := new(Server.Server)
	go func() {
		if err = srv.Run("80", handlers.InitRoutes()); err != nil {
			logrus.Fatalf("error occured while running http server: %s", err.Error())
		}
	}()

	logrus.Info("Server started")

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	logrus.Info("Server shutting down")

	if err = srv.Shutdown(context.Background()); err != nil {
		logrus.Errorf("error occured on server shutting down: %s", err.Error())
	}

	if err = db.Close(); err != nil {
		logrus.Errorf("error occured on db connection close: %s", err.Error())
	}
}
