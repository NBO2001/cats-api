package cats

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type cat struct {
	ID    string `json:"id" binding:"required"`
	Name  string `json:"name" binding:"required"`
	Breed string `json:"breed" binding:"required"`
	Age   int    `json:"age" binding:"required"`
}

var cats = []cat{
	{ID: "1", Name: "Whiskers", Breed: "Siamese", Age: 3},
	{ID: "2", Name: "Tiger", Breed: "Bengal", Age: 7},
	{ID: "3", Name: "Mittens", Breed: "Tabby", Age: 5},
}

func getCats(c *gin.Context) {
	logrus.Info("Fetching all cats")
	c.IndentedJSON(http.StatusOK, cats)
}

func postCat(c *gin.Context) {
	var newCat cat

	if err := c.BindJSON(&newCat); err != nil {
		c.Status(http.StatusBadRequest)
		return
	}

	cats = append(cats, newCat)
	c.IndentedJSON(http.StatusCreated, newCat)
}

func getCatByID(c *gin.Context) {
	id := c.Param("id")

	for _, a := range cats {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}

	c.Status(http.StatusNotFound)
}

func updateCat(c *gin.Context) {
	id := c.Param("id")
	var newCat cat

	if err := c.BindJSON(&newCat); err != nil {
		c.Status(http.StatusBadRequest)
		return
	}

	for i, a := range cats {
		if a.ID == id {
			cats[i] = newCat
			c.IndentedJSON(http.StatusOK, newCat)
			return
		}
	}

	c.Status(http.StatusNotFound)
}

func deleteCat(c *gin.Context) {
	id := c.Param("id")

	for i, a := range cats {
		if a.ID == id {
			cats = append(cats[:i], cats[i+1:]...)
			c.Status(http.StatusNoContent)
			return
		}
	}

	c.Status(http.StatusNotFound)
}
