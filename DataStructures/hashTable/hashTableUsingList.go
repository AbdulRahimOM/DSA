package hashTable

type HashTableList struct {
	size    int
	buckets []*pairList
}

type pairList struct {
	key   string
	value interface{}
	next  *pairList
}

func NewListHashTable(size int) *HashTableList {
	return &HashTableList{
		size:    size,
		buckets: make([]*pairList, size),
	}
}
func (ht *HashTableList) hash(key string) int {
	hash := 0
	for _, char := range key {
		hash += int(char)
	}
	return hash % ht.size
}
func (ht *HashTableList) Insert(key string, value interface{}) {
	index := ht.hash(key)
	pair := pairList{
		key:   key,
		value: value,
	}
	current := ht.buckets[index]
	if current == nil {
		ht.buckets[index] = &pair
		return
	}
	for current.next != nil {
		current = current.next
	}
	current.next = &pair
}

func (ht *HashTableList) Delete(key string) {
	index := ht.hash(key)
	current := ht.buckets[index]
	if current == nil {
		return
	}
	if current.key == key {
		ht.buckets[index] = nil
	}
	for current.next != nil {
		if current.next.key == key {
			current.next = current.next.next
			return
		}
		current = current.next
	}
}

func (ht *HashTableList) Get(key string) interface{} {
	index := ht.hash(key)
	current := ht.buckets[index]
	for current != nil {
		if current.key == key {
			return current.value
		}
		current = current.next
	}
	return nil
}
