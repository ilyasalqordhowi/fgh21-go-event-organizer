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
func ListAllUsers(c *gin.Context){
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
	listUsers,count := models.FindAllUsers(search,page,limit)
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
			Results: listUsers,
		})		
	}
func DetailUsers(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	data := models.FindOneUser(id)
	fmt.Println(id)

	if data.Id == id {
		c.JSON(http.StatusOK, lib.Message{
			Success: true,
			Message: "Users Found",
			Results: data,
		})
		return
	} else {
		c.JSON(http.StatusNotFound, lib.Message{
			Success: false,
			Message: "Users Not Found",
		})
	}
}
func CreateUsers(c *gin.Context) {
    newUser := models.User{}

    if err := c.ShouldBind(&newUser); 
	err != nil {
        c.JSON(http.StatusBadRequest, lib.Message{
            Success:  false,
            Message: "Invalid input data",
        })
        return
    }

    addUser := models.Create(newUser)
	fmt.Println(addUser)
	c.JSON(http.StatusOK, lib.Message{
		Success:  true,
		Message: "User created successfully",
		Results: addUser,
	})

}
func DeleteUsers(c *gin.Context){
	id, err := strconv.Atoi(c.Param("id"))
	dataUser := models.FindOneUser(id)

	if err != nil {
		c.JSON(http.StatusBadRequest, lib.Message{
			Success:  false,
			Message: "Invalid user ID",
		})
		return
	}

	err = models.DeleteUsers(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, lib.Message{
			Success:  false,
			Message: "Id Not Found",
		})
		return
	}

	c.JSON(http.StatusOK, lib.Message{
		Success:  true,
		Message: "User deleted successfully",
		Results: dataUser,
	})

}

func Update(c *gin.Context) {
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
	data,count := models.FindAllUsers(search,page,limit)
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
    user := models.User{}
    err := c.Bind(&user)
    if err != nil {
        fmt.Println(err)
        return
    }

    result := models.User{}
    for _, v := range data {
        if v.Id == id {
            result = v
        }
    }

    if result.Id == 0 {
        c.JSON(http.StatusNotFound, lib.Message{
            Success: false,
            Message: "user with id " + param + " not found",
        })
        return
    }

    idUser := 0
    for _, v := range data {
        idUser = v.Id
    }
    user.Id = idUser

    models.EditUser(user.Email, user.Password, *user.Username ,param)

    c.JSON(http.StatusOK, lib.Message{
        Success: true,
        Message: "user  id " + param + " edit Success",
		ResultsInfo: totalInfo,
        Results: user,
    })
}
func UpdatePassword(ctx *gin.Context) {
	id := ctx.GetInt("userId")
	user := models.FindOneUser(id)
    found := models.FindOneUserByPassword(user.Password)
	fmt.Println(found.Password,"oiiiiiiiiiii")
	isVerified := lib.Verify(user.Password, found.Password)
	fmt.Println(isVerified,"tezzzztttzztztztzt")
	if isVerified {
		ctx.JSON(http.StatusOK,
			lib.Message{
				Success: true,
				Message: "Password success",
			  })
		  
		}else{
			
			ctx.JSON(http.StatusBadRequest,
		 lib.Message{
			 Success: false,
			 Message: "Password tidak sesuai",
		 })
		}
	if user.Id == 0 {
		ctx.JSON(http.StatusNotFound, lib.Message{
			Success: false,
			Message: "User not found",
		})
		return
	}
	
	var dataPassword struct {
		Password string `form:"password" binding:"required,min=8"`
	}
	if err := ctx.ShouldBind(&dataPassword); err != nil {
		ctx.JSON(http.StatusBadRequest, lib.Message{
			Success: false,
			Message: "Invalid input data",
		})
		return
	}
	
	if err := models.Updatepassword(dataPassword.Password, id); err != nil {
		ctx.JSON(http.StatusInternalServerError, lib.Message{
			Success: false,
			Message: "Failed to update password",
		})
		return
	}
	
	ctx.JSON(http.StatusOK, lib.Message{
		Success: true,
		Message: "Password successfully updated",
	})
}