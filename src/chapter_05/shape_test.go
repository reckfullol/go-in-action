package shape

import "testing"

func TestPerimeter(t *testing.T) {
	rectangle := Rectangle{10.0, 10.0}
	got := Perimeter(rectangle)
	want := 40.0

	if got != want {
		t.Errorf("got %.2f want %.2f", got, want)
	}
}

func TestArea(t *testing.T) {

	areaTests := []struct {
		name  string
		shape Shape
		want  float64
	}{
		{"Rectangle", Rectangle{12, 6}, 72.0},
		{"Circle", Circle{10}, 314.1592653589793},
		{"Triangle", Triangle{12, 6}, 36.0},
	}

	for _, tt := range areaTests {
		// using tt.name from the case to use it as the 't.Run' test name
		t.Run(tt.name, func(t *testing.T) {
			got := tt.shape.Area()
			if got != tt.want {
				t.Errorf("got %.2f want %.2f", got, tt.want)
			}
		})

	}

	//checkArea := func(t *testing.T, shape Shape, want float64) {
	//    t.Helper()
	//    got := shape.Area()
	//    if got != want {
	//        t.Errorf("got %.2f want %.2f", got, want)
	//    }
	//}
	//
	//t.Run("rectangles", func(t *testing.T) {
	//    rectangle := Rectangle{12.0, 6.0}
	//    checkArea(t, rectangle, 72.0)
	//})
	//
	//
	//t.Run("circles", func(t *testing.T) {
	//    circle := Circle{10}
	//    checkArea(t, circle, 314.1592653589793)
	//})
}
