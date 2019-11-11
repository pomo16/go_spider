package test

import (
	"fmt"
	"gowatcher/go_spider/service/database"
	"path/filepath"
	"testing"
)

func TestYamlConfig(t *testing.T) {
	path, _ := filepath.Abs("../config/config.yaml")
	mysql, _ := database.ReadYamlConfig(path)
	fmt.Println(mysql)
}
