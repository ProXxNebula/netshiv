# NetShiv

Network Penetration & Analysis Tool - A CLI utility for subnet calculations, IP range expansion, and encoding operations.

## Features

- 🔢 **Subnet Calculator**: Calculate network addresses, broadcast addresses, and IP ranges from CIDR notation
- 📋 **IP Range Expander**: Generate all IP addresses within a CIDR range
- 🔀 **Binary Converter**: Convert IP addresses to binary representation
- 🔐 **Encoding/Decoding**: Support for Base64, Hex, URL, and ROT13 encoding

## Installation

```bash
# Clone the repository
git clone https://github.com/yourusername/netshiv
cd netshiv

# Build from source
go build -o netshiv

# Or use make
make build
```

## Usage

```bash
# Subnet information
./netshiv -subnet 192.168.1.0/24

# Expand IP range
./netshiv -range 192.168.1.0/24

# Save to file
./netshiv -range 192.168.1.0/24 -o ips.txt

# Binary conversion
./netshiv -binary 192.168.1.1

# Encoding
./netshiv -encode "secret" -type base64
./netshiv -encode "hello" -type hex

# Decoding
./netshiv -decode "c2VjcmV0" -type base64
```

## Building Releases

```bash
# Build for all platforms
make release

# Binaries will be in the build/ directory
```

## Requirements

- Go 1.21 or higher

## License

[Add your license here]

## Disclaimer

This tool is for educational and authorized testing purposes only. Users are responsible for complying with applicable laws.