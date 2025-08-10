package utils

import (
	"github.com/bwmarrin/snowflake"
)

func GetTicket(proxyId int64, key string) (string, error) {
	node, err := snowflake.NewNode(1)
	if err != nil {
		return "", err
	}

	// Generate a snowflake ID.
	id := node.Generate()
	return id.String(), nil
}
