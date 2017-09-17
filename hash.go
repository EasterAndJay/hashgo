package hash

import (
	"errors"
	"hash/fnv"
)

// constructor(size): return an instance of the class with pre-allocated space for the given number of objects.

// boolean set(key, value): stores the given key/value pair in the hash map. Returns a boolean value indicating success / failure of the operation.

// get(key): return the value associated with the given key, or null if no value is set.

// delete(key): delete the value associated with the given key, returning the value on success or null if the key has no value.

// float load(): return a float value representing the load factor (`(items in hash map)/(size of hash map)`) of the data structure. Since the size of the dat structure is fixed, this should never be greater than 1.

type Node struct {
	key   string
	value interface{}
}

type Hash struct {
	values [][]Node
	count  uint32
	size   uint32
}

func New(size uint32) (*Hash, error) {
	if size <= 0 {
		return nil, errors.New("Size of hash must be a natural number")
	}
	return &Hash{
		make([][]Node, size),
		0,
		size,
	}, nil
}

func (h *Hash) Set(key string, value interface{}) bool {
	index := index(key, h.size)
	for _, node := range h.values[index] {
		if node.key == key {
			node.value = value
			return true
		}
	}
	if h.count == h.size {
		return false
	}
	h.values[index] = append(h.values[index], Node{key,value})
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
	for i, node := range h.values[index] {
		if node.key == key {
			h.values[index] = append(h.values[index][:i], h.values[index][i+1:]...)
			h.count--
			return node.value
		}
	}
	return nil
}

func (h *Hash) Load() float32 {
	return float32(h.count) / float32(h.size)
}

func (h *Hash) Count() uint32 {
	return h.count
}

func (h *Hash) Size() uint32 {
	return h.size
}

func index(key string, size uint32) uint32 {
	return uint32(hash(key) % size)
}
func hash(key string) uint32 {
	h := fnv.New32a()
	h.Write([]byte(key))
	return h.Sum32()
}
