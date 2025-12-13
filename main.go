package main

import (
	"flag"
	"fmt"
)

var (
	Version = "v1.0.0-dev" // manually updated for releases
	Commit  = "none"       // auto-injected at build time
	Date    = "unknown"    // auto-injected at build time
)

func showInstructions() {
	instructions := fmt.Sprintf(`
%s%sUSAGE:%s
    netshiv [OPTIONS]

%s%sOPTIONS:%s
    %s-range%s <ip/cidr>      Expand IP range
                          Example: -range 192.168.0.1/24

    %s-expand%s <ip/cidr>     Summarize CIDR information
                          Example: -expand 192.168.0.1/24

    %s-binary%s <ip>          Convert IP address to binary
                          Example: -binary 192.168.0.1

    %s-encode%s <string>      Encode a string (requires -type)
                          Example: -encode "Hello" -type base64

    %s-decode%s <string>      Decode a string (requires -type)
                          Example: -decode "SGVsbG8=" -type base64

    %s-type%s <format>        Encoding type: hex, base64, url
                          Example: -type base64

    %s-version%s              Show tool version

%s%sEXAMPLES:%s
    netshiv -range 192.168.1.0/24
    netshiv -binary 192.168.1.1
    netshiv -encode "secret" -type base64
    netshiv -decode "c2VjcmV0" -type base64
`,
		Bold, BrightGreen, Reset,
		Bold, BrightGreen, Reset,
		BrightCyan, Reset,
		BrightCyan, Reset,
		BrightCyan, Reset,
		BrightCyan, Reset,
		BrightCyan, Reset,
		BrightCyan, Reset,
		BrightCyan, Reset,
		Bold, BrightGreen, Reset,
	)
	fmt.Println(instructions)
}

func main() {
	var version bool
	var ipRange string
	var expand string
	var binaryConvert string
	// Flags var for encoder/decoder
	var encoder string
	var decoder string
	var encodingType string

	// Flag to save to file
	var output string

	// Flag var for ipRange
	flag.StringVar(&ipRange, "range", "", "use -range 192.168.0.1/24")
	flag.StringVar(&expand, "expand", "", "use -subnet 192.168.0.1/24 to summarize CIDR info")
	flag.StringVar(&binaryConvert, "binary", "", "use -binary 192.168.0.1 to convert ip to binary")

	// Encoding/Decoding flags
	flag.StringVar(&encoder, "encode", "", "use -encode 'StringHere' -type(hex, base64, url) to encode it")
	flag.StringVar(&decoder, "decode", "", "use -decode 'DecodeThis' -type(hex, base64, url) to decode it")
	flag.StringVar(&encodingType, "type", "", "use -type 'hex' -type(hex, base64, url) to encode/decode it")

	// Output flag
	flag.StringVar(&output, "o", "", "-o use this flag to save to the file for cidr expander only for now")

	// Version flag
	flag.BoolVar(&version, "version", false, "Show tool version")

	flag.Parse()

	// Check if no flags were provided
	if flag.NFlag() == 0 {
		banner()
		showInstructions()
		return
	}

	switch {
	case ipRange != "":
		mode := 0
		Expander(ipRange, output, mode)
	case expand != "":
		mode := 1
		Expander(expand, output, mode)
	case binaryConvert != "":
		toBinary(binaryConvert)
	case encoder != "":
		Codec(encoder, decoder, encodingType)
	case decoder != "":
		Codec(encoder, decoder, encodingType)
	case version != false:
		fmt.Println("Current Version:", Version)
	}
}
