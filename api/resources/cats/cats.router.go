package cats

import (
	"github.com/gin-gonic/gin"
)

func Router(r *gin.Engine) {

	r.GET("/cats", getCats)
	// r.POST("/cats", postCats)
	// r.GET("/cats/:id", getCatByID)
}
