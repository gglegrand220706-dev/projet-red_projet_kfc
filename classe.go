package personnage

type Kryptonien struct {
	Viemax       int
	Inventaire   int
	Puissance    float64
	Faiblaisse   []string
	Agilite      int
	Intelligence int
	CapacityDisplay     []string
	AttaquesSpé			[]Attaques
	Capacity			[]Capacity

}

var Kryptonien1 = Kryptonien{20, 10, 25, []string{"Crypto 🪙", "Magie ✨", "Puissance plus grande 💥"}, 25, 20, []string{"Laser 🔴"}, []Attaques{AttaqueLasser}, []Capacity{}}

type Batfamily struct {
	Viemax       int
	Inventaire   int
	Puissance    float64
	Agilite      int
	Intelligence int
	CapacityDisplay     []string
}

var BatFamily1 = Batfamily{10, 15, 10, 35, 50, []string{"Adaptation 🧠", "Gadget 🛠️", "*2EXP 📈"}}

type Hulk struct {
	Viemax       int
	Inventaire   int
	Puissance    float64
	Faiblaisse   []string
	Agilite      int
	Intelligence int
	CapacityDisplay     []string
}

var Hulk1 = Hulk{30, 4, 50, []string{"Hypnose 🌀", "Self-contrôle 🧘"}, 15, 10, []string{"Clape Sonic 🔊", "Résistance 🧱"}}

type Kantin struct {
	Viemax       int
	Inventaire   int
	Puissance    float64
	Agilite      int
	Intelligence int
	CapacityDisplay     []string
}

var Kantin1 = Kantin{120, 60, 100, 70, 100, []string{"Toutes 🧬", "HO 🧠", "Prout Bruisateurr!! 💨"}}

type Wakandais struct {
	Viemax       		int
	Inventaire   		int
	Puissance    		float64
	Faiblaisse   		[]string
	Agilite      		int
	Intelligence 		int
	CapacityDisplay     []string
}

var Wakanda1 = Wakandais{15, 13, 20, []string{"EMP ⚡"}, 30, 40, []string{"Assimilation des coups pour renvoi synthétique 🛡️🤖"}}
