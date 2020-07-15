package pld_snowflake

import (
	"github.com/bwmarrin/snowflake"
	"github.com/michaelzx/pld/pld_logger"
)

var snowflakeNode *snowflake.Node

func InitNode(number int64) {
	var err error
	snowflakeNode, err = snowflake.NewNode(number)
	if err != nil {
		panic(err)
	}
}

func GetNode() *snowflake.Node {
	if snowflakeNode == nil {
		pld_logger.Fatal("snowflakeNode 未初始化")
	}
	return snowflakeNode
}
