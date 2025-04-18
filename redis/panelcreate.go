package redis

import (
	"encoding/json"

	"github.com/TicketsBot-cloud/database"
	"github.com/apex/log"
)

func (c *RedisClient) PublishPanelCreate(settings database.Panel) {
	encoded, err := json.Marshal(settings)
	if err != nil {
		log.Error(err.Error())
		return
	}

	c.Publish(DefaultContext(), "tickets:panel:create", string(encoded))
}
