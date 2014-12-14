package limitless

//Colors of the light
var Colors map[string]string

//List of commands
var Commands map[string]string

func init() {
	Colors = map[string]string{
		"violet":        "\x00",
		"royal_blue":    "\x10",
		"baby_blue":     "\x20",
		"aqua":          "\x30",
		"mint":          "\x40",
		"seafoam_green": "\x50",
		"green":         "\x60",
		"lime_green":    "\x70",
		"yellow":        "\x80",
		"yellow_orange": "\x90",
		"orange":        "\xA0",
		"red":           "\xB0",
		"pink":          "\xC0",
		"fusia":         "\xD0",
		"lilac":         "\xE0",
		"lavendar":      "\xF0",
	}

	Commands = map[string]string{
		"color":         "\x40",
		"group_all_off": "\x41",
		"group_all_on":  "\x42",
		"disco_slow":    "\x43",
		"disco_faster":  "\x44",
		"group_1_on":    "\x45",
		"group_1_off":   "\x46",
		"group_2_on":    "\x47",
		"group_2_off":   "\x48",
		"group_3_on":    "\x49",
		"group_3_off":   "\x4A",
		"group_4_on":    "\x4B",
		"group_4_off":   "\x4C",
		"disco":         "\x4D",
		"all_white":     "\xC2",
		"1_white":       "\xC5",
		"2_white":       "\xC7",
		"3_white":       "\xC9",
		"4_white":       "\xCB",
	}
}
