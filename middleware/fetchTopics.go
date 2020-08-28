package middleware

import (
	"github.com/Tsuryu/bff-basu/services/topicservice"
	"github.com/gin-gonic/gin"
)

// FetchTopics : fetch topics
func FetchTopics(context *gin.Context) {
	topics, err := topicservice.FetchAll()
	if err != nil {
		context.JSON(500, gin.H{"message": "Internal error"})
		context.Abort()
		return
	}

	mapa := make(map[string]interface{})
	mapa["topics"] = topics
	context.Keys = mapa
}
