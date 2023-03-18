package main

import (
	"bufio"
	"fmt"
	"gtalk/pkg/gpt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Chat with GPT-3:")
	for {
		fmt.Print("You: ")
		scanner.Scan()
		input := scanner.Text()

		response, err := gpt.GenerateResponse(input)
		if err != nil {
			fmt.Println("Error:", err)
			continue
		}

		fmt.Println("GPT-3:", response)
	}
}
