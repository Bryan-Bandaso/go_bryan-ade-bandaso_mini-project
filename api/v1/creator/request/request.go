package request

import (
	spec "project-art-museum/business/content/spec"
)

type CreateContentRequest struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Nationality string `json:"nationality"`
	Description string `json:"description"`
	Biography   string `json:"biography"`
	Birth_year  string `json:"birth_year"`
	Death_year  string `json:"death_year"`
	Version     int    `json:"version" validate:"required"`

	Artworks struct {
		ID               int    `json:"id"`
		Accession_number string `json:"accession_number"`
		Title            string `json:"title"`
		Tombstone        string `json:"tombstone"`
		Url              string `json:"url"`
	} `json:"artworks"`
}

func (req *CreateContentRequest) ToSpec() *spec.UpsertContentSpec {
	var upsertContentSpec spec.UpsertContentSpec
	upsertContentSpec.Name = req.Name
	upsertContentSpec.Nationality = req.Nationality
	upsertContentSpec.Description = req.Description
	upsertContentSpec.Biography = req.Biography
	upsertContentSpec.Birth_year = req.Birth_year
	upsertContentSpec.Death_year = req.Death_year

	return &upsertContentSpec
}
