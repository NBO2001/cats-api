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
