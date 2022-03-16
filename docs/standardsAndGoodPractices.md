# Standards and good practices

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

### Loops
There is no `do` or `while` loop, only a slightly generalized `for` and `switch`.

### Switch
There is no automatic fall through, but cases can be presented in comma-parated lists
