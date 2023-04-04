package confparser

import (
	"fmt"
	"strings"
	"testing"
)

//
// type CiscoConf Tests
//

// TBD
func TestCiscoConf(t *testing.T) {
	client, _ := NewClient(PARSER_CISCO)
	//
	if client.ConfigParser(testCiscoConf).String() != testCiscoConf {
		t.Fatalf("Error with type PaserClient.ConfigParser() func field. Failed to return correct string\n")
	}
}

// TBD
func TestCiscoConfSet(t *testing.T) {
	client, _ := NewClient(PARSER_CISCO)
	confParser := client.NewConfigParser()
	confParser.Set(testCiscoConf)
	//
	if fmt.Sprintf("%s", confParser) != testCiscoConf {
		t.Fatalf("Error with method CiscoConf.Set(). Failed to set string correctly\n")
	}
}

// TBD
func TestCiscoConfHostname(t *testing.T) {
	client, _ := NewClient(PARSER_CISCO)
	//
	result := client.ConfigParser(testCiscoConf).Hostname()
	if result != testCiscoConfResultHostname {
		t.Fatalf("Error with method CiscoConf.Hostname(). Failed to extract Hostname correctly. Expecting (%v). Got (%v)\n", testCiscoConfResultHostname, result)
	}
}

// TBD
func TestCiscoConfInterfaceBlocks(t *testing.T) {
	client, _ := NewClient(PARSER_CISCO)
	//
	resultBlocks := client.ConfigParser(testCiscoConf).InterfaceBlocks()
	requiredBlocks := strings.Split(testCiscoConfResultInterfaceBlocksList, "\n!\n")
	//
	if len(resultBlocks) != len(requiredBlocks) {
		t.Fatalf("Error with method CiscoConf.InterfaceBlocks(). Miss-match in number of interfaces extracted. Expecting (%v). Got (%v)\n", len(requiredBlocks), len(resultBlocks))
	}
	//
	for i, _ := range resultBlocks {
		if resultBlocks[i] != requiredBlocks[i] {
			t.Fatalf("Error with method CiscoConf.InterfaceBlocks(). Interface Block Miss-match\nExpecting (%v)\nGot (%v)\n", requiredBlocks[i], resultBlocks[i])
		}
	}
}

// TBD
func TestCiscoConfInterfaceBlocksEthernet(t *testing.T) {
	client, _ := NewClient(PARSER_CISCO)
	//
	resultBlocks := client.ConfigParser(testCiscoConf).InterfaceBlocksEthernet()
	requiredBlocks := strings.Split(testCiscoConfResultInterfaceEthernetBlocksList, "\n!\n")
	//
	if len(resultBlocks) != len(requiredBlocks) {
		t.Fatalf("Error with method CiscoConf.InterfaceBlocksEthernet(). Miss-match in number of interfaces extracted. Expecting (%v). Got (%v)\n", len(requiredBlocks), len(resultBlocks))
	}
	//
	for i, _ := range resultBlocks {
		if resultBlocks[i] != requiredBlocks[i] {
			t.Fatalf("Error with method CiscoConf.InterfaceBlocksEthernet(). Interface Block Miss-match\nExpecting (%v)\nGot (%v)\n", requiredBlocks[i], resultBlocks[i])
		}
	}
}

// TBD
func TestCiscoConfInterfaceBlocksVlan(t *testing.T) {
	client, _ := NewClient(PARSER_CISCO)
	//
	resultBlocks := client.ConfigParser(testCiscoConf).InterfaceBlocksVlan()
	requiredBlocks := strings.Split(testCiscoConfResultInterfaceVlanBlocksList, "\n!\n")
	//
	if len(resultBlocks) != len(requiredBlocks) {
		t.Fatalf("Error with method CiscoConf.InterfaceBlocksVlan(). Miss-match in number of interfaces extracted. Expecting (%v). Got (%v)\n", len(requiredBlocks), len(resultBlocks))
	}
	//
	for i, _ := range resultBlocks {
		if resultBlocks[i] != requiredBlocks[i] {
			t.Fatalf("Error with method CiscoConf.InterfaceBlocksVlan(). Interface Block Miss-match\nExpecting (%v)\nGot (%v)\n", requiredBlocks[i], resultBlocks[i])
		}
	}
}

// TBD
func TestCiscoConfVlanBlocks(t *testing.T) {
	client, _ := NewClient(PARSER_CISCO)
	//
	resultBlocks := client.ConfigParser(testCiscoConf).VlanBlocks()
	requiredBlocks := strings.Split(testCiscoConfResultVlanBlocksList, "\n!\n")
	//
	if len(resultBlocks) != len(requiredBlocks) {
		t.Fatalf("Error with method CiscoConf.VlanBlocks(). Miss-match in number of vlans extracted. Expecting (%v). Got (%v)\n", len(requiredBlocks), len(resultBlocks))
	}
	//
	for i, _ := range resultBlocks {
		if resultBlocks[i] != requiredBlocks[i] {
			t.Fatalf("Error with method CiscoConf.VlanBlocks(). Interface Block Miss-match\nExpecting (%v)\nGot (%v)\n", requiredBlocks[i], resultBlocks[i])
		}
	}
}

//
// type CiscoConfVlan Tests
//

// TBD
func TestCiscoConfVlan(t *testing.T) {
	client, _ := NewClient(PARSER_CISCO)
	//
	if client.ConfigVlanParser(testCiscoConfVlan).String() != testCiscoConfVlan {
		t.Fatalf("Error with type PaserClient.ConfigVlanParser() func field. Failed to return correct string\n")
	}
}

// TBD
func TestCiscoConfVlanSet(t *testing.T) {
	client, _ := NewClient(PARSER_CISCO)
	confVlanParser := client.NewConfigVlanParser()
	confVlanParser.Set(testCiscoConfVlan)
	//
	if fmt.Sprintf("%s", confVlanParser) != testCiscoConfVlan {
		t.Fatalf("Error with method CiscoConfVlan.Set(). Failed to set string correctly\n")
	}
}

// TBD
func TestCiscoConfVlanNumber(t *testing.T) {
	client, _ := NewClient(PARSER_CISCO)
	//
	result := client.ConfigVlanParser(testCiscoConfVlan).VlanNumber()
	if result != testCiscoConfVlanNumberResult {
		t.Fatalf("Error with method CiscoConfVlan.Number(). Failed to extract Vlan Number correctly. Expecting (%v). Got (%v)\n", testCiscoConfVlanNumberResult, result)
	}
}

// TBD
func TestCiscoConfVlanName(t *testing.T) {
	client, _ := NewClient(PARSER_CISCO)
	//
	result := client.ConfigVlanParser(testCiscoConfVlan).VlanName()
	if result != testCiscoConfVlanNameResult {
		t.Fatalf("Error with method CiscoConfVlan.Name(). Failed to extract Vlan Name correctly. Expecting (%v). Got (%v)\n", testCiscoConfVlanNameResult, result)
	}
}

//
// type CiscoConfInt Tests
//

// TBD
func TestCiscoConfInt(t *testing.T) {
	client, _ := NewClient(PARSER_CISCO)
	//
	if client.ConfigIntParser(testCiscoIntAccessNoPoe).String() != testCiscoIntAccessNoPoe {
		t.Fatalf("Error with type PaserClient.ConfigIntParser() func field. Failed to return correct string\n")
	}
}

// TBD
func TestCiscoConfIntSet(t *testing.T) {
	client, _ := NewClient(PARSER_CISCO)
	confIntParser := client.NewConfigIntParser()
	confIntParser.Set(testCiscoIntAccessNoPoe)
	//
	if confIntParser.String() != testCiscoIntAccessNoPoe {
		t.Fatalf("Error with method CiscoConfInt.Set(). Failed to set string correctly\n")
	}
}

// TBD
func TestCiscoConfIntInterfaceNameFull(t *testing.T) {
	client, _ := NewClient(PARSER_CISCO)
	//
	result := client.ConfigIntParser(testCiscoIntAccessNoPoe).InterfaceNameFull()
	if result != testCiscoIntAccessNoPoeResultInterfaceNameFull {
		t.Fatalf("Error with method CiscoConfInt.InterfaceNameFull(). Failed to extract Interface Name Full correctly. Expecting (%v). Got (%v)\n", testCiscoIntAccessNoPoeResultInterfaceNameFull, result)
	}
}

// TBD
func TestCiscoConfIntInterfaceNameShort(t *testing.T) {
	client, _ := NewClient(PARSER_CISCO)
	//
	result := client.ConfigIntParser(testCiscoIntAccessNoPoe).InterfaceNameShort()
	if result != testCiscoIntAccessNoPoeResultInterfaceNameShort {
		t.Fatalf("Error with method CiscoConfInt.InterfaceNameShort(). Failed to extract Interface Name Short correctly. Expecting (%v). Got (%v)\n", testCiscoIntAccessNoPoeResultInterfaceNameShort, result)
	}
}

// TBD
func TestCiscoConfIntDescription(t *testing.T) {
	client, _ := NewClient(PARSER_CISCO)
	//
	result := client.ConfigIntParser(testCiscoIntAccessNoPoe).Description()
	if result != testCiscoIntAccessNoPoeResultDescription {
		t.Fatalf("Error with method CiscoConfInt.Description(). Failed to extract Interface Description correctly. Expecting (%v). Got (%v)\n", testCiscoIntAccessNoPoeResultDescription, result)
	}
}

// TBD
func TestCiscoConfIntVlanAccess(t *testing.T) {
	client, _ := NewClient(PARSER_CISCO)
	//
	result := client.ConfigIntParser(testCiscoIntAccessNoPoe).VlanAccess()
	if result != testCiscoIntAccessNoPoeResultAccessVlan {
		t.Fatalf("Error with method CiscoConfInt.VlanAccess(). Failed to extract Interface Access Vlan correctly. Expecting (%v). Got (%v)\n", testCiscoIntAccessNoPoeResultAccessVlan, result)
	}
}

// TBD
func TestCiscoConfIntVlanVoice(t *testing.T) {
	client, _ := NewClient(PARSER_CISCO)
	//
	result := client.ConfigIntParser(testCiscoIntAccessVoiceVlan).VlanVoice()
	if result != testCiscoIntAccessVoiceVlanResultVoiceVlan {
		t.Fatalf("Error with method CiscoConfInt.VlanVoice(). Failed to extract Interface Voice Vlan correctly. Expecting (%v). Got (%v)\n", testCiscoIntAccessVoiceVlanResultVoiceVlan, result)
	}
}

// TBD
func TestCiscoConfIntVlanTrunkNative(t *testing.T) {
	client, _ := NewClient(PARSER_CISCO)
	//
	result := client.ConfigIntParser(testCiscoIntTrunkShutdown).VlanTrunkNative()
	if result != testCiscoIntTrunkShutdownResultTrunkNativeVlan {
		t.Fatalf("Error with method CiscoConfInt.VlanTrunkNative(). Failed to extract Interface Trunk Native Vlan correctly. Expecting (%v). Got (%v)\n", testCiscoIntTrunkShutdownResultTrunkNativeVlan, result)
	}
}

// TBD
func TestCiscoConfIntVlansTrunkAllowed(t *testing.T) {
	client, _ := NewClient(PARSER_CISCO)
	//
	result := client.ConfigIntParser(testCiscoIntTrunkShutdown).VlansTrunkAllowed()
	if result != testCiscoIntTrunkShutdownResultTrunkAllowedVlans {
		t.Fatalf("Error with method CiscoConfInt.VlansTrunkAllowed(). Failed to extract Interface Trunk Allowed Vlans correctly. Expecting (%v). Got (%v)\n", testCiscoIntTrunkShutdownResultTrunkAllowedVlans, result)
	}
}

// TestCiscoConfIntVlansTrunkAllowedSlice is deliberately commented out.
// Method VlansTrunkAllowedSlice() is not part of the main parser.ConfIntParser interface type.
func TestCiscoConfIntVlansTrunkAllowedSlice(t *testing.T) {
	/*
	   client, _ := NewClient(PARSER_CISCO)
	   confIntParser := client.NewConfigIntParser()
	   confIntParser.Set(testCiscoIntTrunkShutdown)
	   testDataSlice := strings.Split(testCiscoIntTrunkShutdownResultTrunkAllowedVlans,",")
	   //
	   result := confIntParser.VlansTrunkAllowedSlice()

	   	if !isEqualStringSlice(result, testDataSlice) {
	   		t.Fatalf("Error with method CiscoConfInt.VlansTrunkAllowedSlice(). Failed to extract Interface Trunk Allowed Vlans `[]string` correctly. Expecting (%v). Got (%v)\n", testCiscoIntTrunkShutdownResultTrunkAllowedVlans, result)
	   	}
	*/
}

// TBD
func TestCiscoConfIntIsShutdownFalse(t *testing.T) {
	client, _ := NewClient(PARSER_CISCO)
	confIntParser := client.NewConfigIntParser()
	confIntParser.Set(testCiscoIntAccessNoPoe)
	//
	result := client.ConfigIntParser(testCiscoIntAccessNoPoe).IsShutdown()
	if result {
		t.Fatalf("Error with method CiscoConfInt.IsShutdown(). Shutdown status evaluated incorrectly. Expecting (%v). Got (%v)\n", testCiscoIntAccessNoPoeResultIsShutdown, result)
	}
}

// TBD
func TestCiscoConfIntIsShutdownTrue(t *testing.T) {
	client, _ := NewClient(PARSER_CISCO)
	//
	result := client.ConfigIntParser(testCiscoIntTrunkShutdown).IsShutdown()
	if !result {
		t.Fatalf("Error with method CiscoConfInt.IsShutdown(). Shutdown status evaluated incorrectly. Expecting (%v). Got (%v)\n", testCiscoIntTrunkShutdownResultIsShutdown, result)
	}
}

// TBD
func TestCiscoConfIntIsPoeDisabledTrue(t *testing.T) {
	client, _ := NewClient(PARSER_CISCO)
	confIntParser := client.NewConfigIntParser()
	confIntParser.Set(testCiscoIntAccessNoPoe)
	//
	result := client.ConfigIntParser(testCiscoIntAccessNoPoe).IsPoeDisabled()
	if !result {
		t.Fatalf("Error with method CiscoConfInt.IsPoeDisabled(). PoE status evaluated incorrectly. Expecting (%v). Got (%v)\n", testCiscoIntAccessNoPoeResultIsPoeDisabled, result)
	}
}
