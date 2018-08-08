package main

//TODO
// * Ensure image isn't a 404
// * Check if there are any NSFW markers and remove those
// * Allow support for mulitple subreddits
// * Allow control/start/stop/pause
// * Servce in golang and not python
// * Single binary
// * CLI arguments or env vars for which subreddits

// jq format .data.children[0].data.preview.images[0].source.url

import (
	"encoding/json"
	"github.com/jmoiron/jsonq"
	"github.com/kr/pretty"
	"io/ioutil"
	"math/rand"
	"net/http"
	"os"
	"strings"
	"text/template"
	"time"
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
		pics = append(pics, pic)
	}

	return pics
}
func contains(a []string, x string) bool {
	for _, n := range a {
		if x == n {
			return true
		}
	}
	return false
}

func grab20(pics []string) []string {
	// We want to randomly grab 20 images out of our set
	if len(pics) < 20 {
		return pics
	}
	seed := rand.NewSource(time.Now().Unix())
	r := rand.New(seed)
	var n int
	var con []string
	for i := 0; i < 20; i++ {
		n = r.Intn(len(pics))
		if !contains(con, pics[n]) {
			con = append(con, pics[n])
		} else {
			i--
		}

	}

	return con
}

func main() {
	var base string
	var suffixes []string
	var pics []string
	base = "https://www.reddit.com/r/"
	subreddits := []string{"funnystockpics", "wtfstockphotos", "earthporn", "disneyvacation"}
	//suffixes = []string{"/.json?limit=100", "/top/.json?limit=100", "/new/.json?limit=100", "/top/.json?sort=top&t=month&limit=100", "/top/.json?sort=top&t=all&limit=100"}
	suffixes = []string{"/.json?limit=100"}
	for _, sub := range subreddits {
		for _, suf := range suffixes {
			//pretty.Println(suf)
			uri := base + sub + suf
			//pretty.Println("URI is " + uri)
			pics = append(pics, get_the_pics(uri)...)
		}
	}
	pics = dedupe(pics)
	//fmt.Println(len(pics))

	the20 := grab20(pics)

	tmpl := template.New("test")
	tmpl, err := tmpl.ParseFiles("page.tmpl")
	if err != nil {
		pretty.Println(err)
	}
	err = tmpl.ExecuteTemplate(os.Stdout, "page.tmpl", the20)
	if err != nil {
		pretty.Println(err)
	}
}
