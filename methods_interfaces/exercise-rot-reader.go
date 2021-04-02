package main

import (
    "io"
    "os"
    "strings"
)

type rot13Reader struct {
    r io.Reader
}

func (rotReader rot13Reader) Read(b []byte) (int, error) {
    n, err := rotReader.r.Read(b)
    b = rot13(b[:n])
    return n, err
}

func rot13(b []byte) []byte {
    for i := range(b) {
        var base byte
        if b[i] >= 'A' && b[i] <= 'Z' {
            base = 'A'
        } else if b[i] >= 'a' && b[i] <= 'z' {
            base = 'a'
        } else {
            continue
        }
        b[i] = (b[i] - base + 13) % 26 + base
    }
    return b
}

func main() {
    s := strings.NewReader("Lbh penpxrq gur pbqr!")
    r := rot13Reader{s}
    io.Copy(os.Stdout, &r)
}

