package main

import (
	"fmt"
	"log"
	"os"

	"github.com/PuerkitoBio/goquery"
	"github.com/iain17/go-cfscrape"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Printf("Usage: %s url\n", os.Args[0])
	}

	url := os.Args[1]

	// First get will incur cloudflare challenge
	resp, err := cfscrape.Get(url)
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	// Load the HTML document
	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	s := GetListComics(doc)
	fmt.Println(s)

	// GetChapterImages(doc)

}

type Mangalist struct {
	title string
	url   string
}

func GetListComics(doc *goquery.Document) []Mangalist {
	var komikUrls []Mangalist
	doc.Find(".list-update_items-wrapper a").Each(func(i int, s *goquery.Selection) {
		val, _ := s.Attr("href")

		title := s.Find("a .list-update_item-info h3")
		ret, _ := title.Html()

		komikUrls = append(komikUrls, Mangalist{
			title: ret,
			url:   val,
		})
	})
	return komikUrls
}

func GetChapterImages(doc *goquery.Document) {
	doc.Find(".chapter_body").Each(func(i int, s *goquery.Selection) {
		s.Find(".main-reading-area").Each(func(i int, s *goquery.Selection) {
			imagesUrl, _ := s.Html()
			fmt.Println(imagesUrl)
		})
	})
}
