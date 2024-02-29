package safe

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_Group(t *testing.T) {
	t.Parallel()

	r := require.New(t)

	var wg Group
	for i := 0; i < 10; i++ {
		wg.Go(func() error {
			panic(ErrTest)
		})
	}

	err := wg.Wait()
	r.Error(err)
	r.True(errors.Is(err, ErrTest))
}
