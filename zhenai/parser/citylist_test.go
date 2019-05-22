package parser

import (
	"crawler/fetcher"
	"fmt"
	"testing"
)

func TestParseCityList(t *testing.T) {
	contents, err := fetcher.Fetch("https://www.zhenai.com/zhenghun")

	if err != nil {
		panic(err)
	}
	fmt.Printf("%s", contents)
	ParseCityList(contents)
}
