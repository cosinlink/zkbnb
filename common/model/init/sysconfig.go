package init

import (
	"github.com/zecrey-labs/zecrey-legend/common/model/sysconfig"
	"github.com/zecrey-labs/zecrey-legend/common/sysconfigName"
)

const (
	// network rpc
	BSC_Test_Network_RPC   = "http://tf-dex-preview-validator-nlb-6fd109ac8b9d390a.elb.ap-northeast-1.amazonaws.com:8545"
	Local_Test_Network_RPC = "http://127.0.0.1:8545/"

	ZecreyLegendContractAddr = "0xfdd3253dB0160C923872Ce2a043b8203b972fe05"
	GovernanceContractAddr   = "0xFc17b0D06603bf15357a8DBa9ee0F8fb97b435Ef"
	VerifierContractAddr     = "0xb020FB79123e3493aA4327958cb64529038484AE"
)

func initSysConfig() []*sysconfig.Sysconfig {
	return []*sysconfig.Sysconfig{
		{
			Name:      sysconfigName.SysGasFee,
			Value:     "1",
			ValueType: "float",
			Comment:   "based on ETH",
		},
		{
			Name:      sysconfigName.MaxAssetId,
			Value:     "9",
			ValueType: "int",
			Comment:   "max number of asset id",
		},
		{
			Name:      sysconfigName.TreasuryAccountIndex,
			Value:     "0",
			ValueType: "int",
			Comment:   "treasury index",
		},
		{
			Name:      sysconfigName.GasAccountIndex,
			Value:     "1",
			ValueType: "int",
			Comment:   "gas index",
		},
		{
			Name:      sysconfigName.ZecreyLegendContract,
			Value:     ZecreyLegendContractAddr,
			ValueType: "string",
			Comment:   "Zecrey contract on BSC",
		},
		// Governance Contract
		{
			Name:      sysconfigName.GovernanceContract,
			Value:     GovernanceContractAddr,
			ValueType: "string",
			Comment:   "Governance contract on BSC",
		},

		// Asset_Governance Contract
		//{
		//	Name:      sysconfigName.AssetGovernanceContract,
		//	Value:     AssetGovernanceContractAddr,
		//	ValueType: "string",
		//	Comment:   "Asset_Governance contract on BSC",
		//},

		// Verifier Contract
		//{
		//	Name:      sysconfigName.VerifierContract,
		//	Value:     VerifierContractAddr,
		//	ValueType: "string",
		//	Comment:   "Verifier contract on BSC",
		//},
		// network rpc
		{
			Name:      sysconfigName.BscTestNetworkRpc,
			Value:     BSC_Test_Network_RPC,
			ValueType: "string",
			Comment:   "BSC network rpc",
		},
		// TODO
		{
			Name:      "Local_Test_Network_RPC",
			Value:     Local_Test_Network_RPC,
			ValueType: "string",
			Comment:   "Local network rpc",
		},
	}
}
