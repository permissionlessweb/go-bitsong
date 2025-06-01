// @generated
#[allow(clippy::derive_partial_eq_without_eq)]
#[derive(Clone, PartialEq, ::prost::Message)]
pub struct EventIssue {
    #[prost(string, tag = "1")]
    pub denom: ::prost::alloc::string::String,
}
impl ::prost::Name for EventIssue {
    const NAME: &'static str = "EventIssue";
    const PACKAGE: &'static str = "bitsong.fantoken.v1beta1";
    fn full_name() -> ::prost::alloc::string::String {
        ::prost::alloc::format!("bitsong.fantoken.v1beta1.{}", Self::NAME)
    }
}
#[allow(clippy::derive_partial_eq_without_eq)]
#[derive(Clone, PartialEq, ::prost::Message)]
pub struct EventDisableMint {
    #[prost(string, tag = "1")]
    pub denom: ::prost::alloc::string::String,
}
impl ::prost::Name for EventDisableMint {
    const NAME: &'static str = "EventDisableMint";
    const PACKAGE: &'static str = "bitsong.fantoken.v1beta1";
    fn full_name() -> ::prost::alloc::string::String {
        ::prost::alloc::format!("bitsong.fantoken.v1beta1.{}", Self::NAME)
    }
}
#[allow(clippy::derive_partial_eq_without_eq)]
#[derive(Clone, PartialEq, ::prost::Message)]
pub struct EventMint {
    #[prost(string, tag = "1")]
    pub recipient: ::prost::alloc::string::String,
    #[prost(string, tag = "2")]
    pub coin: ::prost::alloc::string::String,
}
impl ::prost::Name for EventMint {
    const NAME: &'static str = "EventMint";
    const PACKAGE: &'static str = "bitsong.fantoken.v1beta1";
    fn full_name() -> ::prost::alloc::string::String {
        ::prost::alloc::format!("bitsong.fantoken.v1beta1.{}", Self::NAME)
    }
}
#[allow(clippy::derive_partial_eq_without_eq)]
#[derive(Clone, PartialEq, ::prost::Message)]
pub struct EventBurn {
    #[prost(string, tag = "1")]
    pub sender: ::prost::alloc::string::String,
    #[prost(string, tag = "2")]
    pub coin: ::prost::alloc::string::String,
}
impl ::prost::Name for EventBurn {
    const NAME: &'static str = "EventBurn";
    const PACKAGE: &'static str = "bitsong.fantoken.v1beta1";
    fn full_name() -> ::prost::alloc::string::String {
        ::prost::alloc::format!("bitsong.fantoken.v1beta1.{}", Self::NAME)
    }
}
#[allow(clippy::derive_partial_eq_without_eq)]
#[derive(Clone, PartialEq, ::prost::Message)]
pub struct EventSetAuthority {
    #[prost(string, tag = "1")]
    pub denom: ::prost::alloc::string::String,
    #[prost(string, tag = "2")]
    pub old_authority: ::prost::alloc::string::String,
    #[prost(string, tag = "3")]
    pub new_authority: ::prost::alloc::string::String,
}
impl ::prost::Name for EventSetAuthority {
    const NAME: &'static str = "EventSetAuthority";
    const PACKAGE: &'static str = "bitsong.fantoken.v1beta1";
    fn full_name() -> ::prost::alloc::string::String {
        ::prost::alloc::format!("bitsong.fantoken.v1beta1.{}", Self::NAME)
    }
}
#[allow(clippy::derive_partial_eq_without_eq)]
#[derive(Clone, PartialEq, ::prost::Message)]
pub struct EventSetMinter {
    #[prost(string, tag = "1")]
    pub denom: ::prost::alloc::string::String,
    #[prost(string, tag = "2")]
    pub old_minter: ::prost::alloc::string::String,
    #[prost(string, tag = "3")]
    pub new_minter: ::prost::alloc::string::String,
}
impl ::prost::Name for EventSetMinter {
    const NAME: &'static str = "EventSetMinter";
    const PACKAGE: &'static str = "bitsong.fantoken.v1beta1";
    fn full_name() -> ::prost::alloc::string::String {
        ::prost::alloc::format!("bitsong.fantoken.v1beta1.{}", Self::NAME)
    }
}
#[allow(clippy::derive_partial_eq_without_eq)]
#[derive(Clone, PartialEq, ::prost::Message)]
pub struct EventSetUri {
    #[prost(string, tag = "1")]
    pub denom: ::prost::alloc::string::String,
}
impl ::prost::Name for EventSetUri {
    const NAME: &'static str = "EventSetUri";
    const PACKAGE: &'static str = "bitsong.fantoken.v1beta1";
    fn full_name() -> ::prost::alloc::string::String {
        ::prost::alloc::format!("bitsong.fantoken.v1beta1.{}", Self::NAME)
    }
}
#[allow(clippy::derive_partial_eq_without_eq)]
#[derive(Clone, PartialEq, ::prost::Message)]
pub struct Metadata {
    /// name defines the name of the fantoken (eg: Kitty Punk)
    #[prost(string, tag = "1")]
    pub name: ::prost::alloc::string::String,
    /// symbol is the token symbol usually shown on exchanges (eg: KITTY)
    #[prost(string, tag = "2")]
    pub symbol: ::prost::alloc::string::String,
    /// URI to a document (on or off-chain) that contains additional
    /// information.Optional.
    #[prost(string, tag = "3")]
    pub uri: ::prost::alloc::string::String,
    /// sdk.AccAddress allowed to set a new uri
    #[prost(string, tag = "4")]
    pub authority: ::prost::alloc::string::String,
}
impl ::prost::Name for Metadata {
    const NAME: &'static str = "Metadata";
    const PACKAGE: &'static str = "bitsong.fantoken.v1beta1";
    fn full_name() -> ::prost::alloc::string::String {
        ::prost::alloc::format!("bitsong.fantoken.v1beta1.{}", Self::NAME)
    }
}
/// FanToken defines a standard for the fungible token
#[allow(clippy::derive_partial_eq_without_eq)]
#[derive(Clone, PartialEq, ::prost::Message)]
pub struct FanToken {
    /// denom represents the string name of the given denom unit (e.g ft<hash>).
    #[prost(string, tag = "1")]
    pub denom: ::prost::alloc::string::String,
    #[prost(string, tag = "2")]
    pub max_supply: ::prost::alloc::string::String,
    /// sdk.AccAddress allowed to mint new fantoken
    #[prost(string, tag = "3")]
    pub minter: ::prost::alloc::string::String,
    #[prost(message, optional, tag = "4")]
    pub meta_data: ::core::option::Option<Metadata>,
}
impl ::prost::Name for FanToken {
    const NAME: &'static str = "FanToken";
    const PACKAGE: &'static str = "bitsong.fantoken.v1beta1";
    fn full_name() -> ::prost::alloc::string::String {
        ::prost::alloc::format!("bitsong.fantoken.v1beta1.{}", Self::NAME)
    }
}
/// Params defines fantoken module's parameters
#[allow(clippy::derive_partial_eq_without_eq)]
#[derive(Clone, PartialEq, ::prost::Message)]
pub struct Params {
    #[prost(message, optional, tag = "1")]
    pub issue_fee: ::core::option::Option<super::super::super::cosmos::base::v1beta1::Coin>,
    #[prost(message, optional, tag = "2")]
    pub mint_fee: ::core::option::Option<super::super::super::cosmos::base::v1beta1::Coin>,
    #[prost(message, optional, tag = "3")]
    pub burn_fee: ::core::option::Option<super::super::super::cosmos::base::v1beta1::Coin>,
}
impl ::prost::Name for Params {
    const NAME: &'static str = "Params";
    const PACKAGE: &'static str = "bitsong.fantoken.v1beta1";
    fn full_name() -> ::prost::alloc::string::String {
        ::prost::alloc::format!("bitsong.fantoken.v1beta1.{}", Self::NAME)
    }
}
/// GenesisState defines the fantoken module's genesis state
#[allow(clippy::derive_partial_eq_without_eq)]
#[derive(Clone, PartialEq, ::prost::Message)]
pub struct GenesisState {
    #[prost(message, optional, tag = "1")]
    pub params: ::core::option::Option<Params>,
    #[prost(message, repeated, tag = "2")]
    pub fan_tokens: ::prost::alloc::vec::Vec<FanToken>,
}
impl ::prost::Name for GenesisState {
    const NAME: &'static str = "GenesisState";
    const PACKAGE: &'static str = "bitsong.fantoken.v1beta1";
    fn full_name() -> ::prost::alloc::string::String {
        ::prost::alloc::format!("bitsong.fantoken.v1beta1.{}", Self::NAME)
    }
}
#[allow(clippy::derive_partial_eq_without_eq)]
#[derive(Clone, PartialEq, ::prost::Message)]
pub struct UpdateFeesProposal {
    #[prost(string, tag = "1")]
    pub title: ::prost::alloc::string::String,
    #[prost(string, tag = "2")]
    pub description: ::prost::alloc::string::String,
    #[prost(message, optional, tag = "3")]
    pub issue_fee: ::core::option::Option<super::super::super::cosmos::base::v1beta1::Coin>,
    #[prost(message, optional, tag = "4")]
    pub mint_fee: ::core::option::Option<super::super::super::cosmos::base::v1beta1::Coin>,
    #[prost(message, optional, tag = "5")]
    pub burn_fee: ::core::option::Option<super::super::super::cosmos::base::v1beta1::Coin>,
}
impl ::prost::Name for UpdateFeesProposal {
    const NAME: &'static str = "UpdateFeesProposal";
    const PACKAGE: &'static str = "bitsong.fantoken.v1beta1";
    fn full_name() -> ::prost::alloc::string::String {
        ::prost::alloc::format!("bitsong.fantoken.v1beta1.{}", Self::NAME)
    }
}
#[allow(clippy::derive_partial_eq_without_eq)]
#[derive(Clone, PartialEq, ::prost::Message)]
pub struct UpdateFeesProposalWithDeposit {
    #[prost(string, tag = "1")]
    pub title: ::prost::alloc::string::String,
    #[prost(string, tag = "2")]
    pub description: ::prost::alloc::string::String,
    #[prost(string, tag = "3")]
    pub issue_fee: ::prost::alloc::string::String,
    #[prost(string, tag = "4")]
    pub mint_fee: ::prost::alloc::string::String,
    #[prost(string, tag = "5")]
    pub burn_fee: ::prost::alloc::string::String,
    #[prost(string, tag = "7")]
    pub deposit: ::prost::alloc::string::String,
}
impl ::prost::Name for UpdateFeesProposalWithDeposit {
    const NAME: &'static str = "UpdateFeesProposalWithDeposit";
    const PACKAGE: &'static str = "bitsong.fantoken.v1beta1";
    fn full_name() -> ::prost::alloc::string::String {
        ::prost::alloc::format!("bitsong.fantoken.v1beta1.{}", Self::NAME)
    }
}
/// QueryFanTokenRequest is request type for the Query/FanToken RPC method
#[allow(clippy::derive_partial_eq_without_eq)]
#[derive(Clone, PartialEq, ::prost::Message)]
pub struct QueryFanTokenRequest {
    #[prost(string, tag = "1")]
    pub denom: ::prost::alloc::string::String,
}
impl ::prost::Name for QueryFanTokenRequest {
    const NAME: &'static str = "QueryFanTokenRequest";
    const PACKAGE: &'static str = "bitsong.fantoken.v1beta1";
    fn full_name() -> ::prost::alloc::string::String {
        ::prost::alloc::format!("bitsong.fantoken.v1beta1.{}", Self::NAME)
    }
}
/// QueryFanTokenResponse is response type for the Query/FanToken RPC method
#[allow(clippy::derive_partial_eq_without_eq)]
#[derive(Clone, PartialEq, ::prost::Message)]
pub struct QueryFanTokenResponse {
    #[prost(message, optional, tag = "1")]
    pub fantoken: ::core::option::Option<FanToken>,
}
impl ::prost::Name for QueryFanTokenResponse {
    const NAME: &'static str = "QueryFanTokenResponse";
    const PACKAGE: &'static str = "bitsong.fantoken.v1beta1";
    fn full_name() -> ::prost::alloc::string::String {
        ::prost::alloc::format!("bitsong.fantoken.v1beta1.{}", Self::NAME)
    }
}
/// QueryFanTokensRequest is request type for the Query/FanTokens RPC method
#[allow(clippy::derive_partial_eq_without_eq)]
#[derive(Clone, PartialEq, ::prost::Message)]
pub struct QueryFanTokensRequest {
    #[prost(string, tag = "1")]
    pub authority: ::prost::alloc::string::String,
    /// pagination defines an optional pagination for the request.
    #[prost(message, optional, tag = "2")]
    pub pagination:
        ::core::option::Option<super::super::super::cosmos::base::query::v1beta1::PageRequest>,
}
impl ::prost::Name for QueryFanTokensRequest {
    const NAME: &'static str = "QueryFanTokensRequest";
    const PACKAGE: &'static str = "bitsong.fantoken.v1beta1";
    fn full_name() -> ::prost::alloc::string::String {
        ::prost::alloc::format!("bitsong.fantoken.v1beta1.{}", Self::NAME)
    }
}
/// QueryFanTokensResponse is response type for the Query/FanTokens RPC method
#[allow(clippy::derive_partial_eq_without_eq)]
#[derive(Clone, PartialEq, ::prost::Message)]
pub struct QueryFanTokensResponse {
    #[prost(message, repeated, tag = "1")]
    pub fantokens: ::prost::alloc::vec::Vec<FanToken>,
    #[prost(message, optional, tag = "2")]
    pub pagination:
        ::core::option::Option<super::super::super::cosmos::base::query::v1beta1::PageResponse>,
}
impl ::prost::Name for QueryFanTokensResponse {
    const NAME: &'static str = "QueryFanTokensResponse";
    const PACKAGE: &'static str = "bitsong.fantoken.v1beta1";
    fn full_name() -> ::prost::alloc::string::String {
        ::prost::alloc::format!("bitsong.fantoken.v1beta1.{}", Self::NAME)
    }
}
/// QueryParametersRequest is request type for the Query/Parameters RPC method
#[allow(clippy::derive_partial_eq_without_eq)]
#[derive(Clone, PartialEq, ::prost::Message)]
pub struct QueryParamsRequest {}
impl ::prost::Name for QueryParamsRequest {
    const NAME: &'static str = "QueryParamsRequest";
    const PACKAGE: &'static str = "bitsong.fantoken.v1beta1";
    fn full_name() -> ::prost::alloc::string::String {
        ::prost::alloc::format!("bitsong.fantoken.v1beta1.{}", Self::NAME)
    }
}
/// QueryParametersResponse is response type for the Query/Parameters RPC method
#[allow(clippy::derive_partial_eq_without_eq)]
#[derive(Clone, PartialEq, ::prost::Message)]
pub struct QueryParamsResponse {
    #[prost(message, optional, tag = "1")]
    pub params: ::core::option::Option<Params>,
}
impl ::prost::Name for QueryParamsResponse {
    const NAME: &'static str = "QueryParamsResponse";
    const PACKAGE: &'static str = "bitsong.fantoken.v1beta1";
    fn full_name() -> ::prost::alloc::string::String {
        ::prost::alloc::format!("bitsong.fantoken.v1beta1.{}", Self::NAME)
    }
}
/// MsgIssue defines a message for issuing a new fan token
#[allow(clippy::derive_partial_eq_without_eq)]
#[derive(Clone, PartialEq, ::prost::Message)]
pub struct MsgIssue {
    /// symbol which corresponds to the symbol of the fan token. It is a string and
    /// cannot change for the whole life of the fan token
    #[prost(string, tag = "1")]
    pub symbol: ::prost::alloc::string::String,
    /// name which corresponds to the name of the fan token. It is a string and
    /// cannot change for the whole life of the fan token
    #[prost(string, tag = "2")]
    pub name: ::prost::alloc::string::String,
    /// max_supply that represents the maximum number of possible mintable tokens.
    /// It is an integer number, expressed in micro unit 10^6
    #[prost(string, tag = "3")]
    pub max_supply: ::prost::alloc::string::String,
    /// authority which is who can set a new uri metadata
    #[prost(string, tag = "4")]
    pub authority: ::prost::alloc::string::String,
    /// minter who is who can mint new fantoken and disable the minter process, the
    /// minter key also pay the gas fee
    #[prost(string, tag = "5")]
    pub minter: ::prost::alloc::string::String,
    /// URI which is the current uri of the fan token. It is a string can change
    /// during the fan token lifecycle thanks to the MsgEdit
    #[prost(string, tag = "6")]
    pub uri: ::prost::alloc::string::String,
}
impl ::prost::Name for MsgIssue {
    const NAME: &'static str = "MsgIssue";
    const PACKAGE: &'static str = "bitsong.fantoken.v1beta1";
    fn full_name() -> ::prost::alloc::string::String {
        ::prost::alloc::format!("bitsong.fantoken.v1beta1.{}", Self::NAME)
    }
}
/// MsgIssueResponse defines the MsgIssue response type
#[allow(clippy::derive_partial_eq_without_eq)]
#[derive(Clone, PartialEq, ::prost::Message)]
pub struct MsgIssueResponse {
    #[prost(string, tag = "1")]
    pub denom: ::prost::alloc::string::String,
}
impl ::prost::Name for MsgIssueResponse {
    const NAME: &'static str = "MsgIssueResponse";
    const PACKAGE: &'static str = "bitsong.fantoken.v1beta1";
    fn full_name() -> ::prost::alloc::string::String {
        ::prost::alloc::format!("bitsong.fantoken.v1beta1.{}", Self::NAME)
    }
}
/// MsgDisableMint defines a message for disable the mint function
#[allow(clippy::derive_partial_eq_without_eq)]
#[derive(Clone, PartialEq, ::prost::Message)]
pub struct MsgDisableMint {
    #[prost(string, tag = "1")]
    pub denom: ::prost::alloc::string::String,
    #[prost(string, tag = "2")]
    pub minter: ::prost::alloc::string::String,
}
impl ::prost::Name for MsgDisableMint {
    const NAME: &'static str = "MsgDisableMint";
    const PACKAGE: &'static str = "bitsong.fantoken.v1beta1";
    fn full_name() -> ::prost::alloc::string::String {
        ::prost::alloc::format!("bitsong.fantoken.v1beta1.{}", Self::NAME)
    }
}
/// MsgDisableMintResponse defines the MsgDisableMint response type
#[allow(clippy::derive_partial_eq_without_eq)]
#[derive(Clone, PartialEq, ::prost::Message)]
pub struct MsgDisableMintResponse {
    #[prost(string, tag = "1")]
    pub denom: ::prost::alloc::string::String,
}
impl ::prost::Name for MsgDisableMintResponse {
    const NAME: &'static str = "MsgDisableMintResponse";
    const PACKAGE: &'static str = "bitsong.fantoken.v1beta1";
    fn full_name() -> ::prost::alloc::string::String {
        ::prost::alloc::format!("bitsong.fantoken.v1beta1.{}", Self::NAME)
    }
}
/// MsgMint defines a message for minting a new fan token
#[allow(clippy::derive_partial_eq_without_eq)]
#[derive(Clone, PartialEq, ::prost::Message)]
pub struct MsgMint {
    #[prost(string, tag = "1")]
    pub recipient: ::prost::alloc::string::String,
    /// coin mean the amount + denom, eg: 10000ftFADJID34MCDM
    #[prost(message, optional, tag = "2")]
    pub coin: ::core::option::Option<super::super::super::cosmos::base::v1beta1::Coin>,
    #[prost(string, tag = "3")]
    pub minter: ::prost::alloc::string::String,
}
impl ::prost::Name for MsgMint {
    const NAME: &'static str = "MsgMint";
    const PACKAGE: &'static str = "bitsong.fantoken.v1beta1";
    fn full_name() -> ::prost::alloc::string::String {
        ::prost::alloc::format!("bitsong.fantoken.v1beta1.{}", Self::NAME)
    }
}
/// MsgMintResponse defines the MsgMint response type
#[allow(clippy::derive_partial_eq_without_eq)]
#[derive(Clone, PartialEq, ::prost::Message)]
pub struct MsgMintResponse {
    #[prost(string, tag = "1")]
    pub recipient: ::prost::alloc::string::String,
    #[prost(message, optional, tag = "2")]
    pub coin: ::core::option::Option<super::super::super::cosmos::base::v1beta1::Coin>,
}
impl ::prost::Name for MsgMintResponse {
    const NAME: &'static str = "MsgMintResponse";
    const PACKAGE: &'static str = "bitsong.fantoken.v1beta1";
    fn full_name() -> ::prost::alloc::string::String {
        ::prost::alloc::format!("bitsong.fantoken.v1beta1.{}", Self::NAME)
    }
}
/// MsgBurn defines a message for burning some fan tokens
#[allow(clippy::derive_partial_eq_without_eq)]
#[derive(Clone, PartialEq, ::prost::Message)]
pub struct MsgBurn {
    /// coin mean the amount + denom, eg: 10000ftFADJID34MCDM
    #[prost(message, optional, tag = "1")]
    pub coin: ::core::option::Option<super::super::super::cosmos::base::v1beta1::Coin>,
    #[prost(string, tag = "2")]
    pub sender: ::prost::alloc::string::String,
}
impl ::prost::Name for MsgBurn {
    const NAME: &'static str = "MsgBurn";
    const PACKAGE: &'static str = "bitsong.fantoken.v1beta1";
    fn full_name() -> ::prost::alloc::string::String {
        ::prost::alloc::format!("bitsong.fantoken.v1beta1.{}", Self::NAME)
    }
}
/// MsgBurnResponse defines the MsgBurn response type
#[allow(clippy::derive_partial_eq_without_eq)]
#[derive(Clone, PartialEq, ::prost::Message)]
pub struct MsgBurnResponse {
    #[prost(string, tag = "1")]
    pub sender: ::prost::alloc::string::String,
    #[prost(message, optional, tag = "2")]
    pub coin: ::core::option::Option<super::super::super::cosmos::base::v1beta1::Coin>,
}
impl ::prost::Name for MsgBurnResponse {
    const NAME: &'static str = "MsgBurnResponse";
    const PACKAGE: &'static str = "bitsong.fantoken.v1beta1";
    fn full_name() -> ::prost::alloc::string::String {
        ::prost::alloc::format!("bitsong.fantoken.v1beta1.{}", Self::NAME)
    }
}
/// MsgSetMinter defines a message for changing the fan token minter address
#[allow(clippy::derive_partial_eq_without_eq)]
#[derive(Clone, PartialEq, ::prost::Message)]
pub struct MsgSetMinter {
    /// denom the fan token denom
    #[prost(string, tag = "1")]
    pub denom: ::prost::alloc::string::String,
    /// old_minter, the actual minter
    #[prost(string, tag = "2")]
    pub old_minter: ::prost::alloc::string::String,
    /// new_minter, the new fan token minter
    #[prost(string, tag = "3")]
    pub new_minter: ::prost::alloc::string::String,
}
impl ::prost::Name for MsgSetMinter {
    const NAME: &'static str = "MsgSetMinter";
    const PACKAGE: &'static str = "bitsong.fantoken.v1beta1";
    fn full_name() -> ::prost::alloc::string::String {
        ::prost::alloc::format!("bitsong.fantoken.v1beta1.{}", Self::NAME)
    }
}
/// MsgSetMinterResponse defines the MsgTransferAuthority response type
#[allow(clippy::derive_partial_eq_without_eq)]
#[derive(Clone, PartialEq, ::prost::Message)]
pub struct MsgSetMinterResponse {
    /// denom the fan token denom
    #[prost(string, tag = "1")]
    pub denom: ::prost::alloc::string::String,
    /// old_minter, the actual minter
    #[prost(string, tag = "2")]
    pub old_minter: ::prost::alloc::string::String,
    /// new_minter, the new fan token minter
    #[prost(string, tag = "3")]
    pub new_minter: ::prost::alloc::string::String,
}
impl ::prost::Name for MsgSetMinterResponse {
    const NAME: &'static str = "MsgSetMinterResponse";
    const PACKAGE: &'static str = "bitsong.fantoken.v1beta1";
    fn full_name() -> ::prost::alloc::string::String {
        ::prost::alloc::format!("bitsong.fantoken.v1beta1.{}", Self::NAME)
    }
}
/// MsgSetAuthority defines a message for changing the fan token minter address
#[allow(clippy::derive_partial_eq_without_eq)]
#[derive(Clone, PartialEq, ::prost::Message)]
pub struct MsgSetAuthority {
    /// denom the fan token denom
    #[prost(string, tag = "1")]
    pub denom: ::prost::alloc::string::String,
    /// old_authority, the actual metadata authority
    #[prost(string, tag = "2")]
    pub old_authority: ::prost::alloc::string::String,
    /// new_authority, the new fan token metadata authority
    #[prost(string, tag = "3")]
    pub new_authority: ::prost::alloc::string::String,
}
impl ::prost::Name for MsgSetAuthority {
    const NAME: &'static str = "MsgSetAuthority";
    const PACKAGE: &'static str = "bitsong.fantoken.v1beta1";
    fn full_name() -> ::prost::alloc::string::String {
        ::prost::alloc::format!("bitsong.fantoken.v1beta1.{}", Self::NAME)
    }
}
/// MsgSetAuthorityResponse defines the MsgTransferAuthority response type
#[allow(clippy::derive_partial_eq_without_eq)]
#[derive(Clone, PartialEq, ::prost::Message)]
pub struct MsgSetAuthorityResponse {
    /// denom the fan token denom
    #[prost(string, tag = "1")]
    pub denom: ::prost::alloc::string::String,
    /// old_authority, the actual metadata authority
    #[prost(string, tag = "2")]
    pub old_authority: ::prost::alloc::string::String,
    /// new_authority, the new fan token metadata authority
    #[prost(string, tag = "3")]
    pub new_authority: ::prost::alloc::string::String,
}
impl ::prost::Name for MsgSetAuthorityResponse {
    const NAME: &'static str = "MsgSetAuthorityResponse";
    const PACKAGE: &'static str = "bitsong.fantoken.v1beta1";
    fn full_name() -> ::prost::alloc::string::String {
        ::prost::alloc::format!("bitsong.fantoken.v1beta1.{}", Self::NAME)
    }
}
#[allow(clippy::derive_partial_eq_without_eq)]
#[derive(Clone, PartialEq, ::prost::Message)]
pub struct MsgSetUri {
    #[prost(string, tag = "1")]
    pub authority: ::prost::alloc::string::String,
    #[prost(string, tag = "2")]
    pub denom: ::prost::alloc::string::String,
    #[prost(string, tag = "3")]
    pub uri: ::prost::alloc::string::String,
}
impl ::prost::Name for MsgSetUri {
    const NAME: &'static str = "MsgSetUri";
    const PACKAGE: &'static str = "bitsong.fantoken.v1beta1";
    fn full_name() -> ::prost::alloc::string::String {
        ::prost::alloc::format!("bitsong.fantoken.v1beta1.{}", Self::NAME)
    }
}
#[allow(clippy::derive_partial_eq_without_eq)]
#[derive(Clone, PartialEq, ::prost::Message)]
pub struct MsgSetUriResponse {
    #[prost(string, tag = "1")]
    pub denom: ::prost::alloc::string::String,
}
impl ::prost::Name for MsgSetUriResponse {
    const NAME: &'static str = "MsgSetUriResponse";
    const PACKAGE: &'static str = "bitsong.fantoken.v1beta1";
    fn full_name() -> ::prost::alloc::string::String {
        ::prost::alloc::format!("bitsong.fantoken.v1beta1.{}", Self::NAME)
    }
}
include!("bitsong.fantoken.v1beta1.serde.rs");
include!("bitsong.fantoken.v1beta1.tonic.rs");
// @@protoc_insertion_point(module)
