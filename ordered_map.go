package orderedmap

import (
	"container/list"
	"errors"
	"fmt"
)

// ErrNotFound error for element not found
var ErrNotFound = errors.New("collection: element not found")

type T struct {
	keyStorage map[string]*list.Element
	elements   *list.List
}

// NewOrderedMap returnr a new instance of orderedMap
func New() T {
	return T{
		keyStorage: make(map[string]*list.Element),
		elements:   list.New(),
	}
}

type elemKeyValuePair struct {
	Key   string
	Value interface{}
}

// Set sets new elements to list
// time O(1)
// space O(n)
func (m *T) Set(key string, value interface{}) {
	// adds new item
	keyValue := elemKeyValuePair{
		Key:   key,
		Value: value,
	}

	element, ok := m.keyStorage[key]
	if ok {
		element.Value = keyValue
		return
	}

	element = m.elements.PushBack(keyValue)
	m.keyStorage[key] = element
}

// Get gets value by key
// time O(1)
// space O(1)
func (m *T) Get(key string) (interface{}, error) {
	element, ok := m.keyStorage[key]
	if !ok {
		return nil, ErrNotFound
	}

	return element.Value.(elemKeyValuePair).Value, nil
}

// PrevKey returns previous key
// time O(1)
// space O(1)
func (m *T) PrevKey(key string) (string, error) {
	element, ok := m.keyStorage[key]
	if !ok {
		return "", ErrNotFound
	}

	prev := element.Prev()
	if prev == nil {
		return "", nil
	}

	return prev.Value.(elemKeyValuePair).Key, nil
}

// NextKey returns next key
// time O(1)
// space O(1)
func (m *T) NextKey(key string) (string, error) {
	element, ok := m.keyStorage[key]
	if !ok {
		return "", ErrNotFound
	}

	next := element.Next()
	if next == nil {
		return "", nil
	}

	return next.Value.(elemKeyValuePair).Key, nil
}

// LastKey returns last key form list
// time O(1)
// space O(1)
func (m *T) LastKey() string {
	element := m.elements.Back()

	if element == nil {
		return ""
	}

	lastElem := element.Value.(elemKeyValuePair)

	return lastElem.Key
}

// GetFirstKey returns first key form list
// time O(1)
// space O(1)
func (m *T) GetFirstKey() string {
	element := m.elements.Front()

	if element == nil {
		return ""
	}

	firstElem := element.Value.(elemKeyValuePair)

	return firstElem.Key
}

// time O(1)
// space O(1)
func (m *T) Remove(key string) {
	element, ok := m.keyStorage[key]
	if !ok {
		return
	}

	m.elements.Remove(element)

	delete(m.keyStorage, key)

	if len(m.keyStorage) != m.elements.Len() {
		panic(fmt.Sprintf(`
		bug: number of elements in the map is not the same as the number of elements in the list.
		map: %+v
		list: %+v
		`,
			m.keyStorage,
			m.elements,
		))
	}
}

// time O(1)
// space O(1)
func (m *T) Len() int {
	return len(m.keyStorage)
}

// time O(1)
// space O(1)
func (m *T) Has(key string) bool {
	_, ok := m.keyStorage[key]
	return ok
}
