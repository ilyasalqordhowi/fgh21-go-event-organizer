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
	results := models.FindAllWishlist()
	r.JSON(http.StatusOK, lib.Message{
		Success: true,
		Message: "List All Wishlist",
		Results: results,
	})
	fmt.Println(results)
}
func ListOneWishlist(ctx *gin.Context) {

	id := ctx.GetInt("userId")
	
	wishlistItems, err := models.FindOneWishlist(id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, lib.Message{
			Success: false,
			Message: "Wishlist Not Found",
		})
		return
	}

	if len(wishlistItems) == 0 {
		ctx.JSON(http.StatusNotFound, lib.Message{
			Success: false,
			Message: "No wishlist items found for this user",
		})
		return
	}

	
	var results []gin.H

	for _, wishlist := range wishlistItems {
	
		event, err := models.FindOneEventById(wishlist.Event_id)
		if err != nil {
			
			log.Printf("Failed to fetch event with id %d: %v", wishlist.Event_id, err)
			continue
		}


		results = append(results, gin.H{
			"whislist": wishlist,
			"events":    event,
		})
		fmt.Println(results,"ini")
	}

	
	ctx.JSON(http.StatusOK, lib.Message{
		Success: true,
		Message: "Wishlist and events found",
		Results: results,
	})
}

func Createwishlist(ctx *gin.Context) {
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

	err = models.CreateWishlist(eventid, userId)
	if err != nil {

		log.Printf("Createwishlist error: %v", err)

		if err.Error() == "wishlist entry already exists" {
			ctx.JSON(http.StatusConflict, lib.Message{
				Success: false,
				Message: "Event is already in your wishlist",
			})
			return
		}

		ctx.JSON(http.StatusInternalServerError, lib.Message{
			Success: false,
			Message: "Failed to create Wishlist",
		})
		return
	}

	newWish.User_id = userId
	newWish.Event_id = eventid

	ctx.JSON(http.StatusOK, lib.Message{
		Success: true,
		Message: "Wishlist created successfully",
		Results: newWish,
	})
}
func DeleteWishlist(ctx *gin.Context) {

	user_id := ctx.GetInt("userId")


	event_id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, lib.Message{
			Success: false,
			Message: "Invalid event ID",
		})
		return
	}


	err = models.Deletewishlist(user_id, event_id)
	if err != nil {
		if err.Error() == "wishlist item not found" {
			ctx.JSON(http.StatusNotFound, lib.Message{
				Success: false,
				Message: "Wishlist item not found",
			})
			return
		}

		ctx.JSON(http.StatusInternalServerError, lib.Message{
			Success: false,
			Message: "Failed to delete wishlist item",
		})
		return
	}


	ctx.JSON(http.StatusOK, lib.Message{
		Success: true,
		Message: "Wishlist item deleted successfully",
	})
}