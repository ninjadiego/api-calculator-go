// You can edit this code!
// Click here and start typing.
package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

// album represents data about a record album.
type response struct {
	result float64 `json:"result"`
}

type RequestBody struct {
	num1 int
	num2 int
}

func main() {
	router := gin.Default()
	router.GET("/calculator/plus", getPlus)

	router.Run("localhost:8080")
	a := 4
	b := 5
	suma := a + b
	fmt.Println("Suma = ", suma)

	// Resta
	resta := a - b
	fmt.Println("Resta = ", resta)
}

func getPlus(c *gin.Context) {
	var requestBody RequestBody

	if err := c.BindJSON(&requestBody); err != nil {
		fmt.Println("Error")
	}

	fmt.Println(requestBody.num1, requestBody.num2)
	// c.IndentedJSON(http.StatusOK, albums)
}
