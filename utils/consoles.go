package utils

import "fmt"

//goland:noinspection GoUnusedConst
const (
	ColorReset   string = "\033[0m"
	ColorRed     string = "\033[31m"
	ColorGreen   string = "\033[32m"
	ColorYellow  string = "\033[33m"
	ColorBlue    string = "\033[34m"
	ColorMagenta string = "\033[35m"
	ColorCyan    string = "\033[36m"
	ColorWhite   string = "\033[37m"
	ColorGray    string = "\033[90m"
)

func color(val string, col string, startReset string) string {
	return fmt.Sprintf("%s%s%s%s", startReset, col, val, ColorReset)
}

//goland:noinspection GoUnusedExportedFunction
func Gray(val string) string {
	return color(val, ColorGray, "")
}

//goland:noinspection GoUnusedExportedFunction
func Red(val string) string {
	return color(val, ColorRed, "")
}

//goland:noinspection GoUnusedExportedFunction
func Green(val string) string {
	return color(val, ColorGreen, "")
}

//goland:noinspection GoUnusedExportedFunction
func Yellow(val string) string {
	return color(val, ColorYellow, "")
}

//goland:noinspection GoUnusedExportedFunction
func Blue(val string) string {
	return color(val, ColorBlue, "")
}

//goland:noinspection GoUnusedExportedFunction
func Magenta(val string) string {
	return color(val, ColorMagenta, "")
}

//goland:noinspection GoUnusedExportedFunction
func Cyan(val string) string {
	return color(val, ColorCyan, "")
}

//goland:noinspection GoUnusedExportedFunction
func White(val string) string {
	return color(val, ColorWhite, "")
}

//goland:noinspection GoUnusedExportedFunction
func RRed(val string) string {
	return color(val, ColorRed, ColorReset)
}

//goland:noinspection GoUnusedExportedFunction
func RGreen(val string) string {
	return color(val, ColorGreen, ColorReset)
}

//goland:noinspection GoUnusedExportedFunction
func RYellow(val string) string {
	return color(val, ColorYellow, ColorReset)
}

//goland:noinspection GoUnusedExportedFunction
func RBlue(val string) string {
	return color(val, ColorBlue, ColorReset)
}

//goland:noinspection GoUnusedExportedFunction
func RMagenta(val string) string {
	return color(val, ColorMagenta, ColorReset)
}

//goland:noinspection GoUnusedExportedFunction
func RCyan(val string) string {
	return color(val, ColorCyan, ColorReset)
}

//goland:noinspection GoUnusedExportedFunction
func RWhite(val string) string {
	return color(val, ColorWhite, ColorReset)
}
