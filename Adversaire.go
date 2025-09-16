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
	GivenExp			int
	GivenMoney			int
	Drop 				[]Objects
}

var Adversery01 = Adversery{"entrîneur : kentino", 30, 1, 30, 50, []string{}, 50, 50, []string{"coup de point bousté, proute brumisateur, lancé de casquet"}, []Attaques{AttaqueBasique01, AttaqueBasique02, AttaqueBasique03, AttaqueBasique04}, 20, 20, []Objects{ScrapMetal}}
var AdverseryLoki = Adversery{"boose : Loki", 40, 5, 40, 50, []string{}, 60, 65, []string{"coup de sptre", "tourment", "boule de magie"}, []Attaques{AttaqueSeptre, AttaqueMentale, AttaqueMagique}, 30, 30, []Objects{InfinityStoneMind} }
var AdverseryBarry = Adversery{"Barry (-_°)", 10, 1, 10, 10, []string{}, 10, 10, []string{"coup de poing de Barry", "Coup de pied de Barry", "coup de tête de Barry"}, []Attaques{AttaqueBarryPoing, AttaqueBarrypied, AttaqueBarryTete}, 10, 10, []Objects{ScrapMetal}}
var AdverseryWanda = Adversery{"Sorière Wanda", 60, 20, 40, 60, []string{}, 50, 70, []string{"Blockage spatio-temporelle", "Mind contrôle", "ilusion"}, []Attaques{AttaqueBloquage, AttaqueContrôle, AttaqueIlusion}, 50, 50, []Objects{InfinityStoneMind}}