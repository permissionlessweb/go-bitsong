# Messages

Below are type definitions for msgs
 
## x/fantoken
 

## Event Structures
### EventIssue
| Field Name | Type | Tag | Description |
|------------|------|-----|-------------|
| denom | string | 1 | Denomination of the newly issued fan token |

### EventDisableMint
| Field Name | Type | Tag | Description |
|------------|------|-----|-------------|
| denom | string | 1 | Denomination of the fan token that had minting disabled |

### EventMint
| Field Name | Type | Tag | Description |
|------------|------|-----|-------------|
| recipient | string | 1 | Address that received the minted tokens |
| coin | string | 2 | Coin amount and denomination minted |

### EventBurn
| Field Name | Type | Tag | Description |
|------------|------|-----|-------------|
| sender | string | 1 | Address that burned the tokens |
| coin | string | 2 | Coin amount and denomination burned |

### EventSetAuthority
| Field Name | Type | Tag | Description |
|------------|------|-----|-------------|
| denom | string | 1 | Denomination of the fan token |
| old_authority | string | 2 | Previous authority address |
| new_authority | string | 3 | New authority address |

### EventSetMinter
| Field Name | Type | Tag | Description |
|------------|------|-----|-------------|
| denom | string | 1 | Denomination of the fan token |
| old_minter | string | 2 | Previous minter address |
| new_minter | string | 3 | New minter address |

### EventSetUri
| Field Name | Type | Tag | Description |
|------------|------|-----|-------------|
| denom | string | 1 | Denomination of the fan token that had its URI updated |

---

## Main Structures
### Metadata
| Field Name | Type | Tag | Description |
|------------|------|-----|-------------|
| name | string | 1 | Token name (e.g., "Kitty Punk") |
| symbol | string | 2 | Token symbol for exchanges (e.g., "KITTY") |
| uri | string | 3 | Optional URI for additional token information |
| authority | string | 4 | Address allowed to set new URI |

### FanToken
| Field Name | Type | Tag | Description |
|------------|------|-----|-------------|
| denom | string | 1 | Unique token denomination (e.g., ft<hash>) |
| max_supply | string | 2 | Maximum mintable tokens in micro units |
| minter | string | 3 | Address authorized to mint tokens |
| meta_data | optional Metadata | 4 | Optional metadata for the token |

### Params
| Field Name | Type | Tag | Description |
|------------|------|-----|-------------|
| issue_fee | optional Coin | 1 | Fee required to issue new tokens |
| mint_fee | optional Coin | 2 | Fee required to mint tokens |
| burn_fee | optional Coin | 3 | Fee required to burn tokens |

### GenesisState
| Field Name | Type | Tag | Description |
|------------|------|-----|-------------|
| params | optional Params | 1 | Module parameters for Fantoken |
| fan_tokens | repeated FanToken | 2 | List of fan tokens in genesis state |

---

## Query Structures
### QueryFanTokenRequest
| Field Name | Type | Tag | Description |
|------------|------|-----|-------------|
| denom | string | 1 | Token denomination to query |

### QueryFanTokenResponse
| Field Name | Type | Tag | Description |
|------------|------|-----|-------------|
| fantoken | optional FanToken | 1 | Optional fan token in response |

### QueryFanTokensRequest
| Field Name | Type | Tag | Description |
|------------|------|-----|-------------|
| authority | string | 1 | Filter tokens by authority address |
| pagination | optional PageRequest | 2 | Optional pagination parameters |

### QueryFanTokensResponse
| Field Name | Type | Tag | Description |
|------------|------|-----|-------------|
| fantokens | repeated FanToken | 1 | List of fan tokens |
| pagination | optional PageResponse | 2 | Optional pagination metadata |

---

## Message Structures
### MsgIssue
| Field Name | Type | Tag | Description |
|------------|------|-----|-------------|
| symbol | string | 1 | Immutable token symbol |
| name | string | 2 | Immutable token name |
| max_supply | string | 3 | Maximum token supply in micro units |
| authority | string | 4 | URI metadata authority address |
| minter | string | 5 | Mint authority and fee payer |
| uri | string | 6 | Current token URI |

### MsgIssueResponse
| Field Name | Type | Tag | Description |
|------------|------|-----|-------------|
| denom | string | 1 | Denomination of the issued token |

### MsgDisableMint
| Field Name | Type | Tag | Description |
|------------|------|-----|-------------|
| denom | string | 1 | Denomination to disable minting |
| minter | string | 2 | Current minter address |

### MsgDisableMintResponse
| Field Name | Type | Tag | Description |
|------------|------|-----|-------------|
| denom | string | 1 | Denomination that had minting disabled |

### MsgMint
| Field Name | Type | Tag | Description |
|------------|------|-----|-------------|
| recipient | string | 1 | Address to receive minted tokens |
| coin | optional Coin | 2 | Amount and denomination to mint |
| minter | string | 3 | Address performing the minting |

### MsgMintResponse
| Field Name | Type | Tag | Description |
|------------|------|-----|-------------|
| recipient | string | 1 | Address that received tokens |
| coin | optional Coin | 2 | Amount and denomination minted |

### MsgBurn
| Field Name | Type | Tag | Description |
|------------|------|-----|-------------|
| coin | optional Coin | 1 | Amount and denomination to burn |
| sender | string | 2 | Address initiating burn operation |

### MsgBurnResponse
| Field Name | Type | Tag | Description |
|------------|------|-----|-------------|
| sender | string | 1 | Address that burned tokens |
| coin | optional Coin | 2 | Amount and denomination burned |

### MsgSetMinter
| Field Name | Type | Tag | Description |
|------------|------|-----|-------------|
| denom | string | 1 | Token denomination to update |
| old_minter | string | 2 | Current minter address |
| new_minter | string | 3 | New minter address to set |

### MsgSetMinterResponse
| Field Name | Type | Tag | Description |
|------------|------|-----|-------------|
| denom | string | 1 | Token denomination updated |
| old_minter | string | 2 | Previous minter address |
| new_minter | string | 3 |


## x/smart-accounts
<!-- todo: implement hyperlinks to  -->
### Params
| Field Name | Type | Tag | Description |
|------------|------|-----|-------------|
| maximum_unauthenticated_gas | uint64 | 1 | Maximum amount of gas that can be used to authenticate a transaction in ante handler without having fee payer authenticated |
| is_smart_account_active | bool | 2 | Defines the state of the authenticator - if false, classic cosmos sdk authentication will be used instead |
| circuit_breaker_controllers | repeated string | 3 | List of addresses that are allowed to set is_smart_account_active without going through governance |

---

### AccountAuthenticator
| Field Name | Type | Tag | Description |
|------------|------|-----|-------------|
| id | uint64 | 1 | Uniquely identifies the authenticator instance |
| type | string | 2 | Specifies the category of the AccountAuthenticator for differentiation and precise data retrieval |
| config | bytes | 3 | Versatile field that facilitates complex authentication processes with type-specific interpretation |

---

### AuthenticatorData
| Field Name | Type | Tag | Description |
|------------|------|-----|-------------|
| address | string | 1 | Account address key (can have multiple authenticators) |
| authenticators | repeated AccountAuthenticator | 2 | Account's authenticators (SignatureVerification, AllOfs, CosmWasmAuthenticators, etc) |

---

### GenesisState
| Field Name | Type | Tag | Description |
|------------|------|-----|-------------|
| params | optional Params | 1 | Parameters for the authenticator module |
| next_authenticator_id | uint64 | 2 | Next available authenticator ID counter |
| authenticator_data | repeated AuthenticatorData | 3 | Genesis account authenticator data for multiple accounts |

---

### QueryParamsRequest
No fields required.

---

### QueryParamsResponse
| Field Name | Type | Tag | Description |
|------------|------|-----|-------------|
| params | optional Params | 1 | Holds all module parameters |

---

### GetAuthenticatorsRequest
| Field Name | Type | Tag | Description |
|------------|------|-----|-------------|
| account | string | 1 | Account address to query authenticators for |

---

### GetAuthenticatorsResponse
| Field Name | Type | Tag | Description |
|------------|------|-----|-------------|
| account_authenticators | repeated AccountAuthenticator | 1 | List of authenticators associated with the account |

---

### GetAuthenticatorRequest
| Field Name | Type | Tag | Description |
|------------|------|-----|-------------|
| account | string | 1 | Account address to query |
| authenticator_id | uint64 | 2 | Specific authenticator ID to retrieve |

---

### GetAuthenticatorResponse
| Field Name | Type | Tag | Description |
|------------|------|-----|-------------|
| account_authenticator | optional AccountAuthenticator | 1 | Optional field for the requested authenticator |

---

### MsgAddAuthenticator
| Field Name | Type | Tag | Description |
|------------|------|-----|-------------|
| sender | string | 1 | Transaction signer address |
| authenticator_type | string | 2 | Type of authenticator to add |
| data | bytes | 3 | Authenticator-specific configuration data |

---

### MsgAddAuthenticatorResponse
| Field Name | Type | Tag | Description |
|------------|------|-----|-------------|
| success | bool | 1 | Boolean indicating operation success status |

---

### MsgRemoveAuthenticator
| Field Name | Type | Tag | Description |
|------------|------|-----|-------------|
| sender | string | 1 | Transaction signer address |
| id | uint64 | 2 | Authenticator ID to remove |

---

### MsgRemoveAuthenticatorResponse
| Field Name | Type | Tag | Description |
|------------|------|-----|-------------|
| success | bool | 1 | Boolean indicating operation success status |

---

### MsgSetActiveState
| Field Name | Type | Tag | Description |
|------------|------|-----|-------------|
| sender | string | 1 | Transaction signer address |
| active | bool | 2 | Boolean to activate/deactivate the authenticator module |

---

### MsgSetActiveStateResponse
No fields required.

---

### TxExtension
| Field Name | Type | Tag | Description |
|------------|------|-----|-------------|
| selected_authenticators | repeated uint64 | 1 | Authenticator IDs for chosen authenticators per message |

---

 