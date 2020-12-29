# Learn Go with Tests


![https://quii.gitbook.io/learn-go-with-tests/](https://gblobscdn.gitbook.com/assets%2F-L9Tqx5WSaiE4u24Pk05%2F-LmyGcIwYFqlc-kxzIGZ%2F-LXAJRSbtm02phRcFvU4%2Fred-green-blue-gophers-smaller.png?alt=media)

### Resource
https://quii.gitbook.io/learn-go-with-tests/


### Discipline
CYCLE OF TDD (RED GREEN REFACTOR)
* Write a test
* Make the compiler pass
* Run the test, see that it fails and check the error message is meaningful
* Write enough code to make the test pass
* Refactor

On the face of it this may seem tedious but sticking to the feedback loop is important.

Not only does it ensure that you have relevant tests, it helps ensure you design good software by refactoring with the safety of tests.

Seeing the test fail is an important check because it also lets you see what the error message looks like. As a developer it can be very hard to work with a codebase when failing tests do not give a clear idea as to what the problem is.

By ensuring your tests are fast and setting up your tools so that running tests is simple you can get in to a state of flow when writing your code.

By not writing tests you are committing to manually checking your code by running your software which breaks your state of flow and you won't be saving yourself any time, especially in the long run.

Benchmarking
---
_Source : https://quii.gitbook.io/learn-go-with-tests/go-fundamentals/iteration_

Writing benchmarks in Go is another first-class feature of the language and it is very similar to writing tests.

```go
func BenchmarkRepeat(b *testing.B) {
    for i := 0; i < b.N; i++ {
        Repeat("a")
    }
}
```
You'll see the code is very similar to a test.

The testing.B gives you access to the cryptically named b.N.
When the benchmark code is executed, it runs b.N times and measures how long it takes.

The amount of times the code is run shouldn't matter to you, the framework will determine what is a "good" value for that to let you have some decent results.

To run the benchmarks do go test -bench=. (or if you're in Windows Powershell go test -bench=".")

`go test -bench=.`

```bash
goos: darwin
goarch: amd64
pkg: github.com/quii/learn-go-with-tests/for/v4
10000000           136 ns/op
PASS
```
What 136 ns/op means is our function takes on average 136 nanoseconds to run (on my computer). Which is pretty ok! To test this it ran it 10000000 times.

NOTE by default Benchmarks are run sequentially.


 Mocking
--
https://quii.gitbook.io/learn-go-with-tests/go-fundamentals/mocking#mocking

Mutex
--
We've covered a few things from the sync package
* Mutex allows us to add locks to our data
* Waitgroup is a means of waiting for goroutines to finish jobs

When to use locks over channels and goroutines?
--
We've previously covered goroutines in the first concurrency chapter which let us write safe concurrent code so why would you use locks?
The go wiki has a page dedicated to this topic; Mutex Or Channel

"A common Go newbie mistake is to over-use channels and goroutines just because it's possible, and/or because it's fun. Don't be afraid to use a sync.Mutex if that fits your problem best. Go is pragmatic in letting you use the tools that solve your problem best and not forcing you into one style of code."

#### Paraphrasing:
* Use channels when passing ownership of data 
* Use mutexes for managing state

