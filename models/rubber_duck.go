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

func CreateRubberDuck(body *RubberDuck) (string, error) {
	payload, err := json.Marshal(body)
	if err != nil {
		return "", err
	}

	req, err := http.NewRequest(http.MethodPost, ep, bytes.NewReader(payload))
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

func UpdateRubberDuck(id string, body *RubberDuck) error {
	payload, err := json.Marshal(body)
	if err != nil {
		return err
	}

	req, err := http.NewRequest(http.MethodPut, fmt.Sprintf("%s/%s", ep, id), bytes.NewReader(payload))
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("error (%d)", resp.StatusCode)
	}

	return nil
}

func DeleteRubberDuck(id string) error {
	req, err := http.NewRequest(http.MethodDelete, fmt.Sprintf("%s/%s", ep, id), nil)
	if err != nil {
		return err
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusNoContent {
		return fmt.Errorf("error (%d)", resp.StatusCode)
	}

	return nil
}

func GetRubberDuck(id string) (*RubberDuck, error) {
	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("%s", ep), nil)
	if err != nil {
		return nil, err
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("error (%d)", resp.StatusCode)
	}

	var response []RubberDuck
	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		return nil, err
	}

	for _, rubberDuck := range response {
		if rubberDuck.ID == id {
			return &rubberDuck, nil
		}
	}

	return nil, fmt.Errorf("rubber duck with id %s not found", id)
}
