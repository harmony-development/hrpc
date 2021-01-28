package main

import (
	"context"
	"fmt"
	"hrpc/sample/gen/hrpc/jiti"
	"log"
	"net/http"
	"os"

	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/descriptorpb"
)

func main() {
	if os.Args[1] == "server" {
		handler := jiti.NewMuHandler(&MuServ{}, func(err error, w http.ResponseWriter) {
			log.Fatal(err)
		})
		handler.UnaryPre = func(d *descriptorpb.FileDescriptorProto, f func(c context.Context, req proto.Message, headers http.Header) (proto.Message, error)) func(c context.Context, req proto.Message, headers http.Header) (proto.Message, error) {
			println(*d.Name)
			return f
		}
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
			client.Mu(&jiti.Ping{
				Mu: "unary",
			})
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
