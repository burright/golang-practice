package space

// Calculate the constant of orbital period in Earth seconds
var OrbitalPeriod = map[string]float64{
	"Earth":   31557600,
	"Mercury": 7600543.81992,
	"Venus":   19414149.052176,
	"Mars":    59354032.69008,
	"Jupiter": 374355659.124,
	"Saturn":  929292362.8848,
	"Uranus":  2651370019.3296,
	"Neptune": 5200418560.032,
}

type Planet string

func Age(sec float64, planet Planet) float64 {

	// If the planet exists in the map
	if orbit, ok := OrbitalPeriod[string(planet)]; ok {
		// Calculate the age in seconds
		return sec / orbit
	}

	return 0
}
