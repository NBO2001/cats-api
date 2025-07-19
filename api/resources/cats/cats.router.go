package cats

import (
	"github.com/gin-gonic/gin"
)

func Router(r *gin.Engine) {
	r.GET("/cats", getCats)
	r.POST("/cats", postCat)
	r.GET("/cats/:id", getCatByID)
	r.PUT("/cats/:id", updateCat)
	r.DELETE("/cats/:id", deleteCat)
}
