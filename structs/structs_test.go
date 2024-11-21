package structs

import "testing"

func TestPerimeter(t *testing.T) {
	rect := Rectangle{10.0, 20.0}
	got := Perimeter(rect)
	want := 60.0

	if got != want {
		t.Errorf("got %.2f want %.2f", got, want)
	}
}

func TestArea(t *testing.T) {

	areaTests := []struct {
		name  string
		shape Shape
		area  float64
	}{
		{name: "Rectangle", shape: Rectangle{10.0, 20.0}, area: 200.0},
		{name: "Circle", shape: Circle{10.0}, area: 314.1592653589793},
		{name: "Triangle", shape: Triangle{5, 3}, area: 7.5},
	}

	for _, element := range areaTests {
		t.Run(element.name, func(t *testing.T) {
			got := element.shape.Area()
			if got != element.area {
				t.Errorf("got Area %g but want an Area  %g", got, element.area)
			}
		})

	}

}
