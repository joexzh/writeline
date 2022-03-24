package writeline

// https://misc.flogisoft.com/bash/tip_colors_and_formatting

const (
	Bold      = "\033[1m"
	Dim       = "\033[2m"
	Underline = "\033[4m"
	Blink     = "\033[5m"
	// Reverse invert the foreground and background colors
	Reverse = "\033[7m"
	Hidden  = "\033[8m"

	ResetAllStyle = "\033[0m"

	Default      = "\033[39m"
	Black        = "\033[30m"
	Red          = "\033[31m"
	Green        = "\033[32m"
	Yellow       = "\033[33m"
	Blue         = "\033[34m"
	Magenta      = "\033[35m"
	Cyan         = "\033[36m"
	LightGray    = "\033[37m"
	DarkGray     = "\033[90m"
	LightRed     = "\033[91m"
	LightGreen   = "\033[92m"
	LightYellow  = "\033[93m"
	LightBlue    = "\033[94m"
	LightMagenta = "\033[95m"
	LightCyan    = "\033[96m"
	White        = "\033[97m"

	BgDefault      = "\033[49m"
	BgBlack        = "\033[40m"
	BgRed          = "\033[41m"
	BgGreen        = "\033[42m"
	BgYellow       = "\033[43m"
	BgBlue         = "\033[44m"
	BgMagenta      = "\033[45m"
	BgCyan         = "\033[46m"
	BgLightGray    = "\033[47m"
	BgDarkGray     = "\033[100m"
	BgLightRed     = "\033[101m"
	BgLightGreen   = "\033[102m"
	BgLightYellow  = "\033[103m"
	BgLightBlue    = "\033[104m"
	BgLightMagenta = "\033[105m"
	BgLightCyan    = "\033[106m"
	BgWhite        = "\033[107m"
)

func Style(sType string, s string) string {
	return sType + s + ResetAllStyle
}
