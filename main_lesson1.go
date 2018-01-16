package main

import (
	"fmt"
	"net/http"
	"encoding/json"
//	"log"
	"strconv"
)
var m = map[string]string {
	"fruit" : "apple",
	"person" : "farmer",
}
// type == class
type Location struct {
	// 用`json : "lab" `的方式，就是在生成JSON的时候，会变成 {"lat" : xxx}。 By default 是 { "Lat" : xxx }
	Lat float64 `json:"lat"`
	Lon float64 `json:"lon"`
}

type Post struct {
	// `json:"user"` is for the json parsing of this User field. Otherwise, by default it's 'User'.
	User     string `json : "user"`
	Message  string  `json : "message"`
	Location Location `json : "location"`
}

const (
	DISTANCE = "200km"
)


func main() {

	//fmt.Println("started-service")
	//http.HandleFunc("/search", handlerSearch)
	//http.HandleFunc("/searchJSON", handlerSearchJSON)
	//log.Fatal(http.ListenAndServe(":8081", nil))

	//m = make(map[string]string) //make map, initialize and read to use
	////m[“fruit”] = “apple”
	////m[“person”] = “farmer”

	for k, v := range m {
		fmt.Printf("Key: %s --> Value: %s\n", k, v)
	}

	nums := [] int {2, 3, 4}
	for i, num := range nums {
		fmt.Printf("At index %s: %s\n", i, num)
	}


}

//* indicates it's a pointer, as required by the http.Request
//w: writer, write to response; r: request
// := is "define AND assign"
func handlerPost(w http.ResponseWriter, r *http.Request) {
	// Parse from body of request to get a json object.
	fmt.Println("Received one post request")

	decoder := json.NewDecoder(r.Body)

	var p Post //Post p = new Post();

	if err := decoder.Decode(&p); err != nil {
		panic(err)
		return
	}

	// the code block above is equivalent to:
	/*
	err : = decoder.Decode(&p)
	if (err != nil) {
		panic(err)
		return
	}
	 */

	fmt.Fprintf(w, "Post received: %s\n", p.Message)
}

func handlerSearchLesson1(w http.ResponseWriter, r *http.Request) {
	// Parse from body of request to get a json object.
	fmt.Println("Received one search request")
	//http://localhost:8080/search?lat=10.0&lon=20.0
	Lat := r.URL.Query().Get("lat");
	Lon := r.URL.Query().Get("lon");


	fmt.Fprintf(w, "Search request received: Lat = %s & Lon = %s\n", Lat, Lon)
}

func handlerSearchJSONLesson1(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Received one search JSON request")

	//string converted to Float_64, would return 2 variable results
	// _ in place of the second returned result, becuase we are confident we won't use it
	// so we skipped naming it
	lat, _ := strconv.ParseFloat(r.URL.Query().Get("lat"), 64)
	lon, _ := strconv.ParseFloat(r.URL.Query().Get("lon"), 64)

	rang := DISTANCE

	if val := r.URL.Query().Get("range"); val != "" {
		rang = val + "km"
	}

	fmt.Fprintf(w, "Search received: lat = %f ; lon = %f ; rang = %s\n", lat, lon, rang)

	p := &Post {
		User : "1111",
		Message : "一生必去的100个地方",
		Location : Location {
			Lat : lat,
			Lon : lon,
		},
	}

	//js, err := json.Marshal(p); //把 p 拍扁然后写进一个json文本
	if js, err := json.Marshal(p); err != nil {
		panic(err)
		return
	} else {

		w.Header().Set("Content-Type", "application/json")
		w.Write(js)
	}


}



