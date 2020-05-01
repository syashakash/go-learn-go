package _select // underscore added as prefix to package name as select is a keyword
import "testing"

func TestRacer(t *testing.T)  {
	slowURL := "http://www.facebook.com"
	fastURL := "http://www.google.com"

	want := fastURL
	got :=  Racer(slowURL, fastURL)
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}
