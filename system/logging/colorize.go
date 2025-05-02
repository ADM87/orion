package logging

import "fmt"

const (
	// Regular colors
	COLOR_RED     = "\033[31m%s\033[0m"
	COLOR_GREEN   = "\033[32m%s\033[0m"
	COLOR_YELLOW  = "\033[33m%s\033[0m"
	COLOR_BLUE    = "\033[34m%s\033[0m"
	COLOR_MAGENTA = "\033[35m%s\033[0m"
	COLOR_CYAN    = "\033[36m%s\033[0m"
	COLOR_WHITE   = "\033[37m%s\033[0m"
	COLOR_GREY    = "\033[38m%s\033[0m"

	// Bold colors
	COLOR_RED_BOLD     = "\033[1;31m%s\033[0m"
	COLOR_GREEN_BOLD   = "\033[1;32m%s\033[0m"
	COLOR_YELLOW_BOLD  = "\033[1;33m%s\033[0m"
	COLOR_BLUE_BOLD    = "\033[1;34m%s\033[0m"
	COLOR_MAGENTA_BOLD = "\033[1;35m%s\033[0m"
	COLOR_CYAN_BOLD    = "\033[1;36m%s\033[0m"
	COLOR_WHITE_BOLD   = "\033[1;37m%s\033[0m"
	COLOR_GREY_BOLD    = "\033[1;38m%s\033[0m"
)

func Colorize(color string, args ...any) string {
	return fmt.Sprintf(color, args...)
}

func Red(args ...any) string {
	return Colorize(COLOR_RED, args...)
}

func Redf(format string, args ...any) string {
	return Colorize(COLOR_RED, fmt.Sprintf(format, args...))
}

func Green(args ...any) string {
	return Colorize(COLOR_GREEN, args...)
}

func Greenf(format string, args ...any) string {
	return Colorize(COLOR_GREEN, fmt.Sprintf(format, args...))
}

func Yellow(args ...any) string {
	return Colorize(COLOR_YELLOW, args...)
}

func Yellowf(format string, args ...any) string {
	return Colorize(COLOR_YELLOW, fmt.Sprintf(format, args...))
}

func Blue(args ...any) string {
	return Colorize(COLOR_BLUE, args...)
}

func Bluef(format string, args ...any) string {
	return Colorize(COLOR_BLUE, fmt.Sprintf(format, args...))
}

func Magenta(args ...any) string {
	return Colorize(COLOR_MAGENTA, args...)
}

func Magentaf(format string, args ...any) string {
	return Colorize(COLOR_MAGENTA, fmt.Sprintf(format, args...))
}

func Cyan(args ...any) string {
	return Colorize(COLOR_CYAN, args...)
}

func Cyanf(format string, args ...any) string {
	return Colorize(COLOR_CYAN, fmt.Sprintf(format, args...))
}

func White(args ...any) string {
	return Colorize(COLOR_WHITE, args...)
}

func Whitef(format string, args ...any) string {
	return Colorize(COLOR_WHITE, fmt.Sprintf(format, args...))
}

func Grey(args ...any) string {
	return Colorize(COLOR_GREY, args...)
}

func Greyf(format string, args ...any) string {
	return Colorize(COLOR_GREY, fmt.Sprintf(format, args...))
}

func RedBold(args ...any) string {
	return Colorize(COLOR_RED_BOLD, args...)
}

func RedBoldf(format string, args ...any) string {
	return Colorize(COLOR_RED_BOLD, fmt.Sprintf(format, args...))
}

func GreenBold(args ...any) string {
	return Colorize(COLOR_GREEN_BOLD, args...)
}

func GreenBoldf(format string, args ...any) string {
	return Colorize(COLOR_GREEN_BOLD, fmt.Sprintf(format, args...))
}

func YellowBold(args ...any) string {
	return Colorize(COLOR_YELLOW_BOLD, args...)
}

func YellowBoldf(format string, args ...any) string {
	return Colorize(COLOR_YELLOW_BOLD, fmt.Sprintf(format, args...))
}

func BlueBold(args ...any) string {
	return Colorize(COLOR_BLUE_BOLD, args...)
}

func BlueBoldf(format string, args ...any) string {
	return Colorize(COLOR_BLUE_BOLD, fmt.Sprintf(format, args...))
}

func MagentaBold(args ...any) string {
	return Colorize(COLOR_MAGENTA_BOLD, args...)
}

func MagentaBoldf(format string, args ...any) string {
	return Colorize(COLOR_MAGENTA_BOLD, fmt.Sprintf(format, args...))
}

func CyanBold(args ...any) string {
	return Colorize(COLOR_CYAN_BOLD, args...)
}

func CyanBoldf(format string, args ...any) string {
	return Colorize(COLOR_CYAN_BOLD, fmt.Sprintf(format, args...))
}

func WhiteBold(args ...any) string {
	return Colorize(COLOR_WHITE_BOLD, args...)
}

func WhiteBoldf(format string, args ...any) string {
	return Colorize(COLOR_WHITE_BOLD, fmt.Sprintf(format, args...))
}

func GreyBold(args ...any) string {
	return Colorize(COLOR_GREY_BOLD, args...)
}

func GreyBoldf(format string, args ...any) string {
	return Colorize(COLOR_GREY_BOLD, fmt.Sprintf(format, args...))
}
