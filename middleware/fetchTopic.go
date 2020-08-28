package middleware

import (
	"github.com/Tsuryu/bff-basu/services/topicservice"
	"github.com/gin-gonic/gin"
)

// FetchTopic : FetchTopic topics
func FetchTopic(context *gin.Context) {
	topic, statusCode, err := topicservice.FetchOne(context.Param("id"))
	if err != nil {
		if statusCode != 0 {
			context.JSON(statusCode, gin.H{"message": err.Error()})
		} else {
			context.JSON(500, gin.H{"message": "Internal error"})
		}
		context.Abort()
		return
	}

	context.JSON(200, topic)
}
