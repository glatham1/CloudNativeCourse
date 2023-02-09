package lru

import "errors"

type Cacher interface {
	Get(interface{}) (interface{}, error)
	Put(interface{}, interface{}) error
}

type lruCache struct {
	size      int               //size of cache
	remaining int               //remaining capacity
	cache     map[string]string //to store cache data
	queue     []string          //to store order of cache access
}

func NewCache(size int) Cacher {
	return &lruCache{size: size, remaining: size, cache: make(map[string]string), queue: make([]string, 0)}
}

func (lru *lruCache) Get(key interface{}) (interface{}, error) {
	//Get function is responsible for returning the value of a given key from the cache.

	//value being obtained from cache with key
	var val interface{} = lru.cache[key.(string)]
	//by default setting an error value to nil
	var err error = nil

	//if statement that determins if the value was found or not
	if val.(string) == "" {
		return val, errors.New("could not find value")
	} else {
		//if the value was found then the key is appended the tail of the queue and no error is returned
		lru.queue = append(lru.queue, key.(string))

		return val, err
	}
}

func (lru *lruCache) Put(key, val interface{}) error {
	//Put function is responsible for adding a value and corresponding key into the cache and queue

	//if statement determines if the cache is full or not
	if lru.remaining == 0 {
		//if the cache is full the value at the head of the queue is deleted from queue and cache
		delete(lru.cache, lru.queue[0])
		lru.qDel(lru.queue[0])

		//the new value and corresponding key are then added to the cache and tail of the queue
		lru.cache[key.(string)] = val.(string)
		lru.queue = append(lru.queue, key.(string))
	} else {
		//if the cache is not full then the value and key are simply added to the cache and tail of the queue
		lru.cache[key.(string)] = val.(string)
		lru.queue = append(lru.queue, key.(string))
		lru.remaining--
	}

	//returns a nil error if everything runs properly
	var err error = nil
	return err
}

// Delete element from queue
func (lru *lruCache) qDel(ele string) {
	for i := 0; i < len(lru.queue); i++ {
		if lru.queue[i] == ele {
			oldlen := len(lru.queue)
			copy(lru.queue[i:], lru.queue[i+1:])
			lru.queue = lru.queue[:oldlen-1]
			break
		}
	}
}
