package inmemorycache

type Cache struct {
	items map[string]interface{}
}

func New() *Cache {
	return &Cache{
		items: make(map[string]interface{}),
	}
}

func (c *Cache) Set(key string, value interface{}) {
	c.items[key] = value
}

func (c *Cache) Get(key string) interface{} {
	value, found := c.items[key]
	if !found {
		return nil
	}
	return value
}

func (c *Cache) Delete(key string) {
	delete(c.items, key)
}
