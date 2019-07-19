package racer

import (
	"fmt"
	"net/http"
	"time"
)

func measureResponseTime(url string) time.Duration {
	start := time.Now()
	_, _ = http.Get(url)
	return time.Since(start)
}

func ping(url string) chan bool {
	ch := make(chan bool)
	go func() {
		_, _ = http.Get(url)
		ch <- true
	}()

	return ch
}

var tenSecondTimeout = 10 * time.Second

func Racer(a string, b string) (winner string, error error) {
	return ConfigurableRacer(a, b, tenSecondTimeout)
}

func ConfigurableRacer(a string, b string, timeout time.Duration) (winner string, error error) {
	//aDuration := measureResponseTime(a)
	//bDuration := measureResponseTime(b)
	//
	//if aDuration < bDuration {
	//	return a
	//}
	//return b

	select {
	case <-ping(a):
		return a, nil
	case <-ping(b):
		return b, nil
	case <-time.After(timeout):
		return "", fmt.Errorf("timed out waiting for %s and %s", a, b)
	}
}
