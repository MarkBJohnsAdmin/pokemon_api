package frmt

import (
	"fmt"
	"regexp"
	"server/database/fetch"
	"strconv"
)

//	===================================================================================================

// MakePokemonStruct() => combine a TempPkmn and TempSpcs struct into a Pokemon struct

//	-	MakeTempPkmn() => convert the data from PokeApi /pokemon into TempPkmn

//	-	MakeTempSpcs() => convert the data from PokeApi /pokemon-species into TempSpcs

//	---------------------------------------------------------------------------------------------------

// MakeLearnedMoveSlice() => convert the Moves key from /pokemon into a slice of LearnedMove

//	---------------------------------------------------------------------------------------------------

// MakePokedexSlice() => convert the FlavorTextEntries key from /pokemon-species into a slice of Pokedex

//	===================================================================================================

func MakePokemonStruct(pkmn TempPkmn, spcs TempSpcs) Pokemon {
	var name string

	if pkmn.Search == spcs.Search {
		name = spcs.Name
	} else {
		name = PokemonNameMap(pkmn.Search, spcs.Name)
	}

	return Pokemon{
		Id:      pkmn.Id,
		BaseId:  pkmn.BaseId,
		Order:   pkmn.Order,
		Search:  pkmn.Search,
		Name:    name,
		Moniker: spcs.Moniker,
		Ablty1:  pkmn.Ablty1,
		Ablty2:  pkmn.Ablty2,
		Hidden:  pkmn.Hidden,
		Type1:   pkmn.Type1,
		Type2:   pkmn.Type2,
		Egg1:    spcs.Egg1,
		Egg2:    spcs.Egg2,
		Color:   spcs.Color,
		Height:  pkmn.Height,
		Weight:  pkmn.Weight,
		Hp:      pkmn.Hp,
		Atk:     pkmn.Atk,
		Def:     pkmn.Def,
		Spa:     pkmn.Spa,
		Spd:     pkmn.Spd,
		Spe:     pkmn.Spe,
		Sprite:  pkmn.Sprite,
		Offcl:   pkmn.Offcl,
		Shiny:   pkmn.Shiny,
		Home:    pkmn.Home,
		Hshny:   pkmn.Hshny,
	}
}

//	---------------------------------------------------------------------------------------------------

func MakeTempPkmn(pkmnData fetch.Pokemon, order int) TempPkmn {
	a1, a2, hidden := getPkmnAbilities(pkmnData)
	t1, t2 := getPkmnTypes(pkmnData)
	baseId := getPkmnBaseId(pkmnData)
	stats := getPkmnStats(pkmnData)
	imgs := getPkmnImages(pkmnData)

	return TempPkmn{
		Id:     pkmnData.Id,
		Search: pkmnData.Name,
		Order:  order,
		Ablty1: a1,
		Ablty2: a2,
		Hidden: hidden,
		Type1:  t1,
		Type2:  t2,
		Height: pkmnData.Height,
		Weight: pkmnData.Weight,
		BaseId: baseId,
		Hp:     stats.hp,
		Atk:    stats.atk,
		Def:    stats.def,
		Spa:    stats.spa,
		Spd:    stats.spd,
		Spe:    stats.spe,
		Sprite: imgs.sprite,
		Offcl:  imgs.offcl,
		Shiny:  imgs.shiny,
		Home:   imgs.home,
		Hshny:  imgs.hshiny,
	}
}

//	---------------------------------------------------------

func getPkmnAbilities(pkmnData fetch.Pokemon) (ability1 string, ability2, hidden *string) {
	var a1 string
	var a2, hdn *string

	if len(pkmnData.Abilities) != 0 {
		for _, a := range pkmnData.Abilities {
			if a.Slot == 1 {
				a1 = a.Ability.Name
			} else if a.Slot == 2 && !a.IsHidden {
				a2 = &a.Ability.Name
			} else if a.IsHidden {
				hdn = &a.Ability.Name
			}
		}
	}
	return a1, a2, hdn
}

//	---------------------------------------------------------

func getPkmnTypes(pkmnData fetch.Pokemon) (type1 string, type2 *string) {
	if len(pkmnData.Types) == 2 {
		return pkmnData.Types[0].Type.Name,
			&pkmnData.Types[1].Type.Name
	}

	if len(pkmnData.Types) == 1 {
		return pkmnData.Types[0].Type.Name, nil
	}

	return "", nil
}

//	---------------------------------------------------------

func getPkmnBaseId(pkmnData fetch.Pokemon) int {
	re := regexp.MustCompile(`/pokemon-species/(\d+)/`)
	matches := re.FindStringSubmatch(pkmnData.Species.Url)

	if len(matches) < 2 {
		return pkmnData.Id
	}

	id, err := strconv.Atoi(matches[1])
	if err != nil {
		return pkmnData.Id
	}

	return id
}

//	---------------------------------------------------------

func getPkmnStats(pkmnData fetch.Pokemon) struct{ hp, atk, def, spa, spd, spe int } {
	var stats struct{ hp, atk, def, spa, spd, spe int }

	if len(pkmnData.Stats) != 0 {
		for _, s := range pkmnData.Stats {
			switch s.Stat.Name {
			case "hp":
				stats.hp = s.BaseStat
			case "attack":
				stats.atk = s.BaseStat
			case "defense":
				stats.def = s.BaseStat
			case "special-attack":
				stats.spa = s.BaseStat
			case "special-defense":
				stats.spd = s.BaseStat
			case "speed":
				stats.spe = s.BaseStat
			}
		}
		return stats
	}
	return stats
}

//	---------------------------------------------------------

func getPkmnImages(pkmnData fetch.Pokemon) struct{ sprite, offcl, shiny, home, hshiny string } {
	return struct{ sprite, offcl, shiny, home, hshiny string }{
		sprite: pkmnData.Sprites.FrontDefault,
		offcl:  pkmnData.Sprites.Other.OfficialArtwork.FrontDefault,
		shiny:  pkmnData.Sprites.Other.OfficialArtwork.FrontShiny,
		home:   pkmnData.Sprites.Other.Home.FrontDefault,
		hshiny: pkmnData.Sprites.Other.Home.FrontShiny,
	}
}

//	---------------------------------------------------------------------------------------------------

func MakeTempSpcs(spcsData fetch.PokemonSpecies) TempSpcs {
	name := getSpcsName(spcsData)
	moniker := getSpcsMoniker(spcsData)
	e1, e2 := getSpcsEggGroups(spcsData)

	return TempSpcs{
		Id:      spcsData.Id,
		Search:  spcsData.Name,
		Name:    name,
		Moniker: moniker,
		Egg1:    e1,
		Egg2:    e2,
		Color:   spcsData.Color.Name,
	}
}

//	---------------------------------------------------------

func getSpcsName(spcsData fetch.PokemonSpecies) string {
	if len(spcsData.Names) != 0 {
		for _, n := range spcsData.Names {
			if n.Language.Name == "en" {
				return n.Name
			}
		}
	}
	return spcsData.Name
}

//	---------------------------------------------------------

func getSpcsMoniker(spcsData fetch.PokemonSpecies) string {
	if len(spcsData.Genera) != 0 {
		for _, g := range spcsData.Genera {
			if g.Language.Name == "en" {
				return fmt.Sprintf("The %s", g.Genus)
			}
		}
	}
	return ""
}

//	---------------------------------------------------------

func getSpcsEggGroups(spcsData fetch.PokemonSpecies) (egg1 string, egg2 *string) {
	if len(spcsData.EggGroups) == 2 {
		return spcsData.EggGroups[0].Name,
			&spcsData.EggGroups[1].Name
	}

	if len(spcsData.EggGroups) == 1 {
		return spcsData.EggGroups[0].Name, nil
	}

	return "", nil
}

//	===================================================================================================

func MakeLearnedMoveSlice(pkmnData fetch.Pokemon) []LearnedMove {
	var moveSlice []LearnedMove

	for _, m := range pkmnData.Moves {
		versions := getMoveVersions(pkmnData.Id, _learnedMove{
			Move: struct{ Name string }{Name: m.Move.Name},
			VersionGroupDetails: []struct {
				LevelLearnedAt  int
				MoveLearnMethod struct{ Name string }
				VersionGroup    struct{ Name string }
			}(m.VersionGroupDetails),
		})
		moveSlice = append(moveSlice, versions...)
	}
	return moveSlice
}

//	---------------------------------------------------------------------------------------------------

type _learnedMove struct {
	Move struct {
		Name string
	}
	VersionGroupDetails []struct {
		LevelLearnedAt  int
		MoveLearnMethod struct {
			Name string
		}
		VersionGroup struct {
			Name string
		}
	}
}

//	---------------------------------------------------------

func getMoveVersions(pkmnId int, moveData _learnedMove) []LearnedMove {
	var versionSlice []LearnedMove

	for _, vgd := range moveData.VersionGroupDetails {
		version := LearnedMove{
			PkmnId:  pkmnId,
			Move:    moveData.Move.Name,
			Level:   vgd.LevelLearnedAt,
			Method:  vgd.MoveLearnMethod.Name,
			Version: RenameLmoveVersion(vgd.VersionGroup.Name),
		}
		versionSlice = append(versionSlice, version)
	}
	return versionSlice
}

//	===================================================================================================

func MakePokedexSlice(spcsData fetch.PokemonSpecies) []Pokedex {
	var entrySlice []Pokedex

	if len(spcsData.FlavorTextEntries) != 0 {
		for _, fte := range spcsData.FlavorTextEntries {
			if fte.Language.Name == "en" {
				entry := getDexEntry(spcsData.Id, spcsData.Name, _pokedex{
					FlavorText: fte.FlavorText,
					Version:    struct{ Name string }(fte.Version),
				})
				entrySlice = append(entrySlice, entry)
			}
		}
	}
	return entrySlice
}

//	---------------------------------------------------------------------------------------------------

type _pokedex struct {
	FlavorText string
	Version    struct {
		Name string
	}
}

//	---------------------------------------------------------

func getDexEntry(spcsId int, spcsName string, entry _pokedex) Pokedex {
	return Pokedex{
		PkmnId:  spcsId,
		Search:  spcsName,
		Entry:   entry.FlavorText,
		Version: RenamePdexVersion(entry.Version.Name),
	}
}
