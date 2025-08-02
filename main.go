package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"google.golang.org/genai"
)

func main() {
	err := godotenv.Load("env.env")
	if err != nil {
		log.Fatalf("Failed to load env file: %v", err)
	}

	ctx := context.Background()
	client, err := genai.NewClient(ctx, nil)
	if err != nil {
		log.Fatalf("Failed to create new genai client: %v", err)
	}

	config := &genai.GenerateImagesConfig{
		NumberOfImages: 2,
	}

	response, err := client.Models.GenerateImages(
		ctx,
		"imagen-4.0-generate-preview-06-06",
		"Robot coding in a holographic setup",
		config,
	)
	if err != nil {
		log.Fatalf("Failed to get response from model: %v", err)
	}

	for n, image := range response.GeneratedImages {
		fname := fmt.Sprintf("imagen-%d.png", n)
		err := os.WriteFile(fname, image.Image.ImageBytes, 0644)
		if err != nil {
			log.Printf("Failed to save image 'imagen-%d.png'", n)
		}
	}
}
