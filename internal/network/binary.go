package cidr

import (
	"fmt"
	"net"
	"strings"
)

func ToBinary(ip string) {
	// Parse ipv4 from string to 4byte representation
	binaryParse := net.ParseIP(ip).To4()

	var result strings.Builder
	for i, v := range binaryParse {
		fmt.Fprintf(&result, "%08b", v)
		if i < len(binaryParse)-1 {
			result.WriteString(".")
		}
	}

	fmt.Println(result.String())
}
