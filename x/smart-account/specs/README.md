# Smart Account Workflow

This module's standout feature is allowing **an account-by-account customization on how transactions are authenticated.**

Bitsong nodes now make use of their compute resources prior to a block being finalized, in order to handle the minimum processes needed to verify whether an account is valid to be included in the mempool.

**This logic is now programmable via smart contracts & additional configurations, vastly expanding how transactions can be implemented.**

### Minimum Requirements
- **One secp256k1 keypair:** registers authorization methods to, generally funded with gas to pay for transaction fees.
- An authentication method of choice
- `TxExtension`: appended to each message, to specify which authenticator to use (defaults to regular key ECDSA authorization)

## Workflow Module Uses To Authorize Transactions

### 1. Initial Gas Limit  
To safegaurd gas consumtion during authorizations, a maximum gas limit for actions performed by the module is enforced. If any transaction to be authenticated exceeds this limit during the authentication phase, it is rejected before execution.

### 2. Identify Fee Payer:  
**The first signer of the transaction is always the fee payer.** This means that for multiple messages, all fees these messages incur will be paid by the single account. Front end developers should keep this in mind when implementing login workflows.

### 3. Authenticate Each Message  
Multiple messages can be authorized at once. For each message to be authorized:  
- its associated account is identified  
- any authenticator registered for this account is fetched  
- its message is then either validated or rejected.

### 4. Gas Limit Resets  
Once the authorizations are complete, the module resets the gas limits in preparation for the last steps, which is what will be used for the actual transaction execution.

### 5. **Track**:  
After all messages are authenticated, the `Track` function notifies each executed authenticator. This allows authenticators to store any transaction-specific data they might need for the future, such as execution logs or state changes.

### 6. Execute Message  
The transaction is executed

### 7. Confirm Execution  
After all messages are executed, the module calls the `ConfirmExecution` function for each of the authenticators that authenticated the tx.  
This allows authenticators to enforce rules that depend on the outcome of the message execution, like spending limits or transaction frequency caps, even if the initial authorization passed.

---

## Default Authentication Options  
___  
Below we review the available ways transactions sent to nodes can be authenticated as to whether they are included when a block gets finalized.

### AllOf  
AllOf means that you can stack authentication requirements together, requiring 100% of the methods to be valid. An example would be a multi-sig, or even a two-step verification making use of one of the other authentication options.

### AnyOf  
**AnyOf will recognize transactions as valid if `1 of n` authenticators that an account registered to be used is successful.**

### Signature Verification  
The signature verification authenticator is the default authenticator for all accounts. It verifies that the signer of a message is the same as the account associated with the message.

---

## CosmWasm Authenticator  

**This authenticator option allows us to have custom smart contracts made to handle how accounts can have actions authorized for them.**  
When an account registers a contract address to be used as an authentication method, the specific parameters sent by the account registering are **not** stored in the contract state, but rather in the module storage, which keeps things light and keeps the compute resources minimal when making use of the contract.  

### How it works  
Bitsong will make use of the contract `sudo` entry point, which can only be called by the chain itself. This means when an account using CosmWasm authentication submits a tx, the CosmWasmVM is deterministically processed and either validates or rejects the transaction prior to deterministically processing the actual message to perform.

### Module's Go Message Structure  
To register a CosmWasm authenticator to an account, use the following format:  
`MsgAddAuthenticator` arguments:
```text
sender: <bech32_address>  
type: "CosmwasmAuthenticatorV1"  
data: json_bytes({  
    contract: "<contract_address>",  
    params: [<byte_array>]  
})
```

The **params** field is a JSON-encoded value for any parameters to save regarding this authenticator. This contrasts with saving these parameters into the contract state, which is more expensive when retrieving the state at a later date.  

**Contract storage should be used only when the authenticator needs to track dynamic information required for authentication logic.**

---

## MessageFilter  

**The message filter authentication means that you can register to authenticate by default any specific message with a given pattern.**  

**This is a very powerful filter, as it can bypass default authentication for your account to perform actions, so use with care!**  

Recognizing these accounts more as a **permissionless utility account** may help visualize how this authenticator can be used.  

For example, a faucet-like account can be created by allowing any spend messages with specific values:  
```json
{
  "@type": "/cosmos.bank.v1beta1.MsgSend",
  "amount": [
    {
      "denom": "ubtsg",
      "amount": "69"
    }
  ]
}
```

Or a way to mint new tokens during a streaming session:  
```json
{
   "@type":"/bitsong.fantoken.v1beta1.MsgMint",
   "sender":"bitsong1...", 
   // ... other required fields
}
```

---

## Risks & Limitations Present  

### Registration Of Accounts  
Improperly configured authenticators during registration could lead to permanent loss of account access if not carefully managed. Security audits are recommended for custom contracts, and always test locally any deployment workflow first.

### Fees and Gas Consumption  
Custom authentication logic may incur higher gas costs compared to standard signature verification. Gas metering during authentication phases must be strictly enforced to prevent resource exhaustion attacks.

### Composite Authenticators  
Combining multiple authentication methods (AllOf/AnyOf) increases complexity. Ensure logic interactions don't create unintended authorization pathways or denial-of-service vectors.

### Composite IDs  
Account identifiers used in composite authenticators have size limitations due to protobuf encoding constraints. Avoid deeply nested composite structures that exceed chain-specific message size limits.

### Composite Signatures  
When using multiple signature-based authenticators, ensure proper validation of signature uniqueness and replay protection mechanisms to prevent cryptographic vulnerabilities.

---

This module provides unprecedented flexibility in transaction authentication while maintaining the security guarantees of the Cosmos SDK. Developers should carefully consider the tradeoffs between customizability and computational overhead when designing authentication contracts.