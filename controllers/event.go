package controllers

import (
	"math"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/ilyasalqordhowi/fgh21-go-event-organizer/dtos"
	"github.com/ilyasalqordhowi/fgh21-go-event-organizer/lib"
	"github.com/ilyasalqordhowi/fgh21-go-event-organizer/repository"
)
func ListAllEvent(c *gin.Context){
    search := c.Query("search")
	page, _ := strconv.Atoi(c.Query("page"))
	limit, _ := strconv.Atoi(c.Query("limit"))

	if page < 1 {
		page = 1
	}
	if limit < 1 {
		limit = 1000
	}
	if page > 1 {
		page = (page - 1) * limit
	}

	listEvent, count := repository.FindAllEvent(search, page, limit)

	totalPage := math.Ceil(float64(count) / float64(limit))
	next := 0
	prev := 0

	if int(totalPage) > 1 {
		next = int(totalPage) - page
	}
	if int(totalPage) > 1 {
		prev = int(totalPage) - 1
	}

	totalInfo := lib.TotalInfo{
		TotalData: count,
		TotalPage: int(totalPage),
		Page:      page,
		Limit:     limit,
		Next:      next,
		Prev:      prev,
	}

	lib.HandlerOk(c, "success", totalInfo, listEvent)
	}
func DetailEvent(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	data := repository.FindOneEvent(id)

	if data.Id == id {
		lib.HandlerOk(c, "events Found", nil, data)
	} else {
		lib.HandlerNotFound(c, "events Not Found")
	}
}
func DetailCreateEvent(c *gin.Context) {
    id := c.GetInt("userId")
	dataEvent := repository.FindOneByEvent(id)

	lib.HandlerOk(c, "Event Found", nil, dataEvent)

}
func CreateEvent(ctx *gin.Context) {
    var newEvent dtos.Event
	id, exists := ctx.Get("userId")

	if !exists {
		lib.HandlerUnauthorized(ctx, "Unauthorized")
		return
	}

	userId, ok := id.(int)
	if !ok {
		lib.HandlerBadRequest(ctx, "Invalid user ID")
		return
	}

	if err := ctx.ShouldBind(&newEvent); err != nil {
		lib.HandlerBadRequest(ctx, "Invalid input data")
		return
	}

	err := repository.CreateEvents(newEvent, userId)
	if err != nil {
		lib.HandlerBadRequest(ctx, "Failed to create event")
		return
	}

	newEvent.CreateBy = &userId

	lib.HandlerOk(ctx, "Event created successfully", nil, newEvent)

}

func DeleteEvent(c *gin.Context){
	id, err := strconv.Atoi(c.Param("id"))
	dataEvent := repository.FindOneEvent(id)

	if err != nil {
		lib.HandlerBadRequest(c, "Invalid Event ID")
		return
	}

	err = repository.RemoveEvent(id)
	if err != nil {
		lib.HandlerNotFound(c, "Id Not Found")
		return
	}

	lib.HandlerOk(c, "Event deleted successfully", nil, dataEvent)

}

func UpdateEvent(c *gin.Context) {
    param := c.Param("id")
	id, _ := strconv.Atoi(param)
	search := c.Query("search")
	page, _ := strconv.Atoi(c.Query("page"))
	limit, _ := strconv.Atoi(c.Query("limit"))

	if page < 1 {
		page = 1
	}
	if limit < 1 {
		limit = 10
	}
	if page > 1 {
		page = (page - 1) * limit
	}

	data, count := repository.FindAllEvent(search, page, limit)

	totalPage := math.Ceil(float64(count) / float64(limit))
	next := 0
	prev := 0

	if int(totalPage) > 1 {
		next = int(totalPage) - page
	}
	if int(totalPage) > 1 {
		prev = int(totalPage) - 1
	}

	totalInfo := lib.TotalInfo{
		TotalData: count,
		TotalPage: int(totalPage),
		Page:      page,
		Limit:     limit,
		Next:      next,
		Prev:      prev,
	}

	event := dtos.Event{}
	err := c.Bind(&event)
	if err != nil {
		lib.HandlerBadRequest(c, "Failed to bind data")
		return
	}

	result := dtos.Event{}
	for _, v := range data {
		if v.Id == id {
			result = v
		}
	}

	if result.Id == 0 {
		lib.HandlerNotFound(c, "event with id "+param+" not found")
		return
	}

	idEvent := 0
	for _, v := range data {
		idEvent = v.Id
	}
	event.Id = idEvent

	lib.HandlerOk(c, "Event id "+param+" edit Success", totalInfo, event)

}

func DetailEventSections(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	data, err := repository.FindSectionsByEvent(id)

	if err != nil {
		lib.HandlerNotFound(c, "events sections Not Found")
		return
	}

	lib.HandlerOk(c, "events sections Found", nil, data)
}
func ListAllPaymentMethod(c *gin.Context){
    search := c.Query("search")
	page, _ := strconv.Atoi(c.Query("page"))
	limit, _ := strconv.Atoi(c.Query("limit"))

	if page < 1 {
		page = 1
	}
	if limit < 1 {
		limit = 10
	}
	if page > 1 {
		page = (page - 1) * limit
	}
	listPayment, count := repository.FindAllPaymentMethod(search, page, limit)

	totalPage := math.Ceil(float64(count) / float64(limit))
	next := 0
	prev := 0

	if int(totalPage) > 1 {
		next = int(totalPage) - page
	}
	if int(totalPage) > 1 {
		prev = int(totalPage) - 1
	}

	totalInfo := lib.TotalInfo{
		TotalData: count,
		TotalPage: int(totalPage),
		Page:      page,
		Limit:     limit,
		Next:      next,
		Prev:      prev,
	}

	lib.HandlerOk(c, "success", totalInfo, listPayment)
	}