package tables

//	=========================================

type DataTables struct {
	Pokemon      []Pokemon
	LearnedMoves []LearnedMove
	Pokedex      []Pokedex
	Abilities    []Ability
	AbltyDescs   []AbltyDesc
	Moves        []Move
	MoveDescs    []MoveDesc
}

//	=========================================

type Pokemon struct {
	Id      int `gorm:"primaryKey;autoIncrement:false"`
	BaseId  int
	Order   int
	Search  string `gorm:"index"`
	Name    string
	Moniker string
	Ablty1  string
	Ablty2  *string
	AbltyH  *string
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
	Hshiny  string
}

func (Pokemon) TableName() string {
	return "pkmn"
}

//	=========================================

type LearnedMove struct {
	PkmnId  int    `gorm:"primaryKey;autoIncrement:false"`
	Move    string `gorm:"primaryKey;index"`
	Method  string `gorm:"primaryKey"`
	Level   int    `gorm:"primaryKey"`
	Tm      *string
	Version string `gorm:"primaryKey"`
}

func (LearnedMove) TableName() string {
	return "lmoves"
}

//	=========================================

type Pokedex struct {
	PkmnId  int    `gorm:"primaryKey;autoIncrement:false"`
	Search  string `gorm:"primaryKey"`
	Entry   string `gorm:"type:text"`
	Version string `gorm:"primaryKey"`
}

func (Pokedex) TableName() string {
	return "pdex"
}

//	=========================================

type Ability struct {
	Id      int    `gorm:"primaryKey;autoIncrement:false"`
	Search  string `gorm:"index"`
	Name    string
	Effect  string `gorm:"type:text"`
	Summary string `gorm:"type:text"`
}

func (Ability) TableName() string {
	return "abltys"
}

//	=========================================

type AbltyDesc struct {
	AbltyId int    `gorm:"primaryKey;autoIncrement:false"`
	Desc    string `gorm:"type:text"`
	Version string `gorm:"primaryKey"`
}

func (AbltyDesc) TableName() string {
	return "adescs"
}

//	=========================================

type Move struct {
	Id       int    `gorm:"primaryKey;autoIncrement:false"`
	Search   string `gorm:"index"`
	Name     string
	Effect   string `gorm:"type:text"`
	Summary  string `gorm:"type:text"`
	Power    *int
	Acc      *int
	Pp       int
	Priority int
	Class    string
	Target   string
	Type     string
}

func (Move) TableName() string {
	return "moves"
}

//	=========================================

type MoveDesc struct {
	MoveId  int    `gorm:"primaryKey;autoIncrement:false"`
	Desc    string `gorm:"type:text"`
	Version string `gorm:"primaryKey"`
}

func (MoveDesc) TableName() string {
	return "mdescs"
}
