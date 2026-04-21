# go-steghide

Steganographically hide a message in an image's least-signifigant bytes

## Usage

### Encoding

```bash
./go-steghide encode --image some-image.jpg --message secret --output out.png
```

Writes to file and outputs the amount of bits written (necessary for decode)

### Decoding

```bash
./go-steghide decode --image some-image.jpg --bits 48
```

Decode --bits amount of bits from an image

## About

Go beginner project
