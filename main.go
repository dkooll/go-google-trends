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

	for _, item := range r.Channel.ItemList {
		fmt.Println(item.Title)
		fmt.Println(item.Link)
		fmt.Println(item.Traffic)
		fmt.Println("")

		for _, news := range item.NewItems {
			fmt.Println(news.Headline)
			fmt.Println(news.HeadlineLink)
			fmt.Println("")
		}
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
