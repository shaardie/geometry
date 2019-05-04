package geometry

type Point struct {
	X, Y float64
}

func (p Point) Eq(p1 Point) bool {
	return p.X == p1.X && p.Y == p1.Y
}

func (p Point) RayCastingAlgorithm(poly Polygon) int {
	return rayCastingAlgorithm(poly, p)
}

func (p Point) PointInPolygon(poly Polygon) bool {
	return poly.RayCastingAlgorithm(p) == 1
}

func (p Point) PointOnPolygon(poly Polygon) bool {
	return poly.RayCastingAlgorithm(p) == 0
}
