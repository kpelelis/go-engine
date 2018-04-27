package objloader

type Obj struct {
	Name      string
	Vertices  []Vertex
	UVs       []UV
	Normals   []Normal
	Triangles []Triangle
}
