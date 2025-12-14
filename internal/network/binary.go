package cidr

import (
	"fmt"
	"net"
)

func ToBinary(ip string) {
	// Parse ipv4 from string to 4byte representation
	binaryParse := net.ParseIP(ip).To4()

	result := ""
	for i, v := range binaryParse {
		result += fmt.Sprintf("%08b", v)
		if i < len(binaryParse)-1 {
			result += "."
		}
	}

	fmt.Println(result)
}
