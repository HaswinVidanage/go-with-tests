package context

import (
	"context"
	"fmt"
	"net/http"
)

type Store interface {
	Fetch(ctx context.Context) (string, error)
}

/*
context has a method Done() which returns a channel which gets sent a signal when the context is "done" or "cancelled".
We want to listen to that signal and call store.Cancel if we get it but we want to ignore it if our Store manages to Fetch before it.

To manage this we run Fetch in a goroutine and it will write the result into a new channel data. We then use select to effectively race to the two asynchronous processes and then we either write a response or Cancel
*/
func Server(store Store) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		data, err := store.Fetch(r.Context())

		if err != nil {
			return // todo: log error however you like
		}

		fmt.Fprint(w, data)
	}
}

/*
We can see after this that the server code has become simplified as it's no longer explicitly responsible for cancellation, it simply passes through context and relies on the downstream functions to respect any cancellations that may occur.
*/
