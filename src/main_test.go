package main

import (
	"testing"
)

func TestParser(t *testing.T) {

	tests := []struct {
		route   string
		args    string
		want    string
		wantErr bool
	}{{
		route:   "/home",
		args:    `<html><body><h1>Hello, Go!</h1><a href="/home">Home</a></body></html>`,
		want:    "Home",
		wantErr: false,
	},
	}
	for _, tt := range tests {
		t.Run(tt.route, func(t *testing.T) {
			got, err := HtmlParser(tt.args)
			if (err != nil) != tt.wantErr {
				t.Errorf("HtmlParser() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("HtmlParser() = %v, want %v", got, tt.want)
			}
		})
	}
}
