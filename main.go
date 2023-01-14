package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

type RSS struct {
	XMLName xml.Name `xml:"rss"`
	Channel *Channel `xml:"channel"`
}

type Channel struct {
	Title    string `xml:"title"`
	ItemList []Item `xml:"item"`
}

type Item struct {
	Title    string `xml:"title"`
	Link     string `xml:"link"`
	Traffic  string `xml:"approx_traffic"`
	NewItems []News `xml:"news_item"`
}

type News struct {
	Headline     string `xml:"news_item_title"`
	HeadlineLink string `xml:"news_item_url"`
}

func main() {
	var r RSS
	data := readGoogleTrends()

	err := xml.Unmarshal(data, &r)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	for i := range r.Channel.ItemList {
		rank := (i + 1)
		fmt.Println("#", rank)
		fmt.Println("Title: ", r.Channel.ItemList[i].Title)
		fmt.Println("Link: ", r.Channel.ItemList[i].Link)
		fmt.Println("Headline: ", r.Channel.ItemList[i].NewItems[0].Headline)

		// for j := range r.Channel.ItemList[i].NewItems {
		// 	fmt.Println("Headline: ", r.Channel.ItemList[i].NewItems[j].Headline)
		// 	fmt.Println("Headline Link: ", r.Channel.ItemList[i].NewItems[j].HeadlineLink)
		// }
	}
}

func getGoogleTrends() *http.Response {
	resp, err := http.Get("https://trends.google.com/trends/trendingsearches/daily/rss?geo=NL")

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	return resp
}

func readGoogleTrends() []byte {
	resp := getGoogleTrends()
	data, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	return data
}
