package queries

//	===============================================================================================

//	POKEMON STRUCT

//	====================================================

// Each Pokedex entry for a specified Pokemon

type PokedexByPokemon struct {
	Id      int           `json:"id"`
	Search  string        `json:"search"`
	Name    string        `json:"name"`
	Images  pkmnImgs      `json:"images"`
	Pokedex []pdexVersion `json:"pokedex"`
}

// -----------------------------------------------------------------------------------------------
type pdexVersion struct {
	Entry   string `json:"entry"`
	Version string `json:"version"`
}

// -----------------------------------------------------------------------------------------------
type PokedexByVersion struct {
	Version string     `json:"version"`
	Pokemon []pdexPkmn `json:"pokemon"`
}

// -----------------------------------------------------------------------------------------------
type pdexPkmn struct {
	Id     int    `json:"id"`
	Search string `json:"search"`
	Name   string `json:"name"`
	Sprite string `json:"sprite"`
	Entry  string `json:"entry"`
}

//	===============================================================================================

// Temporary flat struct to capture all the data for the database queries

type TempPdexPkmn struct {
	Id     int    `gorm:"column:id"`
	Search string `gorm:"column:search"`
	Name   string `gorm:"column:name"`
	Sprite string `gorm:"column:sprite"`
	Offcl  string `gorm:"column:offcl"`
	Shiny  string `gorm:"column:shiny"`
	Home   string `gorm:"column:home"`
	Hshiny string `gorm:"column:hshiny"`
}

// -----------------------------------------------------------------------------------------------

//	===============================================================================================

// Convert the flat TempPokemon struct into a more structured JSON object with a built in Format() method

func (t TempPdexPkmn) Join(dex []pdexVersion) PokedexByPokemon {
	return PokedexByPokemon{
		Id:     t.Id,
		Search: t.Search,
		Name:   t.Name,
		Images: pkmnImgs{
			Sprite: t.Sprite,
			Offcl: _imgs{
				Default: t.Offcl, Shiny: t.Shiny,
			},
			Home: _imgs{
				Default: t.Home, Shiny: t.Hshiny,
			},
		},
		Pokedex: dex,
	}
}

// -----------------------------------------------------------------------------------------------

func ConvertPdexVersion(version string, pkmnData []pdexPkmn) PokedexByVersion {
	return PokedexByVersion{
		Version: version,
		Pokemon: pkmnData,
	}
}
