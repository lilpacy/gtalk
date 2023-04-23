package gpt

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

const apiEndpoint = "https://api.openai.com/v1/chat/completions"

func NewGPT(accessToken string) *GPT {
	return &GPT{
		accessToken: accessToken,
	}
}

type GPT struct {
	accessToken string
}

func (g *GPT) GenerateResponse(prompt string) (string, error) {
	requestBody := map[string]interface{}{
		"model": "gpt-3.5-turbo",
		"messages": []map[string]interface{}{
			{
				"role":    "user",
				"content": prompt,
			},
		},
		"max_tokens":  200,
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
	request.Header.Set("Authorization", "Bearer "+g.accessToken)

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
	message := choice["message"].(map[string]interface{})
	text := message["content"].(string)

	return text, nil
}
