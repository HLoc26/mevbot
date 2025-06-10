// services/token.go
package services

import (
	"context"
	"fmt"
	"math"
	"math/big"
	"mevbot/contracts/erc20"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

type TokenService struct {
	client *ethclient.Client
}

func NewTokenService(client *ethclient.Client) *TokenService {
	return &TokenService{client: client}
}

// Approve a spender (spenderAddr) to use an amount (amount) of token (tokenAddr) on user's (auth) behalf
func (ts *TokenService) ApproveToken(auth *bind.TransactOpts, tokenAddr, spenderAddr common.Address, amount *big.Int) error {
	tokenInstance, err := erc20.NewErc20(tokenAddr, ts.client)
	if err != nil {
		return fmt.Errorf("failed to create token instance: %v", err)
	}

	fmt.Printf("Approving token %s...\n", tokenAddr.Hex())
	tx, err := tokenInstance.Approve(auth, spenderAddr, amount)
	if err != nil {
		return fmt.Errorf("approve failed: %v", err)
	}

	fmt.Printf("Approve tx: %s\n", tx.Hash().Hex())
	_, err = bind.WaitMined(context.Background(), ts.client, tx)
	return err
}

// Get an address's balance of a token
func (ts *TokenService) GetBalance(tokenAddr, ownerAddr common.Address) (*big.Int, error) {
	tokenInstance, err := erc20.NewErc20(tokenAddr, ts.client)
	if err != nil {
		return nil, fmt.Errorf("failed to create token instance: %v", err)
	}

	balance, err := tokenInstance.BalanceOf(nil, ownerAddr)
	if err != nil {
		return nil, fmt.Errorf("failed to get balance: %v", err)
	}

	return balance, nil
}

func (ts *TokenService) PrintBalance(tokenAddr, ownerAddr common.Address, tokenName string) error {
	balance, err := ts.GetBalance(tokenAddr, ownerAddr)
	if err != nil {
		return err
	}

	tokenFloat := new(big.Float).Quo(new(big.Float).SetInt(balance), big.NewFloat(math.Pow10(18)))
	fmt.Printf("%s balance: %s (%f)\n", tokenName, balance.String(), tokenFloat)
	return nil
}
