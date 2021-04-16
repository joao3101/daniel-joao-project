package domain

import (
	"time"

)

type Athlete struct {
	ID        int       `boil:"id" json:"id" toml:"id" yaml:"id"`
	ClubID    *int  `boil:"club_id" json:"club_id,omitempty" toml:"club_id" yaml:"club_id,omitempty"`
	Name      string    `boil:"name" json:"name" toml:"name" yaml:"name"`
	Age       int       `boil:"age" json:"age" toml:"age" yaml:"age"`
	Position  int       `boil:"position" json:"position" toml:"position" yaml:"position"`
	CreatedAt time.Time `boil:"created_at" json:"created_at" toml:"created_at" yaml:"created_at"`
	DeletedAt *time.Time `boil:"deleted_at" json:"deleted_at,omitempty" toml:"deleted_at" yaml:"deleted_at,omitempty"`
}
