package blockchain

import (
	"fmt"

	"gin-campus-market/contract"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

var (
	EthClient *ethclient.Client
	Instance  *contract.CampusMarketplace
)

func InitBlockchain() {
	url := "https://eth-sepolia.g.alchemy.com/v2/BDzbJa07uo_GrppseQsPi"

	var err error
	EthClient, err = ethclient.Dial(url)
	if err != nil {
		panic(fmt.Sprintf("区块链连接失败: %v", err))
	}

	address := common.HexToAddress("0x3FfAf5E999Fda995b7959249B2F2eFf494427457")
	Instance, err = contract.NewCampusMarketplace(address, EthClient)
	if err != nil {
		panic(fmt.Sprintf("实例化合约失败: %v", err))
	}

	fmt.Println("✅ 区块链服务连接成功！已绑定合约地址:", address.Hex())
}
