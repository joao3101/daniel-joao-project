package domain

type Team struct {
	UUID  string `gorm:"primary_key" json:"uuid"`
	Name  string `json:"name"`
	State int64  `json:"state"`
}
