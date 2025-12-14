package ui

import "fmt"

// ANSI color codes
const (
	Reset       = "\033[0m"
	Green       = "\033[32m"
	BrightGreen = "\033[92m"
	Cyan        = "\033[36m"
	BrightCyan  = "\033[96m"
	Gray        = "\033[90m"
	White       = "\033[97m"
	Bold        = "\033[1m"
	Dim         = "\033[2m"
)

func PrintBanner() {
	banner := fmt.Sprintf(`
%s%s
    ███╗   ██╗███████╗████████╗███████╗██╗  ██╗██╗██╗   ██╗
    ████╗  ██║██╔════╝╚══██╔══╝██╔════╝██║  ██║██║██║   ██║
    ██╔██╗ ██║█████╗     ██║   ███████╗███████║██║██║   ██║
    ██║╚██╗██║██╔══╝     ██║   ╚════██║██╔══██║██║╚██╗ ██╔╝
    ██║ ╚████║███████╗   ██║   ███████║██║  ██║██║ ╚████╔╝ 
    ╚═╝  ╚═══╝╚══════╝   ╚═╝   ╚══════╝╚═╝  ╚═╝╚═╝  ╚═══╝  
%s
    %s[%s═══════════════════════════════════════════════════════%s]%s
    %s│%s          Network Penetration & Analysis Tool          %s│%s
    %s│%s                    %sv1.0.0-alpha%s                       %s│%s
    %s[%s═══════════════════════════════════════════════════════%s]%s

%s[%s>%s]%s CIDR parser ready...
%s[%s>%s]%s Codec modules loaded...
%s[%s✓%s]%s Awaiting target range.
%s`,
		Bold, BrightGreen,
		Reset,
		Gray, Cyan, Gray, Reset,
		Gray, White, Gray, Reset,
		Gray, Dim, Green, Reset, Gray, Reset,
		Gray, Cyan, Gray, Reset,
		Gray, BrightGreen, Gray, Reset,
		Gray, BrightGreen, Gray, Reset,
		Gray, BrightGreen, Gray, Reset,
		Reset,
	)
	fmt.Println(banner)
}

func Banner() {
	PrintBanner()
}
