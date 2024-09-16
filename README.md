# go-rle
RLE (or Run-length encoding) encoder and decoder for golang

## Install
```bash
go install github.com/kirkegaard/go-rle/cmd/rle@latest
```

## Usage
```bash
rle --encode "hello world"
1h1e2l1o1 1w1o1r1l1d

rle --decode "1h1e2l1o1 1w1o1r1l1d"
hello world
```
