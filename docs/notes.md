# Notes

## Installation
* Is possible to have multiple Go versions installed on the same machine?
 Yes, https://go.dev/doc/manage-install

## Naming
The convention in Go is to use MixedCaps or mixedCaps rather than underscores to write multiword names.

### Packages
Lower case, single-word names, no underscores, no mixedCaps. Short, concise and evocative.

### Getters & Setters
It's neither idiomatic nor necessary to put Get into the getter's name. 
If you have a field called owner (lower case, unexported), the getter method should be 
called Owner (upper case, exported), not GetOwner.
The use of upper-case names for export provides the hook to discriminate the field from the method.
A setter function, if needed, will likely be called SetOwner

## Loops
There is no `do` or `while` loop, only a slightly generalized `for` and `switch`.

### Switch
There is no automatic fall through, but cases can be presented in comma-parated lists

## Initialization
`new` is a built-in function. But `new(T)` and `&T{}` are equivalent but `new` exist before `make` 
and `&{}`

The `make(T, args)` function creates slices, maps and channels only, and it returns an 
initialized (not zeroed) value of type T (not `*T`).
The reason for the distinction with `new(T)` is that these three types represent, under the covers, 
references to data structures
that must be initialized before use.

## Arrays
* Arrays are values. Assigning one array to another copies all the elements.
* In particular, if you pass an array to a function, it will receive a copy of the array, not a 
pointer to it.
* The size of an array is part of its type. The types `[10]int` and `[20]int` are distinct.

## Slices
Slices hold references to an underlying array, if you assign one slice to another, both
refer to the same array. If a function takes a slice argument, changes it makes to
the elements of the slice will be visible to the caller, analogous to passing a
pointer to the underlying array.

## Maps
The key can be of any type for which the equality operator is defined, such as integers, floating 
point and complex numbers, strings, pointers, interfaces (as long as the dynamic type supports 
equality), structs and arrays. Slices cannot be used as map keys, because equality is not defined 
on them.

If you pass a map to a function that changes the contents of the map, the changes will be visible
in the caller.

An attempt to fetch a map value with a key that is not present in the map will return the zero 
value for the type of the entries in the map.
