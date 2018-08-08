package main

// .data.children[0].data.preview.images[0].source.url

import (
	//	"encoding/json"
	//	"byte"
	//	"fmt"
	"github.com/kr/pretty"
	"github.com/savaki/jq"
	"io/ioutil"
	"net/http"
)

const uri = "https://www.reddit.com/r/wtfstockphotos/.json"

//const uri = "https://www.reddit.com/r/predators/.json"

func main() {
	client := &http.Client{}
	request, _ := http.NewRequest("GET", uri, nil)
	request.Header.Set("User-Agent", "[stuff]")
	resp, _ := client.Do(request)
	redditposts, _ := ioutil.ReadAll(resp.Body)
	//	for k, _ := range redditposts {
	//	pretty.Println(k)
	//op, _ := jq.Parse(".data.children.[" + k.(string) + "].data.preview.images.[0].source.url")
	//	op, _ := jq.Parse(".data.children.[1].data.preview.images.[0].source.url")
	op, _ := jq.Parse(".data.children")
	value, _ := op.Apply(redditposts)
	for count, _ := range value {
		pretty.Println(count)
	}
	//pretty.Println(string(value))
	//	}
}
