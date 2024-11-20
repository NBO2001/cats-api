package router

import (
	"github.com/gin-gonic/gin"
	"nbo2001.com/cat-api/api/resources/cats"
)

func RouterCreate(r *gin.Engine) {
	router := r
	cats.Router(router)
}
