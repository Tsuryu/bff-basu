package topicservice

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/Tsuryu/bff-basu/models"
)

// FetchAll : fetchs topics from basu
func FetchAll() (topics []models.Topic, err error) {
	var topicList models.TopicList

	resp, err := http.Get("http://192.168.0.11:8080/queueing/topics")
	if err != nil {
		log.Println("No topics found")
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("Failed to read response")
		return nil, err
	}

	err = json.Unmarshal(body, &topicList)
	if err != nil {
		log.Println("Failed to unmarshal")
		return nil, err
	}

	return topicList.List, nil
}
