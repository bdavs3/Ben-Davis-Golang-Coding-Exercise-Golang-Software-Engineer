package robber

// A Street contains a list of houses.
type Street struct {
	houses []int
}

// NewStreet initializes a Street with the given list of houses.
func NewStreet(houses []int) *Street {
	return &Street{
		houses: houses,
	}
}

// Rob calculates the best possible sum for the "House Robber" problem using the given receiver.
func (s *Street) Rob() int {
	numHouses := len(s.houses)

	if s.houses == nil || numHouses == 0 {
		return 0
	}
	if numHouses == 1 {
		return s.houses[0]
	}

	// To account for the circularity of the street, calculate two values and then take the max.
	// max1 includes the first house and not the last, whereas max2 includes the last house and not the first.
	max1 := calculateEarnings(s.houses, 0, numHouses-2)
	max2 := calculateEarnings(s.houses, 1, numHouses-1)

	return max(max1, max2)
}

// calculateEarnings uses dynamic programming to compute an optimal sum from a list of values with the constraint
// that adjacent values cannot both be included in the final sum.
func calculateEarnings(houses []int, start, end int) int {
	numHouses := len(houses)

	if start == end {
		return houses[start]
	}

	bestValues := make([]int, numHouses)
	bestValues[start] = houses[start]
	bestValues[start+1] = max(houses[start], houses[start+1])

	for i := start + 2; i <= end; i++ {
		bestValues[i] = max(houses[i]+bestValues[i-2], bestValues[i-1])
	}

	return bestValues[end]
}

// max calculates the maximum value among a pair of integers, something that is not made available in the Go standard library.
func max(x, y int) int {
	if x < y {
		return y
	}
	return x
}
