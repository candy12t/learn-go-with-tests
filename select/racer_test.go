package racer

import "testing"

func TestRacer(t *testing.T) {
	slowURL := "http://wwww.facebook.com"
	fastURL := "http://wwww.quii.co.uk"

	want := fastURL
	got := Racer(slowURL, fastURL)

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}
