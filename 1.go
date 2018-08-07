package main

// .data.children[0].data.preview.images[0].source.url

import (
	"encoding/json"
	//	"fmt"
	"github.com/kr/pretty"
	"io/ioutil"
	"log"
	"net/http"
)

const uri = "https://www.reddit.com/r/wtfstockphotos/.json"

/*
type Message struct {
	Uri string `json:"uri"`
}

type Bucket struct {
	data    map[string]interface{} `json:"data"`
	Message Message
}
*/

func main() {
	client := &http.Client{}
	request, err := http.NewRequest("GET", uri, nil)
	if err != nil {
		log.Fatalln(err)
	}
	request.Header.Set("User-Agent", "[stuff]")
	resp, err := client.Do(request)

	redditposts, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	//fmt.Printf("%s", redditposts)
	var message map[string]interface{}
	//var message Message
	json.Unmarshal([]byte(redditposts), &message)
	//pretty.Println(message)
	posts := message["data"].(map[string]interface{})["children"].([]interface{})
	for count := range posts {
		pretty.Println(posts[count].(map[string]interface{})["data"].(map[string]interface{})["preview"].(map[string]interface{})["images"].([]interface{})[0].(map[string]interface{})["source"].(map[string]interface{})["url"])
	}
	//	pretty.Println(message["data"].(map[string]interface{})["children"].([]interface{})[0].(map[string]interface{})["data"].(map[string]interface{})["preview"].(map[string]interface{})["images"].([]interface{})[0].(map[string]interface{})["source"].(map[string]interface{})["url"])
	//things := result["data"].(map[string]interface{})
	/*
		for key, value := range result {
			fmt.Println(key, value.(string))
		}
	*/
}
