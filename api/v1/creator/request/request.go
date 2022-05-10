package request

import (
	spec "project-art-museum/business/creator/content/spec"
)

type CreateContentRequest struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Nationality string `json:"nationality"`
	Description string `json:"description"`
	Biography   string `json:"biography"`
	Birth_year  string `json:"birth_year"`
	Death_year  string `json:"death_year"`

	Artworks struct {
		ID               int    `json:"id"`
		Accession_number string `json:"accession_number"`
		Title            string `json:"title"`
		Tombstone        string `json:"tombstone"`
		Url              string `json:"url"`
	} `json:"artworks"`
}

func (req *CreateContentRequest) ToSpec() *spec.UpsertContentSpec {
	return &spec.UpsertContentSpec{
		ID:          req.ID,
		Name:        req.Name,
		Nationality: req.Nationality,
		Description: req.Description,
		Biography:   req.Biography,
		Birth_year:  req.Birth_year,
		Death_year:  req.Death_year,
	}
}
