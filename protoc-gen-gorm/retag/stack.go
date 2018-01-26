package retag

import "container/list"

type stack struct {
	s *list.List
}

func NewStack() *stack {
	return &stack{
		s: list.New(),
	}
}
func (s *stack) PUSH(value string) {
	s.s.PushBack(value)
}

func (s *stack) POP() string {
	res := s.s.Back()
	if res == nil {
		return ""
	}
	s.s.Remove(res)
	return res.Value.(string)
}

func (s *stack) GetPOP() string {
	res := s.s.Back()
	if res == nil {
		return ""
	}

	return res.Value.(string)
}
