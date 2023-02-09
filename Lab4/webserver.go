package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

func main() {
	db := database{"shoes": 50, "socks": 5}
	mux := http.NewServeMux()

	mux.HandleFunc("/list", db.list)
	mux.HandleFunc("/price", db.price)
	mux.HandleFunc("/create", db.create)
	mux.HandleFunc("/update", db.update)
	mux.HandleFunc("/delete", db.delete)

	log.Fatal(http.ListenAndServe("localhost:8000", mux))
}

type dollars float32

func (d dollars) String() string { return fmt.Sprintf("$%.2f", d) }

type database map[string]dollars

func (db database) list(w http.ResponseWriter, req *http.Request) {
	for item, price := range db {
		fmt.Fprintf(w, "%s: %s\n", item, price)
	}
}

func (db database) create(w http.ResponseWriter, req *http.Request) {
	//function creates a brand new item in database with corresponding price
	item := req.URL.Query().Get("item")
	val := req.URL.Query().Get("price")

	//checks to see if item already exists in database
	if _, ok := db[item]; ok {
		w.WriteHeader(http.StatusNotFound) // 404
		fmt.Fprintf(w, "Item already exists: %q\n", item)
	} else {
		price, _ := strconv.ParseFloat(val, 32)
		db[item] = dollars(price)
		fmt.Fprintf(w, "Item Added: %s\nAt Price: %s\n", item, db[item])
	}
}

func (db database) update(w http.ResponseWriter, req *http.Request) {
	//function updates an item in database with a new price
	item := req.URL.Query().Get("item")
	val := req.URL.Query().Get("price")

	//checks to make sure item exists in database
	if _, ok := db[item]; ok {
		price, _ := strconv.ParseFloat(val, 32)
		db[item] = dollars(price)
		fmt.Fprintf(w, "Item Updated: %s\nTo Price: %s\n", item, db[item])
	} else {
		w.WriteHeader(http.StatusNotFound) // 404
		fmt.Fprintf(w, "No such item: %q\n", item)
	}
}

func (db database) delete(w http.ResponseWriter, req *http.Request) {
	//function deletes item from database
	item := req.URL.Query().Get("item")

	//checks to make sure item is in database
	if _, ok := db[item]; ok {
		delete(db, item)
		fmt.Fprintf(w, "Item Deleted: %s\n", item)
	} else {
		w.WriteHeader(http.StatusNotFound) // 404
		fmt.Fprintf(w, "No such item: %q\n", item)
	}
}

func (db database) price(w http.ResponseWriter, req *http.Request) {
	//function to return price of specific item
	item := req.URL.Query().Get("item")
	if price, ok := db[item]; ok {
		fmt.Fprintf(w, "%s\n", price)
	} else {
		w.WriteHeader(http.StatusNotFound) // 404
		fmt.Fprintf(w, "No such item: %q\n", item)
	}
}
