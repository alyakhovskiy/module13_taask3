package lru

import (
	"testing"
)

func TestLRUCache_Add(t *testing.T) {
	cache := NewLRUCache(3)

	// Добавляем элементы
	cache.Add("key1", "value1")
	cache.Add("key2", "value2")
	cache.Add("key3", "value3")

	// Проверяем успешное добавление
	if added := cache.Add("key4", "value4"); !added {
		t.Errorf("Expected Add to return true, got false")
	}

	// Проверяем, что элемент с наименьшим приоритетом был вытеснен
	if _, ok := cache.Get("key1"); ok {
		t.Errorf("Expected key1 to be evicted, but it is still present in the cache")
	}

	// Проверяем, что элемент с наивысшим приоритетом присутствует
	if value, ok := cache.Get("key4"); !ok || value != "value4" {
		t.Errorf("Expected key4 to be present in the cache with updated value, but it is not")
	}

	// Проверяем, что при добавлении существующего ключа возвращается false
	if added := cache.Add("key2", "new_value2"); added {
		t.Errorf("Expected Add to return false for existing key, got true")
	}

	// Проверяем, что при добавлении элемента вместо существующего ключа происходит замена
	if value, ok := cache.Get("key2"); !ok || value != "value2" {
		t.Errorf("Expected value for key2 to be unchanged, got %s", value)
	}
}

func TestLRUCache_Get(t *testing.T) {
	cache := NewLRUCache(3)

	// Добавляем элемент
	cache.Add("key1", "value1")

	// Проверяем успешное получение
	if value, ok := cache.Get("key1"); !ok || value != "value1" {
		t.Errorf("Expected Get to return value1, got %s", value)
	}

	// Проверяем отсутствие элемента
	if value, ok := cache.Get("key2"); ok || value != "" {
		t.Errorf("Expected Get to return (nil, false), got (%s, %t)", value, ok)
	}
}

func TestLRUCache_Remove(t *testing.T) {
	cache := NewLRUCache(3)

	// Добавляем элемент
	cache.Add("key1", "value1")

	// Проверяем успешное удаление
	if removed := cache.Remove("key1"); !removed {
		t.Errorf("Expected Remove to return true, got false")
	}

	// Проверяем, что элемент отсутствует после удаления
	if _, ok := cache.Get("key1"); ok {
		t.Errorf("Expected key1 to be absent after removal, but it is still present in the cache")
	}

	// Проверяем удаление отсутствующего элемента
	if removed := cache.Remove("key2"); removed {
		t.Errorf("Expected Remove to return false for non-existent key, got true")
	}
}
