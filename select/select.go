package _select

import (
	"net/http"
	"time"
)

func Racer(a, b string) (winner string)  {
	startA := time.Now()
	http.Get(a)
	aDuration := time.Since(startA)

	http.Get(b)
	bDuration := time.Since(b)
	if aDuration < bDuration {
		return a
	}
	return b
}
