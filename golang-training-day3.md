## Go Lang Training - Day 3

### Concurrency Patterns - Pooling

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


### Testing

* Tests in the same package - access to all the exported and unexported
* Tests in a separate package - can access only the unexported identifiers
* Goconvey - constant feedback on tests, good for services/products/etc.
* Team needs to standardize on output content and format for tests.
* Example tests - based on the APIs
* Example test - Example + API Func, e.g. ExampleStart, has //Output: at the
end specifiying the desired output
* Example tests also produce documentation

### Benchmarking

* Use the bytes package instead of the strings package for string manipulation,
bytes has the same api as the strings function, much much faster. Note: check
to see if this can work with multibyte characters or not - might not work
when working with multibyte strings.

### Branch Prediction

* You get some mechanical sympathy by following the idioms, but being aware of
the hardware will let you get more performance out of your code.
* How can we write code to let processors make accurate predictions?
* If you can remove the randomness from your code, you can help the hardware
prediction.

### Caching

* Memory is dealt with in cache lines, which move things in chunks of 64K
* This means we want our data to be contiguous, not randomly spread around. Can we
organize our data across 64K blocks?
* Work with things in a predictable way, e.g. try to read within the same cache
line.
* Slices are an important tool for arranging data in a predictable way.

### Profiling

* Note that profiling on the mac requires hacking the kernel - refer to the notes
* Note the http profiling port - https://golang.org/pkg/net/http/pprof/
* Docker container?

### Go Debug Environment Variable

* Use the GODEBUG environment variable to get information about the runtime -
http://golang.org/pkg/runtime/
* Note this kicks in when running go tooling, so set it just when needed.
* Detailed - M threads, P processors, G go routines

### Standard Library

* Study the io and time packages.



### Go Tooling

<pre>
MACLB015803:example1$ go build
../../../../../ArdanStudios/gotraining/09-testing/01-testing/example1/mongodb/mongodb.go:13:2: cannot find package "gopkg.in/mgo.v2" in any of:
	/usr/local/go/src/gopkg.in/mgo.v2 (from $GOROOT)
	/Users/example/goprojects/src/gopkg.in/mgo.v2 (from $GOPATH)
../../../../../ArdanStudios/gotraining/09-testing/01-testing/example1/buoy/buoy.go:15:2: cannot find package "gopkg.in/mgo.v2/bson" in any of:
	/usr/local/go/src/gopkg.in/mgo.v2/bson (from $GOROOT)
	/Users/example/goprojects/src/gopkg.in/mgo.v2/bson (from $GOPATH)
MACLB015803:example1$ go get
MACLB015803:example1$ go build
</pre>
