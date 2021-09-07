package sqlboilerorm

type ComponentRegistry struct {
	DBConn           RDBConnector
	BoilerRepository IBoilerRepository
}

func NewComponentRegistry(
	dbConn RDBConnector,
	boilerRepository IBoilerRepository) *ComponentRegistry {
	return &ComponentRegistry{
		DBConn:           dbConn,
		BoilerRepository: boilerRepository,
	}
}
