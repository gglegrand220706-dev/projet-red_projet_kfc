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
}

var Adversery01 = Adversery{"entrîneur : kentino", 30, 1, 30, 50, []string{}, 50, 50, []string{"coup de point bousté, proute brumisateur, lancé de casquet"}, []Attaques{AttaqueBasique01, AttaqueBasique02, AttaqueBasique03, AttaqueBasique04}, 20, 20}
var AdverseryLoki = Adversery{"boose : Loki", 40, 5, 40, 50, []string{}, 60, 65, []string{"coup de sptre", "tourment", "boule de magie"}, []Attaques{AttaqueSeptre, AttaqueMentale, AttaqueMagique}, 30, 30 }