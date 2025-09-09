package personnage

import "fmt"

	type Kryptonien struct {

	Viemax int = 20
	Inventaire int  = 10 
	Puissance int = 25
    Faiblaisse string ["Crypto", "magie", "puissance plus grande"]
	Agilite int = 25 
	Intelligence int = 20
	Capacity string ["Vole", "Laser"]
}

	type Batfamily  struct {

	Viemax int = 10
	Inventaire int  = 15
    Puissance int = 10
	Agilite int = 35 
	Intelligence int = 50
	Capacity string ["Adaptation", "Gadjet", "*2EXP"]
}

	type Hulk struct {

	Viemax int = 30
	Inventaire int  = 4
    Puissance int = 35
	Faiblaisse string ["Hypnoses", "Self Controle"]
	Agilite int = 15
	Intelligence int = 10 
	Capacity string ["Clape Sonic", "Resistance"]
}

	type Kantin struct {

	Viemax int = 120
	Inventaire int  = 60
    Puissance int = 100
	Agilite int = 70
	Intelligence int = 100
	Capacity string ["Toutes", "HO", "Prout Bruisateurr!!"]
}

	type Wakandais struct {

	Viemax int = 15
	Inventaire int  = 13
	Puissance int = 20
	Faiblaisse string ["emp"]
	Agilite int = 30
	Intelligence int = 40
	Capacity string ["Assimilation des coups pour renvoie synetique"] 
}