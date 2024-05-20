package main

import (
	"github.com/joho/godotenv"
	_ "github.com/plab0n/search-paste/cmd/server/docs"
	"github.com/plab0n/search-paste/config"
	api "github.com/plab0n/search-paste/internal"
	_ "github.com/plab0n/search-paste/internal/workers"
	"github.com/plab0n/search-paste/pkg/logger"
	"github.com/sirupsen/logrus"
	"os"
	"strings"
	"sync"
)

// @title Go Rest Api
// @description Api Endpoints for Go Server
func main() {
	err := godotenv.Load()
	if err != nil {
		logger.Log.WithFields(logrus.Fields{
			"err": err,
		}).Error("Can't load config from .env. Problem with .env, or the server is in production environment.")
		return
	}

	config := config.ApiEnvConfig{
		Env:     strings.ToUpper(os.Getenv("ENV")),
		Port:    os.Getenv("PORT"),
		Version: os.Getenv("VERSION"),
	}

	logger.Log.WithFields(logrus.Fields{
		"env":     config.Env,
		"version": config.Version,
		"port":    config.Port,
	}).Info("Loaded app config")

	var wg sync.WaitGroup
	wg.Add(1)

	// Starting our magnificent server
	go func() {
		defer wg.Done()

		server := api.AppServer{}
		defer func() {
			if r := recover(); r != nil {
				server.OnShutdown()
			}
		}()

		server.Run(config)
	}()
	wg.Wait()

}

// cSpell:ignore logrus godotenv
