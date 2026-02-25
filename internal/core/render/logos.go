package render

var defaultLogo = []string{
	"  ██████╗ ███████╗██████╗ ███████╗",
	"██╔═══██╗██╔════╝██╔══██╗██╔════╝",
	"██║   ██║███████╗██████╔╝███████╗ ",
	"██║   ██║╚════██║██╔══██╗╚════██║",
	" ╚██████╔╝███████║██║  ██║███████║",
	"  ╚═════╝ ╚══════╝╚═╝  ╚═╝╚══════╝",
}

// GetLogo returns the hardcoded logo
func GetLogo() []string {
	logoLines := defaultLogo
	return logoLines
}
