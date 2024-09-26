package frmt

import "server/database/fetch"

//	===================================================================================================

// MakeAbilityStruct() => convert PokeApi /ability to a slice of Ability

//	---------------------------------------------------------------------------------------------------

// MakeAbltyDescSlice() => get each version's description of the ability from a PokeApi /ability response

//	===================================================================================================

func MakeAbilityStruct(abltyData fetch.Ability) Ability {
	name := getAbilityName(abltyData)
	effect, summary := getAbilityEffects(abltyData)

	return Ability{
		Id:      abltyData.Id,
		Search:  abltyData.Name,
		Name:    name,
		Effect:  effect,
		Summary: summary,
	}
}

//	---------------------------------------------------------

func getAbilityName(abltyData fetch.Ability) string {
	if len(abltyData.Names) != 0 {
		for _, n := range abltyData.Names {
			if n.Language.Name == "en" {
				return n.Name
			}
		}
	}
	return abltyData.Name
}

//	---------------------------------------------------------

func getAbilityEffects(abltyData fetch.Ability) (string, string) {
	if abltyData.EffectEntries != nil && len(*abltyData.EffectEntries) != 0 {
		for _, ee := range *abltyData.EffectEntries {
			if ee.Language.Name == "en" {
				return ee.Effect, ee.ShortEffect
			}
		}
	}
	return AbilityEffectsMap(abltyData.Name)
}

//	===================================================================================================

func MakeAbltyDescSlice(abltyData fetch.Ability) []AbltyDesc {
	var descriptions []AbltyDesc

	if len(abltyData.FlavorTextEntries) != 0 {
		for _, fte := range abltyData.FlavorTextEntries {
			if fte.Language.Name == "en" {
				desc := getAbltyDesc(abltyData.Id, _abltyDesc{
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

type _abltyDesc struct {
	FlavorText   string
	VersionGroup struct {
		Name string
	}
}

//	---------------------------------------------------------

func getAbltyDesc(abltyId int, desc _abltyDesc) AbltyDesc {
	return AbltyDesc{
		AbltyId: abltyId,
		Desc:    desc.FlavorText,
		Version: RenameAdescVersion(desc.VersionGroup.Name),
	}
}
