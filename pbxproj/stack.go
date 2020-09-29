package pbxproj

import (
	"fmt"
	"sync"
)

type stack struct {
	lock sync.Mutex
	s    []interface{}
}

func newStack() *stack {
	return &stack{sync.Mutex{}, make([]interface{}, 0)}
}

func (s *stack) push(v interface{}) {
	s.lock.Lock()
	defer s.lock.Unlock()

	s.s = append(s.s, v)
}

func (s *stack) pop() interface{} {
	s.lock.Lock()
	s.lock.Unlock()

	len := len(s.s)
	if len == 0 {
		return nil
	}
	v := s.s[len-1]
	s.s = s.s[:len-1]
	return v
}

func (s *stack) top() interface{} {
	s.lock.Lock()
	s.lock.Unlock()

	len := len(s.s)
	if len == 0 {
		return nil
	}
	return s.s[len-1]
}

func (s *stack) len() int {
	return len(s.s)
}

func (s *stack) dump() {
	fmt.Println("Stack: ", len(s.s))
	for key, v := range s.s {
		fmt.Println(key, v)
	}
}
