# BAS Go SDK
## Requirement
- Go version above 1.21
## Add SDK Dependencies
```sh
$ go get github.com/bnb-attestation-service/bas-go@v0.2.1
```

replace dependencies in **go.mod**

```
cosmossdk.io/api => github.com/bnb-chain/greenfield-cosmos-sdk/api v0.0.0-20230816082903-b48770f5e210
cosmossdk.io/math => github.com/bnb-chain/greenfield-cosmos-sdk/math v0.0.0-20230816082903-b48770f5e210
github.com/cometbft/cometbft => github.com/bnb-chain/greenfield-cometbft v1.1.0
github.com/cometbft/cometbft-db => github.com/bnb-chain/greenfield-cometbft-db v0.8.1-alpha.1
github.com/cosmos/cosmos-sdk => github.com/bnb-chain/greenfield-cosmos-sdk v1.1.0
github.com/cosmos/iavl => github.com/bnb-chain/greenfield-iavl v0.20.1
github.com/syndtr/goleveldb => github.com/syndtr/goleveldb v1.0.1-0.20210819022825-2ae1ddf74ef7
github.com/consensys/gnark-crypto => github.com/consensys/gnark-crypto v0.7.0
```

We recommend that you use the robust version of the SDK for Greenfield: 
```
github.com/bnb-chain/greenfield v1.2.1-0.20231221015040-11071a6ee95b
github.com/bnb-chain/greenfield-go-sdk v1.2.1
```
## Agent

Before performing almost any operation (except for some offchain APIs), it is necessary to create a new agent. Please note that creating an agent requires providing the user's private key, so there is a certain risk involved. Please test your operations before performing any actions on the main network. The Go SDK provides a simple way to create an agent:

You can create an agent using the following approach:

```go
var _agent *Agent
if _agent, err = NewAgentFromKey(privateKey, BNBTESTRPC, BNBTESTCHAINID, GFTESTRPC, GFTESTCHAINID); err != nil {
	panic(err)
}
```

Note that you need to provide the RPC and chainID for the Binance (BNB) and Greenfield (GF) networks, such as:

```go
BNBTESTRPC     = "https://data-seed-prebsc-1-s1.bnbchain.org:8545"
BNBTESTCHAINID = 97

GFTESTRPC     = "https://gnfd-testnet-fullnode-tendermint-us.bnbchain.org:443"
GFTESTCHAINID = "greenfield_5600-1"
```

## Schema

### Create a schema

To create a schema, users first need to determine the data type of the schema. BAS data types use EVM's ABI encoding method. For example, a schema corresponding to a struct like the one below would be "string bas, uint8 nonce":

```go
type example struct {
	Bas string `abi:bas`
	Nonce uint8 `abi:nonce`
}
```

Once users determine the data structure of the schema, whether it can be revoked, and the resolver contract address of the schema, the schema can be created as follows:

```go
schema := "string bas, uint8 nonce"
revocable := true
resolver := ""
var _agent *Agent
var err error
if _agent, err = NewAgentFromKey(privateKey, BNBTESTRPC, BNBTESTCHAINID, GFTESTRPC, GFTESTCHAINID); err != nil {
	panic(err)
}
if attest, err := _agent.CreateSchema(schema, revocable, resolver); err != nil {
	panic(err)
} else {
	fmt.Println(attest)
}
```

### Get a schema

Sometimes users need to retrieve detailed information about a schema. This can be done using the schema's UID through the Go SDK:

```go
schemaUid := "0xabcdefg......"

var _agent *Agent
var err error
if _agent, err = NewAgentFromKey(privateKey, BNBTESTRPC, BNBTESTCHAINID, GFTESTRPC, GFTESTCHAINID); err != nil {
	panic(err)
}
if schema, err := _agent.GetSchema(schemaUid); err != nil {
	panic(err)
} else {
	fmt.Println(schema)
}
```

### Set name & description for a schema

Users can set the name and description for their created schema. Although theoretically anyone can, currently only the name and description set by the creator of the schema will be specially marked by BAS officials. The implementation is as follows:

```go
schemaUid := "0xabcdefg......"
name := "test_name"
var _agent *Agent
var err error
if _agent, err = NewAgentFromKey(privateKey, BNBTESTRPC, BNBTESTCHAINID, GFTESTRPC, GFTESTCHAINID); err != nil {
	panic(err)
}
if schema, err := _agent.SetSchemaName(schemaUid, name); err != nil {
	panic(err)
} else {
	fmt.Println(schema)
}
```

```go
schemaUid := "0xabcdefg......"
descrip := "test name for bas go"
var _agent *Agent
var err error
if _agent, err = NewAgentFromKey(privateKey, BNBTESTRPC, BNBTESTCHAINID, GFTESTRPC, GFTESTCHAINID); err != nil {
	panic(err)
}
if schema, err := _agent.SetSchemaDescription(schemaUid, descrip); err != nil {
	panic(err)
} else {
	fmt.Println(schema)
}
```

## Onchain

### Create an onchain attestation

To create an onchain attestation, users need to provide:

- schemaUid
    - Users can find the UID corresponding to their created or required schema in the browser
- data
    - Data is a byte array encoded in ABI format. You need to use the data structure of the schema and the EVM Go package to generate this encoding, for example:
        
        ```go
        
        schema := "string name"
        data := map[string]interface{}{
        	"name": "bas",
        }
        
        _schema := fmt.Sprintf("tuple(%s)", schema)
        fmt.Println(_schema, data)
        typ := abi.MustNewType(_schema)
        if res, err := typ.Encode(data); err != nil {
        	return nil, err
        } else {
        	return res, nil
        }
        ```
        
        This encoding method is also integrated in offchain.EncodeData.
        
- revocable
    - Needs to be consistent with the revocable property of the schema being used
- expirationTime
    - Unix timestamp accurate to the second, 0 means no expiration time

Users can then create an onchain attestation as follows:

```go
var _agent *Agent
var err error
if _agent, err = NewAgentFromKey(privateKey, BNBTESTRPC, BNBTESTCHAINID, GFTESTRPC, GFTESTCHAINID); err != nil {
	panic(err)
}
_agent.OnchainAttest(
	schemaUid,
	data,
	revocable,
	expirationTime,
)
```

### Get an Attestation onchain

Users can get detailed information about an onchain attestation using OnchainGetAttestation:

```go
uid := "0xabcdefg..."
var _agent *Agent
var err error
if _agent, err = NewAgentFromKey(privateKey, BNBTESTRPC, BNBTESTCHAINID, GFTESTRPC, GFTESTCHAINID); err != nil {
	panic(err)
}
if attest, err := _agent.OnchainGetAttestation(uid); err != nil {
	panic(err)
} else {
	fmt.Println(attest)
}
```

### Revoke attestation onchain

Users can control the attestation they created by revoking it. For onchain attestation, users can declare the revocation of the attestation by the following method. Note that the caller needs to be the attester of the attestation:

```go
schemaUid := "0x85500e806cf1e74844d51a20a6d893fe1ed6f6b0738b50e43d774827d08eca61"
uid := "0x10e25590023c3fcdb0aaa1429712a67399328c53e2cfb8658348658e9f07d694"
var _agent *Agent
var err error
if _agent, err = NewAgentFromKey(privateKey, BNBTESTRPC, BNBTESTCHAINID, GFTESTRPC, GFTESTCHAINID); err != nil {
	panic(err)
}
if attest, err := _agent.OnchainRevoke(schema, uid); err != nil {
	panic(err)
} else {
	fmt.Println(attest)
}
```

## Offchain

### Config greenfield bucket

Before using most offchain-related APIs, you need to configure the bucket corresponding to the account:

```go
var _agent *Agent
var err error
if _agent, err = NewAgentFromKey(privateKey, BNBTESTRPC, BNBTESTCHAINID, GFTESTRPC, GFTESTCHAINID); err != nil {
	panic(err)
}
_agent.ConfigBucket("bas-xxxx")
```

Please note that the name of the bucket is "bas" + hashof(ADDRESS) under the official standard. Users can also generate their own buckets in the bas browser and view them in [dcellar.io](http://dcellar.io/).

### Create an attestation offchain

There are several steps to create an offchain attestation:

1. Generate the offchain attestation.
2. Save locally / upload  to Greenfield.
3. Whether to make an uploaded attestation public on Greenfield.

To create an attestation and get its corresponding JSON file, you can use the following method (note that this operation is only done locally):

```go
var _agent *Agent
var err error
if _agent, err = NewAgentFromKey(privateKey, BNBTESTRPC, BNBTESTCHAINID, GFTESTRPC, GFTESTCHAINID); err != nil {
	panic(err)
}
data := map[string]interface{}{
	"bas":     "bas-go",
	"nonce":  0,
}
schema := "string bas, uint8 nonce"
if res, err := _agent.OffchainNewAttestation(
		schemaUid,
		schema,
		data,
		resolver,
		revocable,
		refUid,
		nonce,
		time,
		expirationTime,
		version,
	); err != nil {
		panic(err)
	} else {
		if _b, err := json.Marshal(res); err != nil {
			panic(err)
		} else {
			fmt.Println(string(_b))
		}

	}
```

After successfully creating an offchain attestation, you can upload it to Greenfield as follows:

```go
if hash, err := _agent.OffchainUploadAttestationToGF(res); err != nil {
		panic(err)
	} else {
		fmt.Println(hash)
	}
```

> Please note that only offchain attestations uploaded to Greenfield will be indexed by the BAS indexer.
> 

If you want to change the public or private status of an offchain attestation uploaded to Greenfield, you can do so as follows:

```go
var _agent *Agent
var err error
if _agent, err = NewAgentFromKey(privateKey, BNBTESTRPC, BNBTESTCHAINID, GFTESTRPC, GFTESTCHAINID); err != nil {
	panic(err)
}
_agent.ConfigBucket("bas-xxxx")
if hash, err := _agent.OffchainChangeAttestationVisible(schemaUid, attestationUid, true); err != nil {
	panic(err)
} else {
	t.Log(hash)
}
```

### Revoke an attestation offchain

Please note that revoking an offchain attestation is actually an onchain operation, which means it also requires relatively expensive gas fees, and the operations involved are similar. Please refer to the OnchainRevokeOffchain function for specific details.