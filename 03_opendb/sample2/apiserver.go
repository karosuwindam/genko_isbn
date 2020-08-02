package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"
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
				start := time.Now()
				data := getOpenBdData(urldata[2])
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
