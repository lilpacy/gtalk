package main

import (
	"bufio"
	"fmt"
	"github.com/joho/godotenv"
	"gtalk/pkg/gpt"
	"log"
	"os"
	"strings"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	accessToken := os.Getenv("API_KEY")

	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Chat with GPT-3:")
	for {
		fmt.Print("You: ")
		scanner.Scan()
		input := scanner.Text()

		gpt := gpt.NewGPT(accessToken)
		textChan, err := gpt.GenerateResponse(input)
		if err != nil {
			fmt.Println("Error:", err)
			continue
		}

		fmt.Print("GPT-3: ")
		var buffer strings.Builder
		for text := range textChan {
			buffer.WriteString(text)
			fmt.Print(text)
			os.Stdout.Sync()
		}
		fmt.Println()
	}
}
