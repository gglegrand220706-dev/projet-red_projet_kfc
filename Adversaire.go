package personnage

type Adversery struct {
	Nom         		string
	Vieactuelle 		int
	Niveau      		int
	Viemax       		int
	Puissance    		int
	Faiblaisse   		[]string
	Agilite      		int
	Intelligence 		int
	CapacityDisplay     []string
	Attaques 			[]Attaques
	Capacity 			[]Capacity
	GivenExp			int
	GivenMoney			int
	Drop 				[]Objects
}

var Adversery01 = Adversery{"Entraîneur : Kentino 🥋", 30, 1, 30, 50, []string{}, 50, 50, []string{"Coup de poing boosté ✊", "Prout brumisateur 💨", "Lancé de casquette 🧢"}, []Attaques{AttaqueBasique01, AttaqueBasique02, AttaqueBasique03, AttaqueBasique04}, []Capacity{}, 20, 20, []Objects{ScrapMetal, CuireBat}}
var AdverseryLoki   = Adversery{"Boss : Loki 🪄", 40, 10, 40, 50, []string{}, 60, 65, []string{"Coup de sceptre 🪄", "Tourment 🌀", "Boule de magie 🔮"}, []Attaques{AttaqueSeptre, AttaqueMentale, AttaqueMagique}, []Capacity{}, 30, 30, []Objects{InfinityStoneMind, CuireBat, ScrapMetal}}
var AdverseryBarry  = Adversery{"Barry (-_°) ⚡", 10, 1, 10, 10, []string{}, 10, 10, []string{"Coup de poing de Barry ⚡✊", "Coup de pied de Barry ⚡🦵", "Coup de tête de Barry ⚡🤯"}, []Attaques{AttaqueBarryPoing, AttaqueBarrypied, AttaqueBarryTete}, []Capacity{}, 10, 10, []Objects{ScrapMetal, CuireBat}}
var AdverseryRicooo = Adversery{"Rico (o°o ) 🥊", 15, 5, 15, 15, []string{}, 15, 15, []string{"Uppercut semi dévastateur 💥", "Talon descendant 🦶", "Coup de coude 💪"}, []Attaques{AttaqueUpperCut, AttaqueHeal, AttaqueCoude}, []Capacity{}, 15, 15, []Objects{ScrapMetal, ScrapMetal, CuireBat}}
var AdverseryWanda  = Adversery{"Sorcière Wanda 🪄", 60, 20, 40, 60, []string{}, 50, 70, []string{"Blocage spatio-temporel ⏳", "Manipulation de réalité 🌌", "Illusion 👁️"}, []Attaques{AttaqueBloquage, AttaqueContrôle, AttaqueIlusion}, []Capacity{}, 50, 50, []Objects{InfinityStoneReality, CuireBat, ScrapMetal}}
var AdverseryJoker  = Adversery{"Le Joker 🎭", 80, 30, 80, 60, []string{}, 60, 70, []string{"Gaz hilarant 😂", "Coup de Marteau 🎯🔨", "Corruption de l'âme 🖤"}, []Attaques{AttaqueGaz, AttaqueMarteau, AttaqueBlague},[]Capacity{}, 60, 60, []Objects{InfinityStoneSoule, CuireBat, ScrapMetal, CuireBat}}
var Dormammu = Adversery{"Dormammu 🌀", 80, 40, 80, 80, []string{}, 65, 75, []string{"Hurlement Laser 🔴", "Distortion Espace-Temps ⏳🌌", "Lances Planètaires 🪐"}, []Attaques{RayonLaserDorm, LancesDorme}, []Capacity{PierreTemps}, 100, 100, []Objects{CuireBat, ScrapMetal, InfinityStoneTime, GrootsBranches}}