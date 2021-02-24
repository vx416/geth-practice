package client

import (
	"context"
	"fmt"
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/assert"
)

func TestBalanceOf(t *testing.T) {
	client, err := NewTokenClient("https://rinkeby.infura.io/v3/247f1fc54d244f1c969dd3bf8f5de22c", "0x56f77c99506656400Daf4F79c34FC3b0Cbf11Fd0")
	assert.NoError(t, err)
	balance, err := client.Contract.BalanceOf(&bind.CallOpts{}, common.HexToAddress("0x5C344C04A590990007623C8Ce8e802B51710F619"))
	assert.NoError(t, err)
	t.Logf("%+v", balance)
}

func TestTransfer(t *testing.T) {
	client, err := NewTokenClient("https://rinkeby.infura.io/v3/247f1fc54d244f1c969dd3bf8f5de22c", "0x56f77c99506656400Daf4F79c34FC3b0Cbf11Fd0")
	assert.NoError(t, err)
	signerFunc, err := client.SignerFunc("1012a797ee3aecb39640dbc390fbd0c61005b0e97936a76da6bbb7f3208f4518")
	assert.NoError(t, err)
	tran, err := client.Contract.Transfer(&bind.TransactOpts{
		Context:  context.Background(),
		From:     common.HexToAddress("0x5C344C04A590990007623C8Ce8e802B51710F619"),
		GasLimit: 100000,
		Signer:   signerFunc,
	}, common.HexToAddress("0x4b4486025A4fb09Fccd6b8245EBa9b459260680D"), big.NewInt(5000000000000000000))
	assert.NoError(t, err)
	t.Logf("%+v", tran)
}

func TestMint(t *testing.T) {
	client, err := NewTokenClient("https://rinkeby.infura.io/v3/247f1fc54d244f1c969dd3bf8f5de22c", "0x56f77c99506656400Daf4F79c34FC3b0Cbf11Fd0")
	assert.NoError(t, err)
	signerFunc, err := client.SignerFunc("1012a797ee3aecb39640dbc390fbd0c61005b0e97936a76da6bbb7f3208f4518")
	assert.NoError(t, err)

	mintAmount := new(big.Int)
	_, err = fmt.Sscan("50000000000000000000000", mintAmount)
	assert.NoError(t, err)

	tran, err := client.Contract.Mint(&bind.TransactOpts{
		Context:  context.Background(),
		From:     common.HexToAddress("0x5C344C04A590990007623C8Ce8e802B51710F619"),
		GasLimit: 100000,
		Signer:   signerFunc,
	}, common.HexToAddress("0x4b4486025A4fb09Fccd6b8245EBa9b459260680D"), mintAmount)
	assert.NoError(t, err)
	t.Logf("%+v", tran)
}
