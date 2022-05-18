package content

import (
	"project-art-museum/business"
	"project-art-museum/business/content/spec"
	"time"

	validator "github.com/go-playground/validator/v10"
)

type Repository interface {
	GetAll() (contents []Content, err error)
	CreateContent(content Content) (ID string, err error)
	UpdateContent(content Content, currentVersion int) (err error)
	FindContentByID(ID string) (content *Content, err error)
	DeleteContent(id int) (content *Content, err error)
}

type Service interface {
	GetAll() (contents []Content, err error)
	CreateContent(upsertcontentSpec spec.UpsertContentSpec) (string, error)
	UpdateContent(ID string, upsertcontentSpec spec.UpsertContentSpec, currentVersion int, modifiedBy string) (err error)
	FindContentByID(ID string) (content *Content, err error)
	DeleteContent(ID int) (content *Content, err error)
}

type service struct {
	repository Repository
	validate   *validator.Validate
}

func NewService(repository Repository) Service {
	return &service{
		repository: repository,
		validate:   validator.New(),
	}
}

func (s *service) GetAll() (contents []Content, err error) {
	return
}

func (s *service) CreateContent(upsertcontentSpec spec.UpsertContentSpec) (string, error) {
	err := s.validate.Struct(upsertcontentSpec)

	if err != nil {
		return "", business.ErrInvalidSpec
	}

	content := NewContent(
		upsertcontentSpec.Name,
		upsertcontentSpec.Nationality,
		upsertcontentSpec.Description,
		upsertcontentSpec.Biography,
		upsertcontentSpec.Birth_year,
		upsertcontentSpec.Death_year,
		time.Now(),
	)
	ID, err := s.repository.CreateContent(content)
	if err != nil {
		return "", err
	}

	return ID, nil
}

func (s *service) FindContentByID(ID string) (content *Content, err error) {
	return s.repository.FindContentByID(ID)
}

func (s *service) UpdateContent(ID string, upsertcontentSpec spec.UpsertContentSpec, currentVersion int, modifiedBy string) error {
	err := s.validate.Struct(upsertcontentSpec)

	if err != nil || len(ID) == 0 {
		return business.ErrInvalidSpec
	}

	//get the content first to make sure data is exist
	content, err := s.repository.FindContentByID(ID)

	if err != nil {
		return err
	} else if content == nil {
		return business.ErrNotFound
		// } else if content.Version != currentVersion {
		// 	return business.ErrHasBeenModified
	}

	newContent := content.ModifyContent(upsertcontentSpec.Name, upsertcontentSpec.Nationality, upsertcontentSpec.Description,
		upsertcontentSpec.Biography, upsertcontentSpec.Birth_year, upsertcontentSpec.Death_year)

	return s.repository.UpdateContent(newContent, currentVersion)
}

func (s *service) DeleteContent(id int) (content *Content, err error) {
	return
}
