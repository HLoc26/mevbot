package client

import (
	"context"
	"crypto/ecdsa"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
)

type EthereumClient struct {
	Client     *ethclient.Client // ETH Client, connect to an Etherium node via RPC URL
	RpcClient  *rpc.Client
	PrivateKey *ecdsa.PrivateKey // Caller's wallet private key, use for signing txs
	PublicKey  *ecdsa.PublicKey  // Caller's wallet public key, generated from PrivateKey, use for signing txs
	Address    common.Address    // ETH;s address, generated from PublicKey
	ChainID    *big.Int          // ETH network's chain ID
}

// privateHex -> Private ECDSA -> Public ECDSA -> Address
func NewEthereumClient(rpcUrl, privateKeyHex string) (*EthereumClient, error) {
	privateKey, err := crypto.HexToECDSA(privateKeyHex)
	if err != nil {
		return nil, err
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("Cannot cast public key")
	}

	address := crypto.PubkeyToAddress(*publicKeyECDSA)

	// Connect to an etherium node
	client, err := ethclient.Dial(rpcUrl)
	if err != nil {
		return nil, err
	}
	chainId, err := client.ChainID(context.Background())

	rpcClient, err := rpc.Dial(rpcUrl)

	return &EthereumClient{
		Client:     client,
		PrivateKey: privateKey,
		RpcClient:  rpcClient,
		PublicKey:  publicKeyECDSA,
		Address:    address,
		ChainID:    chainId,
	}, nil
}

func (ec *EthereumClient) SubscribePendingTxs(ctx context.Context, ch chan common.Hash) (ethereum.Subscription, error) {
	return ec.RpcClient.Subscribe(ctx, "eth", ch, "newPendingTransactions")
}
