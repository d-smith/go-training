Go Lang Training - Day 3

Concurrency Patterns - Pooling

* Design points - what are the guarantees and responsibilities in the API you
are designing using channels, go routines, etc.
* Buffered channels - arbitrary sizing means you are delaying dealing with blocking,
and no clean semantics for handing out work
* Pattern - pool with unbuffered channel, sender blocks but is guaranteed that the
work had been handed off and is being worked on once the send completes
* GOMAXPROCS shouldn't be larger than the numner of CPUs on your machine.
* Work - https://github.com/ArdanStudios/gotraining/blob/master/08-concurrency_patterns/work/work.go
* For range on channel - you are blocked on receive (e.g. <- pool), for range quits
when the channel is closed.
* Pool - use buffered channel to manage a pool of resources - https://github.com/ArdanStudios/gotraining/tree/master/08-concurrency_patterns/pool
* Appropriate aliasing - when you read `func New(fn func() (io.Closer, error), size uint) (*Pool, error)`,
if this was aliased, the reader of the code then has to go look for the base type
definition. Aliasing is appropriate when you are adding behavior, with no added
behavior it reduces readability.
* APIs - code for what it does today, don't speculatively add things you think you
might need in the future.
* Closure example - https://github.com/ArdanStudios/gotraining/blob/master/08-concurrency_patterns/pool/main/main.go -
the query variable is passed to the go routine so they can have their own unique value,
otherwise they would share q which would get set to the max value in the loop


Testing

* Tests in the same package - access to all the exported and unexported
* Tests in a separate package - can access only the unexported identifiers
* Goconvey - constant feedback on tests, good for services/products/etc.
* Team needs to standardize on output content and format for tests.
* Example tests - based on the APIs
* Example test - Example + API Func, e.g. ExampleStart, has //Output: at the
end specifiying the desired output
* Example tests also produce documentation

Benchmarking

* Use the bytes package instead of the strings package for string manipulation,
bytes has the same api as the strings function, much much faster


Go Tooling

<pre>
MACLB015803:example1 a045103$ go build
../../../../../ArdanStudios/gotraining/09-testing/01-testing/example1/mongodb/mongodb.go:13:2: cannot find package "gopkg.in/mgo.v2" in any of:
	/usr/local/go/src/gopkg.in/mgo.v2 (from $GOROOT)
	/Users/a045103/goprojects/src/gopkg.in/mgo.v2 (from $GOPATH)
../../../../../ArdanStudios/gotraining/09-testing/01-testing/example1/buoy/buoy.go:15:2: cannot find package "gopkg.in/mgo.v2/bson" in any of:
	/usr/local/go/src/gopkg.in/mgo.v2/bson (from $GOROOT)
	/Users/a045103/goprojects/src/gopkg.in/mgo.v2/bson (from $GOPATH)
MACLB015803:example1 a045103$ go get
MACLB015803:example1 a045103$ go build
</pre>
