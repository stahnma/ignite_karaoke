package main

// jq format .data.children[0].data.preview.images[0].source.url

import (
	"encoding/json"
	"fmt"
	"github.com/jmoiron/jsonq"
	"github.com/kr/pretty"
	"io/ioutil"
	"net/http"
	"strings"
)

//const uri = "https://www.reddit.com/r/wtfstockphotos/.json?limit=100"

//const uri = "https://www.reddit.com/r/predators/.json"

func dedupe(a []string) []string {
	result := []string{}
	seen := map[string]string{}
	for _, val := range a {
		if _, ok := seen[val]; !ok {
			result = append(result, val)
			seen[val] = val
		}
	}
	return result
}

func get_the_pics(uri string) []string {
	var pics []string
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
		pic, err := jq.String("data", "preview", "images", "0", "source", "url")
		if err != nil {
			continue
		}
		//pretty.Println(pic)
		pics = append(pics, pic+"\n")
	}

	return pics
}

func main() {
	base := "https://www.reddit.com/r/wtfstockphotos"
	suffixes := []string{"/.json?limit=100", "/top/.json?limit=100", "/new/.json?limit=100", "/top/.json?sort=top&t=month&limit=100", "/top/.json?sort=top&t=all&limit=100"}
	//suffixes := []string{"/.json?limit=100"}
	var pics []string
	for _, suf := range suffixes {
		pretty.Println(suf)
		uri := base + suf
		pretty.Println("URI is " + uri)
		pics = append(pics, get_the_pics(uri)...)
	}
	pics = dedupe(pics)
	fmt.Println(len(pics))
}
