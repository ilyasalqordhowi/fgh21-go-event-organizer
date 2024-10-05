package controllers

import (
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
	lib.HandlerOk(r, "List All wishlist", nil, results)
}
func ListOneWishlist(ctx *gin.Context) {
	
	id := ctx.GetInt("userId")

	wishlistItems, err := repository.FindOnewishlist(id)
	if err != nil {
		lib.HandlerNotFound(ctx, "Wishlist Not Found")
		return
	}

	if len(wishlistItems) == 0 {
		lib.HandlerNotFound(ctx, "No wishlist items found for this user")
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
			"whishlist": wishlist,
			"event":     event,
		})
	}

	lib.HandlerOk(ctx, "Wishlist and events found", nil, results)

}

func CreateWishListEvent(ctx *gin.Context) {
	var newWish models.Wishlist

	id, exists := ctx.Get("userId")
	if !exists {
		lib.HandlerUnauthorized(ctx, "Unauthorized")
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
		lib.HandlerBadRequest(ctx, "Invalid event ID")
		return
	}

	log.Printf("userId: %d, eventid: %d", userId, eventid)

	err = repository.Createwishlist(eventid, userId)
	if err != nil {
		log.Printf("Createwishlist error: %v", err)

		if err.Error() == "wishlist entry already exists" {
			ctx.JSON(http.StatusConflict, lib.Message{
				Success: false,
				Message: "Event is already in your wishlist",
			})
			return
		}

		lib.HandlerBadRequest(ctx, "Failed to create Wishlist")
		return
	}

	newWish.User_id = userId
	newWish.Event_id = eventid

	lib.HandlerOk(ctx, "Wishlist created successfully", nil, newWish)
}
func DeleteWishlist(ctx *gin.Context) {

	user_id := ctx.GetInt("userId")

	event_id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		lib.HandlerBadRequest(ctx, "Invalid event ID")
		return
	}

	err = repository.Deletewishlist(user_id, event_id)
	if err != nil {
		if err.Error() == "wishlist item not found" {
			lib.HandlerNotFound(ctx, "Wishlist item not found")
			return
		}

		lib.HandlerBadRequest(ctx, "Failed to delete wishlist item")
		return
	}

	lib.HandlerOk(ctx, "Wishlist item deleted successfully", nil, nil)
}
