package main

import (
	"fmt"

	"github.com/gocolly/colly"
)

// type item struct {
// 	Name   string `json:"name"`
// 	Price  string `json:"price"`
// 	ImgUrl string `json:"imgurl"`
// }

func main() {
	scaperUrl := "https://hcmut.edu.vn/"
	c := colly.NewCollector(
		colly.AllowedDomains("www.hcmut.edu.vn", "hcmut.edu.vn"),
	)
	// c.OnHTML(".head-text", func(h *colly.HTMLElement) {
	// 	fmt.Println(h.Text)
	// })

	c.OnHTML("a[href]", func(h *colly.HTMLElement) {
		h.Request.Visit(h.Attr("href"))
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println(fmt.Sprintf("Visiting %s", r.URL))
	})

	c.OnError(func(r *colly.Response, err error) {
		fmt.Printf("error while scraping: %s/n", err.Error())
	})

	c.Visit(scaperUrl)

}
