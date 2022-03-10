package writeline

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
	"sync"
	"time"
)

const errPrefix = "LineWriter: "

// LineWriter maintains <n> new lines. Each line can be overwritten. Zero value is not ready for use.
type LineWriter struct {
	stop     chan struct{}
	lines    int
	currLine int // from top to down, begin at zero
	wr       *bufio.Writer
	duration time.Duration
	mu       sync.Mutex
}

// New creates a *LineWriter for io.Writer, inits and maintains n lines. Also starts a flush timer with default duration 200*time.Millisecond
// 	Note: This is a buffer method.
func New(lines int, w io.Writer) (*LineWriter, error) {
	wl := &LineWriter{
		wr:       bufio.NewWriter(w),
		lines:    lines,
		stop:     make(chan struct{}),
		duration: 200 * time.Millisecond,
	}

	go wl.startTimer()
	err := wl.initLines()
	if err != nil {
		return nil, err
	}
	return wl, nil
}

// NewWithStdout creates a *LineWriter for os.Stdout, inits and maintains n lines. Also starts a flush timer with default duration 200*time.Millisecond
// 	Note: This is a buffer method.
func NewWithStdout(lines int) (*LineWriter, error) {
	return New(lines, os.Stdout)
}

func (w *LineWriter) SetFlushDuration(d time.Duration) {
	w.mu.Lock()
	defer w.mu.Unlock()

	w.duration = d
	w.stop <- struct{}{}
	go w.startTimer()
}

// NewLine Move cursor to new line at the bottom.
// 	Note: This is a buffer method.
func (w *LineWriter) NewLine(s string) *LineWriter {
	w.mu.Lock()
	defer w.mu.Unlock()

	if err := w.moveCursorToLine(w.lines - 1); err != nil {
		return w
	}
	_, err := w.wr.WriteString("\n\r" + s)
	if err != nil {
		return w
	}
	w.lines += 1
	w.currLine = w.lines - 1
	return nil
}

// WriteLastLine overwrites the last line.
// 	Any position control string will lead to unexpected behavior. Such as \n, \033[K ...
// 	Note: This is a buffer method.
func (w *LineWriter) WriteLastLine(s string) *LineWriter {
	w.mu.Lock()
	defer w.mu.Unlock()

	w.writeLine(w.lines-1, s)
	return w
}

// WriteLine overwrites the nth line.
// 	Any position control string will lead to unexpected behavior. Such as \n, \033[K...
// 	Note: This is a buffer method.
func (w *LineWriter) WriteLine(n int, s string) *LineWriter {
	w.mu.Lock()
	defer w.mu.Unlock()

	w.writeLine(n, s)
	return w
}

// Lines return total maintained lines
func (w *LineWriter) Lines() int {
	return w.lines
}

// Flush all buffered string to the underlying io.Writer
func (w *LineWriter) Flush() *LineWriter {
	w.mu.Lock()
	defer w.mu.Unlock()
	w.wr.Flush()
	return w
}

// Close end with a newline, stop timer, flush all remaining buffered string
func (w *LineWriter) Close() {
	w.mu.Lock()
	defer w.mu.Unlock()
	defer w.wr.Flush()
	defer close(w.stop)

	if err := w.moveCursorToLine(w.lines - 1); err != nil {
		return
	}
	_, err := w.wr.WriteString("\n")
	if err != nil {
		return
	}
	return
}

func (w *LineWriter) startTimer() {
	for {
		select {
		case <-time.After(w.duration):
			w.Flush()
		case <-w.stop:
			w.Flush()
			return
		}
	}
}

func (w *LineWriter) moveCursorToLine(n int) error {
	switch {
	case n > w.lines-1:
		return errors.New(errPrefix + "out of line range")
	case n == w.currLine:
		return nil
	case n > w.currLine:
		return w.moveDownLines(n - w.currLine)
	case n < w.currLine:
		return w.moveUpLines(w.currLine - n)
	}
	return nil
}

func (w *LineWriter) initLines() error {
	for i := 0; i < w.lines-1; i++ {
		_, err := w.wr.WriteString("\n")
		if err != nil {
			return err
		}
	}
	w.currLine = w.lines - 1
	return nil
}

func (w *LineWriter) cursorReturn() error {
	_, err := w.wr.WriteString(carriageReturn)
	return err
}

func (w *LineWriter) moveUpLines(n int) error {
	_, err := w.wr.WriteString(fmt.Sprintf(moveUpLines, n))
	if err != nil {
		return err
	}
	w.currLine -= n
	return nil
}

func (w *LineWriter) moveDownLines(n int) error {
	_, err := w.wr.WriteString(fmt.Sprintf(moveDownLines, n))
	if err != nil {
		return err
	}
	w.currLine += n
	return nil
}

func (w *LineWriter) eraseLine() (err error) {
	_, err = w.wr.WriteString(carriageReturn + eraseToEnd)
	return
}

func (w *LineWriter) writeLine(n int, s string) error {
	if err := w.moveCursorToLine(n); err != nil {
		return errFunc(err)
	}
	if err := w.eraseLine(); err != nil {
		return errFunc(err)
	}
	_, err := w.wr.WriteString(s)
	if err != nil {
		return errFunc(err)
	}
	return nil
}

func errFunc(err error) error {
	return errors.New(errPrefix + err.Error())
}
