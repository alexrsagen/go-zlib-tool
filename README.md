# go-zlib-tool
Tool for compressing and extracting raw zlib compressed files.

Uses [compress/zlib](https://pkg.go.dev/compress/zlib) internally.

Should support any file starting with the following zlib magic headers:
- `78 01` - No Compression/low
- `78 9C` - Default Compression
- `78 DA` - Best Compression 

Can only extract a single compressed file.

## Usage

```
Usage of go-zlib-tool:
  -c    Compress the input
  -x    Extract the input
  -i string
        Input path
  -o string
        Output path
```
