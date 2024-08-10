package controllers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/ilyasalqordhowi/fgh21-go-event-organizer/lib"
	"github.com/ilyasalqordhowi/fgh21-go-event-organizer/models"
)
func ListAllUsers(c *gin.Context){
	listUsers := models.FindAllUsers()
		c.JSON(http.StatusOK, lib.Message{
			Success: true,
			Message: "success",
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

// func UpdateUsers(c *gin.Context) {
// 	param := c.Param("id")
// 	id, err := strconv.Atoi(param)
// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, lib.Message{
// 			Success:  false,
// 			Message: "invalid user id",
// 		})
// 		return
// 	}

// 	var user models.User
// 	if err := c.Bind(&user); err != nil {
// 		c.JSON(http.StatusBadRequest, lib.Message{
// 			Success:  false,
// 			Message: "invalid request body",
// 		})
// 		return
// 	}

// 	if err := models.EditUser(id, user); err != nil {
// 		c.JSON(http.StatusNotFound, lib.Message{
// 			Success:  false,
// 			Message: "User Not Found",
// 		})
// 		return
// 	}

// 	c.JSON(http.StatusOK, lib.Message{
// 		Success:  true,
// 		Message: "User Success",
// 		Results: user,
// 	})
// }