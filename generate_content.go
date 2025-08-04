package main

import (
	"context"
	"fmt"
	"os"

	"google.golang.org/genai"
)

type GenerateParams struct {
	prompt         string
	numberOfImages int32
	ratio          string
}

func generateImages(ctx context.Context, params *GenerateParams) ([]*genai.GeneratedImage, error) {
	client, err := genai.NewClient(
		ctx,
		&genai.ClientConfig{APIKey: os.Getenv("GEMINI_API_KEY")},
	)
	if err != nil {
		return []*genai.GeneratedImage{}, fmt.Errorf("failed to create genai client: %w", err)
	}

	config := &genai.GenerateImagesConfig{
		NumberOfImages:   params.numberOfImages,
		AspectRatio:      params.ratio,
		PersonGeneration: "allow_all",
	}

	response, err := client.Models.GenerateImages(
		ctx,
		os.Getenv("IMAGEN_MODEL"),
		params.prompt,
		config,
	)
	if err != nil {
		return []*genai.GeneratedImage{}, fmt.Errorf("failed to get response from model: %w", err)
	}

	return response.GeneratedImages, nil
}
