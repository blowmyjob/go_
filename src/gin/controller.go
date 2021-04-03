package gin1

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type User struct {
	Name     string `json:"name"`
	Password string `json:"password"`
}

func Login(c *gin.Context) {
	json := User{}

	_ = c.BindJSON(&json)

	fmt.Println("%v", &json)
	c.JSON(http.StatusOK, gin.H{
		"name":     json.Name,
		"password": json.Password,
	})
}
