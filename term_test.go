package term

import "testing"
import "flag"

var dev = flag.String("device", "/dev/tty", "device to use")

// assert that Term implements the same method set across
// all supported platforms
var _ interface {
	Available() (int, error)
	Buffered() (int, error)
	Close() error
	DTR() (bool, error)
	Flush() error
	RTS() (bool, error)
	Read(b []byte) (int, error)
	Restore() error
	SendBreak() error
	SetCbreak() error
	SetDTR(v bool) error
	SetOption(options ...func(*Term) error) error
	SetRTS(v bool) error
	SetRaw() error
	SetSpeed(baud int) error
	Write(b []byte) (int, error)
} = new(Term)

func TestTermSetCbreak(t *testing.T) {
	tt := opendev(t)
	defer tt.Close()
	if err := tt.SetCbreak(); err != nil {
		t.Fatal(err)
	}
}

func TestTermSetRaw(t *testing.T) {
	tt := opendev(t)
	defer tt.Close()
	if err := tt.SetRaw(); err != nil {
		t.Fatal(err)
	}
}

func TestTermSetSpeed(t *testing.T) {
	tt := opendev(t)
	defer tt.Close()
	if err := tt.SetSpeed(57600); err != nil {
		t.Fatal(err)
	}
}

func TestTermRestore(t *testing.T) {
	tt := opendev(t)
	defer tt.Close()
	if err := tt.Restore(); err != nil {
		t.Fatal(err)
	}
}

func opendev(t *testing.T) *Term {
	tt, err := Open(*dev)
	if err != nil {
		t.Fatal(err)
	}
	return tt
}
