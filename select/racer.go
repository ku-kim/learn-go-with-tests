package main

import (
	"net/http"
	"time"
)

func Racer(a, b string) (winner string) {
	aDuration := measureReponseTime(a)
	bDuration := measureReponseTime(b)

	if aDuration < bDuration {
		return a
	}

	return b
}

func measureReponseTime(str string) time.Duration {
	startA := time.Now()
	http.Get(str)
	aDuration := time.Since(startA)
	return aDuration
}
