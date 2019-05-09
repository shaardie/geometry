package geometry

type Segment struct {
	P1, P2 Point
}

func (s *Segment) Sort() {
	if s.P1.X < s.P2.X {
		return
	} else if s.P1.X > s.P1.X {
		s.P1, s.P2 = s.P2, s.P1
		return
	}
}
