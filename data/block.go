package data

/*
Todo:
	- Block should have functions
		- ToBytes()
	- Where do I build the links / Encrypt Blocks?
		- Maybe there should be a builder "class"
		- Remove block from having to know about crypto
	- How to I test block encryption?
		- Is Encrypt = Decrypt enough?
*/

import (
	"bytes"
	"crypto/rand"
	"encoding/binary"
	"fmt"
	"math"

	"github.com/bkidney/ProjectDistorage/crypto"
)

type Block struct {
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
	links       []Link
	content     []byte
	padding     []byte
}

func NewBlock() *Block {
	return &Block{}
}

func (blk *Block) Create(block_size int, links []Link, content []byte) {

	blk.magic = [6]byte{0xF0, 0x07, 0xDA, 0x7A, '\r', '\n'}
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

func Encrypt(blk Block) (name []byte, key_iv []byte, encrypted_blk []byte) {

	plaintext := blk.GetBytes()
	key_iv = crypto.GetHash(plaintext)
	key := key_iv[16:]
	iv := key_iv[:16]

	encrypted_blk = crypto.Encrypt(plaintext, key, iv)
	name = crypto.GetHash(encrypted_blk)

	return
}

func Decrypt(key_iv, blk []byte) (decrypted_blk []byte) {

	key := key_iv[16:]
	iv := key_iv[:16]

	decrypted_blk = crypto.Decrypt(blk, key, iv)

	return
}

func (blk *Block) GetBytes() []byte {
	var bin_buf bytes.Buffer
	binary.Write(&bin_buf, binary.BigEndian, blk)

	return bin_buf.Bytes()
}

func (blk *Block) ContentSize() int32 {
	return blk.content_len
}

func (blk *Block) Offset() int32 {
	return blk.offset
}

func (blk *Block) NumLinks() int8 {
	return blk.num_links
}
