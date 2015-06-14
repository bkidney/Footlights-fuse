package data

/*
	- Encoding?
*/

type Link struct {
	/*
	   magic - 54
	   length - 16
	   uri_length - 8
	   cipher_name_len - 8
	   decryption_key_len - 16
	   uri - uri_len
	   cipher - cipher_name_len
	   decryption_key - decryption_key_len
	*/
	magic              [6]byte
	length             int16
	uri_length         int8
	cipher_name_len    int8
	decryption_key_len int16
	uri                []byte
	cipher             []byte
	decryption_key     []byte
}

func (lnk *Link) create(name []byte, key []byte, cipher_name string) {

	lnk.magic = [6]byte{'L', 'I', 'N', 'K', '\r', '\n'}

	lnk.cipher = []byte(cipher_name)
	lnk.cipher_name_len = int8(len(lnk.cipher))

	lnk.uri = name
	lnk.uri_length = int8(len(lnk.uri))

	lnk.decryption_key = key
	lnk.decryption_key_len = int16(len(lnk.decryption_key))

	lnk.length = 0
}
