# NetShiv

```
    ███╗   ██╗███████╗████████╗███████╗██╗  ██╗██╗██╗   ██╗
    ████╗  ██║██╔════╝╚══██╔══╝██╔════╝██║  ██║██║██║   ██║
    ██╔██╗ ██║█████╗     ██║   ███████╗███████║██║██║   ██║
    ██║╚██╗██║██╔══╝     ██║   ╚════██║██╔══██║██║╚██╗ ██╔╝
    ██║ ╚████║███████╗   ██║   ███████║██║  ██║██║ ╚████╔╝ 
    ╚═╝  ╚═══╝╚══════╝   ╚═╝   ╚══════╝╚═╝  ╚═╝╚═╝  ╚═══╝  
```

**Network Penetration & Analysis Tool** - A powerful CLI utility for subnet calculations, IP range expansion, binary conversion, and encoding operations.

[![Go Version](https://img.shields.io/badge/Go-1.21+-00ADD8?style=flat&logo=go)](https://go.dev/)
[![License](https://img.shields.io/badge/license-MIT-blue.svg)](LICENSE)
[![Platform](https://img.shields.io/badge/platform-Linux%20%7C%20Windows%20%7C%20macOS-lightgrey)](https://github.com/ProXxNebula/netshiv/releases)

---

## 🚀 Features

### 🔢 Subnet Calculator
Calculate network information from CIDR notation:
- Network address
- Broadcast address
- IP range (first and last usable IPs)
- Number of available hosts

### 📋 IP Range Expander
Generate complete lists of IP addresses within a CIDR range:
- Export to file for use with network scanners
- Support for any CIDR notation (/8 to /32)

### 🔀 Binary Converter
Convert IP addresses to binary representation:
- Dotted binary format (e.g., `11000000.10101000.00000001.00000001`)
- Useful for understanding subnet masks and network boundaries

### 🔐 Encoding & Decoding
Multi-format encoding support:
- **Base64**: Standard Base64 encoding/decoding
- **Hex**: Hexadecimal encoding/decoding
- **URL**: URL-safe encoding/decoding
- **ROT13**: Classic Caesar cipher rotation

---

## 📦 Installation

### Download Pre-built Binaries

Download the latest release for your platform from the [Releases](https://github.com/ProXxNebula/netshiv/releases) page:

```bash
# Linux (x64)
wget https://github.com/ProXxNebula/netshiv/releases/download/v1.0.0/netshiv-linux-amd64
chmod +x netshiv-linux-amd64
sudo mv netshiv-linux-amd64 /usr/local/bin/netshiv

# macOS (Apple Silicon)
wget https://github.com/ProXxNebula/netshiv/releases/download/v1.0.0/netshiv-darwin-arm64
chmod +x netshiv-darwin-arm64
sudo mv netshiv-darwin-arm64 /usr/local/bin/netshiv

# Windows
# Download netshiv-windows-amd64.exe and add to PATH
```

### Build from Source

```bash
# Clone the repository
git clone https://github.com/ProXxNebula/netshiv.git
cd netshiv

# Build for your platform
make build

# Or build for all platforms
make release

# Install system-wide (Linux/macOS)
make install
```

**Requirements:**
- Go 1.21 or higher
- Make (optional, for using Makefile commands)

---

## 📖 Usage

### Basic Commands

```bash
# Show help and available options
netshiv

# Show version
netshiv -version
```

### Subnet Operations

```bash
# Get subnet information summary
netshiv -expand 192.168.1.0/24

# Expand CIDR range to list all IPs
netshiv -range 192.168.1.0/24

# Save IP list to file
netshiv -range 10.0.0.0/28 -o targets.txt
```

**Example Output:**
```
Network Address: 192.168.1.0
Broadcast IP: 192.168.1.255
IP Range: First: 192.168.1.0  Last: 192.168.1.255
```

### Binary Conversion

```bash
# Convert IP to binary
netshiv -binary 192.168.1.1
```

**Example Output:**
```
11000000.10101000.00000001.00000001
```

### Encoding & Decoding

```bash
# Base64 encoding
netshiv -encode "Hello World" -type base64
# Output: SGVsbG8gV29ybGQ=

# Base64 decoding
netshiv -decode "SGVsbG8gV29ybGQ=" -type base64
# Output: Hello World

# Hex encoding
netshiv -encode "secret" -type hex
# Output: 736563726574

# URL encoding
netshiv -encode "hello world" -type url
# Output: hello+world

# ROT13 encoding (self-reversing)
netshiv -encode "test" -type rot13
netshiv -decode "grfg" -type rot13
```

---

## 🛠️ Development

### Project Structure

```
netshiv/
├── cmd/
│   └── netshiv/                 # main entry point for your CLI
│       └── main.go
├── internal/                    # non-exported/internal packages
│   ├── codec/                   # encoding/decoding utilities
│   │   └── encoder.go
│   ├── network/                 # subnet, IP expansion, binary conversion
│   │   ├── cidrExpander.go
│   │   └── binary.go
│   └── ui/                      # banner & ANSI output helpers
│       └── banner.go
├── build/                       # release binaries (generated)
├── go.mod
├── go.sum
├── Makefile
├── LICENSE
├── README.md
└── CONTRIBUTING.md
```

### Building

```bash
# Development build (fast, includes debug info)
make dev

# Production build (optimized, stripped)
make build

# Cross-compile for all platforms
make release

# Run tests
make test

# Clean build artifacts
make clean
```

### Makefile Targets

| Command | Description |
|---------|-------------|
| `make build` | Build for current platform |
| `make release` | Build for all platforms |
| `make clean` | Remove build artifacts |
| `make test` | Run Go tests |
| `make install` | Install to /usr/local/bin |
| `make dev` | Quick development build |

---

## 🎯 Use Cases

### Penetration Testing
- Generate target IP lists from CIDR ranges for network scanning
- Quickly calculate subnet boundaries during reconnaissance
- Prepare IP lists for tools like Nmap, Masscan, or custom scripts

### Network Administration
- Plan and document IP address allocation
- Verify subnet configurations
- Calculate network capacity and addressing schemes

### CTF Challenges
- Decode encoded flags and messages
- Analyze network configurations in challenges
- Quick binary/hex conversions for puzzle solving

### Learning & Education
- Understand binary representation of IP addresses
- Practice subnet calculations
- Learn encoding schemes (Base64, Hex, ROT13)

---

## 📋 Examples

### Example 1: Scan Preparation
```bash
# Generate list of IPs in corporate subnet
netshiv -range 10.10.1.0/24 -o corporate-ips.txt

# Use with nmap
nmap -iL corporate-ips.txt -p 80,443 --open
```

### Example 2: Subnet Planning
```bash
# Check how many hosts in /26 network
netshiv -expand 192.168.50.0/26

# Output shows:
# Network: 192.168.50.0
# Broadcast: 192.168.50.63
# Usable IPs: 62 hosts (192.168.50.1 - 192.168.50.62)
```

### Example 3: CTF Flag Decoding
```bash
# Found encoded flag: "U0dWc2JHOGdWMjl5YkdRPQ=="
netshiv -decode "U0dWc2JHOGdWMjl5YkdRPQ==" -type base64
# Output: SGVsbG8gV29ybGQ=

# Decode again
netshiv -decode "SGVsbG8gV29ybGQ=" -type base64
# Output: Hello World
```

---

## 🤝 Contributing

Contributions are welcome! Feel free to:
- Report bugs
- Suggest new features
- Submit pull requests
- Improve documentation

### Development Setup
```bash
git clone https://github.com/ProXxNebula/netshiv.git
cd netshiv
go mod download
make build
```

---

## 📄 License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

---

## ⚠️ Disclaimer

**NetShiv is intended for educational purposes and authorized security testing only.**

Users are solely responsible for ensuring compliance with:
- Applicable laws and regulations
- Network owner permissions
- Organizational security policies
- Ethical hacking guidelines

Unauthorized network scanning, penetration testing, or security assessments may be illegal in your jurisdiction. Always obtain proper authorization before testing networks you do not own.

---

## 🔗 Links

- **GitHub**: [https://github.com/ProXxNebula/netshiv](https://github.com/ProXxNebula/netshiv)
- **Issues**: [Report a bug](https://github.com/ProXxNebula/netshiv/issues)
- **Releases**: [Download binaries](https://github.com/ProXxNebula/netshiv/releases)

---

## 📊 Project Stats

- **Language**: Go
- **Lines of Code**: ~500
- **Dependencies**: Standard library only
- **Platforms**: Linux, Windows, macOS (x64 & ARM64)

---

Made with ❤️ for the cybersecurity community