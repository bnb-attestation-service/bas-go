package agent

import (
	"fmt"
	"testing"
)

func TestCreateSchema(t *testing.T) {
	schema := "string invite_code_test, uint8 nonce"
	revocable := true
	resolver := ""
	var _agent *Agent
	var err error
	if _agent, err = NewAgentFromKey(privateKey, TESTBAS, TESTSCHEMA, BNBTESTRPC, BNBTESTCHAINID, GFTESTRPC, GFTESTCHAINID); err != nil {
		panic(err)
	}
	if attest, err := _agent.CreateSchema(schema, revocable, resolver); err != nil {
		panic(err)
	} else {
		fmt.Println(attest)
	}

}

func TestGetSchema(t *testing.T) {
	schemaUid := "0xacc308075dabd756f3806f0f2a0d919d12b13597ba4791de96283aa646c2c5b5"

	var _agent *Agent
	var err error
	if _agent, err = NewAgentFromKey(privateKey, TESTBAS, TESTSCHEMA, BNBTESTRPC, BNBTESTCHAINID, GFTESTRPC, GFTESTCHAINID); err != nil {
		panic(err)
	}
	if schema, err := _agent.GetSchema(schemaUid); err != nil {
		panic(err)
	} else {
		fmt.Println(schema)
	}

}

func TestSetSchemaName(t *testing.T) {
	schemaUid := "0xacc308075dabd756f3806f0f2a0d919d12b13597ba4791de96283aa646c2c5b5"
	name := "test_name"
	var _agent *Agent
	var err error
	if _agent, err = NewAgentFromKey(privateKey, TESTBAS, TESTSCHEMA, BNBTESTRPC, BNBTESTCHAINID, GFTESTRPC, GFTESTCHAINID); err != nil {
		panic(err)
	}
	if schema, err := _agent.SetSchemaName(schemaUid, name, TESTSCHEMANAME); err != nil {
		panic(err)
	} else {
		fmt.Println(schema)
	}

}

func TestSetSchemaDescription(t *testing.T) {
	schemaUid := "0xacc308075dabd756f3806f0f2a0d919d12b13597ba4791de96283aa646c2c5b5"
	descrip := "test name for bas go"
	var _agent *Agent
	var err error
	if _agent, err = NewAgentFromKey(privateKey, TESTBAS, TESTSCHEMA, BNBTESTRPC, BNBTESTCHAINID, GFTESTRPC, GFTESTCHAINID); err != nil {
		panic(err)
	}
	if schema, err := _agent.SetSchemaDescription(schemaUid, descrip, TESTSCHEMADESCRIPTION); err != nil {
		panic(err)
	} else {
		fmt.Println(schema)
	}

}
