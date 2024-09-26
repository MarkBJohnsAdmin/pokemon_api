package queries

//	===============================================================================================

//	MOVE STRUCT

//	====================================================

// A full Move object with all of the data from the "moves" table.

type Move struct {
	Id         int       `json:"id"`
	Search     string    `json:"search"`
	Name       string    `json:"name"`
	Type       string    `json:"type"`
	FlavorText flvrText  `json:"flavor_text"`
	Meta       moveMeta  `json:"meta"`
	Stats      moveStats `json:"stats"`
}

// -----------------------------------------------------------------------------------------------
type moveMeta struct {
	Class  string `json:"damage_class"`
	Target string `json:"target"`
	// More keys to be added later
}

// -----------------------------------------------------------------------------------------------
type moveStats struct {
	Power    *int `json:"power"`
	Acc      *int `json:"accuracy"`
	Pp       int  `json:"pp"`
	Priority int  `json:"priority"`
}

//	===============================================================================================

// Temporary flat struct to capture all the data for the database queries

type TempMove struct {
	Id       int    `gorm:"column:id"`
	Search   string `gorm:"column:search"`
	Name     string `gorm:"column:name"`
	Type     string `gorm:"column:type"`
	Effect   string `gorm:"column:effect"`
	Summary  string `gorm:"column:summary"`
	Class    string `gorm:"column:class"`
	Target   string `gorm:"column:target"`
	Power    *int   `gorm:"column:power"`
	Acc      *int   `gorm:"column:acc"`
	Pp       int    `gorm:"column:pp"`
	Priority int    `gorm:"column:priority"`
}

//	===============================================================================================

// Convert the flat TempPokemon struct into a more structured JSON object with a built in Format() method

func (t TempMove) Format() Move {
	return Move{
		Id:     t.Id,
		Search: t.Search,
		Name:   t.Name,
		Type:   t.Type,
		FlavorText: flvrText{
			Effect: t.Effect, Summary: t.Summary,
		},
		Meta: moveMeta{
			Class: t.Class, Target: t.Target,
		},
		Stats: moveStats{
			Power: t.Power, Acc: t.Acc,
			Pp: t.Pp, Priority: t.Priority,
		},
	}
}

// ----------------------------------------------------
func ConvertMoves(temp []TempMove) []Move {
	result := make([]Move, len(temp))

	for i, move := range temp {
		result[i] = move.Format()
	}

	return result
}
