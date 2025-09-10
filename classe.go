package personnage

type Kryptonien struct {
	Viemax       int
	Inventaire   int
	Puissance    int
	Faiblaisse   []string
	Agilite      int
	Intelligence int
	Capacity     []string
}

var Kryptonien1 = Kryptonien{20, 10, 25, []string{"Crypto,", "magie,", "puissance plus grande"}, 25, 20, []string{"Vole", "Laser"}}

type Batfamily struct {
	Viemax       int
	Inventaire   int
	Puissance    int
	Agilite      int
	Intelligence int
	Capacity     []string
}

var BatFamily1 = Batfamily{10, 15, 10, 35, 50, []string{"Adaptation", "Gadjet", "*2EXP"}}

type Hulk struct {
	Viemax       int
	Inventaire   int
	Puissance    int
	Faiblaisse   []string
	Agilite      int
	Intelligence int
	Capacity     []string
}

var Hulk1 = Hulk{30, 4, 35, []string{"Hypnoses", "Self Controle"}, 15, 10, []string{"Clape Sonic", "Resistance"}}

type Kantin struct {
	Viemax       int
	Inventaire   int
	Puissance    int
	Agilite      int
	Intelligence int
	Capacity     []string
}

var Kantin1 = Kantin{120, 60, 100, 70, 100, []string{"Toutes", "HO", "Prout Bruisateurr!!"}}

type Wakandais struct {
	Viemax       int
	Inventaire   int
	Puissance    int
	Faiblaisse   []string
	Agilite      int
	Intelligence int
	Capacity     []string
}

var Wakanda1 = Wakandais{15, 13, 20, []string{"emp"}, 30, 40, []string{"Assimilation des coups pour renvoie synetique"}}
