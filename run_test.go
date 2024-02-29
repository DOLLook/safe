package safe

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_Run(t *testing.T) {
	t.Parallel()

	tcs := []struct {
		name string
		fn   RunFn
		exp  string
		err  bool
	}{
		{
			name: "no error/panic",
			fn: func() error {
				return nil
			},
		},
		{
			name: "error",
			fn: func() error {
				return ERRTest
			},
			err: true,
		},
		{
			name: "panic w/ error",
			fn: func() error {
				panic(ERRTest)
			},
			err: true,
		},
		{
			name: "panic w/ string",
			fn: func() error {
				panic(ERRTest.Error())
			},
			err: true,
		},
		{
			name: "panic w/ struct",
			fn: func() error {
				panic(struct {
					foo string
				}{
					foo: "bar",
				})
			},
			err: true,
			exp: "{foo:bar}",
		},
	}

	for _, tc := range tcs {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			r := require.New(t)

			err := Run(tc.fn)
			if !tc.err {
				r.NoError(err)
				return
			}

			if tc.exp == "" {
				tc.exp = ERRTest.Error()
			}

			r.Error(err)
			r.Equal(tc.exp, err.Error())
		})
	}

}
