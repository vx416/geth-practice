package client

import (
	"context"
	"crypto/ecdsa"
	"eth-practice/client/abi"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/pkg/errors"
	"golang.org/x/crypto/sha3"
)

// https://rinkeby.infura.io/v3/247f1fc54d244f1c969dd3bf8f5de22c
func NewTokenClient(nodeURL string, contractAddress string) (*TokenClient, error) {
	ctx := context.Background()
	client, err := ethclient.Dial(nodeURL)
	if err != nil {
		return nil, err
	}

	address := common.HexToAddress(contractAddress)
	contract, err := abi.NewAbi(address, client)

	chainID, err := client.NetworkID(ctx)
	if err != nil {
		err = errors.WithStack(err)
		return nil, err
	}

	signer := types.NewEIP155Signer(chainID)

	return &TokenClient{
		client:          client,
		ContractAddress: address,
		Contract:        contract,
		signer:          signer,
	}, nil
}

type TokenClient struct {
	client          *ethclient.Client
	ContractAddress common.Address
	Contract        *abi.Abi
	signer          types.EIP155Signer
}

func (client *TokenClient) SignerFunc(priHex string) (func(from common.Address, t *types.Transaction) (*types.Transaction, error), error) {
	privateKey, _, err := client.GetAddrFromPriKey(priHex)
	if err != nil {
		return nil, err
	}
	return func(from common.Address, t *types.Transaction) (*types.Transaction, error) {
		signedTx, err := types.SignTx(t, client.signer, privateKey)
		if err != nil {
			return nil, err
		}
		return signedTx, nil
	}, nil
}

// GetAddrFromPriKey 由私鑰計算出公鑰和地址
func (client *TokenClient) GetAddrFromPriKey(priHex string) (privateKey *ecdsa.PrivateKey, addressHex string, err error) {
	privateKey, err = crypto.HexToECDSA(priHex)
	if err != nil {
		return nil, "", err
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		return nil, "", err
	}
	publicKeyBytes := crypto.FromECDSAPub(publicKeyECDSA)

	hash := sha3.NewLegacyKeccak256()
	hash.Write(publicKeyBytes[1:])

	addressHex = hexutil.Encode(hash.Sum(nil)[12:])
	return privateKey, addressHex, nil
}
