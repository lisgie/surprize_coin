package protocol

type Transaction interface {
	Hash() [32]byte
	Encode() []byte
	//Decoding is not listed here, because it returns a different type for each tx (return value Transaction itself
	//is apparently not allowed)
	TxFee() uint64 //just fee is already occupied by the struct fields themselves
	Size() uint8
}
