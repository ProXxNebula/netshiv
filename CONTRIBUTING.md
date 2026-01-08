# Contributing to NetShiv

Thanks for your interest in contributing! 🎉

## 🐛 Reporting Bugs

Found a bug? [Open an issue](https://github.com/ProXxNebula/netshiv/issues) with:
- What you did
- What you expected
- What actually happened
- Your OS and Go version

## 💡 Suggesting Features

Have an idea? [Open an issue](https://github.com/ProXxNebula/netshiv/issues) and describe:
- The problem you want to solve
- Your proposed solution
- Example use cases

## 🔧 Contributing Code

### Quick Start

```bash
# Fork the repo, then:
git clone https://github.com/ProXxNebula/netshiv.git
cd netshiv
make build
```

### Making Changes

1. Create a branch: `git checkout -b feature/your-feature`
2. Make your changes
3. Test it: `make build && ./netshiv`
4. Commit: `git commit -m "feat: add your feature"`
5. Push: `git push origin feature/your-feature`
6. Open a Pull Request

### Code Style

- Use `gofmt` for formatting
- Add comments for complex logic
- Handle errors properly
- Keep it simple and readable

### Commit Messages

```bash
feat: add new feature
fix: fix bug description
docs: update documentation
```

## 📝 Adding New Encodings

Want to add a new encoding format? Here's how:

1. Add encode/decode functions in `codec.go`
2. Update the switch cases in `Codec()` function
3. Update help text in `main.go`
4. Test it works!

Example:
```go
// In codec.go
func encodeBase32(s string) string {
    return base32.StdEncoding.EncodeToString([]byte(s))
}

// In Codec() switch statement
case "base32":
    fmt.Println("Base32:", encodeBase32(encode))
```

## ❓ Questions?

Open a [Discussion](https://github.com/ProXxNebula/netshiv/discussions) or create an issue!

---

**That's it! Keep it simple and have fun coding!** 🚀