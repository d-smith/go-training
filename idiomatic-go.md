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



* Go build, golint, go vet
