package cache

import (
	"fmt"
	"time"
)

type Cache struct {
	items map[string]item
}

type item struct {
	Object     interface{}
	Expiration int64
}

func NewCache() *Cache {
	items := make(map[string]item)
	return &Cache{items: items}
}

func (c *Cache) Get(k string) (interface{}, bool) {
	item, found := c.items[k]
	if !found {
		return nil, false
	}

	if item.Expiration > 0 {
		if time.Now().UnixNano() > item.Expiration {
			fmt.Println(item.Expiration)
			return nil, false
		}
	}

	return item.Object, true
}

func (c *Cache) Set(key string, value interface{}, duration time.Duration) {
	var expiration int64
	if duration > 0 {
		expiration = time.Now().Add(duration).UnixNano()
	}
	c.items[key] = item{
		Object:     value,
		Expiration: expiration,
	}
}
