package cache

import (
	"fmt"
	"sync"
	"time"
)

//нужно получать url.
//удалять после взятия url из кэша
//Чистить весь кэш при завершении работы

type val struct {
	url     string
	counter int
}

type Cache struct {
	m  map[int]val
	mu *sync.Mutex
}

// Чистим кэш
func (c Cache) DeleteCache() {
	c.mu.Lock()
	defer c.mu.Unlock()
	for k := range c.m {
		delete(c.m, k)
	}
}

func (c Cache) Get(key int) (string, error) {
	c.mu.Lock()
	defer c.mu.Unlock()
	v, ok := c.m[key]
	if !ok {
		return "", fmt.Errorf("нет такого ключа")
	}
	v.counter -= 1
	return v.url, nil

}

func (c Cache) Set(key int, url string) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.m[key] = val{url, 1}
}

func (c Cache) New() Cache {
	m := make(map[int]val)
	go func() {
		for {
			c.mu.Lock()
			time.Sleep(1 * time.Second)
			for k, v := range c.m {
				if v.counter == 0 {
					delete(c.m, k)
				}
			}
			c.mu.Unlock()
		}
	}()
	return Cache{m: m}
}
