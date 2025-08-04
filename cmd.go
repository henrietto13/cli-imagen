package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path"
	"slices"
	"strconv"
	"strings"
	"time"
)

func runCMD() {
	ctx := context.Background()
	saveDir, err := getSavedir()
	if err != nil {
		log.Fatalf("Failed to get save directory: %v", err)
	}

	for {
		fmt.Println("Image image generator CLI")
		fmt.Println("Type q at any point to exit...")
		fmt.Println()

		c := exec.Command("clear")
		c.Stdout = os.Stdout
		c.Run()

		prompt := getPrompt()
		numberOfImages := getNumberOfImages()
		ratio := getRatio()
		sufix := getSufix()

		params := GenerateParams{
			prompt:         prompt,
			numberOfImages: numberOfImages,
			ratio:          ratio,
		}

		images, err := generateImages(ctx, &params)
		if err != nil {
			log.Fatalf("Failed to generate images: %v", err)
		}

		for n, image := range images {
			fname := fmt.Sprintf("imagen_%s%s-%d.png", time.Now().Format("20060102"), sufix, n)
			fpath := path.Join(saveDir, fname)
			err := os.WriteFile(fpath, image.Image.ImageBytes, 0644)
			if err != nil {
				log.Printf("Failed to save image 'imagen-%d.png': %v", n, err)
			}
			fmt.Printf("Image %d saved: %s\n", n, fpath)
		}

		reader := bufio.NewReader(os.Stdin)
		fmt.Println("Press enter to do it again!")
		reader.ReadString('\n')
	}
}

func getSavedir() (string, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", fmt.Errorf("could not get user home directory: %v", err)
	}
	saveDir := path.Join(homeDir, os.Getenv("OUTPUT_DIR"))
	infoDir, err := os.Stat(saveDir)
	if os.IsNotExist(err) {
		err := os.Mkdir(saveDir, 0777)
		if err != nil {
			return "", fmt.Errorf("failed to create save directory: %v", err)
		}
		return saveDir, nil
	}
	if err != nil {
		return "", fmt.Errorf("failed to check save directory: %v", err)
	}

	if infoDir.IsDir() {
		return saveDir, nil
	}
	return "", fmt.Errorf("save direcotry is not a valid directory: %v", err)
}

func getSufix() string {
	fmt.Println()
	fmt.Println("Sufix is 1 word used for naming the output files:")
	fmt.Print(" · Sufix ->  ")
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	sufix := strings.ReplaceAll(strings.TrimSpace(input), " ", "_")
	if strings.ToLower(sufix) == "q" {
		fmt.Println("Good bye!")
		os.Exit(0)
	}
	if sufix != "" {
		return fmt.Sprintf("-%s", sufix)
	}
	return sufix
}

func getRatio() string {
	fmt.Println()
	fmt.Println(`### Aspect Ratio

Specify the image's width-to-height ratio.

**Supported Values:**
  * 1:1 (Square)
  * 3:4 (Portrait) (default)
  * 4:3 (Traditional Landscape)
  * 9:16 (Tall Portrait)
  * 16:9 (Widescreen Landscape)

---`)
	ratios := []string{"1:1", "3:4", "4:3", "9:16", "16:9"}

	reader := bufio.NewReader(os.Stdin)
	fmt.Print(" · Ratio ->  ")
	input, _ := reader.ReadString('\n')
	ratio := strings.TrimSpace(input)
	if strings.ToLower(ratio) == "q" {
		fmt.Println("Good bye!")
		os.Exit(0)
	}
	if slices.Contains(ratios, ratio) {
		return ratio
	}
	fmt.Println("Invalid ratio. Using 3:4 ...")
	return "3:4"
}

func getPrompt() string {
	fmt.Println(`## Imagen Prompt Guide

A good prompt is **descriptive** and **clear**. Focus on these three core elements:

  1. **Subject:** What do you want to see? (e.g., "a majestic lion")
  2. **Context:** Where is it? (e.g., "roaming a savanna at sunset")
  3. **Style:** How should it look? (e.g., "photorealistic, National Geographic style")

**Example:** "A majestic lion roaming a savanna at sunset, photorealistic, National Geographic style."

---`)
	fmt.Println()
	fmt.Println()

	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print(" · Promt ->  ")
		input, _ := reader.ReadString('\n')
		prompt := strings.TrimSpace(input)
		if strings.ToLower(prompt) == "q" {
			fmt.Println("Good bye!")
			os.Exit(0)
		}
		if prompt != "" {
			return prompt
		}
		fmt.Println("Prompt cannot be empty. Please try again...")
	}
}

func getNumberOfImages() int32 {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print(" · Number of images ->  ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)
		if strings.ToLower(input) == "q" {
			fmt.Println("Good bye!")
			os.Exit(0)
		}
		num, err := strconv.ParseInt(input, 10, 32)
		if err != nil {
			fmt.Println("Invalid number. Plase try again...")
			continue
		}

		if num <= 0 {
			fmt.Println("Invalid number. Using 1...")
			return 1
		} else if num >= 4 {
			fmt.Println("Invalid number. Using 4...")
			return 4
		}

		return int32(num)

	}
}
