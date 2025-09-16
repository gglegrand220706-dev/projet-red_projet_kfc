package personnage

type Attaques struct {
	Name string
	Damage int
	LandingRate float32
}

var AttaqueLasser = Attaques{"Laser", 10, 70.5}
var CoupDePoing = Attaques{"Coup de poing", 2, 90}
var HighKick = Attaques{"Coup pied", 6, 85.5}
var GutPunch = Attaques{"coup estomac", 8, 80}

var AttaqueSeptre = Attaques{"coup de septre", 4, 90}
var AttaqueMentale =  Attaques{"tourment", 5, 80}
var AttaqueMagique = Attaques{"Boule Magqiue", 6, 90}

var AttaqueBarryPoing = Attaques{"coup de poing de Barry", 5, 90}
var AttaqueBarrypied = Attaques{"Coup de pied de Barry", 6, 80}
var AttaqueBarryTete = Attaques{"coup de tête de Barry", 10, 80}

var AttaqueBloquage = Attaques{"Blockage spatio-temporelle", 10, 80}
var AttaqueContrôle = Attaques{"Mind contrôle", 15, 75}
var AttaqueIlusion = Attaques{"Ilusion", 10, 80}

var AttaqueBasique01 = Attaques{"coup de dague", 2, 100}
var AttaqueBasique02 = Attaques{"coup de poing", 3, 100}
var AttaqueBasique03 = Attaques{"attaque magique", 4, 100}
var AttaqueBasique04 = Attaques{"pied bouche", 5, 100}

type Capacity struct {
	Name string
	Bonus string
	Rate int
}

var Vole = Capacity{"Vole", "attaques ont moins de chance de toucher leurs cibles", 60}
