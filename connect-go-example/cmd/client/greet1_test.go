package main

import (
	"bytes"
	"context"
	"example/gen/greet/v1/greetv1connect"
	"testing"
)

func Test_greet1(t *testing.T) {
	type args struct {
		Client greetv1connect.GreetServiceClient
		ctx    context.Context
		name   string
	}
	tests := []struct {
		name       string
		args       args
		wantWriter string
		wantErr    bool
	}{
		// TODO: Add test cases.
		{
			name: "may write name",
			args:       args{
				Client: nil,
				ctx:    context.Background(),
				name:   "Jane",
			},
			wantWriter: "hello",
			wantErr:    false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			writer := &bytes.Buffer{}
			if err := greet1(tt.args.Client, tt.args.ctx, tt.args.name, writer); (err != nil) != tt.wantErr {
				t.Errorf("greet1() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotWriter := writer.String(); gotWriter != tt.wantWriter {
				t.Errorf("greet1() = %v, want %v", gotWriter, tt.wantWriter)
			}
		})
	}
}
