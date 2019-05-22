package fetcher

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func Fetch(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	if resp.StatusCode == http.StatusOK {
		all, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			panic(err)
		}

		return all, nil
	} else {
		return []byte{}, fmt.Errorf("aawef")
	}
}