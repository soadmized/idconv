# 🔄 idconv

[![Go Reference](https://pkg.go.dev/badge/github.com/soadmized/idconv.svg)](https://pkg.go.dev/github.com/soadmized/idconv)
[![Go Report Card](https://goreportcard.com/badge/github.com/soadmized/idconv)](https://goreportcard.com/report/github.com/soadmized/idconv)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

A lightweight and convenient CLI utility for working with UUIDs. Allows converting UUIDs between various formats:
- Standard string UUID format
- Base64
- Raw Base64 (no padding)

Also includes a function for generating random UUIDs.

## 📥 Installation
```bash
go install github.com/soadmized/idconv@latest
```
## 🚀 Usage

Basic syntax:

```bash
idconv [uuid | base64 | rawBase64] | random
```

## 📝 Examples

### 1️⃣ Converting from standard UUID format
```bash
idconv 9ca5de1c-c01a-4cc1-96d6-1017b40d5843
```
### 2️⃣ Converting from Base64

```bash
idconv nKXeHMAaTMGW1hAXtA1YQw==
```
### 3️⃣ Converting from Raw Base64 (no padding)
```bash
idconv nKXeHMAaTMGW1hAXtA1YQw
```

### 4️⃣ Generating a random UUID
```bash
idconv random
```

## 📋 Output Format

In all cases, the result will be presented in the following format:

```
 ----------------------------------------------------
|   uuid    |  9ca5de1c-c01a-4cc1-96d6-1017b40d5843  |
 ----------------------------------------------------
|  base64   |        nKXeHMAaTMGW1hAXtA1YQw==        |
 ----------------------------------------------------
| rawBase64 |         nKXeHMAaTMGW1hAXtA1YQw         |
 ----------------------------------------------------
```

## 💻 How It Works

The utility takes an input value and tries to determine its format:
1. If it's a valid UUID - converts it to base64 and rawBase64
2. If it's a valid base64 string - decodes it to a UUID and also outputs rawBase64
3. If it's a valid rawBase64 string - decodes it to a UUID and also outputs base64
4. If the `random` key is specified - generates a new random UUID and displays it in all formats

## 📝 TODO

Future development plans:
1. Add conversion of UUIDs to/from various endian formats (native, little, big)
2. Refactor the code using Viper or Cobra for improved CLI functionality

## 🤝 Contributing

Contributions are welcome! Feel free to open issues or create pull requests.

## 📜 License

MIT
