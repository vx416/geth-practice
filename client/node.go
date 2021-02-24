package client

var (
	Rinkeby = &Node{
		URL:       "https://rinkeby.infura.io/v3/247f1fc54d244f1c969dd3bf8f5de22c",
		NetworkID: 4,
	}
)

type Node struct {
	URL       string
	NetworkID int64
}

// 0x478F041ca93674c220328fB6EddfFD3105f8F5ca
