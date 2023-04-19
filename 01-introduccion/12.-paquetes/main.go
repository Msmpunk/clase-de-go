package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type Data struct {
	ID   int
	Name string
}

var dataList = []Data{
	{ID: 1, Name: "John"},
	{ID: 2, Name: "Jane"},
}

func main() {
	http.HandleFunc("/data", dataHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func dataHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		getData(w, r)
	case "POST":
		postData(w, r)
	case "DELETE":
		deleteData(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func getData(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(dataList)
}

func postData(w http.ResponseWriter, r *http.Request) {
	var data Data
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	dataList = append(dataList, data)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}

func deleteData(w http.ResponseWriter, r *http.Request) {
	var data Data
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	for i, d := range dataList {
		if d.ID == data.ID {
			dataList = append(dataList[:i], dataList[i+1:]...)
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(data)
			return
		}
	}
	http.Error(w, "Data not found", http.StatusNotFound)
}
