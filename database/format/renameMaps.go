package frmt

// The various "version" values are all very long and I want them to be
//	condensed before they are uploaded to the database

func RenameLmoveVersion(input string) string {
	versionMap := map[string]string{
		"scarlet-violet":                      "sv",
		"black-2-white-2":                     "b2w2",
		"black-white":                         "bw",
		"brilliant-diamond-and-shining-pearl": "bdsp",
		"diamond-pearl":                       "dp",
		"heartgold-soulsilver":                "hgss",
		"omega-ruby-alpha-sapphire":           "oras",
		"platinum":                            "plat",
		"sun-moon":                            "sm",
		"ultra-sun-ultra-moon":                "usum",
		"x-y":                                 "xy",
		"sword-shield":                        "swsh",
		"colosseum":                           "clsm",
		"crystal":                             "crstl",
		"emerald":                             "emrld",
		"firered-leafgreen":                   "frlg",
		"gold-silver":                         "gs",
		"ruby-sapphire":                       "rs",
		"xd":                                  "xd",
		"red-blue":                            "rb",
		"yellow":                              "yllw",
		"lets-go-pikachu-lets-go-eevee":       "letsgo",
	}

	if match, ok := versionMap[input]; ok {
		return match
	}

	return input
}

func RenamePdexVersion(input string) string {
	versionMap := map[string]string{
		"alpha-sapphire":  "a_sapph",
		"black":           "black",
		"black-2":         "blck2",
		"blue":            "blue",
		"crystal":         "crstl",
		"diamond":         "dimnd",
		"emerald":         "emrld",
		"firered":         "f_red",
		"gold":            "gold",
		"heartgold":       "h_gold",
		"leafgreen":       "l_grn",
		"lets-go-eevee":   "lg_e",
		"lets-go-pikachu": "lg_p",
		"omega-ruby":      "o_rby",
		"pearl":           "prl",
		"platinum":        "plat",
		"red":             "red",
		"ruby":            "ruby",
		"sapphire":        "sapph",
		"shield":          "shld",
		"silver":          "slvr",
		"soulsilver":      "s_slvr",
		"sword":           "swrd",
		"white":           "whte",
		"white-2":         "whte2",
		"x":               "x",
		"y":               "y",
		"yellow":          "yllw",
		"moon":            "moon",
		"sun":             "sun",
		"ultra-moon":      "u_moon",
		"ultra-sun":       "u_sun",
		"legends-arceus":  "lgnd_arc",
		"scarlet":         "scrlt",
		"violet":          "violet",
	}

	if match, ok := versionMap[input]; ok {
		return match
	}

	return input
}

func RenameAdescVersion(input string) string {
	versionMap := map[string]string{
		"black-2-white-2":               "b2w2",
		"black-white":                   "bw",
		"diamond-pearl":                 "dp",
		"emerald":                       "emrld",
		"firered-leafgreen":             "frlg",
		"heartgold-soulsilver":          "hgss",
		"lets-go-pikachu-lets-go-eevee": "letgo",
		"omega-ruby-alpha-sapphire":     "oras",
		"platinum":                      "plat",
		"ruby-sapphire":                 "rs",
		"scarlet-violet":                "sv",
		"sun-moon":                      "sm",
		"sword-shield":                  "swsh",
		"ultra-sun-ultra-moon":          "usum",
		"x-y":                           "xy",
		"the-indigo-disk":               "sv",
		"the-teal-mask":                 "sv",
	}

	if match, ok := versionMap[input]; ok {
		return match
	}

	return input
}

func RenameMdescVersion(input string) string {
	versionMap := map[string]string{
		"black-2-white-2":               "b2w2",
		"black-white":                   "bw",
		"crystal":                       "crstl",
		"diamond-pearl":                 "dp",
		"emerald":                       "emrld",
		"firered-leafgreen":             "frlg",
		"gold-silver":                   "gs",
		"heartgold-soulsilver":          "hgss",
		"lets-go-pikachu-lets-go-eevee": "letsgo",
		"omega-ruby-alpha-sapphire":     "oras",
		"platinum":                      "plat",
		"ruby-sapphire":                 "rs",
		"sun-moon":                      "sm",
		"sword-shield":                  "swsh",
		"ultra-sun-ultra-moon":          "usum",
		"x-y":                           "xy",
	}

	if match, ok := versionMap[input]; ok {
		return match
	}

	return input
}

func RenameMachineVersion(input string) string {
	versionMap := map[string]string{
		"black-2-white-2":                     "b2w2",
		"black-white":                         "bw",
		"brilliant-diamond-and-shining-pearl": "bdsp",
		"colosseum":                           "clsm",
		"crystal":                             "crstl",
		"diamond-pearl":                       "dp",
		"emerald":                             "emrld",
		"firered-leafgreen":                   "frlg",
		"gold-silver":                         "gs",
		"heartgold-soulsilver":                "hgss",
		"lets-go-pikachu-lets-go-eevee":       "letgo",
		"omega-ruby-alpha-sapphire":           "oras",
		"platinum":                            "plat",
		"red-blue":                            "rb",
		"ruby-sapphire":                       "rs",
		"scarlet-violet":                      "sv",
		"sun-moon":                            "sm",
		"sword-shield":                        "swsh",
		"the-indigo-disk":                     "sv",
		"the-teal-mask":                       "sv",
		"ultra-sun-ultra-moon":                "usum",
		"x-y":                                 "xy",
		"xd":                                  "xd",
		"yellow":                              "yllw",
	}

	if match, ok := versionMap[input]; ok {
		return match
	}

	return input
}

func RenameMoveTarget(input string) string {
	targetMap := map[string]string{
		"selected-pokemon":          "selected",
		"all-opponents":             "all-opps",
		"random-opponent":           "rndm-opp",
		"users-field":               "usr-field",
		"all-other-pokemon":         "all-others",
		"specific-move":             "specific",
		"entire-field":              "field",
		"all-pokemon":               "all",
		"selected-pokemon-me-first": "me-first",
		"all-allies":                "allies",
		"fainting-pokemon":          "fainting",
	}

	if match, ok := targetMap[input]; ok {
		return match
	}

	return input
}
