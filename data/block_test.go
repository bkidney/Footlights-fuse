package data_test

import (
	"github.com/bkidney/ProjectDistorage/data"
	"testing"
)

func TestBlock_BlockCreation(t *testing.T) {

	var tests = []struct {
		size    int
		links   []data.Link
		content []byte
	}{
		{
			size:    4096,
			links:   []data.Link{},
			content: []byte{},
		},
	}

	for i, tt := range tests {
		blk := data.NewBlock()
		blk.Create(tt.size, tt.links, tt.content)

		if blk.ContentSize() != 0 {
			t.Errorf("%d. Content size mismatch: \nGot: %d\nExp: %d) ", i, blk.ContentSize(), len(tt.content))
		}
		if blk.NumLinks() != 0 {
			t.Errorf("%d. Content size mismatch: \nGot: %d\nExp: %d) ", i, blk.NumLinks(), len(tt.links))
		}
		if blk.Offset() != int32(len(tt.links))*96+16 {
			t.Errorf("%d. Content size mismatch: \nGot: %d\nExp: %d) ", i, blk.Offset(), len(tt.links)*96+16)
		}
	}
}
