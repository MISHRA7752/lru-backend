package lru

import (
	"container/list"
	"errors"
	"fmt"
	"sync"
	"time"
)

type CacheItem struct {
	Key        string `json:"key"`
	Value      string `json:"value"`
	Expiration int64  `json:"expiration"`
}

type LRUCache struct {
	capacity int
	items    map[string]*list.Element
	ll       *list.List
	mu       sync.Mutex
}

func NewLRUCache(capacity int) *LRUCache {
	return &LRUCache{
		capacity: capacity,
		items:    make(map[string]*list.Element, 0),
		ll:       list.New(),
	}
}

func (_c *LRUCache) GetOne(key string) (string, bool) {
	_c.mu.Lock()
	defer _c.mu.Unlock()

	if ele, found := _c.items[key]; found {
		if time.Now().Unix() > ele.Value.(*CacheItem).Expiration {
			_c.removeElement(ele)
			return "", false
		}
		_c.ll.MoveToFront(ele)
		return ele.Value.(*CacheItem).Value, true
	}
	return "", false
}

func (_c *LRUCache) GetAll() ([]CacheItem, error) {
	finalAns := make([]CacheItem, 0)
	_c.mu.Lock()
	defer _c.mu.Unlock()

	for _, v := range _c.items {
		fmt.Println("Hello GetALL", v.Value.(*CacheItem).Key)
		currTime := time.Now().Unix()
		eleExpTime := v.Value.(*CacheItem).Expiration
		if currTime-eleExpTime < 0 {
			obj := v.Value.(*CacheItem)
			finalAns = append(finalAns, CacheItem{
				Key:        obj.Key,
				Value:      obj.Value,
				Expiration: obj.Expiration,
			})
		}
	}
	fmt.Println("finl", finalAns)
	if finalAns == nil {
		return finalAns, errors.New("empty list")
	}
	return finalAns, nil
}

func (_c *LRUCache) Set(key, value string, expiration int64) {
	_c.mu.Lock()
	defer _c.mu.Unlock()
	if ele, found := _c.items[key]; found {
		_c.ll.MoveToBack(ele)
		ele.Value.(*CacheItem).Value = value
		ele.Value.(*CacheItem).Expiration = time.Now().Unix() + expiration
		return
	}
	item := &CacheItem{key, value, time.Now().Unix() + expiration}
	ele := _c.ll.PushFront(item)
	_c.items[key] = ele
	if _c.ll.Len() > _c.capacity {
		_c.removeOldest()
	}

}

func (_c *LRUCache) Delete(key string) {
	_c.mu.Lock()
	defer _c.mu.Unlock()
	if ele, found := _c.items[key]; found {
		_c.removeElement(ele)
	}
}

func (_c *LRUCache) removeElement(ele *list.Element) {
	_c.ll.Remove(ele)
	delete(_c.items, ele.Value.(*CacheItem).Key)
}

func (_c *LRUCache) removeOldest() {
	ele := _c.ll.Front()
	if ele != nil {
		_c.removeElement(ele)
	}
}
