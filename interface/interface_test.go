package myinterface

import (
	"bytes"
	"testing"
)

func TestHello(t *testing.T) {
	writer := bytes.Buffer{}
	Greet(&writer, "peppe")
	want := "Ciao peppe"

	if writer.String() != want {
		t.Errorf("got: %q is different from want: %q", writer.String(), want)
	}
}
