package personnage

type Attaques struct {
	Name string
	Damage int
	LandingRate float32
	NeededMastery int
	Critique int
}

var AttaqueLasser   = Attaques{"Laser 🔴", 10, 70.5, 0, 10}
var CoupDePoing     = Attaques{"Coup de poing ✊", 2, 90, 0, 20}
var HighKick        = Attaques{"Coup de pied 🦵", 6, 85.5, 0, 10}
var GutPunch        = Attaques{"Coup à l’estomac 🤜", 8, 80, 0, 20}

var ClapSuperSonic = Attaques{"Clape SuperSonic", 10, 80, 0, 10}

var RenvoieSinetique = Attaques{"Renvoie Sinetqiue", 30, 100, 0, 10}

var AttaqueHache    = Attaques{"Lancé de Hache 🪓", 10, 75, 10, 10}
var AttaqueHache02  = Attaques{"Hache de foudre ⚡🪓", 12, 75, 20, 10}

var  RayonLaserDorm = Attaques{"Hurlement Laser", 17, 75, 0, 10}
var LancesDorme = Attaques{"Lances Planètaires", 20, 75, 0, 10} 

var AttaqueGantThanos   = Attaques{"Pierre du pouvoir 💎", 6, 80, 10, 10}
var AttaqueGantThanos02 = Attaques{"Pierre de l'esprit 🧠💎", 7, 80, 10, 10}
var AttaqueGantThanos03 = Attaques{"Pierre de la réalité 🌌💎", 8, 80, 15, 10}
var AttaqueGantThanos04 = Attaques{"Pierre de l'espace 🌍💎", 9, 82, 15, 10}
var AttaqueGantThanos05 = Attaques{"Pierre de l'âme 🕊️💎", 10, 82, 20, 10}
var AttaqueGantSnap     = Attaques{"Inéluctable ✋💥", int(CurrentAdversery[AdverseryChoice].Vieactuelle / 2), 100, 25, 0}

var AttaquesEpeeTraining = Attaques{"Coup d'estoque ⚔️", 5, 75, 0, 10}

var AttaquesMarteauLancé  = Attaques{"Lancé de Marteau 🔨", 5, 65, 5, 10}
var AttaquesMarteauFoudre = Attaques{"Foudre du marteau ⚡🔨", 7, 75, 10, 10}

var AttaqueSeptre  = Attaques{"Coup de sceptre 🪄", 4, 90, 0, 10}
var AttaqueMentale = Attaques{"Tourment 🌀", 5, 80, 0, 10}
var AttaqueMagique = Attaques{"Boule Magique 🔮", 6, 90, 0, 10}

var AttaqueBarryPoing = Attaques{"Coup de poing de Barry ⚡✊", 5, 90, 0, 10}
var AttaqueBarrypied  = Attaques{"Coup de pied de Barry ⚡🦵", 6, 80, 0, 10}
var AttaqueBarryTete  = Attaques{"Coup de tête de Barry ⚡🤯", 10, 80, 0, 10}

var AttaqueUpperCut = Attaques{"Uppercut semi dévastateur 💥", 8, 90, 0, 10}
var AttaqueHeal     = Attaques{"Talon descendant 🦶", 12, 75, 0, 10}
var AttaqueCoude    = Attaques{"Coup de coude 💪", 15, 80, 0, 10}

var AttaqueBloquage = Attaques{"Blocage spatio-temporel ⏳", 10, 80, 0, 10}
var AttaqueContrôle = Attaques{"Manipulation de réalité 🌌", 15, 75, 0, 10}
var AttaqueIlusion  = Attaques{"Illusion 👁️", 10, 80, 0, 10}

var GodSlicer = Attaques{"Trancher divin", 12, 80, 10, 10}
var TrancheMagique = Attaques{"vague magique tranchente", 15, 75, 20, 15}

var AttaqueGaz     = Attaques{"Gaz hilarant 😂", 10, 100, 0, 10}
var AttaqueMarteau = Attaques{"Coup de Marteau 🎯🔨", 20, 80, 0, 10}
var AttaqueBlague  = Attaques{"Corruption de l'âme 🖤", 25, 80, 0, 10}

var AttaqueBasique01 = Attaques{"Coup de dague 🗡️", 2, 100, 0, 10}
var AttaqueBasique02 = Attaques{"Coup de poing ✊", 3, 100, 0, 10}
var AttaqueBasique03 = Attaques{"Attaque magique ✨", 4, 100, 0, 10}
var AttaqueBasique04 = Attaques{"Pied bouche 🦶", 5, 100, 0, 10}

var MarteauRonan = Attaques{"Coup de marteau", 15, 80, 0, 10}
var Onde = Attaques{"Onde sismique", 20, 75, 0, 10}
var PowerRush = Attaques{"Power Rush", 30, 75, 0, 10}

var LaserD = Attaques{"Laser guidé", 20, 90, 0, 10}
var Lune = Attaques{"Lancé de lune", 30, 70, 0, 10}
var EAV = Attaques{"Équation anti-vie", 0, 40, 0, 0}

type Capacity struct {
    Name  string
    Bonus string
	Used bool
}

var PierreTemps = Capacity{"Distortion Espace-Temps ⏳🌌", "Rejoue 🔁", false}

type Array struct {
	Att []Attaques
	Cap []Capacity
	Name []string
}

var PourUnSeulMec = Array{[]Attaques{RayonLaserDorm, LancesDorme}, []Capacity{PierreTemps}, []string{"Dormammu", AllWeapons[1].Name}}