package parser

import (
	"fmt"
	"regexp"
	"testing"
)

func TestParseProfile(t *testing.T) {
	re := regexp.MustCompile("(?=as)awef")
	ar := re.FindAllString("awef", 2)
	fmt.Printf("%+v", ar)
}
