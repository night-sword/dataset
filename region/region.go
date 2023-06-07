package region

import (
	_ "embed"
	"encoding/json"
	"fmt"
	"sync"
)

type Region struct {
	Timezone string `json:"timezone"`
	Language string `json:"language"`
	Code     Code   `json:"code"`
	CName    string `json:"cname"`
}

type RegionDataset struct {
	regions map[Code]*Region
}

func NewRegionDataset() *RegionDataset {
	inst := &RegionDataset{}

	_regionOnce.Do(func() {
		dataset := make(map[string]*Region)
		err := json.Unmarshal([]byte(_regionSet), &dataset)
		if err != nil {
			panic("region set unmarshal fail")
		}

		regions := make(map[Code]*Region, len(dataset))
		for _, code := range CodeValues() {
			if r, ok := dataset[code.String()]; ok {
				regions[code] = r
			} else {
				panic(fmt.Sprintf("region %s has not dataset", code.String()))
			}
		}

		inst.regions = regions
	})

	return inst
}

func (inst *RegionDataset) Get(code Code) (region *Region, err error) {
	if err = code.Check(); err != nil {
		return
	}

	region = inst.regions[code]

	return

}

//go:embed region.json
var _regionSet string
var _regionOnce sync.Once
