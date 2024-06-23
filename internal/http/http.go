package http

import (
	"fmt"
	"io"
	"net/http"
)

func GetJoke(path, accept string) (string, error) {
	client := http.Client{}
	request, err := http.NewRequest("GET", path, nil)
	if err != nil {
		fmt.Printf("Error creating request: %s", err)
		return "", err
	}

	request.Header.Set("Accept", accept)

	response, err := client.Do(request)
	if err != nil {
		fmt.Printf("Error sending request: %s", err)
		return "", err
	}
	defer response.Body.Close()

	bodyBytes, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Printf("Error reading response body: %s", err)
		return "", err
	}
	bodyString := string(bodyBytes)
	return bodyString, nil
}
