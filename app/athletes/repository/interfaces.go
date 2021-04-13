package athlete_repository

type AthleteModel interface {
	GetAtlhete(int) (AthleteModelData, error)
	InsertAtlhete(AthleteModelData) (AthleteModelData, error)
}