package main

import (
	"bufio"
	"fmt"
	"gtalk/pkg/gpt"
	"os"
	"strings"
)

func main() {
	accessToken := os.Getenv("OPENAI_API_KEY")
	if accessToken == "" {
		fmt.Println("Error: OPENAI_API_KEY is not set")
		os.Exit(1)
	}

	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Chat with GPT-3:")
	gpt := gpt.NewGPT(accessToken)
	for {
		fmt.Print("You: ")
		scanner.Scan()
		input := scanner.Text()

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
		gpt.Messages = append(gpt.Messages, map[string]string{
			"role":    "assistant",
			"content": buffer.String(),
		})
	}
}
