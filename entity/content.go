package entity

type Content struct {
	ID          int    `gorm:"primary_key:auto_increment" json:"-"`
	Name        string `gorm:"type:varchar(100)" json:"-"`
	Nationality string `gorm:"type:varchar(100)" json:"-"`
	Description string `gorm:"type:varchar(100)" json:"-"`
	Biography   string `gorm:"type:varchar(100)" json:"-"`
	Birth_year  string `gorm:"type:varchar(100)" json:"-"`
	Death_year  string `gorm:"type:varchar(100)" json:"-"`
}
