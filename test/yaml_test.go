package test

import (
	"fmt"
	"gowatcher/go_spider/service/database"
	"testing"
)

func TestYamlConfig(t *testing.T) {
	mysql, _ := database.ReadYamlConfig()
	fmt.Println(mysql)
}
