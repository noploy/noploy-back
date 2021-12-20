package main

import (
	"os"

	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)

type Environment struct {
	Port string
}

func LoadEnvironment() *Environment {
	err := godotenv.Load()
	if err != nil {
		logrus.WithField("error", err.Error()).Error("Error in loading environment files.")
		return nil
	}
	var environment Environment
	environment.Port = os.Getenv("PORT")
	return &environment
}
