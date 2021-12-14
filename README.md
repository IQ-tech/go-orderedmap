# Non concurrency safe ordered map

This project provides a hash map that maintains the order of insertions.

The time complexity of each operation is O(1) and the space complexity
of the map is O(2 \* n) where n is the number of elements in the map.

# Installation

To install, use `go get`

```terminal
go get github.com/IQ-tech/go-orderedmap
```

## Usage

```go
m := orderedmap.New()

m.Set("one", 1)
m.Set("two", 2)
m.Set("three", 3)

// Returns the value associated with the key
m.Get("two") // => 2

// Returns the key that was added before the target key
m.PrevKey("key_not_in_the_map") // => "", orderedmap.ErrNotFound
m.PrevKey("one") // => "", nil
m.PrevKey("three") // => "two", nil

// Returns the key that was added after the target key
m.NextKey("key_not_in_the_map") // => "", orderedmap.ErrNotFound
m.NextKey("two") // => "three", nil

// Returns the key that was added first
m.GetFirstKey() // => "one"

// Returns the key that was added last
m.LastKey() // => "three"

// Removes key from the map
m.Remove("one")
m.PrevKey("one") // => "", orderedmap.ErrNotFound

// Returns the number of elements in the map
m.Len() // => 2 (keys "two" and "three)

// Checks if key is in the map
m.Has("one") // => false
m.Has("two") // => true
```
