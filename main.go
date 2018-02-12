package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net"
	"net/http"
	"net/http/fcgi"
	"os"
	"runtime/pprof"
	"time"

	"github.com/gorilla/mux"
)

type Test1 struct {
	Query  int `json:"query"`
	Result int `json:"result"`
}

type Test2 struct {
	QueryA int `json:"queryA"`
	QueryB int `json:"queryB"`
	QueryC int `json:"queryC"`
	Result int `json:"result"`
}

func fib(n int) int {
	if n < 2 {
		return n
	}
	return fib(n-1) + fib(n-2)
}

func fibsum(a, b, c int) int {
	fiba := fib(a)
	fibb := fib(b)
	fibc := fib(c)
	return fiba + fibb + fibc
}

func test1Handler(w http.ResponseWriter, r *http.Request) {
	//t := time.Now().Format("20060102-030405")
	cpuprofile := fmt.Sprintf("/root/test1-cpu-%s.pprof", time.Now().Format("030405"))
	f, err := os.Create(cpuprofile)
	if err != nil {
		log.Fatal("could not create CPU profile: ", err)
	}
	log.Print("Start CPU profile")
	if err := pprof.StartCPUProfile(f); err != nil {
		log.Fatal("could not start CPU profile: ", err)
	}
	defer func() {
		pprof.StopCPUProfile()
		log.Print("Stop CPU profile")
	}()

	query := 40
	result := fib(query)
	data := Test1{query, result}

	bytes, err := json.Marshal(data)
	if err != nil {
		log.Print("Error: Cannot Marshal")
		return
	}

	fmt.Fprintf(w, string(bytes))
}

/*
func cpuProfile(a, b, c int) int {
	cpuprofile := fmt.Sprintf("/root/cpuProfile-cpu-%s.pprof", time.Now().Format("030405"))
	f, err := os.Create(cpuprofile)
	if err != nil {
		log.Fatal("could not create CPU profile: ", err)
	}
	log.Print("Start CPU profile")
	if err := pprof.StartCPUProfile(f); err != nil {
		log.Fatal("could not start CPU profile: ", err)
	}
	defer func() {
		pprof.StopCPUProfile()
		log.Print("Stop CPU profile")
	}()

	return fibsum(a, b, c)
}
*/

func test2Handler(w http.ResponseWriter, r *http.Request) {
	cpuprofile := fmt.Sprintf("/root/test2-cpu-%s.pprof", time.Now().Format("030405"))
	f, err := os.Create(cpuprofile)
	if err != nil {
		log.Fatal("could not create CPU profile: ", err)
	}
	log.Print("Start CPU profile")
	if err := pprof.StartCPUProfile(f); err != nil {
		log.Fatal("could not start CPU profile: ", err)
	}
	defer func() {
		pprof.StopCPUProfile()
		log.Print("Stop CPU profile")
	}()

	queryA := 33
	queryB := 34
	queryC := 35
	result := fibsum(queryA, queryB, queryC)
	//result := cpuProfile(queryA, queryB, queryC)
	data := Test2{queryA, queryB, queryC, result}

	bytes, err := json.Marshal(data)
	if err != nil {
		log.Print("Error: Cannot Marshal")
		return
	}

	fmt.Fprintf(w, string(bytes))
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/test1", test1Handler).Methods("GET")
	r.HandleFunc("/test2", test2Handler).Methods("GET")

	fmt.Println("Starting server...")
	l, _ := net.Listen("tcp", "127.0.0.1:9000")
	fcgi.Serve(l, r)
}
