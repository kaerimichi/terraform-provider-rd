package models

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type RubberDuck struct {
	ID       string `json:"id"`
	Color    string `json:"color"`
	Material string `json:"material"`
	Size     string `json:"size"`
}

const (
	ep = "http://localhost:8080/rubberducks"
)

func (r *RubberDuck) CreateRubberDuck(body *RubberDuck) (string, error) {
	payload, err := json.Marshal(body)
	if err != nil {
		return "", err
	}

	req, err := http.NewRequest("POST", ep, bytes.NewReader(payload))
	if err != nil {
		return "", err
	}

	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusCreated {
		return "", fmt.Errorf("error (%d)", resp.StatusCode)
	}

	var responseBody bytes.Buffer

	_, err = responseBody.ReadFrom(resp.Body)
	if err != nil {
		return "", err
	}

	var response RubberDuck

	err = json.Unmarshal(responseBody.Bytes(), &response)
	if err != nil {
		return "", err
	}

	return response.ID, nil
}
