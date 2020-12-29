package racer

import (
	"fmt"
	"net/http"
	"time"
)

/*
You have been asked to make a function called WebsiteRacer which takes two URLs and "races" them by hitting them with an
HTTP GET and returning the URL which returned first. If none of them return within 10 seconds then it should return an error.
*/
var tenSecondTimeout = 10 * time.Second

func Racer(a, b string) (winner string, error error) {
	return ConfigurableRacer(a, b, tenSecondTimeout)
}

func ConfigurableRacer(a, b string, timeout time.Duration) (winner string, error error) {
	select {
	case <-ping(a):
		return a, nil
	case <-ping(b):
		return b, nil
	case <-time.After(timeout):
		return "", fmt.Errorf("timed out waiting for %s and %s", a, b)
	}
}

/*
Always make channels
Notice how we have to use make when creating a channel; rather than say var ch chan struct{}. When you use var the variable will be initialised with the "zero" value of the type. So for string it is "", int it is 0, etc.
For channels the zero value is nil and if you try and send to it with <- it will block forever because you cannot send to nil channels
*/
func ping(url string) chan struct{} {
	ch := make(chan struct{})
	go func() {
		http.Get(url)
		close(ch)
	}()
	return ch
}
