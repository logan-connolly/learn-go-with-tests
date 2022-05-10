package structs

import "testing"

func TestAreas(t *testing.T) {
	testCases := []struct {
		name    string
		shape   Shape
		hasArea float64
	}{
		{"Rectangle", Rectangle{Width: 10.0, Height: 5.0}, 50.0},
		{"Circle", Circle{Radius: 10.0}, 314.1592653589793},
		{"Triangle", Triangle{Base: 2.0, Height: 5.0}, 0.0},
	}

	for _, test := range testCases {
		t.Run(test.name, func(t *testing.T) {
			got := test.shape.Area()
			if test.hasArea != got {
				t.Errorf("shape %#v want %g got %g", test.shape, test.hasArea, got)
			}
		})
	}
}
