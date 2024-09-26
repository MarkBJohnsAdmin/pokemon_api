package frmt

import "server/database/fetch"

//	===================================================================================================

// MakeMoveStruct() => convert PokeApi /move into a slice of Move

//	---------------------------------------------------------------------------------------------------

// MakeMoveDescSlice() => get each version's description of the move from a PokeApi /move response

//	---------------------------------------------------------------------------------------------------

// MakeMachineStruct() => convert PokeApi /machine into a slice of Machine

//	===================================================================================================

func MakeMoveStruct(moveData fetch.Move) Move {
	name := getMoveName(moveData)
	effect, summary := getMoveEffects(moveData)

	return Move{
		Id:       moveData.Id,
		Search:   moveData.Name,
		Name:     name,
		Effect:   effect,
		Summary:  summary,
		Acc:      moveData.Accuracy,
		Power:    moveData.Power,
		Pp:       moveData.Pp,
		Priority: moveData.Priority,
		Class:    moveData.DamageClass.Name,
		Target:   RenameMoveTarget(moveData.Target.Name),
		Type:     moveData.Type.Name,
	}
}

//	---------------------------------------------------------

func getMoveName(moveData fetch.Move) string {
	if len(moveData.Names) != 0 {
		for _, n := range moveData.Names {
			if n.Language.Name == "en" {
				return n.Name
			}
		}
	}
	return moveData.Name
}

//	---------------------------------------------------------

func getMoveEffects(moveData fetch.Move) (string, string) {
	if moveData.EffectEntries != nil && len(*moveData.EffectEntries) != 0 {
		for _, ee := range *moveData.EffectEntries {
			if ee.Language.Name == "en" {
				return ee.Effect, ee.ShortEffect
			}
		}
	}
	return MoveEffectsMap(moveData.Name)
}

//	===================================================================================================

func MakeMoveDescSlice(moveData fetch.Move) []MoveDesc {
	var descriptions []MoveDesc

	if len(moveData.FlavorTextEntries) != 0 {
		for _, fte := range moveData.FlavorTextEntries {
			if fte.Language.Name == "en" {
				desc := getMoveDesc(moveData.Id, _moveDesc{
					FlavorText:   fte.FlavorText,
					VersionGroup: struct{ Name string }(fte.VersionGroup),
				})
				descriptions = append(descriptions, desc)
			}
		}
	}
	return descriptions
}

//	---------------------------------------------------------

type _moveDesc struct {
	FlavorText   string
	VersionGroup struct {
		Name string
	}
}

//	---------------------------------------------------------

func getMoveDesc(moveId int, desc _moveDesc) MoveDesc {
	return MoveDesc{
		MoveId:  moveId,
		Desc:    desc.FlavorText,
		Version: RenameMdescVersion(desc.VersionGroup.Name),
	}
}

//	===================================================================================================

func MakeMachineStruct(mchnData fetch.Machine) Machine {
	return Machine{
		Id:      mchnData.Id,
		TmNo:    mchnData.Item.Name,
		Move:    mchnData.Move.Name,
		Version: RenameMachineVersion(mchnData.VersionGroup.Name),
	}
}
