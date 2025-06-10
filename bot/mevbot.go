package bot

import (
	"context"
	"fmt"
	"log"
	"mevbot/client"
	"mevbot/config"
	"mevbot/services"
	"os"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

type MevBot struct {
	config     *config.Config
	ethClient  *client.EthereumClient
	dexService *services.DexService // Used to call SwapExactInput
	inToken    common.Address
	outToken   common.Address
	logFile    *os.File // File để ghi log
	logger     *log.Logger
}

func NewMevBot(inTokenHex string, outTokenHex string, dexAddressHex string) (*MevBot, error) {
	cfg := config.LoadConfig()

	// Create new ETH client to connect to eth node
	ethClient, err := client.NewEthereumClient(cfg.RpcUrlWss, cfg.PrivateKey)
	if err != nil {
		return nil, fmt.Errorf("failed to create etherium client: %v", err)
	}

	inTokenAddr := common.HexToAddress(inTokenHex)
	outTokenAddr := common.HexToAddress(outTokenHex)

	dexAddr := common.HexToAddress(dexAddressHex)
	dexService, err := services.NewDexService(ethClient.Client, dexAddr)
	if err != nil {
		return nil, fmt.Errorf("failed to create dex service: %v", err)
	}

	// Mở file log
	logFile, err := os.OpenFile("mevbot.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return nil, fmt.Errorf("failed to open log file: %v", err)
	}

	// Tạo logger với file
	logger := log.New(logFile, "", log.LstdFlags)

	return &MevBot{
		config:     cfg,
		ethClient:  ethClient,
		dexService: dexService,
		inToken:    inTokenAddr,
		outToken:   outTokenAddr,
		logFile:    logFile,
		logger:     logger,
	}, nil
}

func (bot *MevBot) Run() error {
	// In thông tin khởi động ra console và file log
	fmt.Println("Starting MEV bot...")
	fmt.Println("RPC_URL (WebSocket):", bot.config.RpcUrlWss)
	fmt.Println("Loaded private key (hidden for security)")
	fmt.Println("Connected to Ethereum client")
	fmt.Println("From address:", bot.ethClient.Address.Hex())
	fmt.Println("ChainID:", bot.ethClient.ChainID)
	fmt.Println("DEX address:", bot.dexService.GetDexAddress().Hex())
	bot.logger.Println("Starting MEV bot...")
	bot.logger.Println("RPC_URL (WebSocket):", bot.config.RpcUrlWss)
	bot.logger.Println("Loaded private key (hidden for security)")
	bot.logger.Println("Connected to Ethereum client")
	bot.logger.Println("From address:", bot.ethClient.Address.Hex())
	bot.logger.Println("ChainID:", bot.ethClient.ChainID)
	bot.logger.Println("DEX address:", bot.dexService.GetDexAddress().Hex())

	tokenIn := bot.inToken
	tokenOut := bot.outToken
	fmt.Printf("TokenIn: %s TokenOut: %s\n", tokenIn.Hex(), tokenOut.Hex())
	bot.logger.Printf("TokenIn: %s TokenOut: %s\n", tokenIn.Hex(), tokenOut.Hex())

	txHashes := make(chan common.Hash)

	sub, err := bot.ethClient.SubscribePendingTxs(context.Background(), txHashes)
	if err != nil {
		bot.logger.Fatal("Failed to subscribe: ", err)
	}

	for {
		select {
		case err := <-sub.Err():
			bot.logger.Println("Subscription error:", err)
		case txHash := <-txHashes:
			bot.logger.Println("Received tx hash:", txHash.Hex()) // Log hash vào file
			var tx *types.Transaction
			var isPending bool
			// Thử lấy thông tin giao dịch tối đa 3 lần
			for attempts := 0; attempts < 3; attempts++ {
				tx, isPending, err = bot.ethClient.Client.TransactionByHash(context.Background(), txHash)
				if err == nil && tx != nil {
					break
				}
				bot.logger.Printf("Attempt %d: Không thể lấy thông tin giao dịch: %v", attempts+1, err)
				time.Sleep(500 * time.Millisecond)
			}
			if err != nil || tx == nil {
				bot.logger.Println("Không thể lấy thông tin giao dịch sau nhiều lần thử:", err)
				continue
			}

			// Kiểm tra xem giao dịch có gửi đến DEX không
			if tx.To() != nil && *tx.To() == common.HexToAddress("0xf924B30Df02e4139225DF419b7efeDF7D676132a") {
				// Ghi thông tin giao dịch vào file log và console
				logOutput := fmt.Sprintf("Phát hiện giao dịch trên DEX:\nHash giao dịch: %s\nĐang chờ xử lý: %v\n", tx.Hash().Hex(), isPending)
				fmt.Println("Phát hiện giao dịch trên DEX:")
				fmt.Println("Hash giao dịch:", tx.Hash().Hex())
				fmt.Println("Đang chờ xử lý:", isPending)

				// Lấy địa chỉ người gửi
				signer := types.NewEIP155Signer(bot.ethClient.ChainID)
				sender, err := signer.Sender(tx)
				if err != nil {
					bot.logger.Println("Không thể lấy địa chỉ người gửi:", err)
					continue
				}
				logOutput += fmt.Sprintf("Từ địa chỉ: %s\nĐến địa chỉ: %s\n", sender.Hex(), tx.To().Hex())
				fmt.Println("Từ địa chỉ:", sender.Hex())
				fmt.Println("Đến địa chỉ:", tx.To().Hex())

				logOutput += fmt.Sprintf("Giá trị: %s wei\nGiới hạn gas: %d\nGiá gas: %s wei\nNonce: %d\n", tx.Value().String(), tx.Gas(), tx.GasPrice().String(), tx.Nonce())
				fmt.Println("Giá trị:", tx.Value().String(), "wei")
				fmt.Println("Giới hạn gas:", tx.Gas())
				fmt.Println("Giá gas:", tx.GasPrice().String(), "wei")
				fmt.Println("Nonce:", tx.Nonce())

				// Xử lý dữ liệu (data)
				data := tx.Data()
				if len(data) > 0 {
					previewLen := 4
					if len(data) < previewLen {
						previewLen = len(data)
					}
					logOutput += fmt.Sprintf("Dữ liệu: %x... (độ dài: %d)\n", data[:previewLen], len(data))
					fmt.Printf("Dữ liệu: %x... (độ dài: %d)\n", data[:previewLen], len(data))
				} else {
					logOutput += "Dữ liệu: không có\n"
					fmt.Println("Dữ liệu: không có")
				}
				logOutput += "---\n"
				fmt.Println("---")
				bot.logger.Print(logOutput)
			} else {
				if tx.To() != nil {
					bot.logger.Println(*tx.To(), common.HexToAddress("0xf924B30Df02e4139225DF419b7efeDF7D676132a"))
				}
				bot.logger.Println("Giao dịch không liên quan đến DEX, bỏ qua:", txHash.Hex())
			}
		}
	}
}

// Đóng file log khi bot dừng (nếu cần)
func (bot *MevBot) Close() {
	if bot.logFile != nil {
		bot.logFile.Close()
	}
}
