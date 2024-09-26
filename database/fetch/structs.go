package fetch

//	ALL RESPONSE SLICES

type ApiData struct {
	Pokemon   []Pokemon
	Species   []PokemonSpecies
	Abilities []Ability
	Moves     []Move
	Machines  []Machine
}

type Pokemon struct {
	Abilities []struct {
		Ability struct {
			Name string `json:"name"`
		} `json:"ability"`
		IsHidden bool `json:"is_hidden"`
		Slot     int  `json:"slot"`
	} `json:"abilities"`
	Height int `json:"height"`
	Id     int `json:"id"`
	Moves  []struct {
		Move struct {
			Name string `json:"name"`
		} `json:"move"`
		VersionGroupDetails []struct {
			LevelLearnedAt  int `json:"level_learned_at"`
			MoveLearnMethod struct {
				Name string `json:"name"`
			} `json:"move_learn_method"`
			VersionGroup struct {
				Name string `json:"name"`
			} `json:"version_group"`
		} `json:"version_group_details"`
	} `json:"moves"`
	Name    string `json:"name"`
	Species struct {
		Url string `json:"url"`
	} `json:"species"`
	Sprites struct {
		FrontDefault string `json:"front_default"`
		Other        struct {
			Home struct {
				FrontDefault string `json:"front_default"`
				FrontShiny   string `json:"front_shiny"`
			} `json:"home"`
			OfficialArtwork struct {
				FrontDefault string `json:"front_default"`
				FrontShiny   string `json:"front_shiny"`
			} `json:"official-artwork"`
		} `json:"other"`
	} `json:"sprites"`
	Stats []struct {
		BaseStat int `json:"base_stat"`
		Effort   int `json:"effort"`
		Stat     struct {
			Name string `json:"name"`
		} `json:"stat"`
	} `json:"stats"`
	Types []struct {
		Slot int `json:"slot"`
		Type struct {
			Name string `json:"name"`
		} `json:"type"`
	} `json:"types"`
	Weight int `json:"weight"`
}

//	https://pokeapi.co/api/v2/pokemon-species/

type PokemonSpecies struct {
	Color struct {
		Name string `json:"name"`
	} `json:"color"`
	EggGroups []struct {
		Name string `json:"name"`
	} `json:"egg_groups"`
	FlavorTextEntries []struct {
		FlavorText string `json:"flavor_text"`
		Language   struct {
			Name string `json:"name"`
		} `json:"language"`
		Version struct {
			Name string `json:"name"`
		} `json:"version"`
	} `json:"flavor_text_entries"`
	Genera []struct {
		Genus    string `json:"genus"`
		Language struct {
			Name string `json:"name"`
		} `json:"language"`
	} `json:"genera"`
	Id    int    `json:"id"`
	Name  string `json:"name"`
	Names []struct {
		Language struct {
			Name string `json:"name"`
		} `json:"language"`
		Name string `json:"name"`
	} `json:"names"`
}

//	https://pokeapi.co/api/v2/ability/

type Ability struct {
	EffectEntries *[]struct {
		Effect   string `json:"effect"`
		Language struct {
			Name string `json:"name"`
		} `json:"language"`
		ShortEffect string `json:"short_effect"`
	} `json:"effect_entries"`
	FlavorTextEntries []struct {
		FlavorText string `json:"flavor_text"`
		Language   struct {
			Name string `json:"name"`
		} `json:"language"`
		VersionGroup struct {
			Name string `json:"name"`
		} `json:"version_group"`
	} `json:"flavor_text_entries"`
	Id    int    `json:"id"`
	Name  string `json:"name"`
	Names []struct {
		Language struct {
			Name string `json:"name"`
		} `json:"language"`
		Name string `json:"name"`
	} `json:"names"`
}

//	https://pokeapi.co/api/v2/move/

type Move struct {
	Accuracy    *int `json:"accuracy"`
	DamageClass struct {
		Name string `json:"name"`
	} `json:"damage_class"`
	EffectEntries *[]struct {
		Effect   string `json:"effect"`
		Language struct {
			Name string `json:"name"`
		} `json:"language"`
		ShortEffect string `json:"short_effect"`
	} `json:"effect_entries"`
	FlavorTextEntries []struct {
		FlavorText string `json:"flavor_text"`
		Language   struct {
			Name string `json:"name"`
		} `json:"language"`
		VersionGroup struct {
			Name string `json:"name"`
		} `json:"version_group"`
	} `json:"flavor_text_entries"`
	Id    int    `json:"id"`
	Name  string `json:"name"`
	Names []struct {
		Language struct {
			Name string `json:"name"`
		} `json:"language"`
		Name string `json:"name"`
	} `json:"names"`
	Power    *int `json:"power"`
	Pp       int  `json:"pp"`
	Priority int  `json:"priority"`
	Target   struct {
		Name string `json:"name"`
	} `json:"target"`
	Type struct {
		Name string `json:"name"`
	} `json:"type"`
}

//	https://pokeapi.co/api/machine/

type Machine struct {
	Id   int `json:"id"`
	Item struct {
		Name string `json:"name"`
	} `json:"item"`
	Move struct {
		Name string `json:"name"`
	} `json:"move"`
	VersionGroup struct {
		Name string `json:"name"`
	} `json:"version_group"`
}
