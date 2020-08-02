package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
)

var (
	HTML_ROOT = "./html"
)

func Exists(filename string) bool {
	_, err := os.Stat(filename)
	return err == nil
}

func viewcgi(w http.ResponseWriter, r *http.Request) {
	upath := r.URL.Path
	tmp := strings.Split(upath, "/")
	if r.Method == "POST" {
		if err := r.ParseForm(); err != nil {
		}
		fmt.Fprintf(w, "POST:")
		for k, v := range r.Form {
			fmt.Fprintf(w, "\n%v : %v", k, v)
		}
	} else {
		fmt.Fprintf(w, "GET:%s", tmp[2])
		if err := r.ParseForm(); err != nil {
		}
		for k, v := range r.Form {
			fmt.Fprintf(w, "\n%v : %v", k, v)
		}
	}
}

//ReadHtmlはpathに入力したファイルパスから読み取る
//pathはテキストパスでそのテキスト値をもとに戻す

func ReadHtml(path string) string {
	var output string
	fp, err := os.Open(path)
	if err != nil {
		log.Panic(err)
		return ""
	}
	defer fp.Close()
	buf := make([]byte, 1024)
	for {
		n, err := fp.Read(buf)
		if err != nil {
			break
		}
		if n == 0 {
			break
		}
		output += string(buf[:n])
	}
	return output
}

//静的HTMLのページを返す
func viewhtml(w http.ResponseWriter, r *http.Request) {
	upath := r.URL.Path
	tmp := map[string]string{}
	if !strings.HasPrefix(upath, "/") {
		upath = "/" + upath
		r.URL.Path = upath
	}
	if upath == "/" {
		upath += "index.html"
	}
	if !Exists(HTML_ROOT + upath) {
		w.WriteHeader(404)
		tmp["urlpath"] = r.URL.Path
		fmt.Fprint(w, ReadHtml(HTML_ROOT+"/404.html"))
		return
	} else {
		fmt.Fprint(w, ReadHtml(HTML_ROOT+upath))
	}

}

func main() {
	http.HandleFunc("/", viewhtml)
	http.HandleFunc("/cgi/", viewcgi)
	http.ListenAndServe(":8080", nil)
}
