package machine

type CodedStopper struct {
	positions []int
	index     int
}

func (s *CodedStopper) Stop() Position {
	pos := s.positions[s.index%len(s.positions)]
	s.index++
	return Position(pos)
}

func NewCodedStopper(positions []int) *CodedStopper {
	return &CodedStopper{positions: positions}
}
