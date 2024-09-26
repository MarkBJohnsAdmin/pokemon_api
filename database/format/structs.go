package frmt

type FormattedData struct {
	Pokemon    []Pokemon
	LrndMoves  []LearnedMove
	Pokedex    []Pokedex
	Abilities  []Ability
	AbltyDescs []AbltyDesc
	Moves      []Move
	MoveDescs  []MoveDesc
	Machines   []Machine
}

//	=========================================
//	FROM https://pokeapi.co/api/v2/pokemon/
//		AND https://pokeapi.co/api/v2/pokemon-species

type Pokemon struct {
	Id      int
	BaseId  int
	Order   int
	Search  string
	Name    string
	Moniker string
	Ablty1  string
	Ablty2  *string
	Hidden  *string
	Type1   string
	Type2   *string
	Egg1    string
	Egg2    *string
	Color   string
	Height  int
	Weight  int
	Hp      int
	Atk     int
	Def     int
	Spa     int
	Spd     int
	Spe     int
	Sprite  string
	Offcl   string
	Shiny   string
	Home    string
	Hshny   string
}

//	-----------------------------------------

type LearnedMove struct {
	PkmnId  int
	Move    string
	Level   int
	Method  string
	Version string
}

//	-----------------------------------------

type Pokedex struct {
	PkmnId  int
	Search  string
	Entry   string
	Version string
}

//	=========================================
//	FROM https://pokeapi.co/api/v2/ability/

type Ability struct {
	Id      int
	Search  string
	Name    string
	Effect  string
	Summary string
}

//	-----------------------------------------

type AbltyDesc struct {
	AbltyId int
	Desc    string
	Version string
}

//	=========================================
//	FROM https://pokeapi.co/api/v2/move/

type Move struct {
	Id       int
	Search   string
	Name     string
	Effect   string
	Summary  string
	Acc      *int
	Power    *int
	Pp       int
	Priority int
	Class    string
	Target   string
	Type     string
}

//	-----------------------------------------

type MoveDesc struct {
	MoveId  int
	Desc    string
	Version string
}

//	=========================================
//	FROM https://pokeapi.co/api/v2/machine/

type Machine struct {
	Id      int
	TmNo    string
	Move    string
	Version string
}

//	=========================================
//	Temporary Structs

type TempPkmn struct {
	Id     int
	Search string
	Order  int
	Ablty1 string
	Ablty2 *string
	Hidden *string
	Type1  string
	Type2  *string
	Height int
	Weight int
	BaseId int
	Hp     int
	Atk    int
	Def    int
	Spa    int
	Spd    int
	Spe    int
	Sprite string
	Offcl  string
	Shiny  string
	Home   string
	Hshny  string
}

//	-----------------------------------------

type TempSpcs struct {
	Id      int
	Search  string
	Name    string
	Moniker string
	Egg1    string
	Egg2    *string
	Color   string
}
