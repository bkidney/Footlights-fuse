open Printf

type block = 
  { magic : string
  ; block_len : int
  ; num_links : int
  ; offset : int
  ; content_len : int
  ; links : string (* Need to handle array of links *)
  ; content : string 
  ; padding : string }

let unpack_block bits link =
  bitmatch bits with
  | { magic : 6 ;                  (* Magic Number *)
      block_len : 2 ;               (* Number of Block *)
      num_links : 2 ;               (* Number of links *)
      offset : 4 ;                 (* Offset of User Content *)
      content_len : 4 ;            (* User Content Length*)
      links : num_links*768 : bitstring ;        (* Link Data *)
      content : content_len ;       (* User Content *)
      padding : -1 : bitstring }    (* Padding to block size multiple of 2^x *)
  -> printf "Found a block."
  | { _ } as unmatched -> 
      eprintf "Block was unmatched:\n";
      Bitstring.hexdump_bitstring stderr unmatched;
      exit 1

