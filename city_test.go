package city

import (
	"fmt"
	"testing"
)

func TestGetDistance(t *testing.T) {
	c, err := NewCityConfig("city_info.csv")
	if err != nil {
		t.Errorf("new city config failed, %v", err)
	}

	var dist float64
	dist, err = c.GetDistance("广东深圳", "广东汕头")
	if err != nil {
		t.Errorf("GetDistance, %v", err)
	}
	fmt.Printf("distance of `广东深圳` and `广东汕头`: %.1f(m)\n", dist)
}
