package queries

//	===============================================================================================

//	ABILITY STRUCT

//	====================================================

// Placeholder text to update when I figure out exactly how I want to handle this

type Ability struct {
	Id         int      `json:"id"`
	Search     string   `json:"search"`
	Name       string   `json:"name"`
	FlavorText flvrText `json:"flavor_text"`
}

// -----------------------------------------------------------------------------------------------
type flvrText struct {
	Effect  string `json:"effect"`
	Summary string `json:"summary"`
}

//	===============================================================================================

//	Temporary flat struct to capture all the data fro the database queries

type TempAbility struct {
	Id      int    `gorm:"column:id"`
	Search  string `gorm:"column:search"`
	Name    string `gorm:"column:name"`
	Effect  string `gorm:"column:effect"`
	Summary string `gorm:"column:summary"`
}

//	===============================================================================================

// Convert the flat TempPokemon struct into a more structured JSON object with a built in Format() method

func (t TempAbility) Format() Ability {
	return Ability{
		Id:     t.Id,
		Search: t.Search,
		Name:   t.Name,
		FlavorText: flvrText{
			Effect: t.Effect, Summary: t.Summary,
		},
	}
}

// ----------------------------------------------------
func ConvertAbilities(temp []TempAbility) []Ability {
	result := make([]Ability, len(temp))

	for i, ablty := range temp {
		result[i] = ablty.Format()
	}

	return result
}
