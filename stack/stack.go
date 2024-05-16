package stack

type Stack struct {
	head *node
}

type node struct {
	v    int
	next *node
}

func NewStack() Stack {
	return Stack{}
}

func (s *Stack) Push(n int) {
	if s == nil {
		return
	}
	s.head = &node{
		v:    n,
		next: s.head,
	}
}

func (s *Stack) Pop() int {
	if s == nil || s.head == nil {
		return -1
	}
	value := s.head.v
	s.head = s.head.next
	return value
}

func (s *Stack) Values() []int {
	var out []int
	if s == nil {
		return out
	}
	i := s.head
	for i != nil {
		out = append(out, i.v)
		i = i.next
	}
	return out
}
