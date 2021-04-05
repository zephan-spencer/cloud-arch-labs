package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"sync"
)

func main() {
	db := database{data: map[string]dollars{"shoes": 50, "socks": 5}}

	http.HandleFunc("/create", db.create)
	http.HandleFunc("/read", db.read)
	http.HandleFunc("/update", db.update)
	http.HandleFunc("/delete", db.delete)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

type dollars float32

func (d dollars) String() string { return fmt.Sprintf("$%.2f", d) }

type database struct {
	data map[string]dollars
	sync.RWMutex
}

func (db database) create(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query()

	name := item.Get("name") 
	price, err := strconv.ParseFloat(item.Get("price"), 32)

	if err != nil {
    	fmt.Fprintln(w, "Price provided is invalid, item will not be created")
    	return
	}
	db.Lock()
	db.data[name] = dollars(price)
	db.Unlock()
}

func (db database) read(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")
	if price, ok := db.data[item]; ok {
		fmt.Fprintf(w, "%s\n", price)
	} else {
		w.WriteHeader(http.StatusNotFound) // 404
		fmt.Fprintf(w, "no such item: %q\n", item)
	}
}

func (db database) update(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query()
	name := item.Get("name")
	price, err := strconv.ParseFloat(item.Get("price"), 32)

	if err != nil {
    	fmt.Fprintln(w, "Price provided is invalid, item will not be updated")
    	return
	}

	if _, ok := db.data[name]; ok {
		db.Lock()
		db.data[name] = dollars(price)
		db.Unlock()
	} else {
		w.WriteHeader(http.StatusNotFound) // 404
		fmt.Fprintf(w, "no such item: %q\n", name)
	}
}

func (db database) delete(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")
	if price, ok := db.data[item]; ok {
		fmt.Fprintf(w, "%s\n", price)
	} else {
		w.WriteHeader(http.StatusNotFound) // 404
		fmt.Fprintf(w, "no such item: %q\n", item)
	}
}
