package topicservice

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/Tsuryu/bff-basu/models"
	"github.com/Tsuryu/bff-basu/utils"
)

// FetchOne : fetch a topic from basu
func FetchOne(id string) (topic models.Topic, statusCode int, err error) {
	var topicLocal models.Topic

	resp, err := http.Get("http://192.168.0.11:8080/queueing/topics/" + id)
	if err != nil {
		log.Println("Failed to fetch")
		return topicLocal, statusCode, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("Failed to read response")
		return topicLocal, statusCode, err
	}

	err = utils.GetResponseBody(body, resp.StatusCode, &topicLocal)
	if err != nil {
		return topicLocal, resp.StatusCode, err
	}

	err = json.Unmarshal(body, &topicLocal)
	if err != nil {
		log.Println("Failed to unmarshal")
		return topicLocal, statusCode, err
	}

	return topicLocal, statusCode, nil
}
