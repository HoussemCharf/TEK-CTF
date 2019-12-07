package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
)

var (
	key   = []byte("V5oSAv5948pmCXp0gNiC3EldjEPdmxRp")
	store = sessions.NewCookieStore(key)
	// generator vars
	lenny  = "( ͡° ͜ʖ ͡°)"
	data   = ""
	data2  = ""
	length = 1000
)

// structs defined for solver

// Direction structure
type Direction struct {
	East  int
	North int
}

// Directions Map
var Directions = []Direction{
	Direction{0, 1},  // east
	Direction{1, 0},  // north
	Direction{0, -1}, // west
	Direction{-1, 0}, // south
}

// Position current
type Position struct {
	Northing int
	Easting  int
}

// State structure
type State struct {
	Position
	Facing int
}

// Turn action
func (s State) Turn(steps int) State {
	s.Facing = (s.Facing + steps + len(Directions)) % len(Directions)
	return s
}

// Walk action
func (s State) Walk() State {
	s.Northing += Directions[s.Facing].North
	s.Easting += Directions[s.Facing].East
	return s
}

// Distance to be walked
func (s State) Distance() int {
	return Abs(s.Northing) + Abs(s.Easting)
}

// Abs returns Absolute value
func Abs(n int) int {
	if n >= 0 {
		return n
	}
	return -n

}

// solver will return the corrent answer for a specific input
func solver(filename string) int {
	s := State{}
	buf, _ := ioutil.ReadFile(filename)
	for _, instruction := range strings.Split(strings.TrimSpace(string(buf)), ", ") {
		rotation := string(instruction[0])
		blocks, _ := strconv.Atoi(instruction[1:])
		if rotation == "L" {
			s = s.Turn(1)
		} else if rotation == "R" {
			s = s.Turn(-1)
		}
		for i := 0; i < blocks; i++ {
			s = s.Walk()
		}
	}
	return s.Distance()
}

// generator creates a specific input for different users
func generator(filename string) {
	rand.Seed(time.Now().UnixNano())
	Directions := []rune{'L', 'R'}
	numbers := []int{1, 2, 3, 4, 5, 188}

	for x := 0; x < length; x++ {
		choosenDirection := string(Directions[rand.Intn(len(Directions))])
		choosenSteps := strconv.Itoa(numbers[rand.Intn(len(numbers))])
		data = data + choosenDirection + choosenSteps + lenny
		data2 = data2 + choosenDirection + choosenSteps + ", "
		if x%25 == 0 {
			data = data + "\n"
		}
	}
	f, _ := os.Create(filename + ".txt")
	f2, _ := os.Create(filename)
	f.WriteString(data)
	f2.WriteString(data2)

}

func flag(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "ZmxhZ3tUcnlfSGFyZGVyIX0K")
}

func home(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "cookie-name")
	// check if session assigned or not
	if auth, ok := session.Values["authenticated"].(bool); !ok || !auth {
		// handle session
		session.Values["authenticated"] = true
		sessionID := rand.Int31()
		session.Values["id"] = int(sessionID)
		session.Save(r, w)
		// create a file with solution here
		generator(strconv.Itoa(int(sessionID)))
		fmt.Fprintf(w, "You've been authentificated")
		return
	}

	// fetching attemp from URL
	attemp := r.URL.Query().Get("attemp")
	// handling simple get server
	if attemp == "" {
		w.Header().Add("Default", "Y2hlY2sgL2QK")
		id, _ := session.Values["id"].(int)
		session.Values["answer"] = solver(strconv.Itoa(id))
		session.Save(r, w)
		fmt.Fprintf(w, lenny)
		return
	}
	
	n, err := strconv.Atoi(attemp)
	if err != nil || n < 0 {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}
	answer, _ := session.Values["answer"].(int)
	if n == answer{
		data, _ := ioutil.ReadFile("flag")
		fmt.Fprintf(w, string(data))
		return
	}
	fmt.Fprintf(w, "ZmxhZ3tUcnlfSGFyZGVyIX0K\n ")
	return
}


// robots will handle the robots rabbit whole
func robots(w http.ResponseWriter, r *http.Request) {
	data, _ := ioutil.ReadFile("robots.txt")
	fmt.Fprintf(w, string(data))
}

// debug this function will hand the task description to competitors
func debug(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "cookie-name")
	if auth, ok := session.Values["authenticated"].(bool); ok && auth {
		session.Values["ok"] = true
		http.ServeFile(w, r, "./instruction.html")
	}
}
func input(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "cookie-name")
	if auth, ok := session.Values["authenticated"].(bool); ok && auth {
		id, _ := session.Values["id"].(int)
		http.ServeFile(w, r, "./"+strconv.Itoa(id)+".txt")
	}
	http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
	return
}

// login function will serve the rabbit whole
func login(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./static/index.html")
}
func cheese(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./static/cheese.gif")
}
func main() {
	// assigning mux
	r := mux.NewRouter()
	r.HandleFunc("/input", input)
	r.HandleFunc("/robots.txt", robots)
	r.HandleFunc("/", home)
	r.HandleFunc("/d", debug)
	r.HandleFunc("/flag", flag)
	r.HandleFunc("/login", login)
	r.HandleFunc("/cheese.gif", cheese)
	fmt.Println("Server is up and running!")
	http.ListenAndServe(":8080", r)
}
