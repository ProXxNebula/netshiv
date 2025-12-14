package cidr

import (
	"fmt"
	"net"
)

func ToBinary(ip string) { // NOT exported
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
