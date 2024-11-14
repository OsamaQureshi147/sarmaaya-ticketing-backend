package controllers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/abcdataorg/sarmaaya-ticketing-backend/config"
	"github.com/abcdataorg/sarmaaya-ticketing-backend/services"
	"github.com/gin-gonic/gin"
)

func GetAllTicketsInList(c *gin.Context) {
	cfg := config.GetEnvConfig()
	listId := "901804064774"
	reqUrl := cfg.ClickUpApiUrl + "/list/" + listId + "/task"
	req, err := http.NewRequest("GET", reqUrl, nil)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Add the required authorization header
	req.Header.Add("Authorization", cfg.SarmaayaClickUpToken)

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	defer res.Body.Close()

	// Read and decode JSON response directly
	var jsonData map[string]interface{}
	if err := json.NewDecoder(res.Body).Decode(&jsonData); err != nil {
		log.Println("Failed to decode JSON:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to parse JSON response"})
		return
	}

	// Send the decoded JSON response
	c.JSON(http.StatusOK, gin.H{"data": jsonData, "message": "Tickets fetched successfully"})
}

// CreateTicket - Handles creating a new ticket in ClickUp
func CreateTicket(c *gin.Context) {
	// Parse request body for ticket details (you may add more details as per ClickUp API requirements)
	var ticketData struct {
		Title       string `json:"title" binding:"required"`
		Description string `json:"description" binding:"required"`
	}

	if err := c.ShouldBindJSON(&ticketData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Call service to create ticket in ClickUp
	ticketID, err := services.CreateClickUpTicket(ticketData.Title, ticketData.Description)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create ticket"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"ticket_id": ticketID, "message": "Ticket created successfully"})
}

// GetUserTickets - Retrieves all tickets for the user from ClickUp
func GetUserTickets(c *gin.Context) {
	userID := c.Query("user_id") // Example: Fetch user ID from query params

	tickets, err := services.GetClickUpTickets(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve tickets"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"tickets": tickets})
}

// AttachFileToTicket - Attach a file to a specific ticket in ClickUp
func AttachFileToTicket(c *gin.Context) {
	ticketID := c.Param("id")

	// Retrieve the file from form-data
	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "File not provided"})
		return
	}

	// Call service to upload the file to ClickUp ticket
	if err := services.AttachFileToClickUpTicket(ticketID, file); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to attach file to ticket"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "File attached successfully"})
}
