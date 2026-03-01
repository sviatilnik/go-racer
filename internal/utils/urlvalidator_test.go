package utils

import "testing"

func TestIsURLValid(t *testing.T) {
	tests := []struct {
		url  string
		want bool
	}{
		{url: "", want: false},
		{url: " ", want: false},
		{url: "https://google.com", want: true},
		{url: "http://google.com", want: true},
		{url: "google.com", want: false},
		{url: "домен-на-русском.рф", want: false},
		{url: "https://домен-на-русском.рф", want: true},
	}

	for _, test := range tests {
		if IsURLValid(test.url) != test.want {
			t.Errorf("Incorrect URL check %s %v", test.url, test.want)
		}
	}
}
