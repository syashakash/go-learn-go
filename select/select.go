package _select

import (
	"net/http"
)

func Racer(a, b string) (winner string)  {
	select {
	case <- ping(a):
		return a
		case <- ping(b):
			return b
	}
}

func ping(url string) chan struct{} {
	ch := make(chan struct{}) // chan is struct as there is no memory allocation required.
	go func() {
		http.Get(url)
		close(ch)
	}()
	return ch
}