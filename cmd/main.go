package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/Keyhenge/PowerFlexChallenge/internal/api"
	"github.com/Keyhenge/PowerFlexChallenge/internal/db"
	"github.com/Keyhenge/PowerFlexChallenge/internal/service"

	"go.uber.org/zap"
)

func main() {
	logger := zap.NewExample().Sugar()

	logger.Info("Starting service...")
	//Instantiate DB layer
	dbConfig := db.DBConfig{
		Username: "user",
		Password: "pass",
		Hostname: "postgres",
		Port:     5432,
		DBname:   "postgres",
		Log:      logger,
	}
	baseDB, err := db.NewDB(dbConfig)
	if err != nil {
		log.Fatal(err)
	}

	factoryDB := &db.FactoryDB{BaseDB: *baseDB}
	sprocketDB := &db.SprocketDB{BaseDB: *baseDB}

	//Instantiate service layer
	factoryService := &service.FactoryService{FactoryDB: factoryDB}
	sprocketService := &service.SprocketService{SprocketDB: sprocketDB}

	//Instantiate API layer
	API := &api.API{
		Factory:  factoryService,
		Sprocket: sprocketService,
	}

	//Setup router and server
	router := API.NewRouter()
	server := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}
	go func() {
		logger.Info("Now serving requests!")
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatal(fmt.Errorf("Server unexpectedly failed: %w", err))
		}
	}()

	//Block until we receive SIGINT (Ctrl+C)
	block := make(chan os.Signal, 1)
	signal.Notify(block, syscall.SIGINT)
	<-block

	//Shut down server after passing block
	logger.Info("Shutting down...")
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()
	server.Shutdown(ctx)
	os.Exit(0)
}
