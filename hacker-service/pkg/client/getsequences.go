package client

import (
	"encoding/json"
	"fmt"
	"hacker-service/internal/models"
	"io"
	"log"
	"net/http"
	"os"
)

func GetSequence(id int) (*models.HackMatrix, error) {
	var HackMatrix models.HackMatrix
	url := fmt.Sprintf("%v/GetSequence?id=%v", os.Getenv("MATRIX_SRV_URL"), id)
	resp, err := http.Get(url)
	if err != nil {
		log.Printf("error making request: %v", err)
		return nil, fmt.Errorf("error making request: %v", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Printf("error reading response: %v", err)
		return nil, fmt.Errorf("error reading response: %v", err)
	}

	if resp.StatusCode != http.StatusOK {
		log.Printf("bad response from server: %s", string(body))
		return nil, fmt.Errorf("bad response from server: %s", string(body))
	}

	if err := json.Unmarshal(body, &HackMatrix); err != nil {
		log.Printf("error unmarshalling response: %v", err)
		return nil, fmt.Errorf("error unmarshalling response: %v", err)
	}

	return &HackMatrix, nil
}
