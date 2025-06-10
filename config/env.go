package config

import (
	"log"
	"os"
	"sync"

	"github.com/joho/godotenv"
)

type Env struct {
	rpcUrlWss          string
	rpcUrlHttp         string
	metamaskPrivateKey string
	baoTokenAddr       string
	locTokenAddr       string
	dexAddr            string
	poolAddr           string
}

var env *Env

var once sync.Once

func GetEnv() *Env {
	once.Do(func() {

		err := godotenv.Load()
		if err != nil {
			log.Fatal("Error loading .env file")
		}

		env = &Env{
			rpcUrlWss:          os.Getenv("RPC_URL_WSS"),
			rpcUrlHttp:         os.Getenv("RPC_URL_HTTP"),
			metamaskPrivateKey: os.Getenv("METAMASK_PRIVATE_KEY"),
			baoTokenAddr:       os.Getenv("BAO_TOKEN"),
			locTokenAddr:       os.Getenv("LOC_TOKEN"),
			dexAddr:            os.Getenv("DEX_ADDRESS"),
			poolAddr:           os.Getenv("POOL_ADDRESS"),
		}
	})
	return env
}

func (e Env) GetRpcUrlWss() string {
	return e.rpcUrlWss
}

func (e Env) GetRpcUrlHttp() string {
	return e.rpcUrlHttp
}

func (e Env) GetMetamaskPrivateKey() string {
	return e.metamaskPrivateKey
}

func (e Env) GetBaoTokenAddr() string {
	return e.baoTokenAddr
}

func (e Env) GetLocTokenAddr() string {
	return e.locTokenAddr
}

func (e Env) GetDexAddr() string {
	return e.dexAddr
}

func (e Env) GetPoolAddr() string {
	return e.poolAddr
}
