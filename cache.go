package cache

import "time"

type Data struct {
	Value    string
	Deadline time.Time
}

type Cache struct {
	Map map[string]Data
}

func NewCache() *Cache {
	return &Cache{
		make(map[string]Data),
	}
}

func (c *Cache) Get(key string) (string, bool) {
	for key, val := range c.Map {
		if time.Now().After(val.Deadline) {
			delete(c.Map, key)
		}
	}
	if _, ok := c.Map[key]; !ok {
		return "", false
	}
	return c.Map[key].Value, true
}

func (c *Cache) Put(key, value string) {
	c.Map[key] = Data{Value: value, Deadline: time.Unix(1<<62-1, 0)}
}

func (c *Cache) Keys() []string {
	validKeys := []string{}
	for key := range c.Map {
		validKeys = append(validKeys, key)
	}
	return validKeys
}

func (c *Cache) PutTill(key, value string, deadline time.Time) {
	c.Map[key] = Data{Value: value, Deadline: deadline}
}
