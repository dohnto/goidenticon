package main

import (
	"os"
	"bufio"
	"io"
	"image/png"
	"io/ioutil"
	"crypto/md5"
	"flag"

	log "github.com/Sirupsen/logrus"
	"github.com/dohnto/goidenticon/identicon"
	logFlag "github.com/dohnto/goidenticon/log"
)


var (
	flagSize int
	flagHash string
	flagInFile string
	flagOutFile string
	flagLogLevel logFlag.Level
)

func init() {
	flag.IntVar(&flagSize, "size", 300, "Size of a square image in pixels")
	flag.StringVar(&flagInFile, "in", "", "File path to read input from, if not passed or empty, stdin is used")
	flag.StringVar(&flagOutFile, "out", "identicon.png", "File path to a result image")
	flag.Var(&flagLogLevel, "log-level", "File path to a result image")
}

func main() {
	flag.Parse()

	log.SetLevel(flagLogLevel.Level)

	var reader io.Reader
	var err error

	if flagInFile == "" {
		reader = bufio.NewReader(os.Stdin)
	} else {
		reader, err = os.Open(flagInFile)
		if err != nil {
			log.Panic(err)
		}
	}
	text, err := ioutil.ReadAll(reader)
	if err != nil {
		log.Panic(err)
	}

	hash := md5.New()
	io.WriteString(hash, string(text))

	d := identicon.NewColorful(flagSize)
	im, _ := d.Build(hash.Sum(nil))

	toimg, _ := os.Create(flagOutFile)
	defer toimg.Close()

	png.Encode(toimg, im)
}
