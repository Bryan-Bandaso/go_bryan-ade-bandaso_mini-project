package response

import (
	"project-art-museum/business/content"
)

type CreateNewContentResponse struct {
	ID string `json:"id"`
}

type ContentResponse struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Nationality string `json:"nationality"`
	Description string `json:"description"`
	Biography   string `json:"biography"`
	Birth_year  string `json:"birth_year"`
	Death_year  string `json:"death_year"`
	Version     int    `json:"version"`

	Artworks struct {
		ID               int    `json:"id"`
		Accession_number string `json:"accession_number"`
		Title            string `json:"title"`
		Tombstone        string `json:"tombstone"`
		Url              string `json:"url"`
	} `json:"artworks"`
}

func NewCreateNewContentResponse(id string) *CreateNewContentResponse {
	return &CreateNewContentResponse{
		id,
	}
}

func NewGetContentByIDResponse(content content.Content) *ContentResponse {
	var contentResponse ContentResponse
	contentResponse.Name = content.Name
	contentResponse.Nationality = content.Nationality
	contentResponse.Description = content.Description
	contentResponse.Biography = content.Biography
	contentResponse.Birth_year = content.Birth_year
	contentResponse.Death_year = content.Death_year

	return &contentResponse
}
