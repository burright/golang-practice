package space

var OrbitalPeriod = map[string]float64 {
	"Earth" : 31557600,
	"Mercury" : 0.2408467,
	"Venus" : 0.61519726,
	"Mars" : 1.8808158,
	"Jupiter" : 11.862615,
	"Saturn" : 29.447498,
	"Uranus" : 84.016846,
	"Neptune" : 164.79132,
}

type Planet string


func Age (sec float64, planet Planet) float64 {

	if orbit, ok := OrbitalPeriod[string(planet)]; ok {
		return sec / orbit 
	}
	
	return 0
}

