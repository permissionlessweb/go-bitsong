package v024

import (
	store "cosmossdk.io/store/types"
	"github.com/bitsongofficial/go-bitsong/app/upgrades"
)

var Upgrade = upgrades.Upgrade{
	UpgradeName:          UpgradeName,
	CreateUpgradeHandler: CreateV024UpgradeHandler,
	StoreUpgrades:        store.StoreUpgrades{Added: []string{}, Deleted: []string{}},
}

const (
	UpgradeName   = "v024"
	ContractAdmin = "bitsong1mxascwuvua9xemxe9k9qxgaexpdnzm098c06np"
)

const (
	OldBs721CodeId      = 1
	OldFactoryCodeId    = 2
	OldRoyaltiesCodeId  = 3
	OldBs721CurveCodeId = 4
	OldLaunchpadCodeId  = 5
)
const (
	NewBs721CodeId      = 92
	NewFactoryCodeId    = 93
	NewRoyaltiesCodeId  = 94
	NewBs721CurveCodeId = 95
	NewLaunchpadCodeId  = 96
)

const (
	CreateNftFee   = 500000000
	ProtocolFeeBps = 30
)

type RawContractInfo struct {
	Contract string `json:"contract"`
	Version  string `json:"version"`
}

type ContractOwnership struct {
	Owner         string `json:"owner"`
	PendingOwner  string `json:"pending_owner"`
	PendingExpiry string `json:"pending_expiry"`
}

type Bs721CurveInitMsg struct {
	Symbol         string `json:"symbol"`
	Name           string `json:"name"`
	URI            string `json:"uri"`
	PaymentDenom   string `json:"payment_denom"`
	MaxPerAddress  *int64 `json:"max_per_address"`
	PaymentAddress string `json:"payment_address"`
	SellerFeeBps   int    `json:"seller_fee_bps"`
	ReferralFeeBps int    `json:"referral_fee_bps"`
	ProtocolFeeBps int    `json:"protocol_fee_bps"`
	StartTime      string `json:"start_time"`
	MaxEdition     *int64 `json:"max_edition,omitempty"`
	Bs721CodeID    int    `json:"bs721_code_id"`
	Ratio          int    `json:"ratio"`
	Bs721Admin     string `json:"bs721_admin"`
}

type UpdateConfigObj struct {
	PaymentAddress string `json:"payment_address"`
	SellerFeeBps   uint32 `json:"seller_fee_bps"`
	ReferralFeeBps uint32 `json:"referral_fee_bps"`
	ProtocolFeeBps uint32 `json:"protocol_fee_bps"`
	PaymentDenom   string `json:"payment_denom"`
}

// QUERIES //

type QueryTotalMintPrice struct {
	Amount int `json:"amount"`
}

type QueryTotalMintPriceResponse struct {
	PubBasePrice   uint `json:"pub_base_price"`
	PubRoyalties   uint `json:"pub_royalties"`
	PubReferral    uint `json:"pub_referral"`
	PubProtocolFee uint `json:"pub_protocol_fee"`
	PubTotalPrice  uint `json:"pub_total_price"`
}

// EXECUTE MSGS //

type Mint struct {
	Amount   uint32 `json:"amount"`
	Referral string `json:"referral,omitempty"`
	MintTo   string `json:"mint_to,omitempty"`
}

type UpdateConfig struct {
	Cfg UpdateConfigObj `json:"cfg"`
}
