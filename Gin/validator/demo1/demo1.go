package main

import (
	"fmt"
	"net/http"
	
	"github.com/gin-gonic/gin"
)

type RegisterRequest struct {
	Username string `json:"username" binding:"required"`
	Nickname string `json:"nickname" binding:"required"`
	Email    string `json:"email"    binding:"required,email"`
	Password string `json:"password" binding:"required"`
	Age      uint8  `json:"age"      binding:"gte=1,lte=120"`
}
   
func Register(c *gin.Context) {
	req := new(RegisterRequest)
	if err := c.ShouldBind(req); err != nil {
		fmt.Println("Register Error!")
		c.JSON(http.StatusOK, gin.H{"msg": err.Error()})
		return
	}
	fmt.Println("Register Success!")
	c.JSON(http.StatusOK, gin.H{"msg": "success"})
}

func main() {
	router := gin.Default()
	router.POST("register",Register)
	router.Run(":8080")
}