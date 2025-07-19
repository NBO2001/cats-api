package main

import (
	"nbo2001.com/cat-api/api/router"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	router.RouterCreate(r)

       if err := r.Run("0.0.0.0:8080"); err != nil {
               panic(err)
       }
}
