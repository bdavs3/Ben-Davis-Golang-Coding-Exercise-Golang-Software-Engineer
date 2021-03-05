package robber

type Street struct {
	houses []int
}

func NewStreet(houses []int) *Street {
	return &Street{
		houses: houses,
	}
}

func (s *Street) Rob() {

}
