# Performing Actions
To perform actions with an account registered with custom authenticators, there are a few modifications to the standard message being broadcasted to be included in the mempool. 



## Rust

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

### Forming Msgs For Account That Enabled Custom Authentication

- `TxExtension`
- `non_critical_extension_options`

```rs
// - A: Define the actions the smart-account is to perform
let wavs_action_msg = Any {
type_url: "/cosmwasm.wasm.v1.MsgExecuteContract".into(),
value: cosmos_sdk_proto::cosmwasm::wasm::v1::MsgExecuteContract {
    sender: WAVS_INFUSER_OPERATOR_ADDR.into(), // bech32 address of smart-account being operated
    contract: cw_infuser_addr.to_string(),
    msg: to_json_binary(&cw_infuser::msg::ExecuteMsg::WavsEntryPoint { infusions })?
        .to_vec(),
    funds: vec![],
}
.to_bytes()?,
};
 
// - B: Define the tx extension, specifying which smart-account to aquire
let tx_extension = Any {
    type_url: "/bitsong.smartaccount.v1beta1.TxExtension".into(),
    value: to_json_binary(&TxExtension { selected_authenticators: vec![1] })?.to_vec(),
};

// - C: Form the main Tx Body 
let wavs_broadcast_msg: TxBody = TxBody {
    messages: vec![wavs_action_msg],
    memo: "Cosmic Wavs Account Action".into(),
    timeout_height: 100u64,
    extension_options: vec![],
    non_critical_extension_options: vec![tx_extension].to_vec(),
};

//  - D: Generate signature using ECDSA method of choice (this example makes use of BLS12-381)

```

## TS/JS
## Go

<!-- Todo: implement go utils libar -->