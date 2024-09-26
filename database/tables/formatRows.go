package tables

import (
	frmt "server/database/format"
)

//	===================================================================================================

//	MakePokemonRow() => for the "pkmn" table

//		MakeLearnedMoveRow() => for the "lmoves" table

//		MakePokedexRow() => for the "pdex" table

//	===================================================================================================

//	MakeAbilityRow() => for the "abltys" table

//		MakeAbltyDescRow() => for the "adescs" table

//	===================================================================================================

//	MakeMoveRow() => for the "moves" table

//		MakeMoveDescRow() => for the "mdescs" table

//	===================================================================================================

func MakePokemonRow(pkmnData *frmt.Pokemon) Pokemon {
	return Pokemon{
		Id:      pkmnData.Id,
		BaseId:  pkmnData.BaseId,
		Order:   pkmnData.Order,
		Search:  pkmnData.Search,
		Name:    pkmnData.Name,
		Moniker: pkmnData.Moniker,
		Ablty1:  pkmnData.Ablty1,
		Ablty2:  pkmnData.Ablty2,
		AbltyH:  pkmnData.Hidden,
		Type1:   pkmnData.Type1,
		Type2:   pkmnData.Type2,
		Egg1:    pkmnData.Egg1,
		Egg2:    pkmnData.Egg2,
		Color:   pkmnData.Color,
		Height:  pkmnData.Height,
		Weight:  pkmnData.Weight,
		Hp:      pkmnData.Hp,
		Atk:     pkmnData.Atk,
		Def:     pkmnData.Def,
		Spa:     pkmnData.Spa,
		Spd:     pkmnData.Spd,
		Spe:     pkmnData.Spe,
		Sprite:  pkmnData.Sprite,
		Offcl:   pkmnData.Offcl,
		Shiny:   pkmnData.Shiny,
		Home:    pkmnData.Home,
		Hshiny:  pkmnData.Hshny,
	}
}

//	---------------------------------------------------------------------------------------------------

func MakeLearnedMoveRow(lrndData *frmt.LearnedMove, tmMap *map[string][]tmName) LearnedMove {
	var tmName *string

	if tmList, exists := (*tmMap)[lrndData.Move]; exists {
		for _, tm := range tmList {
			if tm.Version == lrndData.Version {
				tmName = &tm.Name
				break
			}
		}
	}

	return LearnedMove{
		PkmnId:  lrndData.PkmnId,
		Move:    lrndData.Move,
		Level:   lrndData.Level,
		Method:  lrndData.Method,
		Tm:      tmName,
		Version: lrndData.Version,
	}
}

//	---------------------------------------------------------------------------------------------------

func MakePokedexRow(pdexData *frmt.Pokedex) Pokedex {
	return Pokedex{
		PkmnId:  pdexData.PkmnId,
		Search:  pdexData.Search,
		Entry:   pdexData.Entry,
		Version: pdexData.Version,
	}
}

//	===================================================================================================

func MakeAbilityRow(abltyData *frmt.Ability) Ability {
	return Ability{
		Id:      abltyData.Id,
		Search:  abltyData.Search,
		Name:    abltyData.Name,
		Effect:  abltyData.Effect,
		Summary: abltyData.Summary,
	}
}

//	---------------------------------------------------------------------------------------------------

func MakeAbltyDescRow(descData *frmt.AbltyDesc) AbltyDesc {
	return AbltyDesc{
		AbltyId: descData.AbltyId,
		Desc:    descData.Desc,
		Version: descData.Version,
	}
}

//	===================================================================================================

func MakeMoveRow(moveData *frmt.Move) Move {
	return Move{
		Id:       moveData.Id,
		Search:   moveData.Search,
		Name:     moveData.Name,
		Effect:   moveData.Effect,
		Summary:  moveData.Summary,
		Power:    moveData.Power,
		Acc:      moveData.Acc,
		Pp:       moveData.Pp,
		Priority: moveData.Priority,
		Class:    moveData.Class,
		Target:   moveData.Target,
		Type:     moveData.Type,
	}
}

//	---------------------------------------------------------------------------------------------------

func MakeMoveDescRow(descData *frmt.MoveDesc) MoveDesc {
	return MoveDesc{
		MoveId:  descData.MoveId,
		Desc:    descData.Desc,
		Version: descData.Version,
	}
}
