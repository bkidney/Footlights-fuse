package data

type link struct {
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
	tag                [54]byte
	lenght             int64
	uri_lenght         int64
	cipher_name_len    int64
	decryption_key_len int64
	uri                []byte
	cipher             []byte
	decryption_key     []byte
}
