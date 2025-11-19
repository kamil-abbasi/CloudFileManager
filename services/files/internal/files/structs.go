package files

type File struct {
	Id       string
	Filename string
	Location string
	Size     uint64
}

type FileCreateDto struct {
	Id       string
	Filename string
	Location string
	Size     uint64
}

type FileUpdateDto struct {
	Id     string
	Fields struct {
		Filename string
		Location string
	}
}
