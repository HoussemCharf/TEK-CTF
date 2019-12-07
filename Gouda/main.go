package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
)

var (
	clients = make(map[string]int)
)

func input(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "input")
}
func home(w http.ResponseWriter, r *http.Request) {
	inter := r.PostFormValue("inter")
	if inter == "" {
		http.ServeFile(w, r, "./instruction.html")
		return
	}
	// anti bruteforce

	clients[r.RemoteAddr] = clients[r.RemoteAddr] + 1
	if clients[r.RemoteAddr] > 20 {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}
	// handling conversion
	n, _ := strconv.Atoi(inter)

	if n == 209 {
		data, _ := ioutil.ReadFile("flag")
		fmt.Fprintf(w, string(data))
		return
	}
	fmt.Fprintf(w, "ZmxhZ3tUcnlfSGFyZGVyIX0K")
	return

}
func main() {
	http.HandleFunc("/", home)
	http.HandleFunc("/input", input)
	fmt.Println("Server is up and running!")
	http.ListenAndServe(":8005", nil)

}
