package main

import (
	"fmt"

	"github.com/PuerkitoBio/goquery"
)

func GetPageAmazonBook(isbm string) string {
	output := ""
	url := "https://www.amazon.co.jp/s?k=" + isbm + "&i=stripbooks"
	doc, err := goquery.NewDocument(url)
	if err != nil {
		fmt.Println(err.Error())
		return output
	}
	doc.Find("div.sg-col-4-of-12.sg-col-8-of-16.sg-col-12-of-32.sg-col-12-of-20.sg-col-12-of-36.sg-col.sg-col-12-of-24.sg-col-12-of-28").Each(func(i int, s *goquery.Selection) {
		if i == 0 {
			output = s.Find("span.a-size-medium.a-color-base.a-text-normal").Text()
		}

	})
	return output
}

func main() {
	isbm := "9784040645858"
	out := GetPageAmazonBook(isbm)
	fmt.Println(out)
}
