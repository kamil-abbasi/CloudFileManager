package files

import (
	"database/sql"
	"fmt"

	"github.com/kamil-abbasi/CloudFileOperationsBackend/internal/database"
)

type FilesSQLiteRepository struct {
	db *sql.DB
}

func NewSQLiteRepository() FilesRepository {
	db := database.GetInstance()

	return &FilesSQLiteRepository{
		db: db,
	}
}

func (repository *FilesSQLiteRepository) FindOne(id string) (File, bool, error) {
	return File{}, true, nil
}
func (repository *FilesSQLiteRepository) Create(dto FileCreateDto) (File, error) {

	_, err := repository.db.Exec("INSERT INTO files VALUES(?,?,?,?)", dto.Id, dto.Filename, dto.Location, dto.Size)

	if err != nil {
		return File{}, fmt.Errorf("failed to create file metadata, details: %v", err)
	}

	return File(dto), nil
}
func (repository *FilesSQLiteRepository) Update(dto FileUpdateDto) (File, bool, error) {
	return File{}, true, nil
}
func (repository *FilesSQLiteRepository) Remove(id string) (File, bool, error) {
	return File{}, true, nil
}
