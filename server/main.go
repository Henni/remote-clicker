package server

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/go-vgo/robotgo"
)

type Coords struct {
	X int
	Y int
}

func clickHandler(w http.ResponseWriter, r *http.Request) {
	var payload map[string]Coords
	content, err := ioutil.ReadFile("./server-config.json")
	err = json.Unmarshal(content, &payload)
	if err != nil {
		log.Fatal("Error during Unmarshal(): ", err)
	}
	log.Printf("origin: %+v\n", payload)

	queryParams := r.URL.Query()
	operation := queryParams.Get("op")

	if value, exists := payload[operation]; exists {
		log.Println(value)
		robotgo.MouseSleep = 100
		robotgo.Move(value.X, value.Y)
		robotgo.Click()
		log.Println("clicked")
	}
}

func main() {
	http.HandleFunc("/click", clickHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
