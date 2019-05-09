package geometry

func rayCastingAlgorithm(polygon Polygon, p Point) int {
	t := -1
	l := len(polygon)
	for i := 0; i < l; i++ {
		t = t * rightCross(p, polygon[i], polygon[(i+1)%l])
		if t == 0 {
			break
		}
	}
	return t
}

func rightCross(p1, p2, p3 Point) int {
	if p1.Y == p2.Y && p2.Y == p3.Y {
		if (p1.X <= p3.X && p1.X >= p1.X) || p1.X >= p3.X && p1.X <= p1.X {
			return 0
		}
		return 1
	}
	if p2.Y > p3.Y {
		p2, p3 = p3, p2
	}
	if p1.Y <= p2.Y || p1.Y >= p3.Y {
		return 1
	}
	delta := (p2.X-p1.X)*(p3.Y-p1.Y) - (p2.Y-p1.Y)*(p3.X-p1.X)
	if delta == 0 {
		return 0
	} else if delta < 0 {
		return 1
	}
	return -1
}
