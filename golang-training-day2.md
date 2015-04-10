Go Training - Day 2

Packaging

* Semantic units of code you would find in a folder
* Compiler builds all the packages, then links them together
* Package denotes the boundaries for visibility/encapsulation
* In a package, things are either exported or unexported from the package
* Focus is on the API for the package, API is all the exported stuff, think about
the usability of the interface.
* Main package file name is the same name as the package
* Package management vs version management - use vendoring for version management
* Can create return values of unexported types via a New and the short variable
declaration syntax. Generally frowned upon, golint will complain.
* Package name is the last folder name on the path
* Study the time and the io package

Code Files - Readability

* Understand the intent of the package, then read the code from top to bottom
* Story - the product you are building, what it does. Packages - outline of the
story, reading the code should not require leaving the source code file to
understand the story.
* The standard library is the model for good packaging and API design.
* What do the imports look like?
* Code as prose

Using Pointers

* Types have a nature
* Primitive types - built ins, reference types, typically not shared, you provide
copies
* Complex types - structs, generally you don't make copies, you share them
* Complex types can be designed to be used as primitives, e.g. Time is
a primitive data type - see
https://play.golang.org/p/xD6PCx--GG
* General guideline is reference types use value receivers, and complex
types use pointer receivers
* Exception: unmarshal needs to use a shared value, only exception to the rule
in the standard library.

Error Handling

* Think about this in terms of API design
* Errors in go are just values, can take on any behavior we need
* Errors are a way of giving your users a level of respect and the ability
to make decisions about what to do next when an error occurs.
* Use the error part of your API to do things beyond log and die - recovery,
resiliency.
* Errors are returned as interface values
* nil - Interface type that has no concrete implemenation - two word interface
value
* Use the default error type when you can if it gives enough context. This
pattern shows how to use the default error type with more context - https://play.golang.org/p/iNrGaGYZ-o
* Good to provide the error vars to allow the consumer to discriminate the errors - the
error variable provides the context - the string is for logging. Do
not make consumers reply on the error text as this can create coupling not subject
to impact analysis - don't break consumers based on changing messages.
* Error values also provide a backwards compatibility path as well
* If a default error type and/or error variables aren't sufficient use a custom
error type - http://play.golang.org/p/Eu3X54PnWm also see the JSON package in the
standard library
* Switch type as context - e := err.(type) - conversion of the interface type to the
concrete type
* Extract concrete type from an interface type
`err1, ok := err.(*UnmarshalTypeError)` - ok is true if err was an UnmarshallTypeError plus
err1 can now be worked with as a pointer of that type
* net package - example of complex error handling, error types with methods.
* net package OpErr, has isTemporary method
* This example - returns error with the type of *customError, nil value - fix is to
change the signature to return error, not customerError in fail method - http://play.golang.org/p/czXpjvWWTT
* Use pointer receivers for custom errors - protect the user from overwriting your
values in your API. Error values are unique based the address they are created with (when
  based on the default error type pointer receiever)
* If an error occurs, ignore other parameters. As API provider, return zero values for the
other params.
* Return default error type pointers, not your custom error type.
