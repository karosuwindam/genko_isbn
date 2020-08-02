package main

import (
	"fmt"

	"github.com/seihmd/openbd"
)

func main() {
	o := openbd.New()

	// 一冊の情報を取得
	book, err := o.Get("9784040645858")
	if err != nil {
		return
	}
	fmt.Println(book.GetTitle())
	fmt.Println(book.GetAuthor())
	fmt.Println(book.GetCover())
	fmt.Println(book.Hanmoto)
	fmt.Println(book.Onix.DescriptiveDetail.Contributor[0].PersonName.Content)
	fmt.Println(book.Onix.DescriptiveDetail.Contributor[1].PersonName.Content)
	fmt.Println(book.Onix.DescriptiveDetail)
	fmt.Println(book.Onix.CollateralDetail)

}
