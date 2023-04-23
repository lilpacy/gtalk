package main

import (
	"bufio"
	"fmt"
	"github.com/joho/godotenv"
	"gtalk/pkg/gpt"
	"log"
	"os"
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
		response, err := gpt.GenerateResponse(input)
		if err != nil {
			fmt.Println("Error:", err)
			continue
		}

		fmt.Println("GPT-3:", response)
	}
}
