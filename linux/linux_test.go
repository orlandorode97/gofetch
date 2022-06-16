package linux

import (
	"errors"
	"testing"

	"github.com/OrlandoRomo/gofetch/command"
)

var errUnableCommand = errors.New("unable to execute command")

type mockLinux struct{
	value string
	err error
}

func (m *mockLinux) GetName() (string, error) {
	return m.value, m.err
}

func Test_GetName(t *testing.T) {
	tcs:= []struct{
		Description string
		Linux command.Namer 
		Expected string
		ErrExpected error
	}{
		{
			Description: "success - got linux name",
			Linux: 	&mockLinux{
				value: "linux user name",
			},
			Expected: "linux user name",
		},
		{
			Description: "failed - cannot execute get name command",
			Linux: &mockLinux{
				err: errUnableCommand,
			},
			ErrExpected: errUnableCommand,

		},
	}

	for _, tc:= range tcs {
		t.Run(tc.Description, func(t *testing.T){
			name, err:= tc.Linux.GetName()
			if !errors.Is(err, tc.ErrExpected) {
				t.Errorf("received %s expected %s", err, tc.ErrExpected)
			}

			if tc.Expected != name {
				t.Fatalf("recived %s expected %s", name, tc.Expected)
			}
		})
	}
}
