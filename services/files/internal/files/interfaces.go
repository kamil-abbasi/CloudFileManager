package files

type FilesRepository interface {
	FindOne(id string) (File, bool, error)
	Update(dto FileUpdateDto) (File, bool, error)
	Remove(id string) (File, bool, error)
	Create(dto FileCreateDto) (File, error)
}
