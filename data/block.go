package data

import (
	"crypto/rand"
	"fmt"
	"math"
)

type block struct {
	/*
	   magic - 6
	   block_len - 2
	   num_links - 2
	   offset - 4
	   content_len - 4
	   links - num_links * 768
	   content - content_len
	   padding
	*/
	magic       [6]byte
	block_len   int8
	num_links   int8
	offset      int32
	content_len int32
	links       []link
	content     []byte
	padding     []byte
}

func (blk *block) create(block_size int, links []link, content []byte) {

	blk.magic = [6]byte{0x5, 0x4, 0x3, 0x2, 0x1, 0x0}
	blk.block_len = int8(math.Log2(float64(block_size)))
	blk.num_links = int8(len(links))
	blk.offset = int32(blk.num_links)*96 + 16
	blk.content_len = int32(len(content))

	blk.links = links
	blk.content = content

	padding_len := block_size - 16 - int(blk.content_len) - (int(blk.num_links) * 96)

	padding := make([]byte, padding_len)
	_, err := rand.Read(padding)

	if err != nil {
		fmt.Println(err)
	}

}
