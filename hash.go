package hash

import(
  "hash/fnv"
)
// Using only primitive types, implement a fixed-size hash map that associates
// string keys with arbitrary data object references (you don't need to copy the object).
// Your data structure should be optimized for algorithmic runtime and memory usage.
// You should not import any external libraries, and may not use primitive hash map
// or dictionary types in languages like Python or Ruby.

// The solution should be delivered in one class (or your language's equivalent)
// that provides the following functions:

// constructor(size): return an instance of the class with pre-allocated space for the given number of objects.

// boolean set(key, value): stores the given key/value pair in the hash map. Returns a boolean value indicating success / failure of the operation.

// get(key): return the value associated with the given key, or null if no value is set.

// delete(key): delete the value associated with the given key, returning the value on success or null if the key has no value.

// float load(): return a float value representing the load factor (`(items in hash map)/(size of hash map)`) of the data structure. Since the size of the dat structure is fixed, this should never be greater than 1.

// If your language provides a built-in hashing function for strings
// (ex. `hashCode` in Java or `__hash__` in Python) you are welcome to use that.
// If not, you are welcome to do something naive, or use something you find online
// with proper attribution.

type Node struct {
  key string
  value interface{}
}

type Hash struct {
  values [][]Node
  count int
  size int
}

func New(size int) *Hash {
  return &Hash {
    make([][]Node, size),
    0,
    size
  }
}

func (h *Hash) Set(key string, value interface{}) bool {
  index := index(key, h.size)
  for _, node := range h.values[index] {
    if node.key == key {
      node.value = value
      return true
    }
  }
  if(h.count == h.size) {
    return false
  }
  h.values[index] = append(h.values[index], value)
  h.count++
  return true
}

func (h *Hash) Get(key string) interface{} {
  index := index(key, h.size)
  for _, node := range h.values[index] {
    if node.key == key {
      return node.value
    }
  }
  return nil
}

func (h *Hash) Delete(key string) interface{} {
  index := index(key, h.size)
  for _, node := range h.values[index] {
    if node.key == key {
      h.values[index] = append(h.values[index][:i], h.values[index][i+1:]...)
      count--
      return node.value
    }
  }
  return nil
}

func (h *Hash) Load() float {
  return float(h.count) / float(h.size)
}

func index(key string, size int) {
  return int(hash(key) % size)
}
func hash(key string) uint32 {
  h := fnv.New64a()
  h.Write([]byte(key))
  return h.Sum64()
}