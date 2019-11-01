package main

import (
	"fmt"
	"net/http"
	"runtime/pprof"
)


func main(){
	cache:=make(chan int,4)
	var name chan string
	go func() {
		for i:=0;i< 10;i++ {
			cache<-i
			fmt.Println("1234")
			go func() {
				select {
				case <-cache:
					go func() {
						select {
						case <-name:
						}
					}()
					return
				}
			}()
		}

	}()

	http.HandleFunc("/", handler)
	http.ListenAndServe(":11181", nil)
	select {}
}

func handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain")

	p := pprof.Lookup("goroutine")
	p.WriteTo(w, 1)
}
