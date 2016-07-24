package unsafe2

import (
	"testing"

	"github.com/CodisLabs/codis/pkg/utils/assert"
)

func TestMakeGoSlice(t *testing.T) {
	n := MinOffheapSlice - 1
	s := MakeSlice(n)
	assert.Must(s != nil)
	_, ok := s.(*goSlice)
	assert.Must(ok)
}

func TestMakeJeSlice(t *testing.T) {
	n := MinOffheapSlice * 2
	SetMaxOffheapBytes(n * 2)

	s1 := MakeSlice(n)
	_, ok1 := s1.(*jeSlice)
	assert.Must(ok1 && len(s1.Buffer()) == n)
	defer FreeSlice(s1)

	s2 := MakeSlice(n)
	_, ok2 := s2.(*jeSlice)
	assert.Must(ok2 && len(s2.Buffer()) == n)
	defer FreeSlice(s2)

	assert.Must(OffheapBytes() == n*2)

	s3 := MakeSlice(n)
	_, ok3 := s3.(*goSlice)
	assert.Must(ok3 && len(s3.Buffer()) == n)
	defer FreeSlice(s3)

	assert.Must(OffheapBytes() == n*2)

	FreeSlice(s2)
	assert.Must(OffheapBytes() == n)

	s4 := MakeSlice(n)
	_, ok4 := s4.(*jeSlice)
	assert.Must(ok4 && len(s4.Buffer()) == n)
	defer FreeSlice(s4)

	assert.Must(OffheapBytes() == n*2)

	s5 := MakeOffheapSlice(n)
	assert.Must(s5 != nil && len(s5.Buffer()) == n)
	defer FreeSlice(s5)

	assert.Must(OffheapBytes() == n*3)
}
