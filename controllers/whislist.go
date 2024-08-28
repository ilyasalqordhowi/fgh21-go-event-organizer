package controllers

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/ilyasalqordhowi/fgh21-go-event-organizer/lib"
	"github.com/ilyasalqordhowi/fgh21-go-event-organizer/models"
)

func ListAllWishlist(r *gin.Context) {
	results := models.FindAllwishlist()
	r.JSON(http.StatusOK, lib.Message{
		Success: true,
		Message: "List All whislist",
		Results: results,
	})
}
func ListOneWishlist(ctx *gin.Context) {
	// Extract the user ID from the context
	id := ctx.GetInt("userId")

	// Fetch the wishlist items for the user
	wishlistItems, err := models.FindOnewishlist(id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, lib.Message{
			Success: false,
			Message: "Whislist Not Found",
		})
		return
	}

	if len(wishlistItems) == 0 {
		ctx.JSON(http.StatusNotFound, lib.Message{
			Success: false,
			Message: "No whislist items found for this user",
		})
		return
	}

	// Prepare the response: List of wishlist items and their associated events
	var results []gin.H

	for _, wishlist := range wishlistItems {
		// Fetch event details for each wishlist item using the event_id
		event, err := models.FindOneeventsbyid(wishlist.Event_id)
		if err != nil {
			// Log the error and continue to the next event (don't stop the loop)
			log.Printf("Failed to fetch event with id %d: %v", wishlist.Event_id, err)
			continue
		}

		// Add the wishlist item and event details to the results
		results = append(results, gin.H{
			"whislist": wishlist,
			"event":    event,
		})
	}

	// Return the combined wishlist and event details
	ctx.JSON(http.StatusOK, lib.Message{
		Success: true,
		Message: "Whislist and events found",
		Results: results,
	})
}

func CreateWishListEvent(ctx *gin.Context) {
	var newWish models.Wishlist

	id, exists := ctx.Get("userId")
	if !exists {
		ctx.JSON(http.StatusUnauthorized, lib.Message{
			Success: false,
			Message: "Unauthorized",
		})
		return
	}

	userId, ok := id.(int)
	if !ok {
		ctx.JSON(http.StatusInternalServerError, lib.Message{
			Success: false,
			Message: "Invalid user ID",
		})
		return
	}

	eventid, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, lib.Message{
			Success: false,
			Message: "Invalid event ID",
		})
		return
	}

	log.Printf("userId: %d, eventid: %d", userId, eventid)

	err = models.Createwishlist(eventid, userId)
	fmt.Println(err,"dhdihidhi")
	if err != nil {

		log.Printf("Createwishlist error: %v", err)

		if err.Error() == "whislist entry already exists" {
			ctx.JSON(http.StatusConflict, lib.Message{
				Success: false,
				Message: "Event is already in your whislist",
			})
			return
		}

		ctx.JSON(http.StatusInternalServerError, lib.Message{
			Success: false,
			Message: "Failed to create Whislist",
		})
		return
	}

	newWish.User_id = userId
	newWish.Event_id = eventid

	ctx.JSON(http.StatusOK, lib.Message{
		Success: true,
		Message: "Whislist created successfully",
		Results: newWish,
	})
}
func DeleteWishlist(ctx *gin.Context) {
	// Get the user ID from the context
	user_id := ctx.GetInt("userId")

	// Get the event ID from the URL parameters
	event_id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, lib.Message{
			Success: false,
			Message: "Invalid event ID",
		})
		return
	}

	// Call the model function to delete the wishlist item
	err = models.Deletewishlist(user_id, event_id)
	if err != nil {
		if err.Error() == "whislist item not found" {
			ctx.JSON(http.StatusNotFound, lib.Message{
				Success: false,
				Message: "Whislist item not found",
			})
			return
		}

		ctx.JSON(http.StatusInternalServerError, lib.Message{
			Success: false,
			Message: "Failed to delete whislist item",
		})
		return
	}

	// Return success response
	ctx.JSON(http.StatusOK, lib.Message{
		Success: true,
		Message: "Whislist item deleted successfully",
	})
}
