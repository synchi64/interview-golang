package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

const reqToken = "Bearer secret123"

type animal struct {
	Id      int    `json:"id"`
	Name    string `json:"name"`
	Species string `json:"species"`
}

var animals = []animal{
	{Id: 1, Name: "Otter", Species: "Lutra lutra"},
	{Id: 2, Name: "Ferret", Species: "Mustela putorius furo"},
}

func main() {
	router := gin.Default()

	router.GET("/animals", getAnimals)
	router.GET("/animals/:id", getAnimalsById)
	router.POST("/animals", postAnimals)
	router.PATCH("/animals/:id", patchAnimal)
	router.DELETE("/animals/:id", deletAnimal)
	//could add a PUT

	router.Run("localhost:8080")
}

func checkAuth(bearerToken string) bool {
	return bearerToken == reqToken
}

func getAnimals(c *gin.Context) {
	bearerToken := c.Request.Header.Get("Authorization")
	if checkAuth(bearerToken) {
		c.IndentedJSON(http.StatusOK, animals)
	} else {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "unauthorized",
		})
	}
}

func getAnimalsById(c *gin.Context) {
	bearerToken := c.Request.Header.Get("Authorization")
	if checkAuth(bearerToken) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(400, gin.H{"error": "Invalid ID"})
			return
		}
		for _, a := range animals {
			if a.Id == id {
				c.IndentedJSON(http.StatusOK, a)
				return
			}
		}
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "animal not found"})
		return
	} else {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "unauthorized",
		})
	}
}

func postAnimals(c *gin.Context) {
	bearerToken := c.Request.Header.Get("Authorization")
	if checkAuth(bearerToken) {
		var newAnimal animal
		if err := c.BindJSON(&newAnimal); err != nil {
			return
		}
		newAnimal.Id = len(animals) + 1
		animals = append(animals, newAnimal)
		c.IndentedJSON(http.StatusCreated, newAnimal)
	} else {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "unauthorized",
		})
	}
}

func patchAnimal(c *gin.Context) {
	bearerToken := c.Request.Header.Get("Authorization")
	if checkAuth(bearerToken) {
		var newAnimal animal
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(400, gin.H{"error": "Invalid ID"})
			return
		}
		for _, a := range animals {
			if a.Id == id {
				if newAnimal.Name != "" {
					a.Name = newAnimal.Name
				}
				if newAnimal.Species != "" {
					a.Species = newAnimal.Species
				}
				c.IndentedJSON(http.StatusCreated, a)
				return
			}
		}
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "animal not found"})
	} else {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "unauthorized",
		})
	}
}

func deletAnimal(c *gin.Context) {
	bearerToken := c.Request.Header.Get("Authorization")
	if checkAuth(bearerToken) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(400, gin.H{"error": "Invalid ID"})
			return
		}
		for _, a := range animals {
			if a.Id == id {
				animals = append(animals[:id-1], animals[id:]...)
				count := 1
				for _, b := range animals {
					fmt.Print(b)
					animals[count-1].Id = count
					count++
				}
				return
			}
		}
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "animal not found"})
	} else {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "unauthorized",
		})
	}
}
