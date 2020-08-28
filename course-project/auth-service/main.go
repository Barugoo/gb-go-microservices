package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func main() {
	resp, err := http.Get("https://google.com")
	check(err)
	body, err := ioutil.ReadAll(resp.Body)
	check(err)
	fmt.Println(len(body))
}

func check(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

type DBClient struct {
	dbLatencyMetric prometheus.SummeryVec
	*sqlx.Driver
}

func (db *DBClient) CreateItem(item *Item) (*Item, error) {
	defer func(){
		db.dbLatencyMetric.With(someLabel).Observer(time.Since(start))
	}
	start := time.Now()
	return db.createItem(item)
}

func (db *DBClient) createItem(item *Item) (*Item, error) {
	res, err := db.SQL("INPUT ")
	if err != nil {
		nil, err
	}
	return Item, err
}
