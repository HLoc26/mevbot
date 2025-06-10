// services/dex.go
package services

import (
	"context"
	"fmt"
	"math/big"
	"mevbot/contracts/dex"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

type DexService struct {
	client      *ethclient.Client
	dexInstance *dex.Dex
	dexAddr     common.Address
}

// Load the DEX contract using an ethClient and the DEX's address itself
func NewDexService(client *ethclient.Client, dexAddr common.Address) (*DexService, error) {
	dexInstance, err := dex.NewDex(dexAddr, client)
	if err != nil {
		return nil, fmt.Errorf("failed to load DEX contract: %v", err)
	}

	return &DexService{
		client:      client,
		dexInstance: dexInstance,
		dexAddr:     dexAddr,
	}, nil
}

func (ds *DexService) CreateLiquidityPool(auth *bind.TransactOpts, tokenIn, tokenOut common.Address, feePercentage, maxSwapPercentage *big.Int, poolName, poolSymbol string) error {
	fmt.Println("Creating liquidity pool...")
	tx, err := ds.dexInstance.CreateLiquidityPool(auth, tokenIn, tokenOut, feePercentage, maxSwapPercentage, poolName, poolSymbol)
	if err != nil {
		return fmt.Errorf("CreatePool failed: %v", err)
	}

	fmt.Printf("CreatePool tx: %s\n", tx.Hash().Hex())
	_, err = bind.WaitMined(context.Background(), ds.client, tx)
	return err
}

func (ds *DexService) GetPoolAddress(tokenIn, tokenOut common.Address) (common.Address, error) {
	pool, err := ds.dexInstance.LiquidityPools(nil, tokenIn, tokenOut)
	if err != nil {
		return common.Address{}, fmt.Errorf("failed to get pool address: %v", err)
	}
	return pool, nil
}

func (ds *DexService) AddLiquidity(auth *bind.TransactOpts, tokenIn, tokenOut common.Address, amountTokenIn, amountTokenOut, minLiquidity *big.Int) error {
	fmt.Println("Adding liquidity...")
	tx, err := ds.dexInstance.AddLiquidity(auth, tokenIn, tokenOut, amountTokenIn, amountTokenOut, minLiquidity)
	if err != nil {
		return fmt.Errorf("AddLiquidity failed: %v", err)
	}

	fmt.Printf("AddLiquidity tx: %s\n", tx.Hash().Hex())
	receipt, err := bind.WaitMined(context.Background(), ds.client, tx)
	if err != nil {
		return fmt.Errorf("WaitMined for AddLiquidity failed: %v", err)
	}

	if receipt.Status == 0 {
		return fmt.Errorf("transaction add liquidity failed: %s", tx.Hash().Hex())
	}

	fmt.Printf("Transaction add liquidity succeeded: %s\n", tx.Hash().Hex())
	return nil
}

func (ds *DexService) SwapExactInput(auth *bind.TransactOpts, tokenIn, tokenOut common.Address, amountIn, minAmountOut *big.Int) error {
	fmt.Println("Swapping tokens...")
	tx, err := ds.dexInstance.SwapExactInput(auth, tokenIn, tokenOut, amountIn, minAmountOut)
	if err != nil {
		return fmt.Errorf("SwapExactInput failed: %v", err)
	}

	fmt.Printf("Swap tx: %s\n", tx.Hash().Hex())
	receipt, err := bind.WaitMined(context.Background(), ds.client, tx)
	if err != nil {
		return fmt.Errorf("WaitMined for Swap failed: %v", err)
	}

	if receipt.Status == 0 {
		return fmt.Errorf("transaction failed: %s", tx.Hash().Hex())
	}

	fmt.Printf("Transaction succeeded: %s\n", tx.Hash().Hex())
	fmt.Printf("Tx sent: https://sepolia.arbiscan.io/tx/%s\n", tx.Hash().Hex())
	return nil
}

func (ds *DexService) GetDexAddress() common.Address {
	return ds.dexAddr
}
