package hashTable

type pair struct {
	key   string
	value interface{}
}

type HashTable struct {
	size    int
	buckets [][]pair
}

func NewSliceHashTable(size int) *HashTable {
	return &HashTable{
		size:    size,
		buckets: make([][]pair, size),
	}
}

func (ht *HashTable) hash(key string) int {
	var hash int
	for _, char := range key {
		hash += int(char)
	}
	return hash % ht.size
}

func (ht *HashTable) Insert(key string, value interface{}) {
	index := ht.hash(key)
	pair := pair{key, value}
	ht.buckets[index] = append(ht.buckets[index], pair)
}

func (ht *HashTable) Get(key string) interface{} {
	index := ht.hash(key)
	for _, pair := range ht.buckets[index] {
		if pair.key == key {
			return pair.value
		}
	}
	return nil
}
func (ht *HashTable) Delete(key string) {
	index := ht.hash(key)
	for i, pair := range ht.buckets[index] {
		if pair.key == key {
			ht.buckets[index] = append(ht.buckets[index][:i], ht.buckets[index][i+1:]...)
		}
	}
}
