Go Lang Training - Day 1

Statically typed language

* the more you know about the type system, the better you will do programming
with go
* the most important thing to know to understand a statically typed language
* what's 00001010? depends on the type
* Type - specifies the size and representation, if the compiler knows this it can
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
* string - two word header value data structure, 1st word contains a pointer to an underlying
array of bytes, the second word contains the length in bytes

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

Pointers

* http://play.golang.org/p/cpCcLsdbM6
* Variables have a place in memory and a value
* Where is the memory location?
* Everything is go is passed by value (what you see is what you get)
* Stack vs heap allocation - compiler does escape analysis and decides where to
allocate the memory. See https://github.com/ArdanStudios/gotraining/tree/master/01-language_syntax/03-pointers/example1
* Note escape analysis is also in play when using new
*Escape analysis is an implementation detail - any monkeying you do based on
escape analysis might be invalidated in a language release, for example 1.5 will
make changes to the escape analysis done in 1.4
* Use variable name - value
* Ampersand operator - where is the value
* Params - local variable for the function, sits in the stack frame for the function,
stacks are allocated down in memory.
* Function returns - think of it as a stack pop, allocated memory is 'gone'
* Pointers are about sharing - if we want to share something else where in your
program, use a pointer.

Sharing with Pointers

* http://play.golang.org/p/j4uDMFJqiF
* Everything is about the type, what does * mean? It affects the type information,
means we declare a pointer variable (has a value and an address)
* The value of a pointer variable is always an address
* Size of pointer var is architecture dependent - addresses fit into a word
* star operator in an operation - value the pointer points to
* Can pass the address of composite type parts http://play.golang.org/p/cK1_GFyDOo
* Address that points to a location that contains an int - *int

Performance

* Avoid premature optimization, follow the idioms and keep things simple,
benchmark your app, then go off the reservation based on the benchmark results.
* Idiomatic go is what the langauge team will focus on when optimizing performance.

Misc
* Look at the first sharing instance and comment out the stuff that prevents
optimization (fmt.XXX calls use reflection, which means escape...), looks like
stack grows up
* Secondary memory allocators pattern


Constants

* https://github.com/ArdanStudios/gotraining/tree/master/01-language_syntax/04-constants
* Have a parallel type system, designed around values being mathematically exact
* Minimum 256 bit precision
* Constants can have a type, but can also have a kind (can be implicitly converted
  by the compiler)
* Exist only at compile time
* Literal values in go are un-named constants of a kind

Type Conversions

* API - anything that is exposed is part of your API.
* Named types - mimic size and representation of another type
* Promotion - type overrides the kind
* Constants of a kind - compiler can do implicit conversion, constants
of a type - subject to type checking/enforcement
* API can leverage constants of a kind to allow using operations with constants
to produce things of the right type (think Duration and constances in the time package)

Scope

* A pair of curly brackets declare a new level of scope
* Here err is defined in the scope of the if statement
* If you want the returned value you can't combine the call and the error check (?)

Arrays and Slices

* Linked lists are an artifact from ye olden days - randomized memory access
defeats processor prediction, fetch ahead for data and instructions, etc.
* Mechanical sympathy - keep the data we are working with as contiguous as
possible
* Array size - part of the type information.
* For range avoids going past the array boundary, explicit indexing requires
code to be generated by the compiler to check the index value.
* Arrays are fixed length, hard to use sometimes, especially when working with
arrays of different length, things that might change, etc.

Slices

* Slice - reference type: when we create a value of this type, we create a header
value that points to some other memory
* make - create channels, maps, and slices
* A slice is a 3 word data structure - 1st is pointer to underlying array,
2nd is the length, 3rd is the capacity.
* Nil slice - nil, 0, 0 var s []string
* Empty slice - s := []string{} length and capacity of 0, pointer to an empty array,
  can be used to marshall a json array with nothing in it.
* We never share reference types - we pass the value of the reference type around
* Append - returns a new slice header
* Slice of slices - wart, starting index as expected, end value is index + 1, so
second argument - first argument is the length. Change value in the underlying
data structure when mutiple slice headers are viewing it - all the views see the
change. Unless... you add a third argument to set the capacity of the slice of
slice, e.g. slice1[2:4:4]

Maps

* Map iteration order is random - not just undefined by purposely random. Use
a slice if you need ordering.
* Key - if you can't use it in a conditional expression (equals operator) you can't
use it as a map key. You can use a struct as a key, as long as each field could be
used as a key.
* Make a map with... make. Need to make one or initialize one via a composite literal
before you work with it (key : value). Again, think about zero value/non-zero value when
deciding how to do the initialization.
* Set - via map of string to empty struct map[string]struct{}. Empty struct is a zero
  allocation value. x := string{{}}


Methods and Embedding

* Methods are functions with receivers
* Can have value receivers or pointer receivers - work with a copy of the
value, or work with the shared value (and can thus alter the shared value)
For pointer receivers we don't have to use the * operator.
* Go adjusts calls as needed when caller uses pointer or a value based on the
receiver type.
* Can only declare methods for user types - structs, named types (not built in types).
* Function type - reference type (header value - two parts - first part is pointer
  to the code, second part is any data needed to call it)
* Methods are a compiler construct, we really just have functions. Receiver is
really just the first parameter of the function. See http://play.golang.org/p/MNI1jR8Ets
* Methods are just functions with the first parameter being the receiver.
* Pointer types can be embedded as well (look at File in this example -
  https://play.golang.org/p/xD6PCx--GG)

Interfaces

* Interfaces declare behavior
* Ref type - two word header value - first word is about the type, second
is the value of a concrete type embedded in the interface header.
* Method sets - spec treats them from the value point of view. When I have
a value of type T, only value receivers implement the interface. When I have a
pointer, values and pointers of type T implement the interface. This is
because you can't always get the address of a value - for example see
https://play.golang.org/p/myZaTlp4Qd - expressions are not addressable.
* From the receiver point of view, (*T) *T, (T) &T,T
* Think of it from the receiver's point of view
* Interface compliance is at the value level, not the type level.
* First part of interface header - points to the iTable, which has the type
as well as the method set.

Embedding

* Allows you to take an existing type and extend it
* Can take an existing type and override its behavior
* Embed the type by including an inner type in the struct declaration - https://play.golang.org/p/UkrDXkk-Ch
* The inner type exists in and of itself.
* Type promotion - all the methods and identifiers of the inner type are
promoted up to the outer type.
* Override example - if the outer type implements the interface the the inner type
interface is not promoted, use case is you call the inner impl before adding in the
outer's value.
