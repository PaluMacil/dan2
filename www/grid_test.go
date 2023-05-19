package www_test

import (
	"github.com/PaluMacil/dan2/www"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMakeRows(t *testing.T) {
	rows := www.MakeRows([]string{"bob", "sam", "houston", "greg"}, 2)
	assert.Equal(t, len(rows), 2)
	assert.Equal(t, len(rows[0]), 2)
	assert.Equal(t, rows[0][0], "bob")
	assert.Equal(t, rows[0][1], "sam")
	assert.Equal(t, len(rows[1]), 2)
	assert.Equal(t, rows[1][0], "houston")
	assert.Equal(t, rows[1][1], "greg")
}

func TestMakeRowsZero(t *testing.T) {
	rows := www.MakeRows([]string{"bob", "sam", "houston", "greg"}, 0)
	assert.Equal(t, len(rows), 0)
}

func TestMakePartial(t *testing.T) {
	rows := www.MakeRows([]string{"bob", "sam", "houston", "greg"}, 3)
	assert.Equal(t, len(rows), 2)
	assert.Equal(t, len(rows[0]), 3)
	assert.Equal(t, rows[0][0], "bob")
	assert.Equal(t, rows[0][1], "sam")
	assert.Equal(t, rows[0][2], "houston")
	assert.Equal(t, len(rows[1]), 1)
	assert.Equal(t, rows[0][0], "greg")
}
