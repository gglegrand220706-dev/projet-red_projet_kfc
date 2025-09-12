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
}

var Adversery01 = Adversery{"entrîneur : kentino", 10, 10, 10, 50, []string{}, 50, 50, []string{"coup de point bousté, proute brumisateur, lancé de casquet"}, []Attaques{AttaqueBasique01, AttaqueBasique02, AttaqueBasique03, AttaqueBasique04}, 20}