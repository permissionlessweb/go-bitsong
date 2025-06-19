# Performing Actions
Lets review the workflow to perform on-chain actions with custom authenti

## 0. Registering Authenticators To Accounts 
The available types to register an authenticator are currently:
- `CosmwasmAuthenticatorV1`
- `SignatureVerification`
- `Bls12381V1`
- `AllOf`
- `AnyOf`
- `MessageFilter`
- `PartitionedAnyOf`
- `PartitionedAllOf`

Registering an account requires a default signature from the accounts public key, including any parameters the authentictor may require.
```rs
    let register_smart_account = Any {
        type_url: "/bitsong.smartaccount.v1beta1.MsgAddAuthenticator".into(),
        value: to_json_binary(&MsgAddAuthenticator {
            sender: chain.sender_addr().to_string(),
            authenticator_type: "CosmwasmAuthenticatorV1".into(),
            data: to_json_binary(&CosmwasmAuthenticatorInitData {
                contract: suite.wavs.address()?.to_string(),
                params: vec![],
            })?
            .to_vec(),
        })?
        .to_vec(),
    };
```

## 1. Broadcasting Tx's using custom authenticator

### `TxExtension`
The `TxExtension` object is expected to be included in a transaction's `non_critical_extension_options`. This allows the module to identify the authenticator to use and optionally, any public keys and signatures used for aggregated signature authentication.

### Authenticator IDs
An account can have multiple authenticators, each identified by a unique numerical value. When a transaction includes multiple messages, each message requires an authenticator, and they must be provided in the order of the messages. If no authenticator is provided for a message, the module will attempt to authenticate the transaction using a standard secp256k1 signature from the original public key associated with the address.

### Aggregated Authentication (Optional)
To use aggregated authentication, data must be set in the `ag_auth` object. This signals to the module that the transaction has been authorized by a set of keys and signatures that can be aggregated together, as specified in the `AuthenticationRequest`. This is particularly relevant for authenticators that verify aggregated ECDSA public keys and signatures, such as the BLS12-381 example. It is crucial to note that if an account is not configured to use an aggregated key authenticator and data is still set in this object, any transaction will fail.

The `SignatureV2` and `SingleSignatureData` structures are used to represent signature data. 
```go go/types/signature.go
type SignatureV2 struct {
	// PubKey is the public key to use for verifying the signature
	PubKey cryptotypes.PubKey

	// Data is the actual data of the signature which includes SignMode's and
	// the actual signature from the pubkey. 
	Data SingleSignatureData

	// Sequence is the sequence of this account. Only populated in
	// SIGN_MODE_DIRECT.
	Sequence uint64
}

// SingleSignatureData represents the signature and SignMode of a single (non-multisig) signer
type SingleSignatureData struct {
	// SignMode represents the SignMode of the signature. use SIGN_MODE_DIRECT
	SignMode SignMode

	// Signature is the raw signature.
	Signature []byte
}
```
Currently, an array of JSON marshaled signature data is expected.

#### JSON Marshalling & Unmarshalling
To work with signature data, the following functions are used:
```go go/authenticator/signature.go
var sigs []signing.SignatureV2
// when forming msg to broadcast
signBz, err := authenticator.MarshalSignatureJSON(sigs)
// when writing custom aggregate authenticators
aggAuthData, err := UnmarshalSignatureJSON(cdc, aggSig.GetData())
```
#### Defining `AuthInfo`
A single signature object set in `AuthInfo.SignerInfos`, which is to be the single aggregated signature and pubkey. 

### Examples 
#### Rust
```rs
// - A: Define the actions the smart-account is to perform
let wavs_action = Any {}.to_bytes()?,
// - B: Define the tx extension, specifying which smart-account to authenticate & any aggregate keys used
let non_critical_extension_options = vec![Any {
    type_url: "/bitsong.smartaccount.v1beta1.TxExtension".into(),
    value: to_json_binary(&TxExtension { 
        selected_authenticators: vec![1],
        agg_auth: SmartAccountAuthData{signatures:vec!["==35h3jb63"]}
        })?.to_vec(),
}].to_vec();

// - C: Form the main Tx Body 
let wavs_broadcast_msg: TxBody = TxBody {
    messages: vec![wavs_action],
    memo: "Cosmic Wavs Account Action".into(),
    timeout_height: block_height + 100u64,
    extension_options: vec![],
    non_critical_extension_options,
};

//  - D: Generate signature using ECDSA method of choice (this example makes use of BLS12-381)
// ...
```   

#### TS/JS
```ts
 
// Define the TxExtension
const txExtension = TxExtension.fromPartial({
  selectedAuthenticators: [1],
  aggAuth: {
    signatures: ['==35h3jb63'],
  },
});

// Serialize TxExtension to Any
const nonCriticalExtensionOptions = [
  Any.fromPartial({
    typeUrl: '/bitsong.smartaccount.v1beta1.TxExtension',
    value: TxExtension.encode(txExtension).finish(),
  }),
];

// Form the main TxBody
const txBody = TxBody.fromPartial({
  messages: [wavsAction],
  memo: 'Cosmic Wavs Account Action',
  timeoutHeight: '100',
  extensionOptions: [],
  nonCriticalExtensionOptions: nonCriticalExtensionOptions,
});

```
#### Go

```go
// Create the TxExtension
txExtension := &types.TxExtension{
    SelectedAuthenticators: []uint64{1},
    AggAuth: &pb.SmartAccountAuthData{
        Signatures: []string{"==35h3jb63"},
    },
}

// Marshal TxExtension to Any
txExtensionAny, err := codectypes.NewAnyWithValue(txExtension)
if err != nil {
    fmt.Println(err)
    return
}

// Form the main TxBody
txBody := &tx.TxBody{
    Messages:                  []*codectypes.Any{wavsAction},
    Memo:                      "Cosmic Wavs Account Action",
    TimeoutHeight:             100, // uint64
    ExtensionOptions:          nil,
    NonCriticalExtensionOptions: []*codectypes.Any{txExtensionAny},
}
```

#### Python
```python
def create_tx_body(wavs_action, tx_extension):
    # Serialize TxExtension to Any
    tx_extension_any = any_pb2.Any(
        type_url='/bitsong.smartaccount.v1beta1.TxExtension',
        value=tx_extension.SerializeToString()
    )

    # Form the main TxBody
    tx_body = tx_pb2.TxBody(
        messages=[wavs_action],
        memo='Cosmic Wavs Account Action',
        timeout_height='100',
        extension_options=[],
        nonCriticalExtensionOptions=[tx_extension_any]
    )

    return tx_body

def create_tx_extension(selected_authenticators, agg_auth_signatures):
    # Define the TxExtension
    tx_extension = tx_extension_pb2.TxExtension(
        selected_authenticators=selected_authenticators,
        agg_auth=tx_extension_pb2.AggAuth(
            signatures=agg_auth_signatures
        )
    )

    return tx_extension

# Example usage
wavs_action = ...  # Assuming wavs_action is defined elsewhere
selected_authenticators = [1]
agg_auth_signatures = ['==35h3jb63']

tx_extension = create_tx_extension(selected_authenticators, agg_auth_signatures)
tx_body = create_tx_body(wavs_action, tx_extension)
```


