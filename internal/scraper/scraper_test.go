package scraper

import (
	"strings"
	"testing"
)

func TestParser(t *testing.T) {
	tests := []struct {
		name    string
		args    string
		want    []string // Want a slice of strings for multiple links
		wantErr bool
	}{
		{
			name:    "Retrieve link",
			args:    `<html><body><h1>Hello, Go!</h1><a href="/home">Home</a></body></html>`,
			want:    []string{"/home"},
			wantErr: false,
		},
		{
			name: "Multiple links",
			args: `<a href="#">This link points to nothing</a>
                                        <a href="/about">This link points to an internal page</a>
                                        <a href="https://youtube.com">This link points to an external page</a>`,
			want:    []string{"/about"},
			wantErr: false,
		},
		{
			name: "No valid link",
			args: `<a>This link points to nothing</a>
                                        <a>This link points to nothing</a>`,
			want:    []string{}, // Expect an empty set
			wantErr: false,
		},
		{
			name: "Valid link with class",
			args: `<a href="#">This link points to nothing</a>
                                        <a href="/about">This link points to an internal page</a>
                                        <a class="text-teal-400 hover:text-teal-200" href="https://youtube.com">This link points to an external page</a>`,
			want:    []string{"/about"},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rv := strings.NewReader(tt.args)
			got, err := HtmlParser(rv)
			if (err != nil) != tt.wantErr {
				t.Errorf("HtmlParser() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			// Convert the Set to a slice for easier comparison
			gotSlice := make([]string, 0, len(got))
			for _, k := range got {
				gotSlice = append(gotSlice, k)
			}

			if len(gotSlice) != len(tt.want) {
				t.Errorf("HtmlParser() returned %d links, want %d", len(gotSlice), len(tt.want))
			}

			for _, wantLink := range tt.want {
				found := false
				for _, gotLink := range gotSlice {
					if gotLink == wantLink {
						found = true
						break
					}
				}
				if !found {
					t.Errorf("HtmlParser() did not return expected link: %s", wantLink)
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
			got := IsValidUrl(tt.args)
			if got != tt.want {
				t.Errorf("IsValidUrl(%q) = %v, want %v", tt.args, got, tt.want) // More informative error message
			}
		})
	}
}
