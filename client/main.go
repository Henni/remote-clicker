package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/go-vgo/robotgo"
	hook "github.com/robotn/gohook"
)

type Config struct {
	Address  string
	Mappings map[string]string
}

func hookFunc(operation string, key string, payload Config) {
	robotgo.EventHook(hook.KeyDown, []string{key}, func(e hook.Event) {
		fmt.Println(operation, key)
		url := payload.Address + "/click?op=" + operation
		_, err := http.Get(url)
		if err != nil {
			fmt.Println(err)
		}
	})
}

func main() {
	var payload Config
	content, err := ioutil.ReadFile("./client-config.json")
	fmt.Println(string(content))
	err = json.Unmarshal(content, &payload)
	if err != nil {
		log.Fatal("Error during Unmarshal(): ", err)
	}
	log.Printf("config: %+v\n", payload)

	// fmt.Println("start client")
	// conn, err := net.Dial("tcp", payload.Address)
	// if err != nil {
	// 	log.Fatal("Connection error", err)
	// }

	for operation, key := range payload.Mappings {
		hookFunc(operation, key, payload)
	}

	s := robotgo.EventStart()
	<-robotgo.EventProcess(s)
	// conn.Close()
	fmt.Println("done")
}
