package city

import (
	"encoding/csv"
	"math"
	"os"
	"strconv"
)

var (
	radius = float64(6371000)
	rad    = math.Pi / 180.0
)

// info 城市经纬度信息
type info struct {
	name string
	lng  float64 //经度
	lat  float64 //纬度
}

// Manager 城市管理
type Manager struct {
	citys map[string]*info
}

// NewCityConfig 读取坐标配置，创建管理器
func NewCityConfig(file string) (c *Manager, err error) {
	c = &Manager{}
	c.citys, err = c.parseCityInfo(file)
	if err != nil {
		return
	}

	return
}

// GetDistance 计算两个城市间的距离
func (c *Manager) GetDistance(city1 string, city2 string) (dist float64, err error) {
	c1 := c.citys[city1]
	c2 := c.citys[city2]
	if c1 == nil || c2 == nil {
		return 0.0, nil
	}

	dist = c.earthDistance(c1, c2)

	return
}

// earthDistance 计算距离
func (c *Manager) earthDistance(city1 *info, city2 *info) float64 {
	lng1 := city1.lng * rad
	lat1 := city1.lat * rad
	lng2 := city2.lng * rad
	lat2 := city2.lat * rad
	theta := lng2 - lng1
	dist := math.Acos(math.Sin(lat1)*math.Sin(lat2) + math.Cos(lat1)*math.Cos(lat2)*math.Cos(theta))
	return dist * radius
}

// parseCityInfo 载入城市经纬度
func (c *Manager) parseCityInfo(filepath string) (locations map[string]*info, err error) {
	f, err := os.Open(filepath)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	reader := csv.NewReader(f)
	record, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}

	locations = make(map[string]*info, 0)
	for _, a := range record {
		city := &info{name: a[0]}
		lng, err := strconv.ParseFloat(a[1], 64)
		if err == nil {
			city.lng = lng
		}
		lat, err := strconv.ParseFloat(a[2], 64)
		if err == nil {
			city.lat = lat
		}
		locations[city.name] = city
	}

	return locations, nil
}
