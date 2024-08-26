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
    listEvent,count := models.FindAllEvent(search,page,limit)

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
	data := models.FindOneEvent(id)
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
    dataEvent := models.FindOneByEvent(id)

    c.JSON(http.StatusOK, lib.Message{
        Success: true,
        Message: "Event Found",
        Results: dataEvent,
    })

}
func CreateEvent(ctx *gin.Context) {
    var newEvent models.Event
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

    err := models.CreateEvents(newEvent, userId)
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
// func CreateEvent(c *gin.Context) {
//     event := models.Event{}

//     if err := c.ShouldBind(&event); err != nil {
//         c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
//         return
//     }
//     id, _ := c.Get("userId")
//     result, err := models.CreateEvent(event, id.(int))
//     if err != nil {
//         c.JSON(http.StatusInternalServerError, lib.Message{
//             Success: false,
//             Message: err.Error(),
//         })
//         return
//     }
//     event.CreateBy = id.(int)
//     c.JSON(http.StatusOK, lib.Message{
//         Success: true,
//         Message: "Event created successfully",
//         Results: result,
//     })

// }
func DeleteEvent(c *gin.Context){
	id, err := strconv.Atoi(c.Param("id"))
	dataEvent := models.FindOneEvent(id)

	if err != nil {
		c.JSON(http.StatusBadRequest, lib.Message{
			Success:  false,
			Message: "Invalid Event ID",
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

    data,count := models.FindAllEvent(search,page,limit)


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
	

   event := models.Event{}
    err := c.Bind(&event)
    if err != nil {
        fmt.Println(err)
        return
    }

    result := models.Event{}
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
	data,err := models.FindSectionsByEvent(id)
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
    listPayment,count := models.FindAllPaymentMethod(search,page,limit)

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
// models.EditEvent(*event.Image,*event.Title,event.Date,*event.Descriptions, *event.LocationId, event.CreateBy, param)