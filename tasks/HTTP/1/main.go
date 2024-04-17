package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/jackc/pgx/v5"
)

type Events struct {
	Date     string `json:"date"`
	TypeEv   string `json:"typeEv"`
	Location string `json:"location"`
}

const portNumber = ":8081"

var myMap = make(map[string]Events)

func main() {
	fmt.Println("START")
	event := Events{}
	event2 := Events{"20.12.1999", "birthday", "Dagestan"}
	file, err := os.Open("my.json")
	if err != nil {
		panic(err)
	}
	fileRaw, err := io.ReadAll(file)
	if err != nil {
		panic(err)
	}

	err = json.Unmarshal(fileRaw, &event)
	if err != nil {
		panic(err)
	}
	fmt.Println(event)
	urlExample := "postgresql://postgres:password!@localhost:5432/Events"
	conn, err := pgx.Connect(context.Background(), urlExample)
	if err != nil {
		panic(err)
	}
	defer conn.Close(context.Background())
	insertstr := fmt.Sprintf("INSERT INTO Events VALUES ( '%v','%v','%v' )", event.Date, event.TypeEv, event.Location)
	fmt.Println(insertstr)
	myMap[event2.Date] = event2
	myMap[event.Date] = event

	conn.Exec(context.Background(), insertstr)
	http.HandleFunc("/", HomePageHandler)
	http.HandleFunc("/a", CreateEvHandler)
	err = http.ListenAndServe(":8081", nil)
	if err != nil {
		panic(err)
	}

}
func HomePageHandler(w http.ResponseWriter, r *http.Request) {
	//date := r.URL.Query().Get("25.02.2022")
	for _, ev := range myMap {
		eventJson, err := json.Marshal(ev)
		if err != nil {
			panic(err)
		}
		w.Write(eventJson)
	}
}

func CreateEvent(date, typeEv, location string) (*Events, error) {
	return &Events{
		Date:     date,
		TypeEv:   typeEv,
		Location: location,
	}, nil
}

func CreateEvHandler(w http.ResponseWriter, r *http.Request) {
	date := r.FormValue("date")
	typeEv := r.FormValue("typeEv")
	location := r.FormValue("location")
	event, _ := CreateEvent(date, typeEv, location)
	responseJSON(w, http.StatusOK, map[string]interface{}{"result": event})
}
func responseJSON(w http.ResponseWriter, statusCode int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(data)
}
