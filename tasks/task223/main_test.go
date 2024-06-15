package main

import (
	"fmt"
	"testing"
)

type MockWS struct {
	value int
}

func (m *MockWS) Forecast() int {
	return m.value
}

type testCase struct {
	deg  int
	want string
}

var tests []testCase = []testCase{
	{-10, "холодно"},
	{0, "холодно"},
	{5, "холодно"},
	{10, "прохладно"},
	{15, "идеально"},
	{20, "жарко"},
}

func TestForecast(t *testing.T) {
	for _, test := range tests {
		name := fmt.Sprintf("%v", test.deg)
		service := MockWS{test.deg}
		weather := Weather{&service}
		t.Run(name, func(t *testing.T) {
			got := weather.Forecast()
			if got != test.want {
				t.Errorf("%s: got %s, want %s", name, got, test.want)
			}
		})
	}
}
