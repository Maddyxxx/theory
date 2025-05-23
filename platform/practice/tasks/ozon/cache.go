//  LRU-кеш

package ozon

import (
	"fmt"
	"sync"
)

type Cache interface {
	Set(k, v string)
	Get(k string) (v string, ok bool)
}

type Node struct {
	key  string // Добавлено для удаления из map
	data string
	next *Node
	prev *Node
}

type LRUCache struct {
	mu    sync.RWMutex
	data  map[string]*Node // Храним указатели
	limit int
	head  *Node
	tail  *Node
}

func NewLRUCache(limit int) *LRUCache {
	return &LRUCache{
		data:  make(map[string]*Node, limit),
		limit: limit,
	}
}

func (c *LRUCache) Set(k, v string) {
	c.mu.Lock()
	defer c.mu.Unlock()

	// Если ключ уже есть — обновляем значение и перемещаем в голову
	if node, exists := c.data[k]; exists {
		node.data = v
		c.moveToHead(node)
		return
	}

	// Если кеш заполнен — удаляем последний элемент
	if len(c.data) >= c.limit {
		c.removeTail()
	}

	// Создаем новую ноду
	newNode := &Node{
		key:  k,
		data: v,
		next: c.head,
	}

	// Добавляем в голову списка
	if c.head != nil {
		c.head.prev = newNode
	}
	c.head = newNode

	// Если хвоста нет — это первый элемент
	if c.tail == nil {
		c.tail = newNode
	}

	// Сохраняем в map
	c.data[k] = newNode
}

func (c *LRUCache) Get(k string) (string, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()

	node, ok := c.data[k]
	if !ok {
		return "", false
	}

	// Перемещаем в голову
	c.moveToHead(node)
	return node.data, true
}

func (c *LRUCache) moveToHead(node *Node) {
	// Если уже голова — ничего не делаем
	if node == c.head {
		return
	}

	// Удаляем из текущей позиции
	if node.prev != nil {
		node.prev.next = node.next
	}
	if node.next != nil {
		node.next.prev = node.prev
	}

	// Если это был хвост — обновляем tail
	if node == c.tail {
		c.tail = node.prev
	}

	// Добавляем в голову
	node.next = c.head
	node.prev = nil
	if c.head != nil {
		c.head.prev = node
	}
	c.head = node
}

func (c *LRUCache) removeTail() {
	if c.tail == nil {
		return
	}

	// Удаляем из map
	delete(c.data, c.tail.key)

	// Обновляем хвост
	if c.tail.prev != nil {
		c.tail.prev.next = nil
		c.tail = c.tail.prev
	} else {
		// Это был последний элемент
		c.head = nil
		c.tail = nil
	}
}

func TestCache() {
	cache := NewLRUCache(2)
	cache.Set("a", "1")
	cache.Set("b", "2")

	val, ok := cache.Get("a") // "1", true
	cache.Set("c", "3")       // Вытеснит "b", так как "a" использовали недавно

	_, ok = cache.Get("b") // false, вытеснен
	fmt.Println(cache, val, ok)
}
