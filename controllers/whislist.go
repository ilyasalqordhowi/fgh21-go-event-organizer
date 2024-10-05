package controllers

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/ilyasalqordhowi/fgh21-go-event-organizer/lib"
	"github.com/ilyasalqordhowi/fgh21-go-event-organizer/models"
	"github.com/ilyasalqordhowi/fgh21-go-event-organizer/repository"
)

func ListAllWishlist(r *gin.Context) {
	results := repository.FindAllwishlist()
	r.JSON(http.StatusOK, lib.Message{
		Success: true,
		Message: "List All whislist",
		Results: results,
	})
}
func ListOneWishlist(ctx *gin.Context) {
	
	id := ctx.GetInt("userId")

	
	wishlistItems, err := repository.FindOnewishlist(id)
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


	var results []gin.H

	for _, wishlist := range wishlistItems {

		event, err := repository.FindOneeventsbyid(wishlist.Event_id)
		if err != nil {
		
			log.Printf("Failed to fetch event with id %d: %v", wishlist.Event_id, err)
			continue
		}


		results = append(results, gin.H{
			"whislist": wishlist,
			"event":    event,
		})
	}


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

	err = repository.Createwishlist(eventid, userId)
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

	user_id := ctx.GetInt("userId")


	event_id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, lib.Message{
			Success: false,
			Message: "Invalid event ID",
		})
		return
	}


	err = repository.Deletewishlist(user_id, event_id)
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

	ctx.JSON(http.StatusOK, lib.Message{
		Success: true,
		Message: "Whislist item deleted successfully",
	})
}
