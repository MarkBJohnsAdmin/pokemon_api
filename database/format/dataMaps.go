package frmt

import "strings"

//	PokemonNameMap => gets formatted names for alternate Pokemon forms

//	AbilityEffectsMap => get effects and effect summaries from abilities where data is missing from PokeApi

//	MoveEffectsMap => get effects and effect summaries from moves where data is missing from PokeApi

type _effects struct {
	effect  string
	summary string
}

func PokemonNameMap(search, baseName string) string {
	lookUpMap := map[string]string{
		"mr-mime-galar":               "Galarian Mr. Mime",
		"tauros-paldea-combat-breed":  "Paldean Tauros (Combat Breed)",
		"tauros-paldea-blaze-breed":   "Paldean Tauros (Blaze Breed)",
		"tauros-paldea-aqua-breed":    "Paldean Tauros (Aqua Breed)",
		"ursaluna-bloodmoon":          "Bloodmoon Ursaluna",
		"dudunsparce-three-segment":   "Dudunsparce (Three Segment)",
		"castform-sunny":              "Castform (Sunny Form)",
		"castform-rainy":              "Castform (Rainy Form)",
		"castform-snowy":              "Castform (Snowy Form)",
		"deoxys-attack":               "Deoxys (Attack Forme)",
		"deoxys-defense":              "Deoxys (Defense Forme)",
		"deoxys-speed":                "Deoxys (Speed Forme)",
		"wormadam-sandy":              "Wormadam (Sandy Cloak)",
		"wormadam-trash":              "Wormadam (Trash Cloak)",
		"shaymin-sky":                 "Shaymin (Sky Forme)",
		"basculin-blue-striped":       "Basculin (Blue Striped)",
		"basculin-white-striped":      "Basculin (White Striped)",
		"darmanitan-zen":              "Darmanitan (Zen Mode)",
		"darmanitan-galar-standard":   "Galarian Darmanitan",
		"darmanitan-galar-zen":        "Galarian Darmanitan (Zen Mode)",
		"rotom-heat":                  "Rotom Heat",
		"rotom-wash":                  "Rotom Wash",
		"rotom-frost":                 "Rotom Frost",
		"rotom-fan":                   "Rotom Fan",
		"rotom-mow":                   "Rotom Mow",
		"kyurem-white":                "White Kyurem",
		"kyurem-black":                "Black Kyurem",
		"keldeo-resolute":             "Keldeo (Resolute Form)",
		"meloetta-pirouette":          "Meloetta (Pirouette Forme)",
		"aegislash-blade":             "Aegislash (Blade Forme)",
		"zygarde-10":                  "Zygarde (10% " + "Forme)",
		"zygarde-complete":            "Zygarde (Complete Forme)",
		"hoopa-unbound":               "Hoopa (Unbound)",
		"oricorio-pom-pom":            "Oricorio (Pom-Pom Style)",
		"oricorio-pau":                "Oricorio (Pa'u Style)",
		"oricorio-sensu":              "Oricorio (Sensu Style)",
		"lycanroc-midnight":           "Lycanroc (Midnight Form)",
		"lycanroc-dusk":               "Lycanroc (Dusk Form)",
		"wishiwashi-school":           "Wishiwashi (School Form)",
		"minior-red":                  "Minior (Core Form - Red)",
		"minior-orange":               "Minior (Core Form - Orange)",
		"minior-yellow":               "Minior (Core Form - Yellow)",
		"minior-green":                "Minior (Core Form - Green)",
		"minior-blue":                 "Minior (Core Form - Blue)",
		"minior-indigo":               "Minior (Core Form - Indigo)",
		"minior-violet":               "Minior (Core Form - Violet)",
		"necrozma-dusk":               "Dusk Mane Necrozma",
		"necrozma-dawn":               "Dawn Wings Necrozma",
		"necrozma-ultra":              "Ultra Necrozma",
		"magearna-original":           "Magearna (Original Color)",
		"toxtricity-low-key":          "Toxtricity (Low Key Form)",
		"eiscue-noice":                "Eiscue (Noice Face)",
		"morpeko-hangry":              "Morpeko (Hangry Mode)",
		"calyrex-ice":                 "Calyrex (Ice Rider)",
		"calyrex-shadow":              "Calyrex (Shadow Rider)",
		"eternatus-eternamax":         "Eternamax Eternatus",
		"urshifu-rapid-strike":        "Urshifu (Rapid Strike Style)",
		"squawkabilly-blue-plumage":   "Squawkabilly (Blue Plumage)",
		"squawkabilly-yellow-plumage": "Squawkabilly (Yellow Plumage)",
		"squawkabilly-white-plumage":  "Squawkabilly (White Plumage)",
		"maushold-family-of-three":    "Maushold (Family of Three)",
		"palafin-hero":                "Palafin (Hero Form)",
		"tatsugiri-droopy":            "Tatsugiri (Droopy Form)",
		"tatsugiri-stretchy":          "Tatsugiri (Stretchy Form)",
		"gimmighoul-roaming":          "Gimmighoul (Roaming Form)",
		"ogerpon-wellspring-mask":     "Ogerpon (Wellspring Mask)",
		"ogerpon-hearthflame-mask":    "Ogerpon (Hearthflame Mask)",
		"ogerpon-cornerstone-mask":    "Ogerpon (Cornerstone Mask)",
		"terapagos-terastal":          "Terapagos (Terastal Form)",
		"terapagos-stellar":           "Terapagos (Stellar Form)",
	}

	if name, ok := lookUpMap[search]; ok {
		return name
	}

	if strings.Contains(search, "-mega-x") {
		return "Mega " + baseName + " X"
	}
	if strings.Contains(search, "-mega-y") {
		return "Mega " + baseName + " Y"
	}
	if strings.Contains(search, "-mega") {
		return "Mega " + baseName
	}
	if strings.Contains(search, "-alola") {
		return "Alolan " + baseName
	}
	if strings.Contains(search, "-galar") {
		return "Galarian " + baseName
	}
	if strings.Contains(search, "-hisui") {
		return "Hisuian " + baseName
	}
	if strings.Contains(search, "-paldea") {
		return "Paldean " + baseName
	}
	if strings.Contains(search, "-primal") {
		return "Primal " + baseName
	}
	if strings.Contains(search, "-origin") {
		return baseName + " (Origin Forme)"
	}
	if strings.Contains(search, "-female") {
		return baseName + " (Female)"
	}
	if strings.Contains(search, "-therian") {
		return baseName + " (Therian Forme)"
	}
	if strings.Contains(search, "-crowned") {
		return baseName + " (Crowned)"
	}
	return baseName
}

func AbilityEffectsMap(search string) (string, string) {
	abltyEffctMap := map[string]_effects{
		"intrepid-sword": {
			effect:  "Intrepid Sword raises the Attack stat of the Pokémon with this Ability by one stage when it enters the battle.",
			summary: "Boosts Attack in battle.",
		},
		"dauntless-shield": {
			effect:  "Dauntless Shield raises the Defense stat of the Pokémon with this Ability by one stage when it enters the battle.",
			summary: "Boosts Defense in battle.",
		},
		"libero": {
			effect:  "Libero changes the Pokémon's type to that of its previously used attack.",
			summary: "Libero changes the Pokémon's type to that of its previously used attack.",
		},
		"ball-fetch": {
			effect: "At any time after the first Poké Ball is thrown and fails to catch a Pokémon, at the end of a turn," +
				" if a Pokémon with Ball Fetch is on the field and not holding another item," +
				" it will pick up the same type of ball as the first one thrown. This can only occur once in a battle.",
			summary: "If the Pokémon is not holding an item, it will fetch the Poké Ball from the first failed throw of the battle.",
		},
		"cotton-down": {
			effect: "When a Pokémon with this Ability is hit by a damaging move, the Speed of all the other Pokémon is decreased by one stage. " +
				"If a Pokémon with this Ability is hit by a multistrike move, each hit activates this ability.",
			summary: "When a Pokémon with Cotton Down is hit by an attack, it lowers the attacker's Speed.",
		},
		"propeller-tail": {
			effect:  "Propeller Tail makes its user's moves ignore the target-redirecting effects of moves and Abilities.",
			summary: "Ignores moves and abilities that draw in moves.",
		},
		"mirror-armor": {
			effect:  "Mirror Armor reflects the stat-lowering effects of moves and Abilities back to the user.",
			summary: "Reflects any stat-lowering effects.",
		},
		"gulp-missile": {
			effect: "If a Cramorant with Gulp Missile uses Surf or Dive, it catches prey and changes its form depending on its remaining HP. " +
				"If its HP is over half, it catches an Arrokuda and changes into its Gulping Form. If its HP is under half, it instead catches a Pikachu " +
				"and changes into its Gorging Form. If hit by an attack in either form, Cramorant spits out its catch at the attacker, " +
				"dealing damage equal to ¼ of the attacker's maximum HP and inflicting a secondary effect depending on its catch: Arrokuda lowers the attacker's " +
				"Defense by one stage, while Pikachu paralyzes them. Cramorant returns to its base form after spitting out its catch.",
			summary: "If Cramorant uses the move Surf or Dive, it returns with a fish in its mouth. If Cramorant then takes damage, it retaliates by spitting out the fish, dealing extra damage. " +
				"Cramorant may also return carrying a Pikachu. If Pikachu is launched at the opponent they will become paralyzed.",
		},
		"stalwart": {
			effect:  "Stalwart makes its user's moves ignore the target-redirecting effects of moves and Abilities.",
			summary: "Ignores moves and abilities that draw in moves.",
		},
		"steam-engine": {
			effect: "When a Pokémon with Steam Engine is hit by a Fire or Water type move, its Speed stat is increased by six stages. " +
				"This Ability activates after any stat decreases caused by the move are applied.",
			summary: "Steam Engine raises the user's Speed by three stages when it is hit by a Fire- or Water-type attack. The attack still deals damage.",
		},
		"punk-rock": {
			effect:  "Punk Rock boosts the power of the Pokémon's sound-based moves by 30%. A Pokémon with Punk Rock also takes half damage from damaging sound-based moves.",
			summary: "Boosts sound-based moves and halves damage from the same moves.",
		},
		"sand-spit": {
			effect: "Sand Spit summons a sandstorm in battle for five turns when a Pokémon with Sand Spit is hit by a damaging move, unless sandstorm is already present. " +
				"This can be extended to eight turns if the Pokémon is holding a Smooth Rock.",
			summary: "Creates a sandstorm when hit by an attack.",
		},
		"ice-scales": {
			effect: "Ice Scales halves the damage from special moves that hit Pokémon with this Ability, including from special moves that use the target's Defense stat. " +
				"It has no effect on moves that deal direct damage.",
			summary: "Halves damage from Special moves.",
		},
		"ripen": {
			effect: "Ripen doubles the effects from held Berries when eaten in battle. This affects Berries that restore HP or PP, Berries that reduce damage, " +
				"Berries that raise stats, and Berries that deal damage to Pokémon attacking the holder. The only Berries that have in-battle effects that Ripen does not effect " +
				"are the status condition-healing Berries, the Lansat Berry, the Micle Berry, and the Custap Berry.",
			summary: "Doubles the effect of berries.",
		},
		"ice-face": {
			effect: "When an Eiscue that's in its Ice Face form is hit by a physical move, it takes no damage and changes into its Noice Face form. If a hailstorm or snowstorm begins " +
				"while Eiscue is in its Noice Face form, or if it is sent out in its Noice Face form during hail, it will immediately revert to its Ice Face form. Eiscue will not revert to its " +
				"Ice Face form if there is already hail when it changes into its Noice Face form. This form change persists even when Eiscue is switched out, but is undone at the end of the battle. " +
				"Ice Face does not protect against special moves.",
			summary: "Ice Face is the signature ability of Eiscue. It allows the Pokémon to take a physical attack without taking damage. After the attack, Eiscue changes to its No Ice Face form, " +
				"however, it can be restored in hail.",
		},
		"power-spot": {
			effect:  "The power of moves used by the allies of the Pokémon with Power Spot is increased by 30%.",
			summary: "Just being next to the Pokémon powers up moves.",
		},
		"mimicry": {
			effect: "Depending on the terrain, Mimicry will change the type of a Pokémon with this Ability. During Electric Terrain, it will become a pure Electric type. During Grassy Terrain, it will " +
				"become a pure Grass type. During Misty Terrain, it will become a pure Fairy type. During Psychic Terrain, it will become a pure Psychic type. When terrain ends, the Pokemon reverts back to its " +
				"original typing. Mimicry activates after the Pokémon with Mimicry is switched in while terrain is present, or whenever there is a change in terrain; changing the Pokemon's type by other means " +
				"overrides the type caused by Mimicry. Mimicry will activate regardless of whether the Pokémon with Mimicry is grounded or not.",
			summary: "Changes type depending on the terrain.",
		},
		"screen-cleaner": {
			effect:  "When a Pokémon with Screen Cleaner is sent out, the effects of Light Screen, Reflect, and Aurora Veil end on both sides of the field.",
			summary: "Nullifies effects of Light Screen, Reflect, and Aurora Veil.",
		},
		"steely-spirit": {
			effect:  "The power of Steel-type moves used by a Pokémon with Steely Spirit and its allies is increased by 50%. This effect stacks if multiple Pokémon with this Ability are on the field.",
			summary: "Powers up ally Pokémon's Steel-type moves.",
		},
		"perish-body": {
			effect: "When the Pokémon with Perish Body is hit with a contact move, both the attacker and the Pokémon with Perish Body will faint in three turns. " +
				"If the attacker already has a perish count or is wearing Protective Pads or has Long Reach, the Ability does not affect anyone.",
			summary: "When hit by a move that makes direct contact, the Pokémon and the attacker will faint after three turns unless they switch out of battle.",
		},
		"wandering-spirit": {
			effect:  "When a Pokémon with Wandering Spirit is hit by a contact move, it exchanges Abilities with the attacker. If a Pokémon affected by Wandering Spirit is switched out, it regains its original Ability.",
			summary: "Swaps abilities with opponents on contact.",
		},
		"gorilla-tactics": {
			effect: "Gorilla Tactics boosts the Pokémon's Attack by 50%, but limits the Pokémon to using only one move, similar to the Choice Band. This effect stacks with Choice items, " +
				"meaning a Pokémon with this Ability and a Choice Band will have its Attack boosted to 225%.",
			summary: "Boosts the Pokémon's Attack stat but only allows the use of the first selected move.",
		},
		"neutralizing-gas": {
			effect: "Neutralizing Gas suppresses the effects of Abilities of other Pokémon currently in battle that are not holding an Ability Shield. When a Neutralizing Gas user switches in, " +
				"the Ability displays a message, even before entry hazards take effect. Once Neutralizing Gas is no longer applying, Abilities that activate when gained immediately pop up an activation " +
				"message regardless of whether they have already done so, with the exceptions of Unnerve and Imposter. If Neutralizing Gas stops applying because the Pokémon with Neutralizing Gas is " +
				"about to switch out, those Abilities activate before the switch occurs.",
			summary: "Neutralizes abilities of all Pokémon in battle.",
		},
		"pastel-veil": {
			effect:  "Pastel Veil prevents the Pokémon with this Ability and its allies from being afflicted by poison. It also heals the poisoned status condition of allies if a Pokémon with this Ability is sent out into battle.",
			summary: "Pastel Veil prevents the Pokémon and its teammates from being poisoned. It also cures teammates of poisoning when it enters the battlefield.",
		},
		"hunger-switch": {
			effect: "At the end of each turn, Hunger Switch switches a Morpeko with the Ability between its Full Belly Mode and Hangry Mode. In Full Belly Mode, Aura Wheel is an Electric-type move; in Hangry Mode, it is a Dark-type move. " +
				"If Morpeko's Hunger Switch is negated or replaced, it remains in its current form. If a Pokémon other than Morpeko obtains Hunger Switch with Transform or Imposter, the Pokémon will not change forms. Hunger Switch does " +
				"not activate while Morpeko is Terastallized.",
			summary: "Hunger Switch causes Morpeko to change its form each turn, alternating between Full Belly Mode and Hangry Mode. Morpeko's form also changes the type of its signature move, Aura Wheel.",
		},
		"quick-draw": {
			effect:  "If a Pokémon with this Ability selects a damaging move, it has a 30%" + " chance of going first in its priority bracket. If the Ability activates, this is announced at the start of the turn.",
			summary: "Enables the Pokémon to move first occasionally.",
		},
		"unseen-fist": {
			effect: "When using a contact move, Pokémon with Unseen Fist ignore protection moves in effect for their targets, except Max Guard. This applies even if the user is holding Protective Pads, " +
				"but does not apply if the user is holding a Punching Glove, as it turns punching moves into non-contact moves.",
			summary: "Contact moves can strike through Protect/Detect.",
		},
		"curious-medicine": {
			effect:  "When a Pokémon with this Ability enters the battle, stat stages of ally Pokémon are reset to 0.",
			summary: "Resets all stat changes upon entering battlefield.",
		},
		"transistor": {
			effect:  "If an Electric-type move is used by a Pokémon with this Ability, its Attack or Special Attack stat is multiplied by 1.3 during damage calculation, effectively increasing damage dealt by 30%.",
			summary: "Transistor increases the power of Electric-type moves used by this Pokémon by 30%.",
		},
		"dragons-maw": {
			effect:  "If a Dragon-type move is used by a Pokémon with this Ability, its Attack or Special Attack stat is multiplied by 1.5 during damage calculation, effectively increasing damage dealt by 50%.",
			summary: "Dragon's Maw increases the power of Dragon-type moves used by this Pokémon by 50%.",
		},
		"chilling-neigh": {
			effect:  "When a Pokémon with Chilling Neigh directly causes another Pokémon to faint by using a damaging move, its Attack is increased by one stage. If its Attack is already at +6 stages, Chilling Neigh will not activate.",
			summary: "Boosts Attack after knocking out a Pokémon.",
		},
		"grim-neigh": {
			effect:  "When a Pokémon with Grim Neigh directly causes another Pokémon to faint by using a damaging move, its Special Attack is increased by one stage. If its Special Attack is already at +6 stages, Grim Neigh will not activate.",
			summary: "Boosts Special Attack after knocking out a Pokémon.",
		},
		"as-one-glastrier": {
			effect:  "As Ice Rider Calyrex, this Ability combines the effects of both Calyrex's Unnerve Ability and Glastrier's Chilling Neigh Ability.",
			summary: "Combines Unnerve and Chilling Neigh.",
		},
		"as-one-spectrier": {
			effect:  "As Shadow Rider Calyrex, this Ability combines the effects of both Calyrex's Unnerve Ability and Spectrier's Grim Neigh Ability.",
			summary: "Combines Unnerve and Grim Neigh.",
		},
		"lingering-aroma": {
			effect: "When a Pokémon with Lingering Aroma is hit by a contact move, the attacker's Ability is changed to Lingering Aroma. The attacker regains its original Ability once it is switched out, faints, or the battle ends. " +
				"If a Pokémon with this Ability is hit by a multistrike move that makes contact (such as Fury Swipes), the attacking Pokémon's Ability will become Lingering Aroma after the first strike. " +
				"If a Pokémon with Rock Head uses a recoil contact move on a Pokémon with Lingering Aroma, its Ability will become Lingering Aroma, then subsequently it will take recoil damage.",
			summary: "Contact changes the attacker's Ability to Lingering Aroma.",
		},
		"seed-sower": {
			effect: "When a Pokémon with Seed Sower is hit by a damaging move, it summons Grassy Terrain for five turns, unless Grassy Terrain is already present. This can be extended to eight turns if the Pokémon " +
				"is holding a Terrain Extender. If hit by a multistrike move, the first strike will activate it.",
			summary: "Turns the ground into Grassy Terrain when the Pokémon is hit by an attack.",
		},
		"thermal-exchange": {
			effect: "When a Pokémon with Thermal Exchange is hit by a Fire-type damaging move, its Attack is increased by one stage. It does not provide immunity or reduction to the damage. " +
				"It also prevents the Pokémon from being burned. This Ability does not activate if the substitute is hit",
			summary: "Raises Attack when hit by a Fire-type move. Cannot be burned.",
		},
		"anger-shell": {
			effect: "When the Pokémon's HP falls from over 50%" + " to 50%" + " or lower as a result of an attack from another Pokémon, its Attack, Special Attack, and Speed rise by one stage each, while its Defense and Special Defense " +
				"fall by one stage each. Non-attack damage (e.g., due to a status condition, recoil, or Life Orb) does not activate Anger Shell. If hit by a multistrike move, Anger Shell activates only after all strikes are finished.",
			summary: "When the Pokémon's HP drops below half, Anger Shell lowers its Defense and Special Defense but its Attack, Special Attack and Speed are raised.",
		},
		"purifying-salt": {
			effect: "Purifying Salt grants the Pokémon immunity from all status conditions and halves the Attack and Special Attack of an attacking Pokémon during damage calculation if targeted by a Ghost-type attack. " +
				"Rest fails if used by a Pokémon with Purifying Salt",
			summary: "Protects from status conditions and halves damage from Ghost-type moves.",
		},
		"well-baked-body": {
			effect:  "When a Pokémon with Well-Baked Body is hit by a Fire-type move, its Defense is increased by two stages, and the move will have no effect on that Pokémon",
			summary: "Immune to Fire-type moves, and Defense is sharply boosted.",
		},
		"wind-rider": {
			effect: "Pokémon with Wind Rider are immune to damage from wind moves and its Attack is increased by one stage when hit by one. When Tailwind takes effect on its side or when the Pokémon with Wind Rider is " +
				"sent into battle while Tailwind is in effect on its side, its Attack is increased by one stage in addition to receiving the speed boosting effect of Tailwind",
			summary: "Wind Rider gives immunity to wind moves, and causes the Pokémon's Attack to increase by one stage when hit by one.",
		},
		"guard-dog": {
			effect: "This Ability prevents other Pokémon from forcing the Pokémon with this Ability to switch out or flee with a move or item. Additionally, if a Pokémon with this ability is affected by Intimidate, " +
				"its Attack stat will increase by one stage, rather than decreasing by one stage.",
			summary: "Boosts Attack if intimidated, and prevents being forced to switch out.",
		},
		"rocky-payload": {
			effect:  "If a Rock-type move is used by a Pokémon with this Ability, its Attack or Special Attack stat is multiplied by 1.5 during damage calculation, effectively increasing damage dealt by 50%.",
			summary: "Rocky Payload increases the power of Rock-type moves used by this Pokémon by 50%.",
		},
		"wind-power": {
			effect:  "When a Pokémon with Wind Power is hit by a wind move or when Tailwind takes effect, it becomes charged, boosting the power of the next Electric-type move it uses.",
			summary: "If a Pokémon with Wind Power is hit by a wind move, the power of the next Electric-type move it uses is doubled.",
		},
		"zero-to-hero": {
			effect: "After switching out for the first time in battle, a Palafin with the Ability switches to its Hero Form until the end of battle. This includes through the use of moves that switch the user out, " +
				"and moves used by opposing Pokémon that force it to switch. If Hero Form Palafin faints and revives later, it does not revert to its Zero Form.",
			summary: "Transforms into its Hero Form when switching out.",
		},
		"commander": {
			effect: "When a Tatsugiri with Commander is on the same side of the field as a Dondozo, it will enter the Dondozo's mouth and the Dondozo will have its Attack, Defense, Special Attack, Special Defense, " +
				"and Speed raised by two stages each. While Tatsugiri is inside Dondozo's mouth, the latter will be unable to switch out, and cannot be forced out. The Trainer will not be able to select moves for Tatsugiri, " +
				"and any opponent's moves will miss if Tatsugiri is targeted. All prior status effects will continue normally. If Dondozo is switched in when Tatsugiri is already on the field, Commander will immediately " +
				"activate and cancel Tatsugiri's selected move; any moves targeting Tatsugiri that turn will also miss. When Dondozo faints, Tatsugiri returns to its normal position in battle",
			summary: "Goes inside the mouth of an ally Dondozo if one is on the field.",
		},
		"electromorphosis": {
			effect: "When a Pokémon with this Ability is hit by a move, it becomes charged, boosting the power of the next Electric-type move it uses. " +
				"If hit by a multistrike move, each hit triggers the Ability, but the charges do not stack.",
			summary: "If a Pokémon with Electromorphosis is hit by an attack, the power of the next Electric-type move it uses is doubled.",
		},
		"protosynthesis": {
			effect: "When a Pokémon with Protosynthesis is in harsh sunlight or is holding Booster Energy, its highest stat (other than HP and taking into account stat stages, but not held items) is increased by 30%, or 50% " +
				"if Speed is the highest stat. Upon activation, a message tells which stat is increased to all the players. If multiple stats are tied for highest, Protosynthesis breaks the tie by prioritizing " +
				"Attack, Defense, Special Attack, Special Defense, and Speed, in that order.",
			summary: "Raises highest stat in harsh sunlight, or if holding Booster Energy.",
		},
		"quark-drive": {
			effect: "When a Pokémon with Quark Drive is on the field and Electric Terrain is active, or the Pokémon is holding Booster Energy, its highest stat (other than HP and taking into account stat stages, but not held items) " +
				"is increased by 30%, or 50% if Speed is the highest stat. Upon activation, a message tells which stat is increased to all the players. If multiple stats are tied for highest, Quark Drive breaks the tie by prioritizing " +
				"Attack, Defense, Special Attack, Special Defense, and Speed, in that order.",
			summary: "Raises highest stat on Electric Terrain, or if holding Booster Energy.",
		},
		"good-as-gold": {
			effect: "Pokémon with this Ability are unaffected by other Pokémon's status moves. It does not block status moves that target the user; the other side of the field, or all Pokémon. It does, however, " +
				"block potentially beneficial status moves from ally Pokémon.",
			summary: "Gives immunity to status moves.",
		},
		"vessel-of-ruin": {
			effect:  "Vessel of Ruin decreases the Special Attack stat of all Pokémon on the field other than Pokémon with this Ability by 25%. The effect does not stack if more than one Pokémon with Vessel of Ruin is on the field.",
			summary: "Lowers Special Attack of all Pokémon except itself.",
		},
		"sword-of-ruin": {
			effect:  "Sword of Ruin decreases the Defense stat of all Pokémon on the field other than Pokémon with this Ability by 25%. The effect does not stack if more than one Pokémon with Sword of Ruin is on the field.",
			summary: "Lowers Defense of all Pokémon except itself.",
		},
		"tablets-of-ruin": {
			effect: "Tablets of Ruin decreases the Attack stat of all Pokémon on the field other than Pokémon with this Ability by 25%. The effect does not stack if more than one Pokémon with Tablets of Ruin is on the field. " +
				"This does not affect self-inflicted confusion damage for Pokémon other than the one with this Ability.",
			summary: "Lowers Attack of all Pokémon except itself.",
		},
		"beads-of-ruin": {
			effect:  "Beads of Ruin decreases the Special Defense stat of all Pokémon on the field other than Pokémon with this Ability by 25%. The effect does not stack if more than one Pokémon with Beads of Ruin is on the field.",
			summary: "Lowers Special Defense of all Pokémon except itself.",
		},
		"orichalcum-pulse": {
			effect:  "Orichalcum Pulse summons harsh sunlight in battle as soon as a Pokémon with Orichalcum Pulse enters the battle. Orichalcum Pulse also boosts the Pokémon's Attack stat by roughly 33% in harsh sunlight.",
			summary: "Turns the sunlight harsh when entering battle, and boosts Attack while active.",
		},
		"hadron-engine": {
			effect: "Hadron Engine creates the effect of the move Electric Terrain on the battlefield as soon as a Pokémon with Hadron Engine enters the battle. Hadron Engine also boosts the Pokémon's Special Attack stat " +
				"by roughly 33%" + " on Electric Terrain.",
			summary: "Creates an Electric Terrain when entering battle, and boosts Special Attack while active.",
		},
		"opportunist": {
			effect: "When an opposing Pokémon raises its stats, the Pokémon with this Ability will also raise the same stats by the same number of stages as the opposing Pokémon. This Ability does not activate when an ally " +
				"Pokémon raises its stats, or when an opposing Pokémon raises its stats through Opportunist. Opportunist will trigger if an ally Pokémon uses Swagger or Flatter against an opposing Pokémon, after the confusion is applied. " +
				"If multiple triggers happen at once, Opportunist will sum them up.",
			summary: "Copies stat boosts by the opponent.",
		},
		"cud-chew": {
			effect:  "If a Pokémon with Cud Chew eats a Berry or is affected by a Flung Berry, its effect will activate again at the end of the next turn.",
			summary: "Cud Chew causes the Pokémon to reuse an already consumed Berry at the end of the next turn.",
		},
		"sharpness": {
			effect:  "Sharpness boosts the strength of slicing moves used by a Pokémon with this Ability by 50%.",
			summary: "Sharpness increases the power of all slicing moves by 50%.",
		},
		"supreme-overlord": {
			effect:  "When a Pokémon with Supreme Overlord enters battle, it receives a 10% additive increase to the power of its moves for each ally Pokémon that is fainted in its party.",
			summary: "Attack and Special Attack are boosted for each party Pokémon that has been defeated.",
		},
		"costar": {
			effect:  "When a Pokémon with Costar enters the battle, it copies an ally's stat stages for itself.",
			summary: "Copies ally's stat changes on entering battle.",
		},
		"toxic-debris": {
			effect: "When a Pokémon with this Ability is hit by a physical move, Toxic Spikes are set on the opposing side of the field. If an attack would cause more than two layers of Toxic Spikes to be set or hits the substitute instead, " +
				"the Ability does not activate. If a Pokémon with this Ability is hit by a physical multistrike move, each hit will activate this Ability.",
			summary: "Scatters poison spikes at the feet of the opposing team when the Pokémon takes damage from physical moves.",
		},
		"armor-tail": {
			effect:  "If an opponent uses a priority move that targets the Pokémon with Armor Tail or its allies, Armor Tail prevents the Pokémon from executing that move.",
			summary: "Armor Tail prevents the opponent from using any moves that have priority.",
		},
		"earth-eater": {
			effect: "When a Pokémon with Earth Eater is hit by a Ground-type move, its HP is restored by one quarter of its maximum HP, and the move will have no effect on that Pokémon. " +
				"Earth Eater will not activate if the Pokémon is protected from the Ground-type move.",
			summary: "Restores HP when hit by a Ground-type move.",
		},
		"mycelium-might": {
			effect:  "When a Pokémon with Mycelium Might uses a status move, it always moves last within its priority bracket, and the effects of all Pokémon's ignorable Abilities are ignored for the execution of that move.",
			summary: "Status moves go last, but are not affected by the opponent's ability.",
		},
		"minds-eye": {
			effect: "Mind's Eye enables the Pokémon with this Ability to hit Ghost-type Pokémon with damage-dealing Normal- and Fighting-type moves. It also prevents other Pokémon from lowering the Pokémon's " +
				"accuracy and ignores changes to the opponents' evasion.",
			summary: "The Pokémon ignores changes to opponents' evasiveness, its accuracy can't be lowered, and it can hit Ghost types with Normal- and Fighting-type moves.",
		},
		"supersweet-syrup": {
			effect:  "Once per battle, when a Pokémon with Supersweet Syrup enters the battle, it lowers the evasion stat of all adjacent opponents by one stage.",
			summary: "A sickly sweet scent spreads across the field the first time the Pokémon enters a battle, lowering the evasiveness of opposing Pokémon.",
		},
		"hospitality": {
			effect:  "When the Pokémon enters a battle, it showers its ally with hospitality, restoring a small amount of the ally's HP by 25%.",
			summary: "Heals allies upon switching in.",
		},
		"toxic-chain": {
			effect: "If a Pokémon with this Ability uses a move that does not target the user, there is a 30%" + " chance the target will become badly poisoned. Toxic Chain's chances also apply independently for each hit of a multistrike move. " +
				"If this Ability triggers when using the move Knock Off on a target with a Lum or Pecha Berry, the Berry will be consumed to cure the poisoning instead of being removed by Knock Off's effect.",
			summary: "Toxic Chain may cause bad poisoning when the Pokémon hits an opponent with a move.",
		},
		"embody-aspect": {
			effect: "Upon Terastallizing, Ogerpon's Ability will change to Embody Aspect, and one of its stats is raised by one stage, depending on the mask: Teal Mask: Speed, Wellspring Mask: Sp. Def, Hearthflame Mask: Attack, " +
				"Cornerstone Mask: Defense.",
			summary: "Embody Aspect boosts a particular stat, depending on the form of a Terastallized Ogerpon.",
		},
		"tera-shift": {
			effect:  "When the Pokémon enters a battle, it absorbs the energy around itself and transforms into its Terastal Form.",
			summary: "When Terapagos enters the battle, it transforms into its Tera type.",
		},
		"tera-shell": {
			effect:  "The Pokémon's shell contains the powers of each type. All damage-dealing moves that hit the Pokémon when its HP is full will not be very effective.",
			summary: "All damage dealing moves when Pokémon is at full health are not very effective.",
		},
		"teraform-zero": {
			effect:  "When Terapagos changes into its Stellar Form, it uses its hidden powers to eliminate all effects of weather and terrain, reducing them to zero.",
			summary: "When transforming into Stellar Form, Terapagos clears all weather conditions.",
		},
		"poison-puppeteer": {
			effect:  "When a Pokémon is poisoned by any of Pecharunt's moves, it will also become confused. Poison induced by Toxic Spikes does not activate Poison Puppeteer. Poison Puppeteer does not activate if the user is not Pecharunt.",
			summary: "Pokémon poisoned by Pecharunt's moves will also become confused.",
		},
	}

	if e, found := abltyEffctMap[search]; found {
		return e.effect, e.summary
	}

	return "", ""
}

func MoveEffectsMap(search string) (string, string) {

	moveEffctMap := map[string]_effects{
		"tar-shot": {
			effect:  "Tar Shot lowers the target's Speed stat by one stage and doubles the effectiveness of Fire-type moves used on the target.",
			summary: "Lowers the opponent's Speed and makes them weaker to Fire-type moves.",
		},
		"magic-powder": {
			effect: "Magic Powder changes the target's type to pure Psychic. The effect lasts as long as the target remains " +
				"in battle; however, changing form, transforming or Terastallizing overrides the effect.",
			summary: "Changes target's type to Psychic.",
		},
		"dragon-darts": {
			effect: "Dragon Darts is a multistrike move that deals damage twice. If targeting an opponent when there are " +
				"multiple opponents, Dragon Darts' first strike targets the original target and its second strike targets the " +
				"other opponent. If either the original target or its ally would take no damage from Dragon Darts, both strikes " +
				"are directed at the other Pokémon instead. Wide Guard does not affect Dragon Darts. ",
			summary: "Deals damage and hits twice. In a double battle, it hits each opponent once.",
		},
		"teatime": {
			effect: "Teatime causes all Pokémon on the field to eat their held Berry. This move bypasses substitutes. " +
				"If a Pokémon has eaten its Berry, Teatime fails to work on that Pokémon. The Berry is eaten even if it would " +
				"not have any effect. Teatime causes Pokémon to eat their Berries even in the presence of Unnerve or Magic Room.",
			summary: "Forces all Pokémon on the field to eat their berries.",
		},
		"octolock": {
			effect: "Octolock prevents the target from fleeing or switching out. At the end of each turn, " +
				"the target's Defense and Special Defense are lowered by one stage each. If the user switches out, " +
				"the target will be freed. It cannot be used in Paldea.",
			summary: "Prevents the target from fleeing or switching out, while the user remains in battle. " +
				"It also lowers the target's Defense and Special Defense by one stage each every turn.",
		},
		"bolt-beak": {
			effect: "Bolt Beak deals damage. If the user attacks before the target, or if the target " +
				"switches in during the turn that Bolt Beak is used, its base power doubles to 170.",
			summary: "If the user attacks before the target, the power of this move is doubled.",
		},
		"fishious-rend": {
			effect: "Fishious Rend deals damage. If the user attacks before the target, or if the target switches " +
				"in during the turn that Fishious Rend is used, its base power doubles to 170. Strong Jaw further increases " +
				"power by 50%.",
			summary: "If the user attacks before the target, the power of this move is doubled.",
		},
		"court-change": {
			effect: "Court Change causes any active battle effects to swap sides. This affects barriers like Reflect, " +
				"multi-turn effects like Tailwind, and entry hazards like Spikes. Court Change transfers an entry hazard, " +
				"it will not have any effect until a new Pokémon is switched in.",
			summary: "Swaps the effects on either side of the field.",
		},
		"clangorous-soul": {
			effect: "Clangorous Soul raises the user's Attack, Defense, Special Attack, Special Defense and Speed by one stage each, " +
				"at the cost of 1/3 of its maximum HP",
			summary: "Raises all user's stats but loses HP.",
		},
		"body-press": {
			effect: "Body Press inflicts damage, but it uses the user's Defense stat instead of its Attack stat to calculate damage. " +
				"The user's Defense stat stage-modifiers are applied, but otherwise its Attack modifiers are used. For example, the user's " +
				"Huge Power or a held Choice Band will increase the stat used in calculation, whereas Fur Coat or a held Eviolite will not",
			summary: "The higher the user's Defense, the stronger the attack.",
		},
		"decorate": {
			effect: "Decorate raises the target's Attack and Special Attack stats by two stages each. Unlike most moves, " +
				"it bypasses moves like Protect, Detect, and even Max Guard to affect the target but is blocked by Crafty Shield. " +
				"It cannot affect a target that has a substitute.",
			summary: "Sharply raises target's Attack and Special Attack.",
		},
		"behemoth-blade": {
			effect: "Behemoth Blade deals damage. If the target is Dynamaxed or Gigantamaxed, its damage is doubled. " +
				"The bonus does not apply to Eternamax Eternatus, despite its similarities to a Gigantamaxed Pokémon.",
			summary: "Deals damage, and if the target is Dynamaxed its power doubles to 200.",
		},
		"behemoth-bash": {
			effect: "Behemoth Bash deals damage. If the target is Dynamaxed or Gigantamaxed, its damage is doubled. " +
				"The bonus does not apply to Eternamax Eternatus, despite its similarities to a Gigantamaxed Pokémon.",
			summary: "Deals damage, and if the target is Dynamaxed its power doubles to 200.",
		},
		"aura-wheel": {
			effect: "Aura Wheel inflicts damage and raises the user's Speed stat by one stage. Its type changes depending on " +
				"Morpeko's current mode: Electric type in Full Belly Mode, and Dark type in Hangry Mode.",
			summary: "Deals damage and raises the user's Speed by one stage. Its type changes based on Morpeko's form.",
		},
		"overdrive": {
			effect:  "Overdrive inflicts damage to all opponents.",
			summary: "Hits all adjacent opponents.",
		},
		"apple-acid": {
			effect:  "Apple Acid inflicts damage and lowers the target's Special Defense stat by one stage.",
			summary: "Lowers target's Special Defense.",
		},
		"life-dew": {
			effect: "Life Dew restores the HP of the user and its allies by 25%" + " of their maximum HP. If an ally has Water Absorb, " +
				"Storm Drain, or Dry Skin, it will activate and negate this move's effect on it.",
			summary: "Life Dew recovers up to 25%" + " of the user's and its teammate's maximum HP.",
		},
		"obstruct": {
			effect: "Obstruct protects the user from all effects of physical and special moves that target it during the turn it is " +
				"used. If Obstruct blocks an attack that would make contact with the user, the attacker's Defense stat is lowered by 2 stages. " +
				"Obstruct causes the Defense stat to drop even if the user of Obstruct was immune to that contact move. It cannot be used in " +
				"Paldea.",
			summary: "Prevents any attacks targeted at the user from striking, for the duration of the turn. If a contact move is " +
				"targeted at the user on that turn, the attacker's Defense is lowered by two stages.",
		},
		"dire-claw": {
			effect: "Dire Claw inflicts damage. It has a 50%" + " chance of leaving the target poisoned, paralyzed, or asleep, and " +
				"does not have an increased critical-hit ratio.",
			summary: "Deals damage and has a 50%" + " chance of either poisoning, paralyzing, or putting the target to sleep.",
		},
		"psyshield-bash": {
			effect:  "Psyshield Bash inflicts damage and raises the user's Defense by one stage.",
			summary: "Deals damage and raises the user's Defense by one stage.",
		},
		"power-shift": {
			effect:  "Power Shift switches the user's raw Attack stat with its raw Defense stat. Stat modifiers are not swapped.",
			summary: "Switches the user's Attack stat with its Defense.",
		},
		"stone-axe": {
			effect: "Stone Axe inflicts damage and sets up Stealth Rock on the opposing side of the field. If used by a Pokémon " +
				"with Sharpness, its power is increased by 50%. It does not have an increased critical-hit ratio.",
			summary: "Deals damage and scatters sharp rocks around the opposing field.",
		},
		"springtide-storm": {
			effect:  "Springtide Storm inflicts damage. It has a 30%" + " chance of lowering each target's Attack by one stage.",
			summary: "Deals damage and has a 30%" + " chance of lowering the target's Attack by one stage.",
		},
		"mystical-power": {
			effect:  "Mystical Power inflicts damage and raises the user's Special Attack by one stage.",
			summary: "Deals damage and raises the user's Special Attack by one stage.",
		},
		"raging-fury": {
			effect: "Raging Fury inflicts damage for 2-3 turns, then the user becomes confused due to fatigue. It is disrupted " +
				"if it is not successful due to missing, sleeping, paralysis, freeze, flinching, a protected target, or immunity. If " +
				"a disruption occurs on what would have been the final, confusion-inducing turn of Raging Fury, the user will still " +
				"become confused.",
			summary: "Attacks for 2-3 turns, during which it cannot switch out or change moves, then becomes confused.",
		},
		"wave-crash": {
			effect:  "Wave Crash inflicts damage and the user receives recoil damage equal to 33%" + " of the damage done to the target.",
			summary: "Deals damage, but the user receives 1/3 of the damage it inflicted in recoil.",
		},
		"chloroblast": {
			effect: "Chloroblast inflicts damage, and the user takes damage equal to half of its maximum HP rounded up, unless the " +
				"user has Magic Guard or the target has Sap Sipper.",
			summary: "Deals damage, but the user loses half of its maximum HP.",
		},
		"mountain-gale": {
			effect:  "Mountain Gale inflicts damage and has a 30%" + " chance of causing the target to flinch.",
			summary: "Deals damage and has a 30%" + " chance of causing the target to flinch.",
		},
		"victory-dance": {
			effect:  "Victory Dance raises the user's Attack, Defense, and Speed stats by one stage each.",
			summary: "Raises the user's Attack, Defense and Speed by one stage each.",
		},
		"headlong-rush": {
			effect:  "Headlong Rush inflicts damage and then lowers the user's Defense stat and Special Defense stat by one stage each.",
			summary: "Deals damage, but lowers the user's Defense and Special Defense stats by one stage each after attacking.",
		},
		"barb-barrage": {
			effect:  "Barb Barrage deals damage and has a 50%" + " chance to poison the target. Its power is doubled if the target is already poisoned.",
			summary: "Deals damage and has a 50%" + " chance chance of poisoning the target. If the target is already poisoned, its power is doubled.",
		},
		"esper-wing": {
			effect:  "Esper Wing inflicts damage and has an increased critical hit ratio. It also raises the user's action speed.",
			summary: "Deals damage and raises the user's Speed by one stage. It also has an increased critical hit ratio.",
		},
		"bitter-malice": {
			effect:  "Bitter Malice inflicts damage and lowers the target's Attack stat by one stage.",
			summary: "Deals damage and lowers the target's Attack by one stage.",
		},
		"shelter": {
			effect:  "Shelter raises the user's Defense by two stages.",
			summary: "Raises the user's Defense by two stages.",
		},
		"triple-arrows": {
			effect: "Triple Arrows inflicts damage. For three turns, it increases the user's critical hit ratio and decreases the target's " +
				"defensive stats. It has a 50%" + " chance of lowering the target's Defense by one stage and a 30%" + " chance of causing the target to flinch.",
			summary: "Deals damage and has an increased critical hit ratio. It has a 50%" + " chance of lowering the target's Defense by one stage, " +
				"and has a 30%" + " chance of causing the target to flinch.",
		},
		"infernal-parade": {
			effect: "Infernal Parade inflicts damage and has a 30%" + " chance of leaving a burn. The move's power is doubled " +
				"if the target has a status condition.",
			summary: "Deals damage and has a 30%" + " chance of burning the target. If the target already has a status condition, its power is doubled.",
		},
		"ceaseless-edge": {
			effect: "Ceaseless Edge inflicts damage and sets up one layer of Spikes on the opposing side of the field. " +
				"If used by a Pokémon with Sharpness, its power is increased by 50%. It does not have an increased critical-hit ratio.",
			summary: "Deals damage and lays a trap of thorns on the opposing field.",
		},
		"bleakwind-storm": {
			effect: "Bleakwind Storm inflicts damage and has a 30%" + " chance of lowering the target's Speed stat by one " +
				"stage. When used during rain, Bleakwind Storm bypasses accuracy checks to always hit.",
			summary: "Deals damage and has a 30%" + " chance of lowering the target's Speed by one stage.",
		},
		"wildbolt-storm": {
			effect: "Wildbolt Storm inflicts damage and has a 20%" + " chance of paralyzing the target. When used during " +
				"rain, Wildbolt Storm bypasses accuracy checks to always hit.",
			summary: "Deals damage and has a 20%" + " chance of paralyzing the target.",
		},
		"sandsear-storm": {
			effect: "Sandsear Storm inflicts damage and has a 20%" + " chance of burning the target. When used during rain, " +
				"Sandsear Storm bypasses accuracy checks to always hit.",
			summary: "Deals damage and has a 20%" + " chance of burning the target.",
		},
		"lunar-blessing": {
			effect: "Lunar Blessing restores the HP of the user and its allies by 25%" + " of their maximum HP, and heals any " +
				"non-volatile status conditions they may have.",
			summary: "Heals all status conditions of the user and its teammates, and restores 25%" + " of their maximum HP.",
		},
		"take-heart": {
			effect: "Take Heart raises the user's Special Attack and Special Defense stats by one stage each, and heals any non-volatile " +
				"status conditions the user may have.",
			summary: "Heals the user's status conditions, and raises the user's Special Attack and Special Defense by one stage each.",
		},
		"tera-blast": {
			effect: "Tera Blast deals damage. If the user's Attack stat is higher than its Special Attack stat, Tera Blast becomes a " +
				"physical move when Terastallized; otherwise, it remains a special move. If the user has Terastallized, Tera Blast's type becomes " +
				"the same as their Tera Type, is unaffected by moves and abilities that change its type, and has a different animation depending " +
				"on the move's type.",
			summary: "Deals damage, and its type changes to the user's Tera Type, when the Pokémon is Terastallized.",
		},
		"silk-trap": {
			effect: "Silk Trap protects the user from all effects of physical and special moves that target it during the " +
				"turn it is used. If Silk Trap blocks an attack that would make contact with the user, the attacker's Speed stat " +
				"is lowered by one stage. If the user goes last in the turn, the move will fail.",
			summary: "Protects the user and lowers opponent's Speed on contact.",
		},
		"axe-kick": {
			effect: "Axe Kick inflicts damage and has a 30%" + " chance of confusing the target. If it misses, the user takes " +
				"crash damage equal to half of its maximum HP. The user will also take crash damage if this move is used on a " +
				"protected target, or if the target is immune to this move.",
			summary: "May confuse opponent. If it misses, the user loses HP.",
		},
		"last-respects": {
			effect: "Last Respects inflicts damage. Its base power, which starts at 50, is increased by an additional " +
				"50 each time a Pokémon on the user's side faints during a battle. Revived Pokémon who then faint again also " +
				"count, with the base power capping at 5,050 after Pokémon faint 100 times in the same battle.",
			summary: "Damages increases the more party Pokémon have been defeated.",
		},
		"lumina-crash": {
			effect:  "Lumina Crash inflicts damage and lowers the target's Special Defense by two stages.",
			summary: "Harshly lowers target's Special Defense.",
		},
		"order-up": {
			effect: "If the user has a Tatsugiri in its mouth via the Commander Ability, the move's animation will change and " +
				"the user will gain a one-stage stat boost based on Tatsugiri's form: Curly Form: increases Attack, Droopy Form: " +
				"increases Defense, Stretchy Form: increases Speed.",
			summary: "Deals damage. If the user, Dondozo, has a Tastugiri in its mouth, it also raises a particular stat " +
				"depending on the form of Tatsugiri.",
		},
		"jet-punch": {
			effect:  "Jet Punch inflicts damage. It has a priority of +1, so is used before all moves that do not have increased priority.",
			summary: "Always goes first.",
		},
		"spicy-extract": {
			effect: "Spicy Extract raises the target's Attack and lowers the target's Defense, each by two stages. " +
				"It bypasses accuracy checks to always hit.",
			summary: "Harshly lowers the opponent's Defense and sharply raises their Attack.",
		},
		"spin-out": {
			effect:  "Spin Out inflicts damage and lowers the user's Speed by two stages.",
			summary: "Harshly lowers user's Speed.",
		},
		"population-bomb": {
			effect: "Population Bomb inflicts damage, hitting the target up to ten times per use. Each of " +
				"Population Bomb's strikes has a separate accuracy check. If any strike fails, the move ends. " +
				"If the user is holding Loaded Dice, only one accuracy check is performed, and it will hit between " +
				"four and ten times.  the user has Skill Link, the move will have only one accuracy check and hit " +
				"the maximum ten times.",
			summary: "Hits 1-10 times in a row.",
		},
		"ice-spinner": {
			effect:  "Ice Spinner inflicts damage and removes any terrain present on the battlefield.",
			summary: "Removes effects of Terrain.",
		},
		"glaive-rush": {
			effect: "Glaive Rush inflicts damage. If the move hits, until the next time the user acts, any attacks " +
				"directed at the user deal double the damage and bypass accuracy and evasion checks to always hit.",
			summary: "Attacks from opposing Pokémon during the next turn cannot miss and will inflict double damage.",
		},
		"revival-blessing": {
			effect: "Revival Blessing allows the player to select a fainted Pokémon in the party; the selected " +
				"Pokémon will be revived, and half of its maximum HP will be restored.",
			summary: "Revives a fainted party Pokémon to half HP.",
		},
		"salt-cure": {
			effect: "Salt Cure inflicts 1/8 of the target's maximum HP as damage per turn in addition to the damage " +
				"dealt when it is used. If a Steel and/or Water type is affected by Salt Cure, the amount of damage " +
				"per turn is ¼ of its maximum HP.",
			summary: "Deals damage each turn; Steel and Water types are more affected.",
		},
		"triple-dive": {
			effect:  "Triple Dive inflicts damage, hitting the target three times per use.",
			summary: "Hits 3 times in a row.",
		},
		"mortal-spin": {
			effect: "Mortal Spin inflicts damage and then removes any binding effects and Leech Seed from the user, as well as " +
				"entry hazards from the user's side of the field. It also leaves each target poisoned.",
			summary: "Removes entry hazards and trap move effects, and poisons opposing Pokémon.",
		},
		"doodle": {
			effect:  "Doodle changes the Abilities of the user and any ally to the targeted Pokémon's Ability.",
			summary: "Changes the abilities of the user and its teammates to that of the target.",
		},
		"fillet-away": {
			effect: "Fillet Away deducts half of the user's maximum HP, rounded down, from its current HP and, in " +
				"return, it raises the user's Attack, Special Attack, and Speed stats by two stages each. It fails if the " +
				"user's current HP is less than or equal to half its maximum, or if the user's aforementioned stats are all at +6 stages.",
			summary: "Lowers HP but sharply boosts Attack, Special Attack, and Speed.",
		},
		"kowtow-cleave": {
			effect: "Kowtow Cleave inflicts damage and bypasses accuracy checks to always hit.  used by a Pokémon with " +
				"Sharpness, its power is increased by 50%.",
			summary: "Always hits.",
		},
		"flower-trick": {
			effect:  "Flower Trick inflicts damage and bypasses accuracy checks to always hit, and will always result in a critical hit.",
			summary: "Never misses; always results in a critical hit.",
		},
		"torch-song": {
			effect:  "Torch Song inflicts damage and raises the user's Special Attack by one stage.",
			summary: "Raises user's Special Attack.",
		},
		"aqua-step": {
			effect:  "Aqua Step inflicts damage and raises the user's Speed by one stage if the move is successful.",
			summary: "Raises user's Speed.",
		},
		"raging-bull": {
			effect: "Raging Bull removes Light Screen, Reflect, and Aurora Veil from the target's side of the field, " +
				"then inflicts damage. If the move fails to hit the target, the effects of the screens are not removed. " +
				"If the user is Paldean Tauros, the move's type changes depending on its form: Combat Breed: Fighting type, " +
				"Blaze Breed: Fire type, and Aqua Breed: Water type.",
			summary: "Type depends on the user's form. Breaks through Reflect and Light Screen barriers.",
		},
		"make-it-rain": {
			effect: "Make It Rain inflicts damage and lowers the user's Special Attack by one stage. For each target the move " +
				"hits, coins are scattered on the ground. $5 multiplied by the user's level are scattered each time.",
			summary: "Lowers user's Special Attack. Money is earned after the battle.",
		},
		"psyblade": {
			effect: "Psyblade inflicts damage. If there is Electric Terrain on the field, its base power is increased to 120, " +
				"whether or not the user is grounded. If used by a Pokémon with Sharpness, its power is increased by 50%.",
			summary: "Deals damage, and its power is boosted by 50%" + " on Electric Terrain.",
		},
		"hydro-steam": {
			effect: "Hydro Steam inflicts damage. If used under harsh sunlight, it receives a 50%" + " damage boost, unlike all other " +
				"Water-type moves which would receive a 50%" + " damage reduction in harsh sunlight instead.",
			summary: "Deals damage, and its power is boosted by 50% in harsh sunlight, instead of being reduced.",
		},
		"ruination": {
			effect:  "Ruination deals damage equal to 50%" + " of the opponent's remaining HP, rounded down. It always deals at least 1 HP of damage.",
			summary: "Halves the opponent's HP.",
		},
		"collision-course": {
			effect:  "Collision Course inflicts damage. If the move is super effective on the target, its damage is increased by roughly 33%.",
			summary: "Boosted even more if it's super-effective.",
		},
		"electro-drift": {
			effect:  "Electro Drift inflicts damage. If the move is super effective on the target, its damage is increased by roughly 33%.",
			summary: "Boosted even more if it's super-effective.",
		},
		"shed-tail": {
			effect: "Shed Tail creates a substitute at the cost of half of the user's maximum HP (rounded up), then switches the user out. " +
				"If the user is holding a Berry that would trigger in response to Shed Tail, it eats its held Berry before switching out.",
			summary: "Creates a substitute, then swaps places with a party Pokémon in waiting.",
		},
		"chilly-reception": {
			effect: "Chilly Reception creates snow on the battlefield and switches the user out. The user will still switch out even if " +
				"there is already snow on the field when the move is used. If there are no other teammates to switch to, it will only " +
				"change the weather to snow",
			summary: "Switches out and summons a snowstorm lasting 5 turns.",
		},
		"tidy-up": {
			effect: "Tidy Up raises the user's Attack stat and Speed stat by one stage each. It also clears away substitutes, Spikes, " +
				"Toxic Spikes, Stealth Rock, and Sticky Web on both the user's and target's side of the field. The user's stats will still " +
				"be raised even if there are no traps or substitutes to remove.",
			summary: "Removes the effects of entry hazards and Substitute, and boosts user's Attack and Speed.",
		},
		"snowscape": {
			effect: "When Snowscape is used, a snowstorm will begin on the field. This effect will last for 5 turns. This clears any other type of " +
				"weather, but will fail if there is already a snowstorm. All Ice-type Pokémon have their Defense raised by 50%" + " for the " +
				"duration of the snowstorm.",
			summary: "Raises Defense of Ice types for 5 turns.",
		},
		"pounce": {
			effect:  "Pounce inflicts damage and lowers the target's Speed by one stage.",
			summary: "Lowers opponent's Speed.",
		},
		"trailblaze": {
			effect:  "Trailblaze inflicts damage and raises the user's Speed by one stage.",
			summary: "Deals damage and raises the user's speed by one stage.",
		},
		"chilling-water": {
			effect:  "Chilling Water inflicts damage and lowers the target's Attack by one stage.",
			summary: "Lowers opponent's Attack.",
		},
		"hyper-drill": {
			effect: "Hyper Drill inflicts damage. It bypasses the effects of Protect, Detect, Spiky Shield, King's Shield, Baneful Bunker, Silk Trap, " +
				"and Burning Bulwark, but does not lift the effects of these moves.",
			summary: "Can strike through Protect/Detect.",
		},
		"twin-beam": {
			effect: "Twin Beam inflicts damage, hitting twice per use. Each strike has an equal chance to be a critical hit. Twin Beam will hit again if " +
				"the first strike breaks a substitute. Mirror Coat will only acknowledge the second strike of this move.",
			summary: "Hits twice in one turn.",
		},
		"rage-fist": {
			effect: "Rage Fist inflicts damage. Its power increases by 50 each time the user is hit during the battle by a move that causes direct damage, " +
				"up to a maximum of 350. This includes moves used by either allies or opponents and includes moves that do not actually deplete HP. However, " +
				"the power of Rage Fist does not increase if the user attacks itself due to confusion, nor does it increase when the user's substitute is struck. " +
				"Each strike from a multistrike move counts as one hit, and Transform copies the target's hits counter.",
			summary: "The more times the user has been hit by attacks, the greater the move's power.",
		},
		"armor-cannon": {
			effect:  "Armor Cannon deals damage and lowers the user's Defense and Special Defense stats by one stage each.",
			summary: "Deals damage, but lowers the user's Defense and Special Defense stats by one stage each after attacking.",
		},
		"bitter-blade": {
			effect: "Bitter Blade inflicts damage, and it restores the user's HP by half of the damage dealt. " +
				"If used by a Pokémon with Sharpness, its power is increased by 50%.",
			summary: "Deals damage and the user will recover 50%" + " of the damage taken by the target.",
		},
		"double-shock": {
			effect: "If used by an Electric-type Pokémon, Double Shock will inflict damage to the target and cause the user to lose its Electric type. " +
				"Double Shock fails if the user is not Electric-type, making the move always fail after the first use until the Pokémon regains its Electric type.",
			summary: "After using this move, the user will no longer be Electric type.",
		},
		"gigaton-hammer": {
			effect: "Gigaton Hammer inflicts damage. The user cannot select Gigaton Hammer in succession unless the move failed. " +
				"If the user cannot select another move, it will use Struggle every other turn.",
			summary: "Cannot be used twice in a row.",
		},
		"comeuppance": {
			effect: "Comeuppance returns 1.5 times the damage dealt by the foe's last attack. Unlike Counter or Mirror Coat, Comeuppance does not require " +
				"specifically physical or special damage sources and Comeuppance does not have decreased priority. If the user acts before it is hit by an " +
				"opponent's damaging move, Comeuppance will fail. The move also fails if the user's substitute is hit instead.",
			summary: "Deals more damage to the opponent that last inflicted damage on it.",
		},
		"aqua-cutter": {
			effect:  "Aqua Cutter inflicts damage and has an increased critical hit ratio. If used by a Pokémon with Sharpness, its power is increased by 50%.",
			summary: "Deals damage and has an increased critical-hit ratio.",
		},
		"blazing-torque": {
			effect:  "Deals damage and has a 30%" + " chance of burning the target.",
			summary: "Deals damage and has a 30%" + " chance of burning the target.",
		},
		"wicked-torque": {
			effect:  "Deals damage and has a 10%" + " chance of putting the target to sleep.",
			summary: "Deals damage and has a 10%" + " chance of putting the target to sleep.",
		},
		"noxious-torque": {
			effect:  "Deals damage and has a 30%" + " chance of poisoning the target.",
			summary: "Deals damage and has a 30%" + " chance of poisoning the target.",
		},
		"combat-torque": {
			effect:  "Deals damage and has a 30%" + " chance of paralyzing the target.",
			summary: "Deals damage and has a 30%" + " chance of paralyzing the target.",
		},
		"magical-torque": {
			effect:  "Deals damage and has a 30%" + " chance of confusing the target.",
			summary: "Deals damage and has a 30%" + " chance of confusing the target.",
		},
		"blood-moon": {
			effect: "Blood Moon inflicts damage. The user cannot select Blood Moon in succession unless the move failed. If the user cannot select another move, " +
				"it will use Struggle every other turn.",
			summary: "Cannot be used twice in a row.",
		},
		"matcha-gotcha": {
			effect:  "Matcha Gotcha deals damage and restores the user's HP equal to up to half of the damage dealt. It also has a 20%" + " chance of burning targets.",
			summary: "The user's HP is restored by up to half the damage taken by the target. This may also leave the target with a burn.",
		},
		"syrup-bomb": {
			effect: "Syrup Bomb inflicts damage and lowers the target's Speed stat by one stage each turn, for the next three turns. " +
				"If the move's user leaves the field, the effect will end immediately",
			summary: "Lowers opponent's Speed each turn for 3 turns.",
		},
		"ivy-cudgel": {
			effect: "Ivy Cudgel deals damage and has an increased critical hit ratio. The move's type changes depending on Ogerpon's held mask: " +
				"Grass type in Teal Mask Form, Water type in Wellspring Mask Form, Fire type in Hearthflame Mask Form, and Rock type in Cornerstone Mask Form. " +
				"If used by a Pokémon other than Ogerpon, it will always be Grass type.",
			summary: "deals damage and has an increased critical hit ratio. Its type changes based on the mask worn by Ogerpon.",
		},
		"electro-shot": {
			effect: "On the turn it is selected, Electro Shot will raise the user's Special Attack stat by one stage, and attack on the next turn. " +
				"Once Electro Shot is selected, the user won't be able to switch out until it is disrupted or fully executed. Sleep, freezing, and flinching " +
				"will pause but not disrupt the duration of Electro Shot. It does not need a turn to charge in the rain.",
			summary: "Charges on the first turn, raising Special Attack, and attacks on the second turn. It does not need to charge in rain.",
		},
		"tera-starstorm": {
			effect:  "When used by Terapagos, it becomes Stellar type and damages all opponents.",
			summary: "When used by Terapagos, it becomes Stellar type and damages all opponents.",
		},
		"fickle-beam": {
			effect: "Fickle Beam inflicts damage. There is a 70%" + " chance for the move to inflict damage at 80 base power, and a 30&%" + " chance " +
				"for it to double and inflict damage at 160 base power.",
			summary: "Fickle Beam deals damage with 80 base power, with a 30%" + " chance for the power to double to 160.",
		},
		"burning-bulwark": {
			effect: "Burning Bulwark protects the user from all effects of physical and special moves that target it during the turn it is used. " +
				"If Burning Bulwark blocks an attack that would make contact with the user, the attacker becomes burned.",
			summary: "The user is protected from all damage, and burns any opponents who would have made contact.",
		},
		"thunderclap": {
			effect: "Thunderclap has a priority of +1. If Thunderclap's target has selected a physical or special attack and has not yet had " +
				"a chance to execute its move, Thunderclap will damage the target; otherwise, Thunderclap will fail.",
			summary: "Strikes before a target's move, fails if the target doesn't attack.",
		},
		"mighty-cleave": {
			effect: "Mighty Cleave inflicts damage. It bypasses the effects of Protect, Detect, Spiky Shield, King's Shield, Baneful Bunker, " +
				"Silk Trap, and Burning Bulwark, but does not lift the effects of these moves.",
			summary: "Bypasses Protect and other protection moves.",
		},
		"tachyon-cutter": {
			effect:  "Tachyon Cutter inflicts damage, hitting twice per use. It bypasses accuracy checks to always hit.",
			summary: "Guaranteed to hit twice in a row, never missing.",
		},
		"hard-press": {
			effect: "Hard Press inflicts damage. Its power varies between 1 and 100, and is greater the more HP the target has. The power is the target's " +
				"current health, divided by its total health, and multiplied by 100.",
			summary: "The target is crushed with an arm, a claw, or the like to inflict damage. The more HP the target has left, the greater the move's power.",
		},
		"dragon-cheer": {
			effect: "Dragon Cheer increases the user's allies' critical hit ratio by one stage. If the affected Pokémon are Dragon-type, their critical " +
				"hit ratio is increased by two stages instead.",
			summary: "The user raises morale and gives allies a higher critical hit chance. This rouses Dragons even more.",
		},
		"alluring-voice": {
			effect: "Alluring Voice inflicts damage. The move also confuses the target if the target's stats were raised during the same turn before " +
				"the execution of Alluring Voice.",
			summary: "Confuses the target if their stats have changed earlier during the turn.",
		},
		"temper-flare": {
			effect: "If, on the previous turn, the user's last move missed, failed to affect each of its targets, or was prevented from being used due to " +
				"an effect such as paralysis, Temper Flare's power is doubled from 75 to 150.",
			summary: "Doubles in power if the previous move failed.",
		},
		"supercell-slam": {
			effect: "Supercell Slam deals damage. If the user misses, the user will take crash damage equal to half of the user's max HP, rounded down. " +
				"If the target is immune to Supercell Slam by type or Ability (such as Volt Absorb) or is protected, the move counts as a miss.",
			summary: "If the move misses, the user takes damage instead.",
		},
		"psychic-noise": {
			effect: "Psychic Noise deals damage and prevents the target from recovering HP for two turns. During these two turns, if the target uses a move " +
				"with a healing property, such as Recover or Drain Punch, it will fail.",
			summary: "Deals damage and prevents target from healing.",
		},
		"upper-hand": {
			effect: "If the target has selected a physical or special move with increased priority from +1 to +3 and has not yet had a chance to execute its " +
				"move, Upper Hand will strike first and cause the target to flinch; otherwise, Upper Hand will fail.",
			summary: "Strikes before a target's priority move.",
		},
		"malignant-chain": {
			effect:  "Malignant Chain inflicts damage and has a 50%" + " chance of badly poisoning the target.",
			summary: "May also leave the target badly poisoned.",
		},
	}

	if e, found := moveEffctMap[search]; found {
		return e.effect, e.summary
	}

	return "", ""
}
