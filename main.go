package main

import "errors"

func main() {

}

type Cache interface {
	Set(key, value string) error
	Get(key string) (string, error)
	Delete(key string) error
}

type SimpleCache struct {
	storage map[string]string
}

func NewCache() Cache {
	return &SimpleCache{
		storage: make(map[string]string),
	}
}

func (c *SimpleCache) Set(key, value string) error {
	c.storage[key] = value
	return nil // TODO
}

func (c *SimpleCache) Get(key string) (string, error) {
	value, ok := c.storage[key]
	if !ok {
		return value, ErrNotFound
	}
	return value, nil
}

func (c *SimpleCache) Delete(key string) error {
	delete(c.storage, key)
    return nil // TODO
}

var ErrNotFound = errors.New("value not found")
