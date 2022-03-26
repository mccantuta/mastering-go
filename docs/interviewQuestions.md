# Interview questions

## Data structrures
* What happens when we create a map.
`m := make(map[string]string)`
Internally a structure called map header is created and the variable m receives a pointer to this 
structure, the map header contains all the meta information about the map, like:
- The number of entries that are currently in the map
- The number of buckets in a map is always equal to power of two hence the log(buckets) stored to keep the value small
- Pointer to the bucket array that is stored in contiguous memory location,
- Hash seed which is random to create each map differently

* What happens when we insert a new value in the map

Hash function is called and a hash code is generated for the given key, based on a part of the 
hash code a bucket is determined to store the key value pair. Once the bucket is selected the entry 
needs to be stored in that bucket. The complete hash code of the incoming key is compared with all 
the hashes from the initial array of hash codes i.e h1, h2, h3…. if no hash code matches that means 
this is a new entry. Now if the bucket contains an empty slot then the new entry is stored at the 
end of the list of key value pairs, else a new bucket is created and the entry is stored in the new 
bucket and the overflow pointer of old bucket points to this new bucket

* Map Iteration
Map iterator in golang is random. If you print the same map multiple times you can see that each 
output will be different.

* What happens when map grows
Every time the number of elements in a bucket reaches a certain limit, i.e the load factor which is 
6.5, the map will grow in size by doubling the number of buckets. This is done by creating a new 
bucket array consisting of twice the number of buckets than the old array, and then copying all the 
buckets from old to new array but this is done very efficiently over time and not at once. During 
this the map maintains a pointer to the old bucket which is stored at the end of the new bucket 
array.

* What happens when we delete an entry from the map
Maps in golang only grow in size. Even if we delete all the entries from the map, the number of 
buckets will remain the same
