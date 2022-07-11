package main

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	cors "github.com/rs/cors/wrapper/gin"
	"io/ioutil"
	"net/http"
)

type detail struct {
	Name        string `json:"name"`
	Company     string `json:"company"`
	PhoneNumber string `json:"phone_number"`
}

var details detail

func getDetails(c *gin.Context) {
	body, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		fmt.Println(err)
		c.IndentedJSON(http.StatusBadRequest, err)
		return
	}
	fmt.Println(body)
	c.IndentedJSON(http.StatusOK, details)
	return
}

func postDetails(c *gin.Context) {

	body, err := ioutil.ReadAll(c.Request.Body)
	err = json.Unmarshal(body, &details)
	if err != nil {
		fmt.Println(err)
		c.IndentedJSON(http.StatusBadRequest, err)
		return
	}
	fmt.Println(details)
	c.IndentedJSON(http.StatusOK, details)
	return
}

func main() {
	router := gin.Default()
	router.Use(cors.Default())
	router.GET("/home", getDetails)
	router.POST("/post", postDetails)
	//router.GET("/home", generatePassword)

	router.Run("localhost:8082")
}
