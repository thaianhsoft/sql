package utils

import (
	"log"
	"sync"
	"testing"
)
type Case struct {
	mu *sync.Mutex
	internal map[string]int
}
func (c *Case) getNumberCase(name string) int {
	v := c.internal[name]
	c.mu.Lock()
	c.internal[name]++
	c.mu.Unlock()
	return v
}

func (c *Case) createNotExist(name string) {
	c.mu.Lock()
	if _, exist := c.internal[name]; !exist {
		c.internal[name] = 0
		c.mu.Unlock()
		return
	}
	c.mu.Unlock()
	return
}

var test_assert *Case = &Case{
	mu: &sync.Mutex{},
	internal: make(map[string]int),
}

func Assert(t *testing.T, have interface{}, want interface{}) {
	name := t.Name()
	test_assert.createNotExist(name)
	c := test_assert.getNumberCase(name)
	if have != want {
		log.Printf("[%v] [CASE(%v)]: FAILED -> have(%v) != want(%v)\n", name, c, have, want)
		t.FailNow()
		return
	}
	log.Printf("[%v] [CASE(%v)]: PASSED -> have(%v) == want(%v)\n", name, c, have, want)
}
