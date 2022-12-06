package main

type logger struct {
	buffer []byte
}

func newLogger() *logger {
	return &logger{
		buffer: make([]byte, 0),
	}
}

func (l *logger) Write(p []byte) (n int, err error) {
	l.buffer = append(l.buffer, p...)
	return len(p), nil
}

func (l *logger) reset() {
	l.buffer = make([]byte, 0)
}

func (l *logger) dump() string {
	return string(l.buffer)
}
