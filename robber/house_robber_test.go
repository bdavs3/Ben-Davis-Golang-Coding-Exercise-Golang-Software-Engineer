package robber

import "testing"

func TestRobber(t *testing.T) {
	tests := []struct {
		comment string
		houses  []int
		want    int
	}{
		{
			comment: "nil houses",
			houses:  nil,
			want:    0,
		},
		{
			comment: "Houses (1)",
			houses:  []int{1},
			want:    1,
		},
		{
			comment: "Houses (7, 3, 1, 5, 3)",
			houses:  []int{7, 3, 1, 5, 3},
			want:    12,
		},
		{
			comment: "Houses (8, 11, 6, 40)",
			houses:  []int{8, 11, 6, 40},
			want:    51,
		},
		{
			comment: "Houses (100, 5, 1, 103, 244, 24, 3, 77, 32, 100)",
			houses:  []int{100, 5, 1, 103, 244, 24, 3, 77, 32, 100},
			want:    426,
		},
	}

	for _, test := range tests {
		t.Run(test.comment, func(t *testing.T) {
			street := NewStreet(test.houses)
			got := street.Rob()

			if got != test.want {
				t.Errorf("got %v, want %v", got, test.want)
			}
		})
	}
}
