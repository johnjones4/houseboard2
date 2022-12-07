package core

import "io"

type multiwriter struct {
	writers []io.Writer
}

func (mw *multiwriter) Write(p []byte) (n int, err error) {
	for _, w := range mw.writers {
		n, err = w.Write(p)
	}
	return n, err
}
