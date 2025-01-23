package storage

import (
	"errors"
	"fmt"

	"github.com/google/uuid"
	"github.com/vas-ide/storage/internal/file"
)



type Storage struct{
	files map[uuid.UUID]*file.File

}


func NewStorage() *Storage {
	return &Storage{
		files: make(map[uuid.UUID]*file.File),
	}
}



 
func(s Storage) Upload(filename string, blob []byte) (*file.File, error){
	newFile, err := file.NewFile(filename,blob)
	if err != nil {
		return nil, err
	}

	s.files[newFile.ID] = newFile
	return newFile,err
}

func (s *Storage) GetByID(fileID uuid.UUID) (*file.File, error) {
	reqFile, ok := s.files[fileID]
	if !ok {
		return nil, errors.New(fmt.Sprintf("File %v not found", fileID))
	}

	return reqFile,nil
}




