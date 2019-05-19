package main

import (
	"fmt"
	"os"

	"golang.org/x/xerrors"
)

// Preparation
//  export GO111MODULE=on
//  go mod init project-name

type ApplicationError struct {
	level string
	code  int
	msg   string
	err   error
	frame xerrors.Frame
}

type Option func(*ApplicationError)

//func NewApplicationError(level string, code int, msg string) *ApplicationError {
//func NewApplicationError(level string, code int, msg string, opts ...Option) *ApplicationError {
func NewApplicationError(err error, level string, code int, msg string) *ApplicationError {
	appErr := &ApplicationError{
		level: level,
		code:  code,
		msg:   msg,
		//err:   nil, // Pattern1, 2
		err:   err, // Pattern3
		frame: xerrors.Caller(1),
	}
	/*
		for _, opt := range opts {
			opt(appErr) // Pattern2 stack frame
		}
	*/
	return appErr
}

// Wrap Pattern 2
func Wrap(err error) Option {
	return func(a *ApplicationError) {
		if err != nil {
			a.err = err
			a.frame = xerrors.Caller(1)
		}
	}
}

// Wrap Pattern 1
/*
func (e ApplicationError) Wrap(next error) error {
	e.err = next
	e.frame = xerrors.Caller(1)
	return &e
}
*/
func (e *ApplicationError) Error() string {
	return fmt.Sprintf("%s: code=%d, msg=%s", e.level, e.code, e.msg)
}
func (e *ApplicationError) Unwrap() error {
	return e.err
}

func (e *ApplicationError) Format(s fmt.State, v rune) { xerrors.FormatError(e, s, v) }

func (e *ApplicationError) FormatError(p xerrors.Printer) (next error) {
	p.Print(e.Error())
	e.frame.Format(p)
	return e.err
}

func fileOpen(fname string) error {
	file, err := os.Open(fname)
	if err != nil {
		switch e := err.(type) {
		case *os.PathError:
			return xerrors.Errorf("Error in fileOpen(\"%v\"): %w", e.Path, e.Err)
		default:
			return xerrors.Errorf("Error in fileOpen(): %w", err)
		}
	}
	defer file.Close()
	return nil
}

func main() {
	fmt.Print("Result: ")
	if err := fileOpen("null.txt"); err != nil {
		// Wrap Pattern 1
		//err = NewApplicationError("err", 100, "open err").Wrap(err)
		// Wrap Pattern 2
		//err = NewApplicationError("err", 100, "open err", Wrap(err))
		// Wrap Pattern 3
		err = NewApplicationError(err, "err", 100, "open err")
		fmt.Printf("%+v\n", err)
	} else {
		fmt.Println("OK")
	}
}

/*

// Wrap Pattern 1
$ go run main.go
Result: err: code=100, msg=open err:
    main.main
        /Users/home/github/golang-sample/error-wrap-pattern/main.go:73
  - Error in fileOpen("null.txt"):
    main.fileOpen
        /Users/home/github/golang-sample/error-wrap-pattern/main.go:61
  - no such file or directory

// Wrap Pattern 2
$ go run main.go
Result: err: code=100, msg=open err:
    main.NewApplicationError
        /Users/home/github/golang-sample/error-wrap-pattern/main.go:35
  - Error in fileOpen("null.txt"):
    main.fileOpen
        /Users/home/github/golang-sample/error-wrap-pattern/main.go:78
  - no such file or directory

// Wrap Pattern 3
$ go run main.go
Result: err: code=100, msg=open err:
    main.main
        /Users/home/github/golang-sample/error-wrap-pattern/main.go:98
  - Error in fileOpen("null.txt"):
    main.fileOpen
        /Users/home/github/golang-sample/error-wrap-pattern/main.go:81
  - no such file or directory
*/
