package main

// .data.children[0].data.preview.images[0].source.url

import (
	"encoding/json"
	"github.com/jmoiron/jsonq"
	"github.com/kr/pretty"
	"io/ioutil"
	"net/http"
	"strings"
)

const uri = "https://www.reddit.com/r/wtfstockphotos/.json?limit=100"

//const uri = "https://www.reddit.com/r/predators/.json"

func main() {
	client := &http.Client{}
	request, _ := http.NewRequest("GET", uri, nil)
	request.Header.Set("User-Agent", "[stuff]")
	resp, _ := client.Do(request)
	redditposts, _ := ioutil.ReadAll(resp.Body)
	data := map[string]interface{}{}
	dec := json.NewDecoder(strings.NewReader(string(redditposts)))
	dec.Decode(&data)
	jq := jsonq.NewQuery(data)
	posts, err := jq.ArrayOfObjects("data", "children")
	if err != nil {
		pretty.Println(err)
	}
	for _, val := range posts {
		jq = jsonq.NewQuery(val)
		pics, err := jq.String("data", "preview", "images", "0", "source", "url")
		if err != nil {
			continue
		}
		pretty.Println(pics)
	}

}
