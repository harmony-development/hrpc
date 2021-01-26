package main

import (
	"hrpc/sample/gen/hrpc/jiti"
	"log"
	"net/http"
	"os"
)

func main() {
	if os.Args[1] == "server" {
		handler := jiti.NewMuHandler(&MuServ{}, func(err error, w http.ResponseWriter) {
			log.Fatal(err)
		})
		err := http.ListenAndServe(":9999", handler)
		if err != nil {
			log.Fatal(err)
		}
		return
	}

	client := jiti.NewMuClient("http://localhost:9999")
	resp, err := client.Mu(&jiti.Ping{
		Mu: "mu",
	})
	if err != nil {
		log.Fatal(err)
	}

	println(resp.Mu)
}
