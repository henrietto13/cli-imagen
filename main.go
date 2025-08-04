package main

import (
	"log"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load("env.env")
	if err != nil {
		log.Fatalf("Failed to load env file: %v", err)
	}

	runCMD()
}
