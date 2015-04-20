open Printf

type link = 
  { tag : string
  ; len : int
  ; uri_len : int
  ; cipher_name_len : int
  ; decryption_key_len : int
  ; uri : string 
  ; cipher : string
  ; decryption_key : string }

let unpack_link bits link =
  bitmatch bits with
  | { tag : 54 ;                            (* Link Tag *)
      len : 8 ;                             (* Link Length *)
      uri_len : 8 ;                         (* URI Length *)
      cipher_name_len : 8 ;                 (* Cipher Name Length *)
      decryption_key_len : 16 ;             (* Decryption Key Length *)
      uri : uri_len ;                       (* URI Value *)
      cipher : cipher_name_len ;            (* Cipher Name *)
      decryption_key : decryption_key_len } (* Decryption key*)
  -> link
  | { _ } as unmatched -> 
      eprintf "Link was unmatched:\n";
      Bitstring.hexdump_bitstring stderr unmatched;
      exit 1

