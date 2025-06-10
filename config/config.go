// config/config.go
package config

type Config struct {
	RpcUrlWss    string
	RpcUrlHttp   string
	PrivateKey   string
	GasLimit     uint64
	BaoTokenAddr string
	LocTokenAddr string
	DexAddr      string
	PoolAddr     string
}

func LoadConfig() *Config {
	env := GetEnv()
	return &Config{
		RpcUrlWss:    env.GetRpcUrlWss(),
		RpcUrlHttp:   env.GetRpcUrlHttp(),
		PrivateKey:   env.GetMetamaskPrivateKey(),
		BaoTokenAddr: env.GetBaoTokenAddr(),
		LocTokenAddr: env.GetLocTokenAddr(),
		DexAddr:      env.GetDexAddr(),
		PoolAddr:     env.GetPoolAddr(),
		GasLimit:     800000,
	}
}
