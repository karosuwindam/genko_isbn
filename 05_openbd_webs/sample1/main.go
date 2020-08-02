package main

import (
	"log"
)

func main() {
	var web WebSetupData
	err := web.websetup()
	if err == nil {
		web.webstart()
	} else {
		log.Fatal(err)
	}
}
