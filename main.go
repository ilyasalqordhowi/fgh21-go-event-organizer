package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)
type User struct{
	Id int `json:"id"`
	Name string `json:"name" form:"name"`
	Email string `json:"email" form:"email"`
	Password string `json:"-" form:"password"`
	

}
type Message struct{
	Success bool `json:"success"` 
	Message string `json:"message"`
	Results interface{} `json:"results"`

}


func main() {
	r := gin.Default()
	   r.Use(corsMiddleware())
	data := []User{
		{
		Id: 1, 
		Name:"ilyas",
		Email:"ilyas@mail.com",
		Password:"1234",	
},
	}
	
	r.GET("/users",func (c *gin.Context)  {
		c.JSON(http.StatusOK, Message{
			
			Success: true,
			Message: "success",
			Results: data,
			
		})
			
	})
	r.POST("/users", func(c *gin.Context) {
		user:= User{}

		c.Bind(&user)

		user.Id =len(data)+1

		data = append(data,user)

		c.JSON(http.StatusOK, Message{
			Success : true,
			Message: "create data",
			Results: user,
		})
		
	})
	r.POST("/auth/login", func(c *gin.Context) {
		Auth:= User{}

		c.Bind(&Auth)

		email := Auth.Email
		password := Auth.Password
		
		
		dataResults := true
		if dataResults{
			for dataResults {
				for i := 0; i <len(data); i++ {
					resultsEmail := data[i].Email
					resultsPassword := data[i].Password
					if email == resultsEmail && password == resultsPassword {
						c.JSON(http.StatusOK, Message{
							Success : true,
							Message: "Login success",
						})
						return
					}
				}
				
				dataResults = false
			}

		c.JSON(http.StatusUnauthorized, Message{
		Success: false,
		Message: "email dan password invalid",
	})
			
			} 
	})
r.PATCH("/users/:id", func(c *gin.Context) {
		id, _:= strconv.Atoi(c.Param("id"))
		
		
		updatedUser := User{}

		c.Bind(&updatedUser);

		for i, updated := range data {
			if updated.Id == id {
				data[i].Name = updatedUser.Name
				data[i].Email = updatedUser.Email
				data[i].Password = updatedUser.Password
				c.JSON(http.StatusOK, Message{
					Success: true,
					Message: "User Update",
					Results: data,
				})
				return
			}
		}

		c.JSON(http.StatusNotFound, Message{
			Success: false,
			Message: "Users Not Found",
		})
	})
	r.GET("/users/:id", func(c *gin.Context) {
		idString := c.Param("id") 
		id,_ := strconv.Atoi(idString) 

		for _, getId := range data {
			if getId.Id == id {
				c.JSON(http.StatusOK, Message{
					Success: true,
					Message: "Users Found",
					Results: []User{getId},
				})
				return
			}
		}

		c.JSON(http.StatusNotFound, Message{
			Success: false,
			Message: "Users Not Found",
		})
	})
	r.DELETE("/users/:id", func(c *gin.Context) {
		id, _:= strconv.Atoi(c.Param("id"))
		
		
		updatedUser := User{}

		c.Bind(&updatedUser);

		for i, updated := range data {
			if updated.Id == id {
				data = append(data[:i],data[i+1:]...)
				c.JSON(http.StatusOK, Message{
					Success: true,
					Message: "Delete Success",
					Results: data,
				})
				return
			}
		}

		c.JSON(http.StatusNotFound, Message{
			Success: false,
			Message: "Users Not Found",
		})
	})


		r.Run("localhost:8888")
}
func corsMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
        c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
        c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PATCH, DELETE")
        c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

        if c.Request.Method == "OPTIONS" {
            c.AbortWithStatus(204)
            return
        }

        c.Next()
    }
}