package main

import (
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
	//could add a PUT and DELETE

	router.Run("localhost:8080")
}

func getAnimals(c *gin.Context) {
	bearerToken := c.Request.Header.Get("Authorization")
	if bearerToken == reqToken {
		c.JSON(http.StatusOK, gin.H{
			"data": "resource data",
		})
		return
	}
	c.JSON(http.StatusUnauthorized, gin.H{
		"message": "unauthorized",
	})

	c.IndentedJSON(http.StatusOK, animals)
}

func getAnimalsById(c *gin.Context) {
	bearerToken := c.Request.Header.Get("Authorization")
	if bearerToken == reqToken {
		c.JSON(http.StatusOK, gin.H{
			"data": "resource data",
		})
		return
	}
	c.JSON(http.StatusUnauthorized, gin.H{
		"message": "unauthorized",
	})
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		panic(err)
	}

	for _, a := range animals {
		if a.Id == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "animal not found"})
}

func postAnimals(c *gin.Context) {
	bearerToken := c.Request.Header.Get("Authorization")
	if bearerToken == reqToken {
		c.JSON(http.StatusOK, gin.H{
			"data": "resource data",
		})
		return
	}
	c.JSON(http.StatusUnauthorized, gin.H{
		"message": "unauthorized",
	})
	var newAnimal animal

	if err := c.BindJSON(&newAnimal); err != nil {
		return
	}

	newAnimal.Id = len(animals) + 1

	animals = append(animals, newAnimal)
	c.IndentedJSON(http.StatusCreated, newAnimal)
}

func patchAnimal(c *gin.Context) {
	bearerToken := c.Request.Header.Get("Authorization")
	if bearerToken == reqToken {
		c.JSON(http.StatusOK, gin.H{
			"data": "resource data",
		})
		return
	}
	c.JSON(http.StatusUnauthorized, gin.H{
		"message": "unauthorized",
	})
	var newAnimal animal
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		panic(err)
	}
	if err := c.BindJSON(&newAnimal); err != nil {
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
}
