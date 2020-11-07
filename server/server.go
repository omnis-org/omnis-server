package server

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func LaunchServer() {
	http.HandleFunc("/", HelloServer)
	http.ListenAndServe(":8000", nil)
}

func HelloServer(rw http.ResponseWriter, req *http.Request) {
	decoder := json.NewDecoder(req.Body)
	var o Output
	err := decoder.Decode(&o)
	if err != nil {
		fmt.Errorf("Error during decoding of json input", err)
	}
	log.Println(o)
}
