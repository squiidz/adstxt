package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-zoo/bone"
	"github.com/go-zoo/claw"
	mw "github.com/go-zoo/claw/middleware"
)

var (
	dbpath   = flag.String("dbpath", "./ads.txt.db", "ads.txt database path")
	httpPort = flag.String("http", "8080", "http port")
)

func main() {
	flag.Parse()
	mux := bone.New()
	clw := claw.New(mw.NewLogger(os.Stdout, "||ADSD|| ", 2))
	srvr := NewServer(*dbpath)

	mux.PostFunc("/ads/populate", srvr.addPublisher)
	mux.GetFunc("/ads/:publisher", srvr.getPublisher)

	log.Printf("Server running on port %s\n", *httpPort)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", *httpPort), clw.Merge(mux)))
}
