// @generated
/// This object is used to store the contract address and the
/// jail status of the contract.
#[allow(clippy::derive_partial_eq_without_eq)]
#[derive(Clone, PartialEq, ::prost::Message)]
pub struct CadenceContract {
    /// The address of the contract.
    #[prost(string, tag = "1")]
    pub contract_address: ::prost::alloc::string::String,
    /// The jail status of the contract.
    #[prost(bool, tag = "2")]
    pub is_jailed: bool,
}
impl ::prost::Name for CadenceContract {
    const NAME: &'static str = "CadenceContract";
    const PACKAGE: &'static str = "bitsong.cadence.v1";
    fn full_name() -> ::prost::alloc::string::String {
        ::prost::alloc::format!("bitsong.cadence.v1.{}", Self::NAME)
    }
}
/// GenesisState - initial state of module
#[allow(clippy::derive_partial_eq_without_eq)]
#[derive(Clone, PartialEq, ::prost::Message)]
pub struct GenesisState {
    /// Params of this module
    #[prost(message, optional, tag = "1")]
    pub params: ::core::option::Option<Params>,
}
impl ::prost::Name for GenesisState {
    const NAME: &'static str = "GenesisState";
    const PACKAGE: &'static str = "bitsong.cadence.v1";
    fn full_name() -> ::prost::alloc::string::String {
        ::prost::alloc::format!("bitsong.cadence.v1.{}", Self::NAME)
    }
}
/// Params defines the set of module parameters.
#[allow(clippy::derive_partial_eq_without_eq)]
#[derive(Clone, PartialEq, ::prost::Message)]
pub struct Params {
    /// contract_gas_limit defines the maximum amount of gas that can be used by a contract.
    #[prost(uint64, tag = "1")]
    pub contract_gas_limit: u64,
}
impl ::prost::Name for Params {
    const NAME: &'static str = "Params";
    const PACKAGE: &'static str = "bitsong.cadence.v1";
    fn full_name() -> ::prost::alloc::string::String {
        ::prost::alloc::format!("bitsong.cadence.v1.{}", Self::NAME)
    }
}
/// QueryCadenceContracts is the request type to get all contracts.
#[allow(clippy::derive_partial_eq_without_eq)]
#[derive(Clone, PartialEq, ::prost::Message)]
pub struct QueryCadenceContracts {
    /// pagination defines an optional pagination for the request.
    #[prost(message, optional, tag = "1")]
    pub pagination:
        ::core::option::Option<super::super::super::cosmos::base::query::v1beta1::PageRequest>,
}
impl ::prost::Name for QueryCadenceContracts {
    const NAME: &'static str = "QueryCadenceContracts";
    const PACKAGE: &'static str = "bitsong.cadence.v1";
    fn full_name() -> ::prost::alloc::string::String {
        ::prost::alloc::format!("bitsong.cadence.v1.{}", Self::NAME)
    }
}
/// QueryCadenceContractsResponse is the response type for the Query/CadenceContracts RPC method.
#[allow(clippy::derive_partial_eq_without_eq)]
#[derive(Clone, PartialEq, ::prost::Message)]
pub struct QueryCadenceContractsResponse {
    /// cadence_contracts are the cadence contract s.
    #[prost(message, repeated, tag = "1")]
    pub cadence_contracts: ::prost::alloc::vec::Vec<CadenceContract>,
    /// pagination defines the pagination in the response.
    #[prost(message, optional, tag = "2")]
    pub pagination:
        ::core::option::Option<super::super::super::cosmos::base::query::v1beta1::PageResponse>,
}
impl ::prost::Name for QueryCadenceContractsResponse {
    const NAME: &'static str = "QueryCadenceContractsResponse";
    const PACKAGE: &'static str = "bitsong.cadence.v1";
    fn full_name() -> ::prost::alloc::string::String {
        ::prost::alloc::format!("bitsong.cadence.v1.{}", Self::NAME)
    }
}
/// QueryCadenceContract is the request type to get a single contract.
#[allow(clippy::derive_partial_eq_without_eq)]
#[derive(Clone, PartialEq, ::prost::Message)]
pub struct QueryCadenceContract {
    /// contract_address is the address of the contract to query.
    #[prost(string, tag = "1")]
    pub contract_address: ::prost::alloc::string::String,
}
impl ::prost::Name for QueryCadenceContract {
    const NAME: &'static str = "QueryCadenceContract";
    const PACKAGE: &'static str = "bitsong.cadence.v1";
    fn full_name() -> ::prost::alloc::string::String {
        ::prost::alloc::format!("bitsong.cadence.v1.{}", Self::NAME)
    }
}
/// QueryCadenceContractResponse is the response type for the Query/CadenceContract RPC method.
#[allow(clippy::derive_partial_eq_without_eq)]
#[derive(Clone, PartialEq, ::prost::Message)]
pub struct QueryCadenceContractResponse {
    /// contract is the cadence contract .
    #[prost(message, optional, tag = "1")]
    pub cadence_contract: ::core::option::Option<CadenceContract>,
}
impl ::prost::Name for QueryCadenceContractResponse {
    const NAME: &'static str = "QueryCadenceContractResponse";
    const PACKAGE: &'static str = "bitsong.cadence.v1";
    fn full_name() -> ::prost::alloc::string::String {
        ::prost::alloc::format!("bitsong.cadence.v1.{}", Self::NAME)
    }
}
/// QueryParams is the request type to get all module params.
#[allow(clippy::derive_partial_eq_without_eq)]
#[derive(Clone, PartialEq, ::prost::Message)]
pub struct QueryParamsRequest {}
impl ::prost::Name for QueryParamsRequest {
    const NAME: &'static str = "QueryParamsRequest";
    const PACKAGE: &'static str = "bitsong.cadence.v1";
    fn full_name() -> ::prost::alloc::string::String {
        ::prost::alloc::format!("bitsong.cadence.v1.{}", Self::NAME)
    }
}
/// QueryCadenceContractsResponse is the response type for the Query/CadenceContracts RPC method.
#[allow(clippy::derive_partial_eq_without_eq)]
#[derive(Clone, PartialEq, ::prost::Message)]
pub struct QueryParamsResponse {
    #[prost(message, optional, tag = "1")]
    pub params: ::core::option::Option<Params>,
}
impl ::prost::Name for QueryParamsResponse {
    const NAME: &'static str = "QueryParamsResponse";
    const PACKAGE: &'static str = "bitsong.cadence.v1";
    fn full_name() -> ::prost::alloc::string::String {
        ::prost::alloc::format!("bitsong.cadence.v1.{}", Self::NAME)
    }
}
/// MsgRegisterCadenceContract is the Msg/RegisterCadenceContract request type.
#[allow(clippy::derive_partial_eq_without_eq)]
#[derive(Clone, PartialEq, ::prost::Message)]
pub struct MsgRegisterCadenceContract {
    /// The address of the sender.
    #[prost(string, tag = "1")]
    pub sender_address: ::prost::alloc::string::String,
    /// The address of the contract to register.
    #[prost(string, tag = "2")]
    pub contract_address: ::prost::alloc::string::String,
}
impl ::prost::Name for MsgRegisterCadenceContract {
    const NAME: &'static str = "MsgRegisterCadenceContract";
    const PACKAGE: &'static str = "bitsong.cadence.v1";
    fn full_name() -> ::prost::alloc::string::String {
        ::prost::alloc::format!("bitsong.cadence.v1.{}", Self::NAME)
    }
}
/// MsgRegisterCadenceContractResponse defines the response structure for executing a
/// MsgRegisterCadenceContract message.
#[allow(clippy::derive_partial_eq_without_eq)]
#[derive(Clone, PartialEq, ::prost::Message)]
pub struct MsgRegisterCadenceContractResponse {}
impl ::prost::Name for MsgRegisterCadenceContractResponse {
    const NAME: &'static str = "MsgRegisterCadenceContractResponse";
    const PACKAGE: &'static str = "bitsong.cadence.v1";
    fn full_name() -> ::prost::alloc::string::String {
        ::prost::alloc::format!("bitsong.cadence.v1.{}", Self::NAME)
    }
}
/// MsgUnregisterCadenceContract is the Msg/UnregisterCadenceContract request type.
#[allow(clippy::derive_partial_eq_without_eq)]
#[derive(Clone, PartialEq, ::prost::Message)]
pub struct MsgUnregisterCadenceContract {
    /// The address of the sender.
    #[prost(string, tag = "1")]
    pub sender_address: ::prost::alloc::string::String,
    /// The address of the contract to unregister.
    #[prost(string, tag = "2")]
    pub contract_address: ::prost::alloc::string::String,
}
impl ::prost::Name for MsgUnregisterCadenceContract {
    const NAME: &'static str = "MsgUnregisterCadenceContract";
    const PACKAGE: &'static str = "bitsong.cadence.v1";
    fn full_name() -> ::prost::alloc::string::String {
        ::prost::alloc::format!("bitsong.cadence.v1.{}", Self::NAME)
    }
}
/// MsgUnregisterCadenceContractResponse defines the response structure for executing a
/// MsgUnregisterCadenceContract message.
#[allow(clippy::derive_partial_eq_without_eq)]
#[derive(Clone, PartialEq, ::prost::Message)]
pub struct MsgUnregisterCadenceContractResponse {}
impl ::prost::Name for MsgUnregisterCadenceContractResponse {
    const NAME: &'static str = "MsgUnregisterCadenceContractResponse";
    const PACKAGE: &'static str = "bitsong.cadence.v1";
    fn full_name() -> ::prost::alloc::string::String {
        ::prost::alloc::format!("bitsong.cadence.v1.{}", Self::NAME)
    }
}
/// MsgUnjailCadenceContract is the Msg/UnjailCadenceContract request type.
#[allow(clippy::derive_partial_eq_without_eq)]
#[derive(Clone, PartialEq, ::prost::Message)]
pub struct MsgUnjailCadenceContract {
    /// The address of the sender.
    #[prost(string, tag = "1")]
    pub sender_address: ::prost::alloc::string::String,
    /// The address of the contract to unjail.
    #[prost(string, tag = "2")]
    pub contract_address: ::prost::alloc::string::String,
}
impl ::prost::Name for MsgUnjailCadenceContract {
    const NAME: &'static str = "MsgUnjailCadenceContract";
    const PACKAGE: &'static str = "bitsong.cadence.v1";
    fn full_name() -> ::prost::alloc::string::String {
        ::prost::alloc::format!("bitsong.cadence.v1.{}", Self::NAME)
    }
}
/// MsgUnjailCadenceContractResponse defines the response structure for executing a
/// MsgUnjailCadenceContract message.
#[allow(clippy::derive_partial_eq_without_eq)]
#[derive(Clone, PartialEq, ::prost::Message)]
pub struct MsgUnjailCadenceContractResponse {}
impl ::prost::Name for MsgUnjailCadenceContractResponse {
    const NAME: &'static str = "MsgUnjailCadenceContractResponse";
    const PACKAGE: &'static str = "bitsong.cadence.v1";
    fn full_name() -> ::prost::alloc::string::String {
        ::prost::alloc::format!("bitsong.cadence.v1.{}", Self::NAME)
    }
}
/// MsgUpdateParams is the Msg/UpdateParams request type.
///
/// Since: cosmos-sdk 0.47
#[allow(clippy::derive_partial_eq_without_eq)]
#[derive(Clone, PartialEq, ::prost::Message)]
pub struct MsgUpdateParams {
    /// authority is the address of the governance account.
    #[prost(string, tag = "1")]
    pub authority: ::prost::alloc::string::String,
    /// params defines the x/cadence parameters to update.
    ///
    /// NOTE: All parameters must be supplied.
    #[prost(message, optional, tag = "2")]
    pub params: ::core::option::Option<Params>,
}
impl ::prost::Name for MsgUpdateParams {
    const NAME: &'static str = "MsgUpdateParams";
    const PACKAGE: &'static str = "bitsong.cadence.v1";
    fn full_name() -> ::prost::alloc::string::String {
        ::prost::alloc::format!("bitsong.cadence.v1.{}", Self::NAME)
    }
}
/// MsgUpdateParamsResponse defines the response structure for executing a
/// MsgUpdateParams message.
///
/// Since: cosmos-sdk 0.47
#[allow(clippy::derive_partial_eq_without_eq)]
#[derive(Clone, PartialEq, ::prost::Message)]
pub struct MsgUpdateParamsResponse {}
impl ::prost::Name for MsgUpdateParamsResponse {
    const NAME: &'static str = "MsgUpdateParamsResponse";
    const PACKAGE: &'static str = "bitsong.cadence.v1";
    fn full_name() -> ::prost::alloc::string::String {
        ::prost::alloc::format!("bitsong.cadence.v1.{}", Self::NAME)
    }
}
include!("bitsong.cadence.v1.serde.rs");
include!("bitsong.cadence.v1.tonic.rs");
// @@protoc_insertion_point(module)
