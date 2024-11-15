package services

import (
	"encoding/json"
	"log"
	"mime/multipart"
	"net/http"

	"github.com/abcdataorg/sarmaaya-ticketing-backend/config"
)

// CreateClickUpTicket - Creates a ticket in ClickUp
func CreateClickUpTicket(title, description string) (string, error) {
	// Integrate with ClickUp API to create a ticket
	// Example: API request to ClickUp
	return "ticketID", nil // Return ticket ID if successful
}

func GetAllTicketsInList(listId string) (map[string]interface{}, error) {
	cfg := config.GetEnvConfig()
	reqUrl := cfg.ClickUpApiUrl + "/list/" + listId + "/task"
	req, err := http.NewRequest("GET", reqUrl, nil)
	if err != nil {
		return nil, err
	}

	// Add the required authorization header
	req.Header.Add("Authorization", cfg.SarmaayaClickUpToken)

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	// Read and decode JSON response directly
	var jsonData map[string]interface{}
	if err := json.NewDecoder(res.Body).Decode(&jsonData); err != nil {
		log.Println("Failed to decode JSON:", err)
		return nil, err
	}
	return jsonData, nil
}

// GetClickUpTickets - Retrieves user tickets from ClickUp
func GetClickUpTickets(userID string) ([]map[string]interface{}, error) {
	// Example: API request to ClickUp to get all tickets for a user
	return []map[string]interface{}{}, nil // Return list of tickets
}

// AttachFileToClickUpTicket - Attaches a file to a specific ticket in ClickUp
func AttachFileToClickUpTicket(ticketID string, file *multipart.FileHeader) error {
	// Example: API request to upload the file to ClickUp
	return nil
}
