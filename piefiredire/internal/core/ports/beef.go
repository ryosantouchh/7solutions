package ports

type BeefService interface {
	// in this case, I'm going to fetch data from mock url instead of query from db
	Get() (string, error)

	// TODO : define other methods below
}
