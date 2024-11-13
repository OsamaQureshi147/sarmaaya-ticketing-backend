package controllers

import (
	"net/http"

	"github.com/abcdataorg/sarmaaya-ticketing-backend/services"
	"github.com/gin-gonic/gin"
)

// CreateTicket - Handles creating a new ticket in ClickUp
func CreateTicket(c *gin.Context) {
	// Parse request body for ticket details (you may add more details as per ClickUp API requirements)
	var ticketData struct {
		Title       string `json:"title" binding:"required"`
		Description string `json:"description"`
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
