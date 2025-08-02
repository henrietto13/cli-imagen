package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"slices"
	"strconv"
	"strings"
)

func runCMD() {
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

		// TODO Process input

		fmt.Println(prompt, numberOfImages, ratio)

		reader := bufio.NewReader(os.Stdin)
		fmt.Println("Press enter to do it again!")
		reader.ReadString('\n')
	}
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
