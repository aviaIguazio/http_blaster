package data_generator

import (
	"strconv"
	"testing"
)

func Test_igz_tsdb_item_v2_init(t *testing.T) {
	gen := MemoryGenerator{}
	gen.GenerateRandomData(strconv.Itoa(1))
}
