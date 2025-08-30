# HEIC Converter

I bought an iPhone and have been miserable ever since.

This converts `.heic` images into different formats (for now, `.png` or `.webp`)

## How to use it
1. Put the .heic files inside the `./input` folder

2. `go mod tidy` to install dependencies

3. `go build -o heicc main.go` to build the binary

4. `./heicc <format>` to convert files inside `./input` to the specified format (converted files goes into `./output`)
   - Example: `./heicc png` or `./heicc webp`


## Under the hood
This converter uses:
- [github.com/jdeng/goheif](github.com/jdeng/goheif) as Heif Decoder. \
This is based on libde265.
- [github.com/HugoSmits86/nativewebp](github.com/HugoSmits86/nativewebp) as Webp Encoder. \
This is a native Go encoder. No dependencies on external libs.