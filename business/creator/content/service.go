package content

import (
	"project-art-museum/business/creator/content/spec"

	validator "github.com/go-playground/validator/v10"
)

type Repository interface {
	GetAll() (contents []Content, err error)
	CreateContent(spec spec.UpsertContentSpec) (err error)
	UpdateContent(content Content, currentVersion int) (err error)
	FindContentByID(id int) (content *Content, err error)
	DeleteContent(id int) (content *Content, err error)
}

type Service interface {
	GetAll() (contents []Content, err error)
	CreateContent(spec spec.UpsertContentSpec) (err error)
	UpdateContent(content Content, currentVersion int) (err error)
	FindContentByID(id int) (content *Content, err error)
	DeleteContent(id int) (content *Content, err error)
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

func (s *service) CreateContent(spec spec.UpsertContentSpec) (err error) {
	return
}

func (s *service) FindContentByID(id int) (content *Content, err error) {
	return
}

func (s *service) DeleteContent(id int) (content *Content, err error) {
	return
}

func (s *service) UpdateContent(content Content, currentVersion int) (err error) {
	return
}
