package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
	"strconv"
)

func main() {
	db := database{"shoes": 50, "socks": 5}
	http.HandleFunc("/create", db.create)
	http.HandleFunc("/read", db.read)
	http.HandleFunc("/update", db.update)
	http.HandleFunc("/delete", db.delete)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

type dollars float32

func (d dollars) String() string { return fmt.Sprintf("$%.2f", d) }

type database map[string]dollars

func (db database) create(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")
	parsedString := strings.SplitN(item, ",", 2)

	name := ""

	var price float32

	for index, element := range parsedString {
		if index == 0 {
			name = element
		} else {
			tempPrice, err := strconv.ParseFloat(element, 32)
			price = float32(tempPrice)
			if err != nil {
    			fmt.Fprintln(w, "Price provided is invalid, item will not be created")
    			return
			}
		}
	}
	db[name] = dollars(price)
}

func (db database) read(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")
	if price, ok := db[item]; ok {
		fmt.Fprintf(w, "%s\n", price)
	} else {
		w.WriteHeader(http.StatusNotFound) // 404
		fmt.Fprintf(w, "no such item: %q\n", item)
	}
}

func (db database) update(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")
	parsedString := strings.SplitN(item, ",", 2)

	name := ""

	var price float32

	for index, element := range parsedString {
		if index == 0 {
			name = element
		} else {
			tempPrice, err := strconv.ParseFloat(element, 32)
			price = float32(tempPrice)
			if err != nil {
    			fmt.Fprintln(w, "Price provided is invalid, item will not be created")
    			return
			}
		}
	}
	if _, ok := db[item]; ok {
		db[name] = dollars(price)
	} else {
		w.WriteHeader(http.StatusNotFound) // 404
		fmt.Fprintf(w, "no such item: %q\n", name)
	}
}

func (db database) delete(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")
	if price, ok := db[item]; ok {
		fmt.Fprintf(w, "%s\n", price)
	} else {
		w.WriteHeader(http.StatusNotFound) // 404
		fmt.Fprintf(w, "no such item: %q\n", item)
	}
}
