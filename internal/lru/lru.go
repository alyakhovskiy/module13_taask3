package lru

import (
	"container/list"
)

type LRUCache interface {
	// Add Добавляет новое значение с ключом в кеш (с наивысшим приоритетом), возвращает true, если все прошло успешно
	// В случае дублирования ключа вернуть false
	// В случае превышения размера - вытесняется наименее приоритетный элемент
	Add(key, value string) bool
	// Get Возвращает значение под ключом и флаг его наличия в кеше
	// В случае наличия в кеше элемента повышает его приоритет
	Get(key string) (value string, ok bool)
	// Remove Удаляет элемент из кеша, в случае успеха возвращает true, в случае отсутствия элемента - false
	Remove(key string) (ok bool)
}

type lruCache struct {
	capacity int
	cache    map[string]*list.Element
	queue    *list.List
}

func NewLRUCache(n int) LRUCache {
	return &lruCache{
		capacity: n,
		cache:    make(map[string]*list.Element),
		queue:    list.New(),
	}
}

type entry struct {
	key   string
	value string
}

func (c *lruCache) Add(key, value string) bool {
	// Проверяем наличие ключа в кеше
	if _, ok := c.cache[key]; ok {
		return false // Ключ уже существует, возвращаем false
	}

	// Проверяем, достигнут ли лимит capacity
	if len(c.cache) >= c.capacity {
		// Вытесняем наименее приоритетный элемент из кеша
		oldest := c.queue.Back()
		delete(c.cache, oldest.Value.(*entry).key)
		c.queue.Remove(oldest)
	}

	// Добавляем новый элемент в кеш (с наивысшим приоритетом)
	newEntry := &entry{key, value}
	element := c.queue.PushFront(newEntry)
	c.cache[key] = element

	return true
}

func (c *lruCache) Get(key string) (value string, ok bool) {
	// Проверяем наличие ключа в кеше
	if element, exists := c.cache[key]; exists {
		// Повышаем приоритет элемента, перемещая его в начало очереди
		c.queue.MoveToFront(element)
		return element.Value.(*entry).value, true
	}

	return "", false
}

func (c *lruCache) Remove(key string) (ok bool) {
	// Проверяем наличие ключа в кеше
	if element, exists := c.cache[key]; exists {
		// Удаляем элемент из кеша и из очереди
		delete(c.cache, key)
		c.queue.Remove(element)
		return true
	}

	return false
}
