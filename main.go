package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	http.HandleFunc("/", recordingHandler)

	log.Fatal(http.ListenAndServe(":10080", nil))
}

func recordingHandler(w http.ResponseWriter, r *http.Request) {
	body, _ := ioutil.ReadAll(r.Body)
	defer r.Body.Close()

	os.Mkdir("./log", 0777)
	filename := "log/req_" + fmt.Sprint(time.Now().Unix()) + ".log"
	logfile, err := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}
	defer logfile.Close()

	io.WriteString(io.MultiWriter(logfile, os.Stdout), string(body))
}
