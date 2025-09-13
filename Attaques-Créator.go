package personnage

type Attaques struct {
	Name string
	Damage int
	LandingRate float32
}

var AttaqueLasser = Attaques{"Laser", 10, 50.5}
var CoupDePoing = Attaques{"Coup", 2, 70}
var HighKick = Attaques{"Coup pied", 5, 60.5}
var GutPunch = Attaques{"coup estomac", 8, 50}

var AttaqueBasique01 = Attaques{"coup de dague", 2, 100}
var AttaqueBasique02 = Attaques{"coup de dague", 3, 100}
var AttaqueBasique03 = Attaques{"coup de dague", 4, 100}
var AttaqueBasique04 = Attaques{"coup de dague", 5, 100}

type Capacity struct {
	Name string
	Bonus string
	Rate int
}

var Vole = Capacity{"Vole", "attaques ont moins de chance de toucher leurs cibles", 60}