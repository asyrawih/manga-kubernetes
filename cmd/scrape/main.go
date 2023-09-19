package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"sync"
	"time"

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

	m := GetListComics(doc)

	var wg sync.WaitGroup

	start := time.Now()
	for _, list := range m[:1] {
		wg.Add(1)
		go func(mg Mangalist) {
			// First get will incur cloudflare challenge
			resp := Fetch(mg.url)
			defer resp.Body.Close()
			rc := GetChapterList(resp)
			b, _ := json.Marshal(rc)
			_ = os.WriteFile("some.json", b, 0644)
			wg.Done()
		}(list)
	}

	wg.Wait()

	fmt.Printf("time.Since(start).String(): %v\n", time.Since(start).String())
}

func Fetch(url string) *http.Response {
	// First get will incur cloudflare challenge
	resp, err := cfscrape.Get(url)
	if err != nil {
		panic(err)
	}

	return resp
}

type Mangalist struct {
	title string
	url   string
}

type ReadChapter struct {
	Chapter string   `json:"chapter,omitempty"`
	Images  []string `json:"images,omitempty"`
}

func GetChapterList(resp *http.Response) []ReadChapter {
	var readChapter []ReadChapter
	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	doc.Find(".komik_info-chapters-wrapper li").Each(func(_ int, s *goquery.Selection) {
		nodes := s.Find("a")
		link, _ := nodes.Attr("href")
		chap := nodes.Text()
		r := Fetch(link)
		images := GetChapterImages(r)
		readChapter = append(readChapter, ReadChapter{
			Chapter: chap,
			Images:  images,
		})
	})
	return readChapter
}

func GetListComics(doc *goquery.Document) []Mangalist {
	var komikUrls []Mangalist
	doc.Find(".list-update_items-wrapper a").Each(func(_ int, s *goquery.Selection) {
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

func GetChapterImages(resp *http.Response) []string {
	var imageURI []string
	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	doc.Find(".chapter_body").Each(func(_ int, s *goquery.Selection) {
		s.Find(".main-reading-area img").Each(func(_ int, s *goquery.Selection) {
			imagesURL, _ := s.Attr("src")
			imageURI = append(imageURI, imagesURL)
		})
	})
	return imageURI
}
