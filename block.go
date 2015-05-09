package data

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
	block_len   int16
	num_links   int16
	offset      int32
	content_len int32
	link        []byte
	content     []byte
	padding     []byte
}
