package routers

import (
	"github.com/Tsuryu/bff-basu/middleware"
	"github.com/gin-gonic/gin"
)

// Router : init paths handlers
func Router() {
	router := gin.Default()
	router.GET("/topics", middleware.FetchTopics, middleware.FilterDefaultTopics)
	router.GET("/topics/:id", middleware.FetchTopic)
	router.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
