package utils

import (
	"net/http"
	"fmt"
	"io/ioutil"
)

func HttpGet(url string) []byte{
	resp, err := http.Get(url)
	defer resp.Body.Close()
	data ,_ :=ioutil.ReadAll(resp.Body)
	if err!=nil {
		fmt.Print(err)
	}
	return data
}