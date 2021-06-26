package space

// Planet - the planet type
type Planet string

const secondsInEY float64 = 31557600

var orbitalPeriodMap = map[Planet]float64{
	"Mercury": 0.2408467,
	"Venus":   0.61519726,
	"Earth":   1.0,
	"Mars":    1.8808158,
	"Jupiter": 11.862615,
	"Saturn":  29.447498,
	"Uranus":  84.016846,
	"Neptune": 164.79132,
}

// Age will calculate the space age on the given planet based on the age provided (in seconds).
func Age(seconds float64, planet Planet) float64 {
	earthYears := seconds / secondsInEY
	return earthYears / orbitalPeriodMap[planet]

}
