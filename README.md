# sekai
Kira Hub

## Create order book
```sh
# command
sekaid tx kiraHub createOrderBook base quote mnemonic --from validator --keyring-backend=test --chain-id testing
{"body":{"messages":[{"@type":"/kira.kiraHub.MsgCreateOrderBook","Base":"base","Quote":"quote","Mnemonic":"mnemonic","Curator":"kira17pzan0q4d5acyykygwqass8z0crjvflhjq3qvm"}],"memo":"","timeout_height":"0","extension_options":[],"non_critical_extension_options":[]},"auth_info":{"signer_infos":[],"fee":{"amount":[],"gas_limit":"200000"}},"signatures":[]}

confirm transaction before signing and broadcasting [y/N]: y

# response
{"height":"3","txhash":"0C727E276E9808167BEE7C8048704F59EAB81CC78F785D91FFDB1C22B286FD57","codespace":"","code":0,"data":"0A110A0F6372656174656F72646572626F6F6B","raw_log":"[{\"events\":[{\"type\":\"message\",\"attributes\":[{\"key\":\"action\",\"value\":\"createorderbook\"}]}]}]","logs":[{"msg_index":0,"log":"","events":[{"type":"message","attributes":[{"key":"action","value":"createorderbook"}]}]}],"info":"","gas_wanted":"200000","gas_used":"50086","tx":null,"timestamp":""}
```

## Query order book
- By ID
Ex1.
```sh
# command
sekaid query kiraHub listorderbooks ID e6a8fc6cf92e157f9f03580291b6e5db --chain-id testing

# response
{"orderbooks":[{"ID":"e6a8fc6cf92e157f9f03580291b6e5db","Index":0,"Base":"base","Quote":"quote","Mnemonic":"mnemonic","Curator":"kira17pzan0q4d5acyykygwqass8z0crjvflhjq3qvm"}]}
```
Ex2.
```sh
# command
sekaid query kiraHub listorderbooks ID e6a8 --chain-id testing

# response
{"orderbooks":[{"ID":"","Index":0,"Base":"","Quote":"","Mnemonic":"","Curator":""}]}
```
- By curator
```sh
# command
sekaid query kiraHub listorderbooks Curator $(sekaid keys show -a validator --keyring-backend=test) --chain-id testing

# response
{"orderbooks":[{"ID":"e6a8fc6cf92e157f9f03580291b6e5db","Index":0,"Base":"base","Quote":"quote","Mnemonic":"mnemonic","Curator":"kira17pzan0q4d5acyykygwqass8z0crjvflhjq3qvm"}]}
```
- By Index

Ex1.
```sh
# command
sekaid query kiraHub listorderbooks Index 0 --chain-id testing

# response
{"orderbooks":[{"ID":"e6a8fc6cf92e157f9f03580291b6e5db","Index":0,"Base":"base","Quote":"quote","Mnemonic":"mnemonic","Curator":"kira17pzan0q4d5acyykygwqass8z0crjvflhjq3qvm"}]}
```

Ex2.
```sh
# command
sekaid query kiraHub listorderbooks Index 10 --chain-id testing

# response
{"orderbooks":[]}
```

- By Quote
Ex1.
```sh
# command
sekaid query kiraHub listorderbooks Quote quote --chain-id testing

# response
{"orderbooks":[{"ID":"e6a8fc6cf92e157f9f03580291b6e5db","Index":0,"Base":"base","Quote":"quote","Mnemonic":"mnemonic","Curator":"kira17pzan0q4d5acyykygwqass8z0crjvflhjq3qvm"}]}
```
Ex2.
```sh
# command
sekaid query kiraHub listorderbooks Quote q --chain-id testing

# response
{"orderbooks":[]}
```
- By Base
Ex1.
```sh
# command
sekaid query kiraHub listorderbooks Base base --chain-id testing

# response
{"orderbooks":[{"ID":"e6a8fc6cf92e157f9f03580291b6e5db","Index":0,"Base":"base","Quote":"quote","Mnemonic":"mnemonic","Curator":"kira17pzan0q4d5acyykygwqass8z0crjvflhjq3qvm"}]}
```
Ex2.
```sh
# command
sekaid query kiraHub listorderbooks Base b --chain-id testing

# response
{"orderbooks":[]}
```

- By trading pair
```sh
# command
sekaid query kiraHub listorderbooks_tradingpair base quote --chain-id testing

# response
{"orderbooks":[{"ID":"36e0cd99f92e157f9f03580291b6e5db","Index":0,"Base":"base","Quote":"quote","Mnemonic":"mnemonic","Curator":"kira1hqvm6nup0kkntgfq9xe0hk7gxnm4wjakmmpae0"}]}
```
```sh
# command
sekaid query kiraHub listorderbooks_tradingpair base quote1 --chain-id testing

# response
{"orderbooks":[]}
```
## Create order
```sh
# command
sekaid tx kiraHub createOrder 36e0cd99f92e157f9f03580291b6e5db 0 0 0 --from validator --keyring-backend=test --chain-id=testing
{"body":{"messages":[{"@type":"/kira.kiraHub.MsgCreateOrder","OrderBookID":"36e0cd99f92e157f9f03580291b6e5db","OrderType":"limitBuy","Amount":"0","LimitPrice":"0","ExpiryTime":"0","Curator":"kira13n3el7pzynrd3d4mn92k49djktf75kngl2xlc7"}],"memo":"","timeout_height":"0","extension_options":[],"non_critical_extension_options":[]},"auth_info":{"signer_infos":[],"fee":{"amount":[],"gas_limit":"200000"}},"signatures":[]}

confirm transaction before signing and broadcasting [y/N]: y

# response
{"height":"8","txhash":"97F57E6C73053F3FF0F6838D556F91DCCC133110F30D4B8A72AB87A96B5FFAF0","codespace":"","code":0,"data":"0A0D0A0B6372656174656F72646572","raw_log":"[{\"events\":[{\"type\":\"message\",\"attributes\":[{\"key\":\"action\",\"value\":\"createorder\"}]}]}]","logs":[{"msg_index":0,"log":"","events":[{"type":"message","attributes":[{"key":"action","value":"createorder"}]}]}],"info":"","gas_wanted":"200000","gas_used":"45716","tx":null,"timestamp":""}
```
## Cancel order
```sh
# command
sekaid tx kiraHub cancelOrder a991645f855efbb8855efbb891b6e5db --from validator --keyring-backend=test --chain-id=testing
{"body":{"messages":[{"@type":"/kira.kiraHub.MsgCancelOrder","OrderID":"a991645f855efbb8855efbb891b6e5db","Curator":""}],"memo":"","timeout_height":"0","extension_options":[],"non_critical_extension_options":[]},"auth_info":{"signer_infos":[],"fee":{"amount":[],"gas_limit":"200000"}},"signatures":[]}

confirm transaction before signing and broadcasting [y/N]: y

# response
{"height":"49","txhash":"D906B3669842655EA834978F9BBA8ADAFCD533AAD4E95B68BAB7435F7C2BEFBC","codespace":"","code":0,"data":"0A0D0A0B6372656174656F72646572","raw_log":"[{\"events\":[{\"type\":\"message\",\"attributes\":[{\"key\":\"action\",\"value\":\"createorder\"}]}]}]","logs":[{"msg_index":0,"log":"","events":[{"type":"message","attributes":[{"key":"action","value":"createorder"}]}]}],"info":"","gas_wanted":"200000","gas_used":"40936","tx":null,"timestamp":""}
```
## Query order
```sh
# command
sekaid query kiraHub listorders 36e0cd99f92e157f9f03580291b6e5db 0 0

# response
{"orders":[{"ID":"","Index":0,"OrderBookID":"36e0cd99f92e157f9f03580291b6e5db","OrderType":"limitBuy","Amount":"0","LimitPrice":"0","ExpiryTime":"0","IsCancelled":false,"Curator":"kira13n3el7pzynrd3d4mn92k49djktf75kngl2xlc7"}]}
```
```sh
# command
sekaid query kiraHub listorders 36e0cd99f92e157f9f03580291b6e5b 0 0

# response
{"orders":[]}
```
## upsertSignerKey

- CLI

```sh
# Secp256k1 type
sekaid tx kiraHub upsertSignerKey AnzIM9IcLb07Cvwq3hdMJuuRofAgxfDekkD3nJUPPw0w
sekaid tx kiraHub upsertSignerKey AnzIM9IcLb07Cvwq3hdMJuuRofAgxfDekkD3nJUPPw0w --key-type=Secp256k1

# Ed25519 type
sekaid tx kiraHub upsertSignerKey TXgDkmTYpPRwU/PvDbfbhbwiYA7jXMwQgNffHVey1dC644OBBI4OQdf4Tro6hzimT1dHYzPiGZB0aYWJBC2keQ== --key-type=Ed25519

# enabled false
sekaid tx kiraHub upsertSignerKey AnzIM9IcLb07Cvwq3hdMJuuRofAgxfDekkD3nJUPPw0w --enabled=false

# permissions
sekaid tx kiraHub upsertSignerKey AnzIM9IcLb07Cvwq3hdMJuuRofAgxfDekkD3nJUPPw0w --permissions=1,2

# expiry-time (set when this key expire if does not set it's automatically set to 10 days after current timestamp)
sekaid tx kiraHub upsertSignerKey AnzIM9IcLb07Cvwq3hdMJuuRofAgxfDekkD3nJUPPw0w --expiry-time=1598247750
```

Ex1.
```sh
# command
sekaid tx kiraHub upsertSignerKey AnzIM9IcLb07Cvwq3hdMJuuRofAgxfDekkD3nJUPPw0w --from validator --keyring-backend=test --chain-id=testing
{"body":{"messages":[{"@type":"/kira.kiraHub.MsgUpsertSignerKey","PubKey":"AnzIM9IcLb07Cvwq3hdMJuuRofAgxfDekkD3nJUPPw0w","KeyType":"Secp256k1","ExpiryTime":"1599186595","Enabled":true,"Permissions":[],"Curator":"kira13n3el7pzynrd3d4mn92k49djktf75kngl2xlc7"}],"memo":"","timeout_height":"0","extension_options":[],"non_critical_extension_options":[]},"auth_info":{"signer_infos":[],"fee":{"amount":[],"gas_limit":"200000"}},"signatures":[]}

confirm transaction before signing and broadcasting [y/N]: y

# response
{"height":"11543","txhash":"34B0D49958A881D280392644D3CBC358100DCF8353F1DCC26899D3B039CAAF02","codespace":"","code":0,"data":"0A110A0F7570736572747369676E65726B6579","raw_log":"[{\"events\":[{\"type\":\"message\",\"attributes\":[{\"key\":\"action\",\"value\":\"upsertsignerkey\"}]}]}]","logs":[{"msg_index":0,"log":"","events":[{"type":"message","attributes":[{"key":"action","value":"upsertsignerkey"}]}]}],"info":"","gas_wanted":"200000","gas_used":"44086","tx":null,"timestamp":""}
```

## Query signer keys

```sh
# command
sekaid query kiraHub getsignerkeys --curator $(sekaid keys show -a validator --keyring-backend=test)

# response
{"signerkeys":[{"PubKey":"AnzIM9IcLb07Cvwq3hdMJuuRofAgxfDekkD3nJUPPw0w","KeyType":"Secp256k1","ExpiryTime":"1599187090","Enabled":true,"Permissions":[],"Curator":"kira1xsq0wapm5t975k3hn2rj4y2zhnm5up9d59uhpy"}]}
```
- Rest

```sh
# TODO: should add readme for sample upsertSignerKey with same examples of CLI
```
---
`dev` branch
