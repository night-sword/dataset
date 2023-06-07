package region

import (
	"fmt"
	"testing"
)

func TestRegionDataset_Random(t *testing.T) {
	rd := NewRegionDataset()
	r, err := rd.Get(CodeEC)
	fmt.Println(r, err)
}
