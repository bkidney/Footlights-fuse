package data_test

import (
	"bytes"
	"testing"

	"github.com/bkidney/ProjectDistorage/data"
	"github.com/spf13/viper"
)

type TestVector struct {
	size    int
	links   []data.Link
	content []byte
}

func testVectors() (tests []TestVector) {

	tests = []TestVector{
		{ // An empty block
			size:    4096,
			links:   []data.Link{},
			content: []byte{},
		},
	}

	return
}

func testVectorsAsBlocks() (blocks []data.Block) {
	tests := testVectors()

	blocks = make([]data.Block, len(tests))

	for _, testV := range tests {
		blk := data.NewBlock()
		blk.Create(testV.size, testV.links, testV.content)
		blocks = append(blocks, *blk)
	}

	return
}

func loadDefaultSettings() {
	viper.SetDefault("cipher", "AES")
	viper.SetDefault("hash", "SHA256")
	viper.SetDefault("cipher-mode", "CBC")
}

func TestBlock_BlockCreation(t *testing.T) {

	loadDefaultSettings()

	tests := testVectors()

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

func TestBlock_BlockEncryption(t *testing.T) {

	loadDefaultSettings()

	tests := testVectorsAsBlocks()

	for i, blk := range tests {
		_, key_iv, encrypted_blk := data.Encrypt(blk)
		decrypted_blk := data.Decrypt(key_iv, encrypted_blk)

		if !bytes.Equal(blk.Bytes(), decrypted_blk) {
			t.Errorf("%d. Decrypted content does not match original ", i)
		}
	}

}
