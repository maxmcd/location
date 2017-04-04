package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/api/place/board-bitmap", bitmapHandler)
	http.ListenAndServe(":4040", handlers.LoggingHandler(os.Stdout, r))
}

func bitmapHandler(w http.ResponseWriter, req *http.Request) {
	client := &http.Client{}
	w.Header().Add("Access-Control-Allow-Origin", "*")
	request, _ := http.NewRequest("GET",
		"https://www.reddit.com/api/place/board-bitmap", nil)
	resp, err := client.Do(request)
	if err != nil {
		panic(err)
	}
	bodyBytes, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(bodyBytes[4] >> 4)
	fmt.Println(bodyBytes[4] % 16)
	w.Write(bodyBytes)
}
