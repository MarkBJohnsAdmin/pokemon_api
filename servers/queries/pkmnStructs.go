package queries

import (
	"fmt"
)

//	===============================================================================================

//	POKEMON STRUCT

//	====================================================

// A full Pokemon object with all of the data from the "pkmn" table. The pkmnAblty structs are fetched from the ablty1, ablty2,
//	and ablty_h values inserted into the "abltys" table.

type Pokemon struct {
	Id         int           `json:"id"`
	BaseId     int           `json:"base_id"`
	Order      int           `json:"order"`
	Search     string        `json:"search"`
	Name       string        `json:"name"`
	Moniker    string        `json:"moniker"`
	Types      pkmnTypes     `json:"types"`
	Abilities  pkmnAbilities `json:"abilities"`
	Stats      pkmnStats     `json:"stats"`
	Images     pkmnImgs      `json:"images"`
	Attributes pkmnAttr      `json:"attributes"`
}

// -----------------------------------------------------------------------------------------------
type pkmnTypes struct {
	Primary   string  `json:"primary"`
	Secondary *string `json:"secondary"`
}

// -----------------------------------------------------------------------------------------------
type pkmnAbilities struct {
	Ablty1 pkmnAblty  `json:"primary"`
	Ablty2 *pkmnAblty `json:"secondary"`
	AbltyH *pkmnAblty `json:"hidden"`
}

// ----------------------------------------------------
type pkmnAblty struct {
	Search  string `json:"search"`
	Name    string `json:"name"`
	Summary string `json:"summary"`
}

// -----------------------------------------------------------------------------------------------
type pkmnStats struct {
	Hp  int `json:"hp"`
	Atk int `json:"attack"`
	Def int `json:"defense"`
	Spa int `json:"sp_atk"`
	Spd int `json:"sp_def"`
	Spe int `json:"speed"`
}

// -----------------------------------------------------------------------------------------------
type pkmnImgs struct {
	Sprite string `json:"sprite"`
	Offcl  _imgs  `json:"official"`
	Home   _imgs  `json:"home"`
}

// ----------------------------------------------------
type _imgs struct {
	Default string `json:"default"`
	Shiny   string `json:"shiny"`
}

// -----------------------------------------------------------------------------------------------
type pkmnAttr struct {
	Height string `json:"height"`
	Weight string `json:"weight"`
	Color  string `json:"color"`
}

//	===============================================================================================

type PokemonRow struct {
	Id     int    `json:"id"`
	Search string `json:"search"`
	Name   string `json:"name"`
	Sprite string `json:"sprite"`
	Types  struct {
		Type1 string  `json:"primary"`
		Type2 *string `json:"secondary"`
	} `json:"types"`
}

//	===============================================================================================

// Temporary flat struct to capture all the data for the database queries

type TempPokemon struct {
	Search  string  `gorm:"column:search"`
	Name    string  `gorm:"column:name"`
	Moniker string  `gorm:"column:moniker"`
	Type1   string  `gorm:"column:type1"`
	Ablty1  string  `gorm:"column:ablty1"`
	A1Name  string  `gorm:"column:a1Name"`
	A1Sum   string  `gorm:"column:a1Sum"`
	Sprite  string  `gorm:"column:sprite"`
	Offcl   string  `gorm:"column:offcl"`
	Shiny   string  `gorm:"column:shiny"`
	Home    string  `gorm:"column:home"`
	Hshiny  string  `gorm:"column:hshiny"`
	Color   string  `gorm:"column:color"`
	Egg1    string  `gorm:"column:egg1"`
	Id      int     `gorm:"column:id"`
	BaseId  int     `gorm:"column:base_id"`
	Order   int     `gorm:"column:order"`
	Hp      int     `gorm:"column:hp"`
	Atk     int     `gorm:"column:atk"`
	Def     int     `gorm:"column:def"`
	Spa     int     `gorm:"column:spa"`
	Spd     int     `gorm:"column:spd"`
	Spe     int     `gorm:"column:spe"`
	Height  int     `gorm:"column:height"`
	Weight  int     `gorm:"column:weight"`
	Type2   *string `gorm:"column:type2"`
	Ablty2  *string `gorm:"column:ablty2"`
	AbltyH  *string `gorm:"column:ablty_h"`
	Egg2    *string `gorm:"column:egg2"`
	A2Name  *string `gorm:"column:a2Name"`
	A2Sum   *string `gorm:"column:a2Sum"`
	AhName  *string `gorm:"column:ahName"`
	AhSum   *string `gorm:"column:ahSum"`
}

// ----------------------------------------------------
type TempPkmnRow struct {
	Id     int     `gorm:"column:id"`
	Search string  `gorm:"column:search"`
	Name   string  `gorm:"column:name"`
	Sprite string  `gorm:"column:sprite"`
	Type1  string  `gorm:"column:type1"`
	Type2  *string `gorm:"column:type2"`
}

//	===============================================================================================

// Convert the flat TempPokemon struct into a more structured JSON object with a built in Format() method

func (t TempPokemon) Format() Pokemon {
	var scndary, hidden *pkmnAblty

	if t.Ablty2 != nil && t.A2Name != nil && t.A2Sum != nil {
		scndary = &pkmnAblty{
			Search: *t.Ablty2, Name: *t.A2Name, Summary: *t.A2Sum,
		}
	}

	if t.AbltyH != nil && t.AhName != nil && t.AhSum != nil {
		hidden = &pkmnAblty{
			Search: *t.AbltyH, Name: *t.AhName, Summary: *t.AhSum,
		}
	}

	return Pokemon{
		Id: t.Id, BaseId: t.BaseId, Order: t.Order,
		Search: t.Search, Name: t.Name, Moniker: t.Moniker,
		Types: pkmnTypes{
			Primary: t.Type1, Secondary: t.Type2,
		},
		Abilities: pkmnAbilities{
			Ablty1: pkmnAblty{
				Search: t.Ablty1, Name: t.A1Name, Summary: t.A1Sum,
			},
			Ablty2: scndary,
			AbltyH: hidden,
		},
		Stats: pkmnStats{
			Hp: t.Hp, Atk: t.Atk, Def: t.Def,
			Spa: t.Spa, Spd: t.Spd, Spe: t.Spe,
		},
		Images: pkmnImgs{
			Sprite: t.Sprite,
			Offcl:  _imgs{Default: t.Offcl, Shiny: t.Shiny},
			Home:   _imgs{Default: t.Home, Shiny: t.Hshiny},
		},
		Attributes: pkmnAttr{
			Height: fmt.Sprintf("%v m", float64(t.Height)/10),
			Weight: fmt.Sprintf("%v kg", float64(t.Weight)/10),
			Color:  t.Color,
		},
	}
}

// ----------------------------------------------------
func (t TempPkmnRow) Format() PokemonRow {
	return PokemonRow{
		Id:     t.Id,
		Search: t.Search,
		Name:   t.Name,
		Sprite: t.Sprite,
		Types: struct {
			Type1 string  `json:"primary"`
			Type2 *string `json:"secondary"`
		}{
			Type1: t.Type1,
			Type2: t.Type2,
		},
	}
}

// ----------------------------------------------------
func ConvertPokemon(temp []TempPokemon) []Pokemon {
	result := make([]Pokemon, len(temp))

	for i, pkmn := range temp {
		result[i] = pkmn.Format()
	}

	return result
}

// -----------------------------------------------------------------------------------------------
func ConvertPkmnRow(temp []TempPkmnRow) []PokemonRow {
	result := make([]PokemonRow, len(temp))

	for i, pkmn := range temp {
		result[i] = pkmn.Format()
	}

	return result
}
