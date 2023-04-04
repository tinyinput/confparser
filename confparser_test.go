package confparser

import (
	"testing"
)

// Test `NewClient()` creationg for Cisco Parser
func TestNewClientUnsupportedMake(t *testing.T) {
	make := ""
	_, err := NewClient(make)
	if err == nil {
		t.Fatalf("Error in `NewClient(%v)`: Missing or invalid make (%[1]v) should return error, but did not\n", make, err)
	}
}

// Test `NewClient()` creation for Cisco Parser
func TestNewClientCisco(t *testing.T) {
	make := PARSER_CISCO
	_, err := NewClient(make)
	if err != nil {
		t.Fatalf("Error in `NewClient(%v)`: %v\n", make, err)
	}
}

// Test `NewConfigParser()` creation for Cisco Parser
func TestNewConfigParserCisco(t *testing.T) {
	make := PARSER_CISCO
	client, _ := NewClient(make)
	confParser := client.NewConfigParser()
	if _, ok := confParser.(*CiscoConf); !ok {
		t.Fatalf("Config Parser created by `NewConfigParser(%v)` is not the correct Type\n", make)
	}
}

// Test `NewConfigParser()` creation for Cisco Parser
func TestNewConfVlanParserCisco(t *testing.T) {
	make := PARSER_CISCO
	client, _ := NewClient(make)
	confVlanParser := client.NewConfigVlanParser()
	if _, ok := confVlanParser.(*CiscoConfVlan); !ok {
		t.Fatalf("Config Vlan Parser created by `NewConfigVlanParser(%v)` is not the correct Type\n", make)
	}
}

// Test `NewConfigParser()` creation for Cisco Parser
func TestNewConfIntParserCisco(t *testing.T) {
	make := PARSER_CISCO
	client, _ := NewClient(make)
	confIntParser := client.NewConfigIntParser()
	if _, ok := confIntParser.(*CiscoConfInt); !ok {
		t.Fatalf("Config Interface Parser created by `NewConfigIntParser(%v)` is not the correct Type\n", make)
	}
}

//
// Utility Functions
//

// isEqualStringSlice compares 2 `[]string`. It is used in some of the tests
func isEqualStringSlice(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}
