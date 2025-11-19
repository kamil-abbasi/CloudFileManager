package files

type FileService struct {
	repository FilesRepository
}

func NewService() *FileService {

	repository := NewSQLiteRepository()

	return &FileService{
		repository: repository,
	}
}

func (service *FileService) FindOne(id string) (File, bool, error) {
	return service.repository.FindOne(id)
}

func (service *FileService) Create(dto FileCreateDto) (File, error) {
	return service.repository.Create(dto)
}

func (service *FileService) Update(dto FileUpdateDto) (File, bool, error) {
	return service.repository.Update(dto)
}

func (service *FileService) Remove(id string) (File, bool, error) {
	return service.repository.Remove(id)
}
