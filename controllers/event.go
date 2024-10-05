package controllers

import (
	"fmt"
	"math"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/ilyasalqordhowi/fgh21-go-event-organizer/dtos"
	"github.com/ilyasalqordhowi/fgh21-go-event-organizer/lib"
	"github.com/ilyasalqordhowi/fgh21-go-event-organizer/repository"
)
func ListAllEvent(c *gin.Context){
    search := c.Query("search")
    page,_ := strconv.Atoi(c.Query("page"))
	limit,_ := strconv.Atoi(c.Query("limit"))
	
	if page < 1 {
        page = 1
	}
	if limit < 1 {
        limit = 1000
	}
    if page > 1 {
        page = (page - 1)*limit
    }
    listEvent,count := repository.FindAllEvent(search,page,limit)

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
			Results: listEvent,
		})	
	}
func DetailEvent(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	data := repository.FindOneEvent(id)
	fmt.Println(id)

	if data.Id == id {
		c.JSON(http.StatusOK, lib.Message{
			Success: true,
			Message: "events Found",
			Results: data,
		})
		return
	} else {
		c.JSON(http.StatusNotFound, lib.Message{
			Success: false,
			Message: "events Not Found",
		})
	}
}
func DetailCreateEvent(c *gin.Context) {
    id := c.GetInt("userId")
    fmt.Println(id)
    dataEvent := repository.FindOneByEvent(id)

    c.JSON(http.StatusOK, lib.Message{
        Success: true,
        Message: "Event Found",
        Results: dataEvent,
    })

}
func CreateEvent(ctx *gin.Context) {
    var newEvent dtos.Event
    id, exists := ctx.Get("userId")
    fmt.Println(newEvent)
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

    // Bind input data ke struct event
    if err := ctx.ShouldBind(&newEvent); err != nil {
        ctx.JSON(http.StatusBadRequest, lib.Message{
            Success: false,
            Message: "Invalid input data",
        })
        return
    }

    err := repository.CreateEvents(newEvent, userId)
    if err != nil {
        ctx.JSON(http.StatusInternalServerError, lib.Message{
            Success: false,
            Message: "Failed to create event",
        })
        return
    }

    newEvent.CreateBy = &userId

    ctx.JSON(http.StatusOK, lib.Message{
        Success: true,
        Message: "Event created successfully",
        Results: newEvent,
    })
}

func DeleteEvent(c *gin.Context){
	id, err := strconv.Atoi(c.Param("id"))
	dataEvent := repository.FindOneEvent(id)

	if err != nil {
		c.JSON(http.StatusBadRequest, lib.Message{
			Success:  false,
			Message: "Invalid Event ID",
		})
		return
	}

	err = repository.RemoveEvent(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, lib.Message{
			Success:  false,
			Message: "Id Not Found",
		})
		return
	}

	c.JSON(http.StatusOK, lib.Message{
		Success:  true,
		Message: "Event deleted successfully",
		Results: dataEvent,
	})

}

func UpdateEvent(c *gin.Context) {
    param := c.Param("id")
    id,_  := strconv.Atoi(param)
    search := c.Query("search")
    page,_ := strconv.Atoi(c.Query("page"))
	limit,_ := strconv.Atoi(c.Query("limit"))
	
	if page < 1 {
        page = 1
	}
	if limit < 1 {
        limit = 10
	}
     if page > 1 {
        page = (page - 1)*limit
    }

    data,count := repository.FindAllEvent(search,page,limit)


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
	

   event := dtos.Event{}
    err := c.Bind(&event)
    if err != nil {
        fmt.Println(err)
        return
    }

    result := dtos.Event{}
    for _, v := range data {
        if v.Id == id {
            result = v
        }
    }

    if result.Id == 0 {
        c.JSON(http.StatusNotFound, lib.Message{
            Success: false,
            Message: "event with id " + param + " not found",
        })
        return
    }

    idEvent := 0
    for _, v := range data {
        idEvent = v.Id
    }
    event.Id = idEvent
    
    c.JSON(http.StatusOK, lib.Message{
        Success: true,
        Message: "Event  id " + param + " edit Success",
        ResultsInfo: totalInfo,
        Results:event,
    })
}

func DetailEventSections(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	data,err := repository.FindSectionsByEvent(id)
	fmt.Println(id)

    if err != nil {
        c.JSON(http.StatusNotFound, lib.Message{
			Success: false,
			Message: "events sections Not Found",
		})
    }

    c.JSON(http.StatusOK, lib.Message{
        Success: true,
        Message: "events sections Found",
        Results: data,
    })
}
func ListAllPaymentMethod(c *gin.Context){
    search := c.Query("search")
    page,_ := strconv.Atoi(c.Query("page"))
	limit,_ := strconv.Atoi(c.Query("limit"))
	
	if page < 1 {
        page = 1
	}
	if limit < 1 {
        limit = 10
	}
    if page > 1 {
        page = (page - 1)*limit
    }
    listPayment,count := repository.FindAllPaymentMethod(search,page,limit)

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
			Results: listPayment,
		})	
	}
// repository.EditEvent(*event.Image,*event.Title,event.Date,*event.Descriptions, *event.LocationId, event.CreateBy, param)