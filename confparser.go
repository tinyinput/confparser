// The package confparser is used to extract element and logical configuration state from network device configuration files.
//
// It was written for a small personal project, but feel free to use or fork as required.
//
// At the time of writing, only Cisco IOS configurations are supported. But the library was written with the intention of
// being extensible at a later date.
//
// # Creating a Client
//
// The eaiest way to use the package is to create a client for the make of networking device configuration you want to parse.
//
// For example:
//
//	package main
//
//	import (
//		"fmt"
//		".../parser"
//	)
//
//	func main() {
//		// Create the Parser Client
//		client, err := parser.NewClient(parser.PARSER_CISCO)
//		if err != nil {
//			// Handle the error
//		}
//		// Check the Parser Client make
//		fmt.Printf("This Parser Client is for \"%v\" Configuration Files\n", client.Make)
//	}
//
// You then have the choice to parse either:
//
// - The whole configuration file
// - A single block of vlan configuration
// - A single block of Interface configuration
//
// # Parsing an Interface Configuration Block
//
// The most common use-case for me was the need to parse Interface configuration blocks and extract the relavent elements.
//
// For example:
//
//	package main
//
//	import (
//		"fmt"
//		".../confparser"
//	)
//
//	const (
//		myInterfaceConfig = `interface GigabitEthernet1/0/1
//	 description STANDARD ACCESS INTERFACE
//	 switchport access vlan 101
//	 switchport mode access
//	 power inline never`
//	)
//
//	func main() {
//		// Create the Conf Parser Client
//		client, err := confparser.NewClient(confparser.PARSER_CISCO)
//		if err != nil {
//			// Handle the error
//		}
//
//		// Create a Config Interface Parser and Set the Interface Configuration:
//		confIntParser := client.NewConfigIntParser()
//		Set(myInterfaceConfig)
//
//		// Evaluate aspects of the Interface Configuration
//		fmt.Printf("Full Name for Interface: %v\n", confIntParser.InterfaceNameFull())
//		fmt.Printf("Short Name for Interface: %v\n", confIntParser.InterfaceNameShort())
//		fmt.Printf("Access Vlan for Interface: %v\n", confIntParser.VlanAccess())
//	}
//
// See the code and testing packages for me details.
package confparser

import (
	"bufio"
	"fmt"
	"regexp"
	"strings"
)

const (
	// Supprted Parser Makes
	PARSER_CISCO = "cisco"
	PARSER_HP    = "hp"
)

// TBD
type ShowCommands interface {
	IsConfigCommand() bool
	IsVersionCommand() bool
}

// ConfParser is an interface which all device configuration parsers should fulfill.
type ConfParser interface {
	Set(s string)
	String() string
	Hostname() string
	InterfaceBlocks() []string
	InterfaceBlocksEthernet() []string
	InterfaceBlocksVlan() []string
	VlanBlocks() []string
}

// VlanParser is an interface which all device vlan configuration parsers should fulfill.
type ConfVlanParser interface {
	Set(s string)
	String() string
	VlanNumber() string
	VlanName() string
}

// ConfIntParser is an interface which all device interface configuration parsers should fulfill.
type ConfIntParser interface {
	Set(s string)
	String() string
	InterfaceNameFull() string
	InterfaceNameShort() string
	Description() string
	VlanAccess() string
	VlanVoice() string
	VlanTrunkNative() string
	VlansTrunkAllowed() string
	//	VlansTrunkAllowedSlice() []string // Not implemented yet
	IpAddressAndMask() string
	IpAddress() string
	IpMask() string
	IsShutdown() bool
	IsSwAccess() bool
	IsSwTrunk() bool
	IsPoeDisabled() bool
	IsEthernet() bool
	IsEthernetFast() bool
	IsEthernetGigabit() bool
	IsEthernetTenGigabit() bool
	IsVlan() bool
}

// TBD
type confParserFunc func(s string) ConfParser

// TBD
type confVlanParserFunc func(s string) ConfVlanParser

// TBD
type confIntParserFunc func(s string) ConfIntParser

// TBD
type ParserClient struct {
	Make             string `json:"make"`
	ConfigParser     confParserFunc
	ConfigVlanParser confVlanParserFunc
	ConfigIntParser  confIntParserFunc
}

// func
func NewClient(make string) (*ParserClient, error) {
	switch make {
	case PARSER_CISCO:
		return ciscoParserClient, nil
	}
	return nil, fmt.Errorf("Unsupported Parser Make (%v)", make)
}

// TBD
func (p *ParserClient) NewConfigParser() ConfParser {
	switch p.Make {
	case PARSER_CISCO:
		var c CiscoConf
		return &c
	}
	return nil
}

// TBD
func (p *ParserClient) NewConfigVlanParser() ConfVlanParser {
	switch p.Make {
	case PARSER_CISCO:
		var c CiscoConfVlan
		return &c
	}
	return nil
}

// TBD
func (p *ParserClient) NewConfigIntParser() ConfIntParser {
	switch p.Make {
	case PARSER_CISCO:
		var c CiscoConfInt
		return &c
	}
	return nil
}

//
//
//

// TBD
func splitStringBlocksByPrefix(s, p string) ([]string, error) {
	//
	scanner := bufio.NewScanner(strings.NewReader(s))
	//
	var block, blocks []string
	//
	for scanner.Scan() {
		//
		line := scanner.Text()
		// If the line matches the Prefix, then this is a new block
		if block == nil && strings.HasPrefix(line, p) {
			block = append(block, line)
			// If the block is not nil (then we're in a block) and the line has an indent
			// The this is the continuation of a block
		} else if block != nil && strings.HasPrefix(line, " ") {
			block = append(block, line)
			// If the block is not nil (then we're in a block) and the line starts with the Prefix
			// Then this is the start of a new block
		} else if block != nil && strings.HasPrefix(line, p) {
			blocks = append(blocks, strings.Join(block, "\n"))
			block = nil
			block = append(block, line)
			// If the block is not nil (then we're in a block) and the line does not have an indent
			// Then this is the end of a block and not a line we want
		} else if block != nil && !strings.HasPrefix(line, " ") {
			blocks = append(blocks, strings.Join(block, "\n"))
			block = nil
		}
	}
	//
	if err := scanner.Err(); err != nil {
		return nil, err
	}
	//
	return blocks, nil
}

// TBD
func splitStringBlocksByRegex(s string, re *regexp.Regexp) ([]string, error) {
	//
	scanner := bufio.NewScanner(strings.NewReader(s))
	//
	var block, blocks []string
	//
	for scanner.Scan() {
		//
		line := scanner.Text()
		// If the line matches the Prefix, then this is a new block
		if block == nil && re.MatchString(line) {
			block = append(block, line)
			// If the block is not nil (then we're in a block) and the line has an indent
			// The this is the continuation of a block
		} else if block != nil && strings.HasPrefix(line, " ") {
			block = append(block, line)
			// If the block is not nil (then we're in a block) and the line starts with the Prefix
			// Then this is the start of a new block
		} else if block != nil && re.MatchString(line) {
			blocks = append(blocks, strings.Join(block, "\n"))
			block = nil
			block = append(block, line)
			// If the block is not nil (then we're in a block) and the line does not have an indent
			// Then this is the end of a block and not a line we want
		} else if block != nil && !strings.HasPrefix(line, " ") {
			blocks = append(blocks, strings.Join(block, "\n"))
			block = nil
		}
	}
	//
	if err := scanner.Err(); err != nil {
		return nil, err
	}
	//
	return blocks, nil
}
