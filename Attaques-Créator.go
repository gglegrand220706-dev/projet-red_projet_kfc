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

var AttaqueHache = Attaques{"Lancé de Hache", 10, 75}
var AttaqueHache02 = Attaques{"Hache de foudre", 12, 75}

var AttaqueGantThanos = Attaques{"Pierre du pouvoir", 6, 80}
var AttaqueGantThanos02 = Attaques{"Pierre de l'esprit", 7, 80}
var AttaqueGantThanos03 = Attaques{"Pierre de la réalité", 8, 80}
var AttaqueGantThanos04 = Attaques{"Pierre de l'espace", 9, 82}
var AttaqueGantThanos05 = Attaques{"Pierre de l'âme", 10, 82}
var AttaqueGantThanos06 = Attaques{"Pierre du temps", 11, 82}
var AttaqueGantSnap = Attaques{"Ineluctable", int(CurrentAdversery[AdverseryChoice].Vieactuelle/2), 100}

var AttaquesEpeeTraining = Attaques{"Coup d'estoque", 5, 75}

var AttaquesMarteauLancé = Attaques{"Lancé Marteau", 5, 65}
var AttaquesMarteauFoudre = Attaques{"Foudre du marteau", 7, 75}

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
