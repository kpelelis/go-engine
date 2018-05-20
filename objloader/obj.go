package objloader

type Obj struct {
	Name     string
	Material string
	Vertices []Vertex
	UVs      []UV
	Normals  []Normal
	Faces    []Face
}
