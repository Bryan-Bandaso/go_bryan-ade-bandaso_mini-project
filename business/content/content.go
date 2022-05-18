package content

import "time"

type Content struct {
	ID          string
	Name        string
	Nationality string
	Description string
	Biography   string
	Birth_year  string
	Death_year  string
	CreatedAt   time.Time
	Version     int
	//ModifiedAt  time.Time
}

func NewContent(
	name string,
	nationality string,
	description string,
	biography string,
	birth_year string,
	death_year string,
	//ModifiedAt time.Time,
	createdAt time.Time) Content {

	return Content{
		Name:        name,
		Nationality: nationality,
		Description: description,
		Biography:   biography,
		Birth_year:  birth_year,
		Death_year:  death_year,
		CreatedAt:   createdAt,
		Version:     1,
		//ModifiedAt:  createdAt,
	}
}

func (oldContent *Content) ModifyContent(newName string, newNationality string, newDescription string, newBiography string, newBirth_year string, newDead_year string) Content {
	return Content{
		ID:          oldContent.ID,
		Name:        newName,
		Nationality: newNationality,
		Description: newDescription,
		Biography:   newBiography,
		Birth_year:  newBirth_year,
		Death_year:  newDead_year,
		CreatedAt:   oldContent.CreatedAt,
		Version:     oldContent.Version + 1,
	}
}
