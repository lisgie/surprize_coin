package p2p

const HEADER_LEN = 5

type Header struct {
	Len    uint32
	TypeID uint8
}

const (
	FUNDSTX_BRDCST  = 1
	ACCTX_BRDCST    = 2
	CONFIGTX_BRDCST = 3
	BLOCK_BRDCST    = 4

	FUNDSTX_REQ  = 10
	ACCTX_REQ    = 11
	CONFIGTX_REQ = 12
	BLOCK_REQ    = 13
	ACC_REQ      = 14

	FUNDSTX_RES  = 20
	ACCTX_RES    = 21
	CONFIGTX_RES = 22
	BLOCK_RES    = 23
	ACC_RES      = 24

	NEIGHBOR_REQ = 30

	NEIGHBOR_RES = 40

	TIME_BRDCST = 50

	MINER_PING = 100
	MINER_PONG = 101

	//Error codes
	NOT_FOUND = 110
)