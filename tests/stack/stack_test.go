package stack_test

import (
	"testing"

	"github.com/dkvka/gopkg/dsa/stack"
)

func TestStack(t *testing.T) {
	s := stack.New[uint8]()

	t.Run("Push, pop, clear", func(t *testing.T) {
		const PUSH_COUNT = 10
		for i := uint8(0); i < PUSH_COUNT; i++ {
			s.Push(i)
		}
		if s.Size() != PUSH_COUNT {
			t.Errorf("Expected stack size %d, got %d", PUSH_COUNT, s.Size())
		}

		item, ok := s.Pop()
		if !ok || item != 9 {
			t.Errorf("Expected to pop 9, got %d, ok: %v", item, ok)
		}

		s.Clear()
		if len(s.items) != 0 {
			t.Errorf("Expected empty stack after clear, got size %d", s.Size())
		}
	})
}
