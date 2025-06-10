# Performing Actions
To perform actions with an account registered with custom authenticators, there are a few modifications to the standard message being broadcasted to be included in the mempool. 

## Forming Msgs For Account That Enable Custom Authentication
The primary focus for ui integration with smart accounts is the use of the cosmos-sdk msg 
- `TxExtension`
- `non_critical_extension_options`


### Rust
```rs
let tx_extension = Any {
    type_url: "/bitsong.smartaccount.v1beta1.TxExtension".into(),
    value: to_json_binary(&TxExtension { 
        selected_authenticators: vec![1],
        smart_account: SmartAccountAuth{
            public_keys:vec!["autautesu546hwh4"],
            signatures:vec!["==35h3jb63"]}
        })?.to_vec(),
};
```
If there is any data set in `smart_account` object of the TxExtension that is used to trigger custom account authentication, then the module expects that this account owner is using an aggregated consensus method, and forms the `AuthenticationRequest` object in a specfic way for custom authenticators to be built that can authenticate aggregated ECDSA, such as BLS12-381. **If an account is not making use of an aggregated key authenticator, and data is set in this object, it will always fail.**

### Registering Authenticators For An Account 
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
The params are the values your custom smart contract expects to be set inside the go-module for this specific account. This will vary from authenticator to authenticator, and may not be needed at all.

```rs
// - A: Define the actions the smart-account is to perform
let wavs_action_msg = Any { }.to_bytes()?,

// - B: Define the tx extension, specifying which smart-account to aquire
  
// - C: Form the main Tx Body 
let wavs_broadcast_msg: TxBody = TxBody {
    messages: vec![wavs_action_msg],
    memo: "Cosmic Wavs Account Action".into(),
    timeout_height: block_height + 100u64,
    extension_options: vec![],
    non_critical_extension_options: vec![ Any{}].to_vec(),
};

//  - D: Generate signature using ECDSA method of choice (this example makes use of BLS12-381)
```   
## TS/JS
## Go

<!-- Todo: implement go utils libar -->