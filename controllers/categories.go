package controllers

import (
	"fmt"
	"math"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/ilyasalqordhowi/fgh21-go-event-organizer/lib"
	"github.com/ilyasalqordhowi/fgh21-go-event-organizer/models"
)
func ListAllCategory(c *gin.Context){
    search := c.Query("search")
    page,_ := strconv.Atoi(c.Query("page"))
	limit,_ := strconv.Atoi(c.Query("limit"))
	
	if page < 1 {
        page = 1
	}
	if limit < 1 {
        limit = 5
	}
	 if page > 1 {
        page = (page - 1)*limit
    }
    listCategory,count := models.FindAllCategories(search,page,limit)
	totalPage := math.Ceil(float64(count)/float64(limit))
    next := 0 
    prev := 0

    if int(totalPage)> 1 {
        next = int(totalPage) - page
    }
    if int(totalPage)> 1 {
        prev = int(totalPage) - 1
    }
     totalInfo := lib.TotalInfo{
        TotalData: count,
        TotalPage: int(totalPage),
        Page: page,
        Limit: limit,
        Next: next,
        Prev: prev,
    }
		c.JSON(http.StatusOK, lib.Message{
			Success: true,
			Message: "success",
			ResultsInfo: totalInfo,
			Results: listCategory,
		})		
	}
func DetailCategory(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	data := models.FindOneCategories(id)
	fmt.Println(id)

	if data.Id == id {
		c.JSON(http.StatusOK, lib.Message{
			Success: true,
			Message: "categories Found",
			Results: data,
		})
		return
	} else {
		c.JSON(http.StatusNotFound, lib.Message{
			Success: false,
			Message: "categories Not Found",
		})
	}
}
func CreateCategory(c *gin.Context) {
    newCategory := models.Categories{}
    id, _ := c.Get("userId")
    err := models.CreateCategories(newCategory, id.(int))
    
    if err := c.ShouldBind(&newCategory); err != nil {
        c.JSON(http.StatusBadRequest, lib.Message{
            Success: false,
            Message: "Invalid input data",
        })
        return
    }
    if err != nil {
        c.JSON(http.StatusInternalServerError, lib.Message{
            Success: false,
            Message: "Failed to create Profile",
        })
        return
    }
    
    // newCategory.CreateBy = id.(int)
    c.JSON(http.StatusOK, lib.Message{
        Success: true,
        Message: "Events created successfully",
        Results: newCategory,
    })

}
func DeleteCategory(c *gin.Context){
	id, err := strconv.Atoi(c.Param("id"))
	dataCategory := models.FindOneCategories(id)

	if err != nil {
		c.JSON(http.StatusBadRequest, lib.Message{
			Success:  false,
			Message: "Invalid Category ID",
		})
		return
	}

	err = models.RemoveEvent(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, lib.Message{
			Success:  false,
			Message: "Id Not Found",
		})
		return
	}

	c.JSON(http.StatusOK, lib.Message{
		Success:  true,
		Message: "Category deleted successfully",
		Results: dataCategory,
	})

}

func UpdateCategory(c *gin.Context) {
    param := c.Param("id")
    id,_  := strconv.Atoi(param)
    search := c.Query("search")
    page,_ := strconv.Atoi(c.Query("page"))
	limit,_ := strconv.Atoi(c.Query("limit"))
	
	if page < 1 {
        page = 1
	}
	if limit < 1 {
        limit = 5
	} 
	if page > 1 {
        page = (page - 1)*limit
    }
    data,count := models.FindAllCategories(search,page,limit)
	totalPage := math.Ceil(float64(count)/float64(limit))
    next := 0 
    prev := 0

    if int(totalPage)> 1 {
        next = int(totalPage) - page
    }
    if int(totalPage)> 1 {
        prev = int(totalPage) - 1
    }
     totalInfo := lib.TotalInfo{
        TotalData: count,
        TotalPage: int(totalPage),
        Page: page,
        Limit: limit,
        Next: next,
        Prev: prev,
    }

   category := models.Categories{}
    err := c.Bind(&category)
    if err != nil {
        fmt.Println(err)
        return
    }

    result := models.Categories{}
    for _, v := range data {
        if v.Id == id {
            result = v
        }
    }

    if result.Id == 0 {
        c.JSON(http.StatusNotFound, lib.Message{
            Success: false,
            Message: "category with id " + param + " not found",
        })
        return
    }

    idCategory := 0
    for _, v := range data {
        idCategory = v.Id
    }
    category.Id = idCategory

    // models.EditEvent(*event.Image,*event.Title,event.Date,*event.Descriptions, *event.LocationId, event.CreateBy, param)

    c.JSON(http.StatusOK, lib.Message{
        Success: true,
        Message: "Event  id " + param + " edit Success",
		ResultsInfo: totalInfo,
        Results:category,
    })
}