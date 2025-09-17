package personnage

type Attaques struct {
	Name string
	Damage int
	LandingRate float32
	NeededMastery int
}

var AttaqueLasser = Attaques{"Laser", 10, 70.5, 0}
var CoupDePoing = Attaques{"Coup de poing", 2, 90, 0}
var HighKick = Attaques{"Coup pied", 6, 85.5, 0}
var GutPunch = Attaques{"coup estomac", 8, 80, 0}

var AttaqueHache = Attaques{"Lancé de Hache", 10, 75, 10}
var AttaqueHache02 = Attaques{"Hache de foudre", 12, 75, 20}

var AttaqueGantThanos = Attaques{"Pierre du pouvoir", 6, 80, 10}
var AttaqueGantThanos02 = Attaques{"Pierre de l'esprit", 7, 80, 10}
var AttaqueGantThanos03 = Attaques{"Pierre de la réalité", 8, 80, 15}
var AttaqueGantThanos04 = Attaques{"Pierre de l'espace", 9, 82, 15}
var AttaqueGantThanos05 = Attaques{"Pierre de l'âme", 10, 82, 20}
var AttaqueGantThanos06 = Attaques{"Pierre du temps", 11, 82, 20}
var AttaqueGantSnap = Attaques{"Ineluctable", int(CurrentAdversery[AdverseryChoice].Vieactuelle/2), 100, 25}

var AttaquesEpeeTraining = Attaques{"Coup d'estoque", 5, 75, 5}

var AttaquesMarteauLancé = Attaques{"Lancé Marteau", 5, 65, 5}
var AttaquesMarteauFoudre = Attaques{"Foudre du marteau", 7, 75, 10}

var AttaqueSeptre = Attaques{"coup de septre", 4, 90, 0}
var AttaqueMentale =  Attaques{"tourment", 5, 80, 0}
var AttaqueMagique = Attaques{"Boule Magqiue", 6, 90, 0}

var AttaqueBarryPoing = Attaques{"coup de poing de Barry", 5, 90, 0}
var AttaqueBarrypied = Attaques{"Coup de pied de Barry", 6, 80, 0}
var AttaqueBarryTete = Attaques{"coup de tête de Barry", 10, 80, 0}

var AttaqueUpperCut = Attaques{"Upper cut semi devastateur", 8, 90, 0}
var AttaqueHeal = Attaques{"talon descandant", 12, 75, 0}
var AttaqueCoude = Attaques{"Coup de coude", 15, 80, 0}

var AttaqueBloquage = Attaques{"Blockage spatio-temporelle", 10, 80, 0}
var AttaqueContrôle = Attaques{"Manipulation de realitée", 15, 75, 0}
var AttaqueIlusion = Attaques{"Ilusion", 10, 80, 0}

var AttaqueGaz = Attaques{"Gaz Hilarant", 10, 100, 0}
var AttaqueMarteau = Attaques{"Coup de Marteau", 20, 80, 0}
var AttaqueBlague = Attaques{"Coruption de l'ame", 25, 80, 0}

var AttaqueBasique01 = Attaques{"coup de dague", 2, 100, 0}
var AttaqueBasique02 = Attaques{"coup de poing", 3, 100, 0}
var AttaqueBasique03 = Attaques{"attaque magique", 4, 100, 0}
var AttaqueBasique04 = Attaques{"pied bouche", 5, 100, 0}

type Capacity struct {
	Name string
	Bonus string
	Rate int
}

var Vole = Capacity{"Vole", "attaques ont moins de chance de toucher leurs cibles", 60}
