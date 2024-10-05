package controllers

import (
	"math"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/ilyasalqordhowi/fgh21-go-event-organizer/lib"
	"github.com/ilyasalqordhowi/fgh21-go-event-organizer/models"
	"github.com/ilyasalqordhowi/fgh21-go-event-organizer/repository"
)
func ListAllCategory(c *gin.Context){
    search := c.Query("search")
	page, _ := strconv.Atoi(c.Query("page"))
	limit, _ := strconv.Atoi(c.Query("limit"))

	if page < 1 {
		page = 1
	}
	if limit < 1 {
		limit = 5
	}
	if page > 1 {
		page = (page - 1) * limit
	}

	listCategory, count := repository.FindAllCategories(search, page, limit)
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

	lib.HandlerOk(c, "success", totalInfo, listCategory)

	}
func DetailCategory(c *gin.Context) {
    id, _ := strconv.Atoi(c.Param("id"))
	data := repository.FindOneCategories(id)

	if data.Id == id {
		lib.HandlerOk(c, "categories Found", nil, data)
	} else {
		lib.HandlerNotFound(c, "categories Not Found")
	}
}
func CreateCategory(c *gin.Context) {
    newCategory := models.Categories{}
	id, _ := c.Get("userId")
	err := repository.CreateCategories(newCategory, id.(int))

	if err := c.ShouldBind(&newCategory); err != nil {
		lib.HandlerBadRequest(c, "Invalid input data")
		return
	}
	if err != nil {
		lib.HandlerBadRequest(c, "Failed to create Category")
		return
	}

	lib.HandlerOk(c, "Category created successfully", nil, newCategory)

}
func DeleteCategory(c *gin.Context){
    id, err := strconv.Atoi(c.Param("id"))
	dataCategory := repository.FindOneCategories(id)

	if err != nil {
		lib.HandlerBadRequest(c, "Invalid Category ID")
		return
	}

	err = repository.RemoveEvent(id)
	if err != nil {
		lib.HandlerNotFound(c, "Id Not Found")
		return
	}

	lib.HandlerOk(c, "Category deleted successfully", nil, dataCategory)

}

func UpdateCategory(c *gin.Context) {
    param := c.Param("id")
	id, _ := strconv.Atoi(param)

	search := c.Query("search")
	page, _ := strconv.Atoi(c.Query("page"))
	limit, _ := strconv.Atoi(c.Query("limit"))

	if page < 1 {
		page = 1
	}
	if limit < 1 {
		limit = 5
	}
	if page > 1 {
		page = (page - 1) * limit
	}

	data, count := repository.FindAllCategories(search, page, limit)
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

	category := models.Categories{}
	err := c.Bind(&category)
	if err != nil {
		lib.HandlerBadRequest(c, "Failed to bind data")
		return
	}

	result := models.Categories{}
	for _, v := range data {
		if v.Id == id {
			result = v
		}
	}

	if result.Id == 0 {
		lib.HandlerNotFound(c, "category with id "+param+" not found")
		return
	}

	category.Id = result.Id
	lib.HandlerOk(c, "Category id "+param+" edit success", totalInfo, category)
}