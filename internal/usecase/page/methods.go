package page

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/sandronister/go_broker/pkg/broker/types"
	"github.com/sandronister/webcrawler/internal/dto"
)

func (u *Model) GetPage(request *dto.PageDTO) error {
	_, err := http.Get(request.URL)

	if err != nil {
		u.logger.Error("Error getting the page", request.URL)
		return fmt.Errorf("invalid url")
	}

	jsonRequest, err := json.Marshal(request)

	if err != nil {
		u.logger.Error("Error marshalling the request", request)
		return fmt.Errorf("error marshalling the request")
	}

	message := &types.Message{
		Topic: u.env.BrokerTopic,
		Value: jsonRequest,
	}

	return u.broker.Publish(message)

}
