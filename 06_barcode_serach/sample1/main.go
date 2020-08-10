package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

type DbsetupData struct {
	Dbtype     string `json:type`
	Dbuser     string `json:user`
	Dbpassword string `json:password`
	Ipaddr     string `json:host`
	Port       string `json:port`
	Databs     string `json:database`
}

const dbtype = "mysql"
const dbuser = "bookserver"
const dbpassword = "bookserver"
const conecttype = "tcp"
const ipaddr = "mysql_host"
const port = "3306"
const databs = "isbn_bookbase"

var bookdatatable BookData

// var userstable UserData

func dbsetup() DbsetupData {
	var tmp = DbsetupData{}
	var buf bytes.Buffer

	config_json := "config/dbsetup.json"
	raw, err := ioutil.ReadFile(config_json)
	if err != nil {
		tmp.Dbtype = dbtype
		tmp.Dbuser = dbuser
		tmp.Dbpassword = dbpassword
		tmp.Ipaddr = ipaddr
		tmp.Port = port
		tmp.Databs = databs
		fp, err := os.Create(config_json)
		if err != nil {
			panic(err)
		}
		// jsonエンコード
		outputJson, err := json.Marshal(&tmp)
		if err != nil {
			panic(err)
		}
		json.Indent(&buf, outputJson, "", "  ")
		fp.Write(buf.Bytes())
		fp.Close()

	} else {
		var fc DbsetupData
		json.Unmarshal(raw, &fc)
		if fc.Dbtype == "" {
			fc.Dbtype = dbtype
		}
		if fc.Databs == "" {
			fc.Dbuser = dbuser
		}
		if fc.Dbpassword == "" {
			fc.Dbpassword = dbpassword
		}
		if fc.Ipaddr == "" {
			fc.Ipaddr = ipaddr
		}
		if fc.Port == "" {
			fc.Port = port
		}
		if fc.Databs == "" {
			fc.Databs = databs
		}
		// fmt.Println(fc)
		tmp = fc

	}

	return tmp
}

var sqltype string

func main() {
	var web WebSetupData
	dbdata := dbsetup()
	var err error

	err = bookdatatable.bookdatas_sql.SqlSetup(dbdata.Dbtype, dbdata.Dbuser, dbdata.Dbpassword, conecttype, dbdata.Ipaddr, dbdata.Port, dbdata.Databs)
	// sqltype = dbdata.Dbtype
	// sqltype = "sqlite3"
	// err = blogstable.blogs_sql.SqlSetup(sqltype, "./test.db")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer bookdatatable.bookdatas_sql.CloseDB()

	// userstable.users_sql = blogstable.blogs_sql

	bookdatatable.Sqlsetup("bookdatabase")

	err = web.websetup()
	if err == nil {
		web.webstart()
	} else {
		log.Fatal(err)
	}
}
