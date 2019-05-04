package geometry

type Polygon []Point

func (poly Polygon) RayCastingAlgorithm(p Point) int {
	return rayCastingAlgorithm(poly, p)
}

func (poly Polygon) PointInPolygon(p Point) bool {
	return poly.RayCastingAlgorithm(p) == 1
}

func (poly Polygon) PointOnPolygon(p Point) bool {
	return poly.RayCastingAlgorithm(p) == 0
}
