package main

import (
	"log"

	"github.com/joho/godotenv"
)

type GenerateParams struct {
	prompt         string
	numberOfImages int32
	ratio          string
}

func main() {
	err := godotenv.Load("env.env")
	if err != nil {
		log.Fatalf("Failed to load env file: %v", err)
	}

	runCMD()
	// params := GenerateParams{}
	//
	// ctx := context.Background()
	//
	// images, err := generateImages(ctx, "", 2)
	// if err != nil {
	// 	log.Fatalf("Failed to generate images: %v", err)
	// }
	//
	// for n, image := range images {
	// 	fname := fmt.Sprintf("imagen-log-%d.png", n)
	// 	err := os.WriteFile(fname, image.Image.ImageBytes, 0644)
	// 	if err != nil {
	// 		log.Printf("Failed to save image 'imagen-%d.png'", n)
	// 	}
	// }
}
