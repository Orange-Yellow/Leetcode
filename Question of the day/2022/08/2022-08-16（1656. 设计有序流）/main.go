package main

/*
	name: 1656. 设计有序流
	link: https://leetcode.cn/problems/design-an-ordered-stream/
*/

func main() {

}

type OrderedStream struct {
	stream []string
	ptr    int
}

func Constructor(n int) OrderedStream {
	return OrderedStream{make([]string, n+1), 1}
}

func (s *OrderedStream) Insert(idKey int, value string) []string {
	s.stream[idKey] = value
	start := s.ptr
	for s.ptr < len(s.stream) && s.stream[s.ptr] != "" {
		s.ptr++
	}
	return s.stream[start:s.ptr]
}
