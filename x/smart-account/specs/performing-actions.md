# Performing Actions
To perform actions with an account registered with custom authenticators, there are a few modifications to the standard message being broadcasted to be included in the mempool. 

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

An account can have multiple authenticators, each identified by a unique numerical value.


## 1. Defining `TxExtension`

`TxExtension` is the value expected by this module, in order to identify which authenticator to use, also optionally any pubkeys & signatures used aggregated signature authentication.

If there is any data set in `smart_account` object of this modules `TxExtension` accepted, then module forms the `AuthenticationRequest` object in a specfic way for custom authenticators to be built that can authenticate aggregated ECDSA, such as BLS12-381. **If an account is not making use of an aggregated key authenticator, and data is set in this object, any transaction will always fail.**


The params are the values your custom smart contract expects to be set inside the go-module for this specific account. This will vary from authenticator to authenticator, and may not be needed at all.


## 2. Using custom auth via *`non_critical_extension_options`*
 ---
### Forming Msgs For Account That Enable Custom Authentication

 
```rs
// - A: Define the actions the smart-account is to perform
let wavs_action = Any {}.to_bytes()?,
// - B: Define the tx extension, specifying which smart-account to authenticate & any aggregate keys used
let non_critical_extension_options = vec![Any {
    type_url: "/bitsong.smartaccount.v1beta1.TxExtension".into(),
    value: to_json_binary(&TxExtension { 
        selected_authenticators: vec![1],
        smart_account: AgAuthData{signatures:vec!["==35h3jb63"]}
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

`AgAuthData` expects an array of `signing.SignatureV2` interfaced marshalled into their bytes, and is only usef for aggregated signature authentication methods (such as bls12-381)
## TS/JS
## Go

```
- NewBls12381Account()
- GenSimpleTxBls12381()
    - GenTxBls12381()
    - MakeTxBuilderBls381()
    - 
```


<!-- Todo: implement go utils libar -->
