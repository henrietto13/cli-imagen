package main

import (
	"context"
	"fmt"
	"os"

	"google.golang.org/genai"
)

type AspectRatio int

func generateImages(ctx context.Context, prompt string, numberOfImages int32, ratio string) ([]*genai.GeneratedImage, error) {
	client, err := genai.NewClient(
		ctx,
		&genai.ClientConfig{APIKey: os.Getenv("GEMINI_API_KEY")},
	)
	if err != nil {
		return []*genai.GeneratedImage{}, fmt.Errorf("failed to create genai client: %w", err)
	}

	config := &genai.GenerateImagesConfig{
		NumberOfImages:   numberOfImages,
		AspectRatio:      ratio,
		PersonGeneration: "allow_all",
	}

	response, err := client.Models.GenerateImages(
		ctx,
		os.Getenv("IMAGEN_MODEL"),
		prompt,
		config,
	)
	if err != nil {
		return []*genai.GeneratedImage{}, fmt.Errorf("failed to get response from model: %w", err)
	}

	return response.GeneratedImages, nil
}
