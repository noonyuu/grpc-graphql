package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	greetv1 "example/gen/greet/v1"
	"example/gen/greet/v1/greetv1connect"

	"connectrpc.com/connect"
	"github.com/spf13/cobra"
)

func main() {
	fmt.Println("Starting client")

	cmd := cobra.Command{
		Use: "client",
	}
	cmd.AddCommand(&cobra.Command{
		Use: "greet1",
		RunE: func(cmd *cobra.Command, args []string) error {
			client := greetv1connect.NewGreetServiceClient(
				http.DefaultClient,
				"http://localhost:3001",
			)
			res, err := client.Greet(
				context.Background(),
				connect.NewRequest(&greetv1.GreetRequest{Name: "Jane"}),
			)
			if err != nil {
				log.Println(err)
				return err
			}
			log.Println(res.Msg.Greeting)
			return nil
		},
	})

	cmd.AddCommand(&cobra.Command{
		Use: "greet2",
		RunE: func(cmd *cobra.Command, args []string) error {
			client := greetv1connect.NewGreetServiceClient(
				http.DefaultClient,
				"http://localhost:3001",
			)
			res, err := client.Greet(
				context.Background(),
				connect.NewRequest(&greetv1.GreetRequest{Name: "わあーーー"}),
			)
			if err != nil {
				log.Println(err)
				return err
			}
			log.Println(res.Msg.Greeting)
			return nil
		},
	})

	cmd.Execute()

	// client := greetv1connect.NewGreetServiceClient(
	// 	http.DefaultClient,
	// 	"http://localhost:3001",
	// )
	// res, err := client.Greet(
	// 	context.Background(),
	// 	connect.NewRequest(&greetv1.GreetRequest{Name: "Jane"}),
	// )
	// if err != nil {
	// 	log.Println(err)
	// 	return
	// }
	// log.Println(res.Msg.Greeting)
}
