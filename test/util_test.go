package test

import (
	"fmt"
	"gowatcher/go_spider/utils"
	"testing"
)

func TestUtil(t *testing.T) {
	fmt.Println(utils.FillLastCrawlTime())
}
