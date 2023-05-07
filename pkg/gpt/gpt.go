package gpt

import (
	"bufio"
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"strings"
)

const apiEndpoint = "https://api.openai.com/v1/chat/completions"

func NewGPT(accessToken string) *GPT {
	return &GPT{
		accessToken: accessToken,
	}
}

type GPT struct {
	accessToken string
	Messages    []map[string]string
}

func (g *GPT) GenerateResponse(prompt string) (<-chan string, error) {
	g.Messages = append(g.Messages, map[string]string{
		"role":    "user",
		"content": prompt,
	})

	requestBody := map[string]interface{}{
		"model":       "gpt-3.5-turbo",
		"messages":    g.Messages,
		"temperature": 1.0,
		"stream":      true,
	}

	requestBytes, err := json.Marshal(requestBody)
	if err != nil {
		return nil, err
	}

	requestReader := bytes.NewReader(requestBytes)

	request, err := http.NewRequest(http.MethodPost, apiEndpoint, requestReader)
	if err != nil {
		return nil, err
	}

	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Authorization", "Bearer "+g.accessToken)

	client := &http.Client{}

	response, err := client.Do(request)
	if err != nil {
		return nil, err
	}

	textChan := make(chan string)
	go func() {
		defer response.Body.Close()

		reader := bufio.NewReader(response.Body)
		for {
			line, err := reader.ReadString('\n')
			if err != nil {
				if err == io.EOF {
					close(textChan)
				}
				break
			}

			if strings.HasPrefix(line, "data:") {
				data := strings.TrimSpace(line[5:])
				if data == "[DONE]" {
					close(textChan)
					break
				} else {
					var responseMap map[string]interface{}
					if err := json.Unmarshal([]byte(data), &responseMap); err != nil {
						continue
					}
					if choices, ok := responseMap["choices"].([]interface{}); ok {
						choice := choices[0].(map[string]interface{})
						if delta, ok := choice["delta"].(map[string]interface{}); ok {
							if content, ok := delta["content"].(string); ok {
								textChan <- content
							}
						}
					}
				}
			}
		}
	}()
	return textChan, nil
}
