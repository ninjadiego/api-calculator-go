// You can edit this code!
// Click here and start typing.
package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type RequestBody struct {
	Num1 int `json:"num1"`
	Num2 int `json:"num2"`
}

type ResponseBody struct {
	Result int `json:"result"`
}

func main() {
	router := gin.Default()
	router.GET("/calculator/plus", getPlus)

	router.Run("localhost:8080")
}

func getPlus(c *gin.Context) {
	requestBody := RequestBody{}
	responseBody := ResponseBody{}

	if err := c.BindJSON(&requestBody); err != nil {
		fmt.Println("Error")
	}
	println("Request Body: ", &requestBody)
	a := requestBody.Num1
	b := requestBody.Num2
	println("a: ", a)
	println("b: ", b)
	responseBody.Result = a + b
	println("resultado: ", responseBody.Result)

	c.IndentedJSON(http.StatusOK, responseBody)
}
