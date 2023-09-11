package utils

type triangle struct {
	Vertices [9]float32
}

func NewTriangle(vertices [9]float32) *triangle {
	return &triangle{Vertices: vertices}
}

type Quad struct {
	Tris [2]*triangle
}

func NewQuad(tris [2]*triangle) *Quad {
	return &Quad{Tris: tris}
}

func (quad Quad) PointsFromTris() []float32 {
	points := make([]float32, 0)
	for _, tri := range quad.Tris {
		points = append(points, tri.Vertices[:]...)
	}
	return points
}

type GameObject struct {
	drawable uint32
	ID       string
	Pos      Vector2
	Quads    []*Quad
}

func NewGameObject(ID string, pos Vector2, QuadPoints [][4][3]float32) *GameObject {
	quads := QuadsFromPoints(QuadPoints)

	return &GameObject{ID: ID, Pos: pos, Quads: quads}
}

func (gameobject GameObject) TrisFromQuads() []triangle {
	tris := make([]triangle, 0)
	for _, quad := range gameobject.Quads {
		for _, tri := range quad.Tris {
			tris = append(tris, *tri)
		}
	}
	return tris
}

func (gameobject GameObject) PointsFromQuads() []float32 {
	points := make([]float32, 0)
	for _, quad := range gameobject.Quads {
		points = append(points, quad.PointsFromTris()[:]...)
	}
	return points
}

func QuadsFromPoints(QuadPoints [][4][3]float32) []*Quad {
	quads := make([]*Quad, len(QuadPoints))
	for _, quad := range QuadPoints {
		tripoints := append([]float32{}, quad[0][:]...)
		tripoints = append(tripoints, quad[1][:]...)
		tripoints = append(tripoints, quad[2][:]...)
		tri1 := NewTriangle([9]float32(tripoints))

		tripoints = append([]float32{}, quad[0][:]...)
		tripoints = append(tripoints, quad[2][:]...)
		tripoints = append(tripoints, quad[3][:]...)
		tri2 := NewTriangle([9]float32(tripoints))

		tris := [2]*triangle{tri1, tri2}

		quads = append(quads, NewQuad(tris))
	}

	return quads[1:]
}
