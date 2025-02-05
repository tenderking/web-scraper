package main

import (
	"strings"
	"testing"
	"web-scraper/internal/scraper"
)

func TestParser(t *testing.T) {

	tests := []struct {
		name    string
		args    string
		want    string
		wantErr bool
	}{
		{
			name:    "Retrieve link",
			args:    `<html><body><h1>Hello, Go!</h1><a href="/home">Home</a></body></html>`,
			want:    "/home",
			wantErr: false,
		},
		{name: "Return the valid link",
			args: `<a href="#">This link points to nothing</a>
						<a href="/about">This link points to an internal page</a>
						<a href="https://youtube.com">This link points to an external page</a>
						`,
			want:    "/about",
			wantErr: false,
		},
		{name: "Return the valid link",
			args: `<a>This link points to nothing</a>
						<a href="/about">This link points to an internal page</a>
						<a href="https://youtube.com">This link points to an external page</a>
						`,
			want:    "/about",
			wantErr: false,
		},
		{name: "Return the valid link when class",
			args: `<a>This link points to nothing</a>
						<a href="/about">This link points to an internal page</a>
						<a class="text-teal-400 hover:text-teal-200" href="https://youtube.com">This link points to an external page</a>
						`,
			want:    "/about",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rv := strings.NewReader(tt.args)
			got, err := scraper.HtmlParser(rv)
			if (err != nil) != tt.wantErr {
				t.Errorf("HtmlParser() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			for k := range got {

				if k != tt.want {
					t.Errorf("HtmlParser() = %v, want %v", got, tt.want)
				}
			}
		})
	}
}

func TestValidUrl(t *testing.T) {

	tests := []struct {
		name    string
		args    string
		want    bool
		wantErr bool
	}{
		{
			name:    "Check invalid",
			args:    "#",
			want:    false,
			wantErr: false,
		},
		{
			name:    "Check valid",
			args:    "/about",
			want:    true,
			wantErr: false,
		},
		{
			name:    "Check query params",
			args:    "/anime?name=mha",
			want:    true,
			wantErr: false,
		},
		{
			name:    "Check external",
			args:    "https://youtube.com",
			want:    false,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := scraper.IsValidUrl(tt.args)

			if got != tt.want {
				t.Errorf("HtmlParser() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSetCount(t *testing.T) {

	tests := []struct {
		name    string
		args    scraper.Set
		want    int
		wantErr bool
	}{
		{
			name:    "items count",
			args:    scraper.Set{"/a": {}, "/b": {}},
			want:    2,
			wantErr: false,
		}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.args.Count()

			if got != tt.want {
				t.Errorf("HtmlParser() = %v, want %v", got, tt.want)
			}
		})
	}
	tests_2 := []struct {
		name    string
		args    [2]string
		want    int
		wantErr bool
	}{
		{
			name:    "items count",
			args:    [...]string{"/a", "b"},
			want:    2,
			wantErr: false,
		}}

	for _, tt := range tests_2 {
		new_set := scraper.NewSet()
		for _, k := range tt.args {
			new_set.Add(k)
		}

		t.Run(tt.name, func(t *testing.T) {
			got := new_set.Count()
			if got != tt.want {
				t.Errorf("HtmlParser() = %v, want %v", got, tt.want)
			}
		})
	}

}
