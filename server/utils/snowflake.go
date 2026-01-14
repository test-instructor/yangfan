package utils

import (
	"os"
	"strconv"
	"sync"

	"github.com/bwmarrin/snowflake"
)

var (
	snowflakeOnce sync.Once
	snowflakeNode *snowflake.Node
	snowflakeErr  error
)

func SnowflakeIDString() (string, error) {
	snowflakeOnce.Do(func() {
		nodeID := int64(1)
		if raw := os.Getenv("YF_SNOWFLAKE_NODE"); raw != "" {
			if parsed, err := strconv.ParseInt(raw, 10, 64); err == nil {
				nodeID = parsed
			}
		}
		if nodeID < 0 {
			nodeID = -nodeID
		}
		nodeID = nodeID % 1024
		snowflakeNode, snowflakeErr = snowflake.NewNode(nodeID)
	})
	if snowflakeErr != nil {
		return "", snowflakeErr
	}
	return snowflakeNode.Generate().String(), nil
}
