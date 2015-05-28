package data

type Link struct {
	/*
	   tag - 54
	   len - 8
	   uri_len - 8
	   cipher_name_len - 8
	   decryption_key_len - 16
	   uri - uri_len
	   cipher - cipher_name_len
	   decryption_key - decryption_key_len
	*/
	magic              [6]byte
	length             int64
	uri_length         int64
	cipher_name_len    int64
	decryption_key_len int64
	uri                []byte
	cipher             []byte
	decryption_key     []byte
}

func (lnk *Link) create() {

	lnk.magic = [6]byte{'L', 'I', 'N', 'K', '\r', '\n'}
	lnk.length = 0

}
