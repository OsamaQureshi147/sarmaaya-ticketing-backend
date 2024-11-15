package routes

import (
	"github.com/abcdataorg/sarmaaya-ticketing-backend/controllers"
	"github.com/gin-gonic/gin"
)

func TicketRoutes(router *gin.Engine) {
	tickets := router.Group("/tickets")
	tickets.POST("/", controllers.CreateTicket)                      // Create a new ticket
	tickets.GET("/", controllers.GetUserTickets)                     // Get all tickets for a user
	tickets.GET("/:list_id/task", controllers.GetAllTicketsInList)   // Get all tickets in list
	tickets.POST("/:id/attachments", controllers.AttachFileToTicket) // Attach a file to a specific ticket
}
