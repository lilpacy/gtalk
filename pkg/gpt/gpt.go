package gpt

import (
	"bytes"
	"encoding/json"
	"github.com/joho/godotenv"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

const apiEndpoint = "https://api.openai.com/v1/engines/text-davinci-003/completions"

func GenerateResponse(prompt string) (string, error) {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	accessToken := os.Getenv("API_KEY")

	requestBody := map[string]interface{}{
		"prompt":      prompt,
		"max_tokens":  50,
		"temperature": 0.7,
	}

	requestBytes, err := json.Marshal(requestBody)
	if err != nil {
		return "", err
	}

	requestReader := bytes.NewReader(requestBytes)

	request, err := http.NewRequest(http.MethodPost, apiEndpoint, requestReader)
	if err != nil {
		return "", err
	}

	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Authorization", "Bearer "+accessToken)

	client := &http.Client{}

	response, err := client.Do(request)
	if err != nil {
		return "", err
	}

	defer response.Body.Close()

	responseBody, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return "", err
	}

	var responseMap map[string]interface{}
	if err := json.Unmarshal(responseBody, &responseMap); err != nil {
		return "", err
	}

	choices := responseMap["choices"].([]interface{})
	choice := choices[0].(map[string]interface{})
	text := choice["text"].(string)

	return text, nil
}
