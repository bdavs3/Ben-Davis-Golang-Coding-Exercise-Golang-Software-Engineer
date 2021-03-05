package tree

import "testing"

func TestTree(t *testing.T) {
	var tests = []struct {
		comment string
		want    string
	}{
		{
			comment: "Tall",
			want:    "Tall",
		},
		{
			comment: "Beautiful",
			want:    "Beautiful",
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
