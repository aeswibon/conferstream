package cmd

import (
	"log"

	"github.com/spf13/viper"
)

// Env struct holds the environment variables
type Env struct {
	// Addr is the address to listen on
	Addr string `json:"ADDR"`
	// Cert is the path to the cert file
	Cert string `json:"CERT"`
	// Key is the path to the key file
	Key string `json:"KEY"`
}

// NewEnv returns a new Env
func NewEnv() *Env {
	env := &Env{}
	viper.SetConfigFile(".env")
	viper.AutomaticEnv()
	if err := viper.ReadInConfig(); err != nil {
		log.Fatal("Can't find the .env file, ", err)
	}
	if err := viper.Unmarshal(env); err != nil {
		log.Fatal("Failed to read .env file, ", err)
	}
	return env
}