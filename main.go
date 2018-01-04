package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	_ "github.com/lib/pq"
	"github.com/rachmatprabowo/redsfin2/conf"
	"github.com/rachmatprabowo/redsfin2/core"
)

func main() {
	t := time.Now()
	fmt.Print(t, "\n")
	fmt.Print("Initializing...\n")
	conf.InitDB()
	conf.InitRoute()
	fmt.Print("Runnig...\n")

	router := core.NewRouter()
	log.Fatal(http.ListenAndServe(":8000", router))
}
