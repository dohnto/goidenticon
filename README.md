# goidenticon

`goidenticon` is a simple command line utility to generate identicon from an arbitrary input.

[Identicon](https://en.wikipedia.org/wiki/Identicon) is a visual representation of a hash value,
commonly used to visualize IP addresses or user avatars.
This implementation provides so far only single hash decoder and it uses only different colors to
visualize the hash.

## Examples

Sample icons generated for different usernames

```bash
echo Antonette | goidenticon -out doc/images/Antonette.png
```

![Antonette](https://github.com/dohnto/goidenticon/raw/master/doc/images/Antonette.png "Antonette")

```bash
echo Glen | goidenticon -out doc/images/Glen.png
```

![Glen](https://github.com/dohnto/goidenticon/raw/master/doc/images/Glen.png "Glen")

```bash
echo Willy | goidenticon -out doc/images/Willy.png
```

![Willy](https://github.com/dohnto/goidenticon/raw/master/doc/images/Willy.png "Willy")

## Installation

```bash
go get github.com/dohnto/goidenticon/...
```

## Usage

```bash
# goidenticon -help
Usage of goidenticon:
  -in string
    	File path to read input from, if not passed or empty, stdin is used
  -log-level value
    	File path to a result image
  -out string
    	File path to a result image (default "identicon.png")
  -size int
    	Size of a square image in pixels (default 300)
 ```