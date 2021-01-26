package main

import (
	"fmt"
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

	client := jiti.NewMuClient("localhost:9999")
	in, out, err := client.MuMute()
	if err != nil {
		log.Fatal(err)
	}

	go func() {
		for {
			mu := ""
			fmt.Scanln(&mu)
			fmt.Printf("sending %s\n", mu)
			in <- &jiti.Ping{
				Mu: mu,
			}
		}
	}()

	for {
		data, ok := <-out
		if !ok {
			break
		}

		fmt.Printf("got %s\n", data.Mu)
	}
}
