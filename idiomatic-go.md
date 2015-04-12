Go idioms

* Use var for zero value declarations, use short variable decl syntax in all
other cases.
* Struct literal declaration - only include the fields that you are initializing to
non-zero values.
* Leverage constants of a kind
* For-range : idiomatic way to iterate through an array or a slice
* Interface names - er in the interface name for single method interface, e.g.
notifier interface, notify method.
* New - idiomatic for creating stuff. If you see a New in a package, use it
as a factory, and don't make copies of the returned value.
* Error variables start with Err
* Customer error type names end with Error
* Type as context pattern
* Channel with an empty struct type - typically used for notification, for
instance system shutdown.
* If you hand off a channel to sig.notify you should not close that channel, as a
send on it later will cause a panic
* Work with values until you need to share; create and work with the value until you
need to share it. Exception - just creating the value to share as a pointer.
* If the standard library has an interface with the method names you need, use that
instead of creating your own interface.
* Named return values w/naked returns reduce the readability of the code.
* Keep acquisition and release close to each other, typically done using defer.
* fmt - "fumpt"
* Web site - source graph



Misc

* Go build, golint, go vet
* Slack community - sign up via a icon on the http://blog.gopheracademy.com/ website.
