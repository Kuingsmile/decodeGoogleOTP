# decodeGoogleOTP

This is a cli tool to decode one time password (OTP) secrets from QR codes exported by Google Authenticator. The exported QR codes can be read as image files and decoded result can be exported to JSON or CSV or saved as QR code images or printed as QR code to terminal.

## How to use

Just download the binary from the [release page](https://github.com/kuingsmile/decodeGoogleOTP/releases) and run it in your terminal.

```shell
$ decodeGoogleOTP -i <input file> -c <csv file path>
```

## Parameters

```shell
$ decodeGoogleOTP -h

decodeGoogleOTP is a command line tool to decode Google OTP QR codes. Output can be json, csv, qrcode or plain text.

Usage:
  decodeGoogleOTP [flags]
  decodeGoogleOTP [command]

Available Commands:
  completion  Generate the autocompletion script for the specified shell
  help        Help about any command
  version     Print current version of the application

Flags:
  -c, --csv string      Output in CSV format and specify the output file
  -d, --debug           Enable debug mode
  -h, --help            help for decodeGoogleOTP
  -i, --input string    Input file path
  -j, --json string     Output in JSON format and specify the output file
  -p, --print-qr        Print QR code to terminal
  -q, --qrcode string   Output in QR code image format and specify the output directory
  -s, --silent          Enable silent mode
  -t, --text string     Output url list in plain text format and specify the output file
  -u, --url string      Output in URL format and specify the output file
  -v, --version         Print version information

Use "decodeGoogleOTP [command] --help" for more information about a command.
```

## Examples

### Decode QR code and save as JSON

```shell
$ decodeGoogleOTP -i <input file> -j <output file>
```

### Decode QR code and save as CSV

```shell
$ decodeGoogleOTP -i <input file> -c <output file>
```

### Decode QR code and save as QR code images

```shell
$ decodeGoogleOTP -i <input file> -q <output directory>
```

### Decode QR code and print as QR code to terminal

```shell
$ decodeGoogleOTP -i <input file> -p
```

### Decode QR code and save as plain text

```shell
$ decodeGoogleOTP -i <input file> -t <output file>
```

### Decode QR code and save as URL

```shell
$ decodeGoogleOTP -i <input file> -u <output file>
```

## License

This project is open source under the MIT license. Welcome everyone to use and contribute code.

[MIT License](https://opensource.org/licenses/MIT)

Copyright (c) 2024-present Kuingsmile
