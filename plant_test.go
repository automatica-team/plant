package plant

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNew(t *testing.T) {
	p, err := New("example.yml")
	require.NoError(t, err)
	require.NotEmpty(t, p.Config)
}
