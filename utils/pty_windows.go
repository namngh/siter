//go:build windows
// +build windows

package utils

import (
	"io"
	"os"
	"unicode/utf8"

	"github.com/UserExistsError/conpty"
)

type PTYWindows struct {
	cpty         conpty.ConPty
	shellCommand string
}

func NewPTYUnix(shellCommand string) (IPTY, error) {
	return nil, errors.New("MISS_MATCH_OS")
}

func NewPTYWindows(shellCommand string) (p PTYWindows, err error) {
	cpty, err := conpty.Start(shellCommand)
	if err != nil {
		return p, err
	}

	return PTYWindows{cpty: cpty, shellCommand: shellCommand}, nil
}

func (p PTYWindows) Read(buffer [][]rune) {
	go func() {
		line := []byte{}
		buffer = append(buffer, line)

		for {
			n, err := p.cpty.Read(line)
			if err != nil {
				// if err == io.EOF {
				// 	return
				// }
				os.Exit(0)
			}

			buffer[len(buffer)-1] = line
			if n > 0 {
				if len(buffer) > MaxBufferSize {
					buffer = buffer[1:]
				}

				line = []rune{}
				buffer = append(buffer, utf8.DecodeRune(line[:n]))
			}
		}
	}()
}

func (p PTYWindows) Write(text []byte) (int, error) {
	return r.cpty.Write(text)
}

func (p PTYWindows) Close() {
	p.cpty.Close()
}