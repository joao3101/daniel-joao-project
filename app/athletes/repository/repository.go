package athlete_repository

type AthleteModelData struct {
	ID        int    `json:"id"`
	ClubID    int    `json:"club_id"`
	Name      string `json:"name"`
	Age       int64  `json:"age"`
	Position  int    `json:"position"`
	CreatedAt int64  `json:"created_at"`
	DeletedAt int    `json:"deleted_at"`
}
