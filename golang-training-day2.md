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
