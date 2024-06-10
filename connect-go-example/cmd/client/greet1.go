package main

import (
	"context"
	greetv1 "example/gen/greet/v1"
	"example/gen/greet/v1/greetv1connect"
	"io"
	"log"

	"connectrpc.com/connect"
)

func greet1(Client greetv1connect.GreetServiceClient, ctx context.Context, name string, writer io.Writer) error {
	res, err := Client.Greet(
		ctx,
		connect.NewRequest(&greetv1.GreetRequest{Name: name}),
	)
	if err != nil {
		log.Println(err)
		return err
	}
	log.Println(res.Msg.Greeting)
	return nil
}
