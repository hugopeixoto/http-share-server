package main

import "errors"
import "net/http"
import "log"
import "fmt"
import "io/ioutil"
import "crypto/sha1"
import "flag"

var path = flag.String("path", "uploads", "where to store uploaded files")
var port = flag.Int("port", 4444, "http server port")
var domain = flag.String("domain", "", "")

var username = flag.String("user", "", "")
var password = flag.String("pass", "", "")

func UploadContent(w http.ResponseWriter, r *http.Request) {
	user, pass, ok := r.BasicAuth()
	if !ok || user != *username || pass != *password {
		Error(w, http.StatusUnauthorized, errors.New("wrong/missing credentials"))
		return
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil || len(body) == 0 {
		Error(w, http.StatusBadRequest, errors.New("Unable to read contents"))
		return
	}

	id := fmt.Sprintf("%x", sha1.Sum(body))[0:7]

	filename := "/" + id + "-" + r.URL.Query().Get("name")

	err = ioutil.WriteFile(*path+filename, body, 0644)

	if err != nil {
		Error(w, http.StatusInternalServerError, nil)
		return
	}

	Data(w, http.StatusOK, *domain+filename)
}

func main() {
	flag.Parse()
	http.HandleFunc("/upload", UploadContent)
	log.Fatal(http.ListenAndServe(fmt.Sprintf("localhost:%v", *port), nil))
}
