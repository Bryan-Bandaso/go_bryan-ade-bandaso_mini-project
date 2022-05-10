package spec

type UpsertContentSpec struct {
	ID          int    `validate:"required"`
	Name        string `validate:"required"`
	Nationality string `validate:"required"`
	Description string `validate:"required"`
	Biography   string `validate:"required"`
	Birth_year  string `validate:"required"`
	Death_year  string `validate:"required"`
}
