package personnage

type Kryptonien struct {
	Viemax          int
	Inventaire      int
	Puissance       float64
	Agilite         int
	Intelligence    int
	CapacityDisplay []string
	AttaquesSpé     []Attaques
	Capacity        []Capacity
}

var Kryptonien1 = Kryptonien{20, 10, 25, 25, 20, []string{"Laser 🔴"}, []Attaques{AttaqueLasser}, []Capacity{}}

type Batfamily struct {
	Viemax          int
	Inventaire      int
	Puissance       float64
	Agilite         int
	Intelligence    int
	CapacityDisplay []string
}

var BatFamily1 = Batfamily{10, 15, 10, 35, 50, []string{"Argent", "*2EXP 📈"}}

type Hulk struct {
	Viemax          int
	Inventaire      int
	Puissance       float64
	Agilite         int
	Intelligence    int
	CapacityDisplay []string
	AttaquesSpé     []Attaques
}

var Hulk1 = Hulk{30, 4, 50, 15, 10, []string{"Clape Sonic 🔊"}, []Attaques{ClapSuperSonic}}

type Kantin struct {
	Viemax          int
	Inventaire      int
	Puissance       float64
	Agilite         int
	Intelligence    int
	CapacityDisplay []string
}

var Kantin1 = Kantin{120, 60, 100, 70, 100, []string{"Toutes 🧬", "HO 🧠", "Prout Bruisateurr!! 💨"}}

type Wakandais struct {
	Viemax          int
	Inventaire      int
	Puissance       float64
	Agilite         int
	Intelligence    int
	CapacityDisplay []string
	AttaquesSpé     []Attaques
}

var Wakanda1 = Wakandais{15, 13, 20, 30, 40, []string{"Assimilation des coups pour renvoi synthétique 🛡️🤖"}, []Attaques{RenvoieSinetique}}

type SpiderMan struct {
	Viemax 			int
	Inventaire      int
	Puissance       float64
	Agilite         int
	Intelligence    int
	CapacityDisplay []string
	AttaquesSpé     []Attaques
}

var SpiderMan1 = SpiderMan{20, 10, 15, 40, 35, []string{"Entoiler"}, []Attaques{}}