package main

import (
	"flag"
	"fmt"
	"image"
	"image/draw"
	_ "image/jpeg"
	"image/png"
	_ "image/png"
	"log"
	"os"
	"path/filepath"
)



func main() {
	// parse CLI args. positional args are "decode" and "encode"
	// flag.Usage = CLIUsage

	if len(os.Args) > 1 {
		// subcommand provided
		switch os.Args[1] {
		case "encode":
			// parse encode-subcommand's argument flags
			encodeCmd := flag.NewFlagSet("encode", flag.ExitOnError)
			encodeImage := encodeCmd.String("image", "", "input image to encode a message into")
			encodeMessage := encodeCmd.String("message", "", "the message to encode")
			encodeOutput := encodeCmd.String("output", "out.png", "output file (RGBA format)")
			encodeCmd.Parse(os.Args[2:])

			// input validation
			if len(*encodeImage) == 0 || len(*encodeMessage) == 0 {
				encodeCmd.Usage()
				os.Exit(0)
			}

			rgba, bits := encode(getImage(encodeImage), encodeMessage)
			fmt.Printf("Encoded the message as %d bits into the image\n", bits)
			output(rgba, *encodeOutput)

		case "decode":
			decodeCmd := flag.NewFlagSet("decode", flag.ExitOnError)
			decodeImage := decodeCmd.String("image", "", "input image to decode a message from")
			decodeLength := decodeCmd.Int("bits", 0, "the amount of bits to decode (divisible by 8)")
			decodeCmd.Parse(os.Args[2:])

			// input validation
			if len(*decodeImage) == 0 || *decodeLength == 0 {
				decodeCmd.Usage()
				os.Exit(0)
			}

			message := decode(getImage(decodeImage), *decodeLength)
			fmt.Printf("Decoded: %s\n", message)

		default:
			exit("Usage: go-steghide encode|decode [OPTIONS]\n")
		}
	} else {
		exit("Usage: go-steghide encode|decode [OPTIONS]\n")
	}

	// make a new RGBA image object with a writable .Pix slice
	// rgba := invert(src)
}

func exit(message string) {
	fmt.Println(message)
	os.Exit(0)
}

func getImage(imagePath *string) image.Image {
	// assert that "image" exists
	if _, err := os.Stat(filepath.Clean(*imagePath)); err != nil {
		log.Fatal(err)
	}

	// open image
	f, err := os.Open(filepath.Clean(*imagePath))
	if err != nil {
		log.Fatal(err)
	}
	// TODO catch file does not exist
	// TODO final catch (permission denied, etc)
	defer f.Close()

	// parse to image object
	src, _, err := image.Decode(f)
	if err != nil {
		log.Fatal(err)
	}
	return src
}

func decode(img image.Image, bitsToDecode int) string {
	// make a copy of img
	bounds := img.Bounds()
	rgba := image.NewRGBA(bounds)
	draw.Draw(rgba, bounds, img, bounds.Min, draw.Src)

	bytes := make([]byte, bitsToDecode / 8)
	for bytei := range bitsToDecode / 8 {
		// map a byte
		var b uint8 = 0
		for biti := range 8 {
			bit := rgba.Pix[pixI(bytei*8+biti)] & 1//(1 << biti)
			b |= bit << biti
		}
		// fmt.Printf("Found byte %d\n", b)
		bytes[bytei] = b
	}
	return string(bytes)
}

func encode(img image.Image, msg *string) (*image.RGBA, int) {
	// modify least-signifigant bit of pixels
	// return encoded message, and print the amount of bytes modified (use to decode precise length)

	// make a copy of img
	bounds := img.Bounds()
	rgba := image.NewRGBA(bounds)
	draw.Draw(rgba, bounds, img, bounds.Min, draw.Src)

	// iterate bytes (not runes) of msg
	for bytei := range len(*msg) {
		b := (*msg)[bytei]
		// fmt.Printf("Encoding byte %d\n", b)
		// iterate all bits
		for biti := range(8) {
			// get bit of byte b[biti]
			bit := (b >> biti) & 1
			// modify LSB of Pix[bytei*8+biti]
			modifyPixel := pixI(bytei*8 + biti)
			if bit == 1 {
				rgba.Pix[modifyPixel] |= 1  // set bit
			} else {
				rgba.Pix[modifyPixel] &^= 1 // clear bit
			}
		}
	}
	return rgba, 8*len(*msg)
}

func pixI(i int) int {
	// get Pix[i] of the R, G, B channels but skip alpha (not lossless)
	return (i/3)*4 + (i%3)
}

func output(img *image.RGBA, filepath string) {
	// create file
	out, err := os.Create(filepath)
	if err != nil {
		log.Fatal(err)
	}
	defer out.Close()

	// write to file
	if err := png.Encode(out, img); err != nil {
		log.Fatal(err)
	}
}
