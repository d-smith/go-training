Go Lang Training - Day 1

Statically typed language

* the more you know about the type system, the better you will do programming
with go
* the most important thing to know to understand a statically typed language
* what's 00001010? depends on the type
* Type - need the size and representation, if the compiler knows this it can
check things for safety, then generate assembler/machine code
* Can make performance optimizations based on type info, e.g. size used for
memory optimization

Idioms

* embrace the idioms
* mechanically sympathetic language
* understand how the hardware works, idioms align with how the hardware works
* idiomatic - the preferred way to write the code

Specification

* go spec os locked in - all version 1 language code and libraries have backwards
compatibility guarantees

Implementation

* constantly changing
* understanding the implementation helps you debug, understand how things work,
etc.
* language is pretty stable, current focus is on performance
* Good to understand one level below where you work, understand what's going
on under the hood

Types

* int is architecture dependent - it's a word (might be 32 bits, might be 64),
size is taken from the architecture (for example go playground runs on 32 it hardware).
* int - more efficient to work with the natural word size, use int instead of int32
and int64. Use the sized for interoperability with other libraries, protocols,
writing records to disk, etc.
* float - no machine architecture standard
* string - immutable reference type

Variables

* http://play.golang.org/p/6w6hBNE75a
* var - declare a variable, creates a value set to its zero value. Idiomatic
use is to use var when creating a variable set to its zero value
* use short variable declaration operator when setting a new var to a non-zero value
* think zero value vs not zero value decl - not zero means short decl syntax
* No casting in go - there is conversion.
* No implicit conversion in go lang.

Structs

* type: size and representation information
* zero value for a struct type - each item in the struct is set to its zero value
* struct - composite type
* can use short decl with struct literal syntax
* types in go are allocated on alignement boundaries (4 - 8 bytes), hardware optimization
(think hardware caching)
* can look at the internals using some of the unsafe routines
* Can declare anonymous structs
<pre>
e := struct {
  flag bool
  counter int64
}{
  flag: true,
  counter: 10
}
</pre>
* Use case for anonymous struct: type localization, maybe deal with a web request
where we don't need to share that type around.
