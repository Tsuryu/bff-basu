package middleware

import (
	"github.com/Tsuryu/bff-basu/models"
	"github.com/gin-gonic/gin"
)

// FilterDefaultTopics : remove default topics
func FilterDefaultTopics(context *gin.Context) {
	var topics []models.Topic

	for _, topic := range context.Keys["topics"].([]models.Topic) {
		if topic.LastUpdate != "rmq-internal" {
			topic.LastUpdate = ""
			topics = append(topics, topic)
		}
	}

	context.JSON(200, gin.H{"list": topics})
}
