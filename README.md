## Description

JsMole is a Go script designed to rapidly and faithfully download debugging sources equivalent to those presented by Google Web browsers (such as Chrome or Chromium).

## Features

- Fast Download: Utilizes advanced techniques to accelerate the download of debugging sources.
- Fidelity to Original: Downloaded files are identical to the sources presented by Google Web browsers.
- Easy to Use: Simple and intuitive command-line interface.

## Installation

Ensure you have Go installed on your system. If not, you can download and install it from golang.org.
Clone this GitHub repository into your local directory:

```bash
❯ git clone https://github.com/your-username/jsmole.git
❯ cd jsmole
```

Compile the Go script:

```bash
go build
```

## Usage

Use the command-line script with the link of the target website to download:

```bash
❯ ./jsmole -u https://target.com
```

Help usage:

```bash
❯ ./jsmole -h
NAME:
   jsmole - Google debugger but locally

USAGE:
   jsmole [global options] command [command options]

COMMANDS:
   help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --url value, -u value     Website's url to be scanned
   --output value, -o value  Output folder to be selected (default: "./output")
   --help, -h                show help
```

## Disclaimer

This script is intended to be used ethically and legally. Make sure you have the necessary permissions to download debugging sources from the provided links.

## Contributions

Contributions in the form of pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change.

## License

This project is licensed under the MIT License. See the LICENSE file for more information.
