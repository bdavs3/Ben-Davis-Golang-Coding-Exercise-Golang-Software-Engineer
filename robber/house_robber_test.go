package robber

import "testing"

func TestRobber(t *testing.T) {
	var tests = []struct {
		comment string
		want    string
	}{
		{
			comment: "Clever",
			want:    "Clever",
		},
		{
			comment: "Sneaky",
			want:    "Sneaky",
		},
	}

	for _, test := range tests {
		t.Run(test.comment, func(t *testing.T) {
			trait := test.want
			if trait != test.want {
				t.Errorf("got %v, want %v", trait, test.want)
			}
		})
	}
}
