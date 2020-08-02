package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	"./searchapi"
)

func urlAnalysis(url string) []string {
	tmp := []string{}
	for _, str := range strings.Split(url[1:], "/") {
		tmp = append(tmp, str)
	}
	return tmp
}

func apiserver(w http.ResponseWriter, r *http.Request) {
	urldata := urlAnalysis(r.URL.Path)
	log.Printf("request:%v\n", r.URL.Path)
	if len(urldata) > 1 {
		switch urldata[1] {
		case "":
			w.WriteHeader(400)
			fmt.Fprintf(w, "Err API request")
		case "isbn":
			if len(urldata) > 2 {
				var adata searchapi.AmazonNameType
				start := time.Now()
				ch1 := make(chan bool)
				if len(urldata[2]) == 13 {
					go func() {
						adata = searchapi.GetPageAmazonURL(urldata[2])
						ch1 <- true
					}()
				} else {
					go func() {
						ch1 <- true
					}()
				}
				data := searchapi.GetOpenBdData(urldata[2])
				if data.Title == "" || data.Image == "" {
					<-ch1
					if data.Title == "" {
						data.Title = adata.Title
					}
					if data.Image == "" {
						data.Image = adata.Image
					}

				}
				jsondata, err := json.Marshal(data)
				end := time.Now()
				if err == nil {
					fmt.Fprintf(w, "{\"Data\":%s,\"Time\":%f}", jsondata, (end.Sub(start)).Seconds())
				} else {
					w.WriteHeader(400)
					fmt.Fprintf(w, "Err API request")
				}
			} else {
				w.WriteHeader(400)
				fmt.Fprintf(w, "Err API request")
			}
		default:
			fmt.Fprintf(w, "%s", r.URL.Path)
		}
	} else {
		w.WriteHeader(400)
		fmt.Fprintf(w, "Err API request")
	}
	return

}
