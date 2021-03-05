package robber

type Street struct {
	houses []int
}

func NewStreet(houses []int) *Street {
	return &Street{
		houses: houses,
	}
}

func (s *Street) Rob() int {
	numHouses := len(s.houses)
	if numHouses == 0 {
		return s.houses[0]
	} else if numHouses == 0 {
		return max(s.houses[0], s.houses[1])
	} else {
		maxAccumulatedValues := make([]int, numHouses)
		maxAccumulatedValues[0] = s.houses[0]
		maxAccumulatedValues[1] = max(s.houses[0], s.houses[1])

		for i := 2; i < numHouses; i++ {
			maxAccumulatedValues[i] = max(s.houses[i]+maxAccumulatedValues[i-2], maxAccumulatedValues[i-1])
		}

		return maxAccumulatedValues[numHouses-1]
	}
}

func max(x, y int) int {
	if x < y {
		return y
	}
	return x
}
