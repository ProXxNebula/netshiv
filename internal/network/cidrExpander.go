package cidr

import (
	"encoding/binary"
	"fmt"
	"net"
	"os"
)

func Expander(ipAddress string, output string, mode int) {

	// Parse IP/CIDR
	ip, ipNet, err := net.ParseCIDR(ipAddress)
	if err != nil {
		fmt.Println("CIDR input error")
		return
	}

	// Convert ipv4 to 4 bytes slice
	ip = ip.To4()
	mask := ipNet.Mask
	mask32 := binary.BigEndian.Uint32(mask)

	// Network address
	networkIP := ip.Mask(mask)
	net32 := binary.BigEndian.Uint32(networkIP)

	// Broadcast: network OR inverted mask
	broadcast32 := net32 | ^mask32

	// Print summary
	bcast := make([]byte, 4)
	binary.BigEndian.PutUint32(bcast, broadcast32)

	switch mode {
	case 0:
		fmt.Println("Network Address:", networkIP)
		fmt.Println("Broadcast IP:", net.IP(bcast))
		fmt.Printf("IP Range: First: %v  Last: %v\n", networkIP, net.IP(bcast))
	case 1:
		var ips []net.IP
		// Loop the entire range
		for x := net32; x <= broadcast32; x++ {
			buf := make([]byte, 4)
			binary.BigEndian.PutUint32(buf, x)
			ips = append(ips, net.IP(buf))
		}

		// If output file is specified, save to file
		if output != "" {
			f, err := os.Create(output)
			if err != nil {
				fmt.Println("Error creating file:", err)
				return
			}
			defer f.Close()

			for _, ip := range ips {
				f.WriteString(ip.String() + "\n")
			}
			fmt.Println("Saved", len(ips), "IPs to", output)
		} else {
			// Otherwise, just print
			for _, ip := range ips {
				fmt.Println(ip)
			}
		}

	}
}
