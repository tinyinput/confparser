/*
The package parser contains functions to extact information from network device configurations

Example Csico Switch configuration would be:

	interface GigabitEthernet1/0/14
	 description [A-001:OB.1.001]
	 switchport access vlan 2
	 switchport mode access
	 switchport voice vlan 917
	 srr-queue bandwidth share 1 30 35 5
	 priority-queue out
	 mls qos trust cos
	 auto qos trust
	 spanning-tree portfast

	interface GigabitEthernet1/0/46
	 description [DP:EXT.J05]-ENV1-RACK26
	 switchport access vlan 137
	 switchport mode access
	 power inline never

	interface GigabitEthernet2/0/11
	 description [26-T-011:EXT.J05]-BUS TURNAROUND:CHAMPS AP:EXT.J05-AP2
	 switchport trunk native vlan 2101
	 switchport trunk allowed vlan 2101,2109,2117,2121,2122
	 switchport mode trunk
	 shutdown

	interface GigabitEthernet1/0/9
	 description [H-014:MU.2.007]
	 switchport access vlan 2110
	 switchport mode access
	 power inline consumption 9000

The package is intended for internal use only.
*/

package confparser

import (
	"regexp"
	"strings"
)

const (
	CommandsCiscoShowConfig     = "show run,show running-config"
	CommandsCiscoShowInterfaces = "show interfaces"
	CommandsCiscoShowVersion    = "show version"
)

const (
	PrefixCiscoConfInterface = "interface"
	PrefixCiscoConfVlan      = "vlan"
)

const (
	RegexCiscoConfMainHostname              = `^\s*hostname (.+)$`
	RegexCiscoConfVlan                      = `^\s*vlan \d+$`
	RegexCiscoConfVlanNumber                = `^\s*vlan (\d+)$`
	RegexCiscoConfVlanName                  = `^\s*name (.+)$`
	RegexCiscoConfInterface                 = `^\s*interface .+$`
	RegexCiscoConfIntNameFull               = `^\s*interface (.+)$`
	RegexCiscoConfIntNameShort              = `^\s*interface (.{2})[^\d]+(.+)$`
	RegexCiscoConfIntDescription            = `^\s*description\s(.+)$`
	RegexCiscoConfIntShutdown               = `^\s*shutdown\s*$`
	RegexCiscoConfIntNoShutdown             = `^\s*no\sshutdown\s*$`
	RegexCiscoConfIntSwitchPortModeAccess   = `^\s*switchport mode access\s*$`
	RegexCiscoConfIntSwitchPortModeTrunk    = `^\s*switchport mode trunk\s*$`
	RegexCiscoConfIntAccessVlan             = `^\s*switchport access vlan (\d+)\s*$`
	RegexCiscoConfIntVoiceVlan              = `^\s*switchport voice vlan (\d+)\s*$`
	RegexCiscoConfIntTrunkNativeVlan        = `^\s*switchport trunk native vlan (\d+)\s*$`
	RegexCiscoConfIntTrunkAllowedVlans      = `^\s*switchport trunk allowed vlan ([\d,]+)\s*$`
	RegexCiscoConfIntPowerInlineNever       = `^\s*power inline never\s*$`
	RegexCiscoConfIntNoIpV4Address          = `^\s*no ip address\s*$`
	RegexCiscoConfIntIpV4AddressAndMask     = `^\s*ip address (\d{1,3}\.\d{1,3}\.\d{1,3}\.\d{1,3}\s\d{1,3}\.\d{1,3}\.\d{1,3}\.\d{1,3})\s*$`
	RegexCiscoConfIntPowerInlineConsumption = `^\s*power inline consumption (\d+)\s*$`
	RegexCiscoConfIntTypeEthernet           = `^\s*interface .+Ethernet\d+`
	RegexCiscoConfIntTypeVlan               = `^\s*interface Vlan\d+`
)

var (
	reCiscoConfMainHostname              = regexp.MustCompile(RegexCiscoConfMainHostname)
	reCiscoConfVlan                      = regexp.MustCompile(RegexCiscoConfVlan)
	reCiscoConfVlanNumber                = regexp.MustCompile(RegexCiscoConfVlanNumber)
	reCiscoConfVlanName                  = regexp.MustCompile(RegexCiscoConfVlanName)
	reCiscoConfInterface                 = regexp.MustCompile(RegexCiscoConfInterface)
	reCiscoConfIntNameFull               = regexp.MustCompile(RegexCiscoConfIntNameFull)
	reCiscoConfIntNameShort              = regexp.MustCompile(RegexCiscoConfIntNameShort)
	reCiscoConfIntDescription            = regexp.MustCompile(RegexCiscoConfIntDescription)
	reCiscoConfIntShutdown               = regexp.MustCompile(RegexCiscoConfIntShutdown)
	reCiscoConfIntNoShutdown             = regexp.MustCompile(RegexCiscoConfIntNoShutdown)
	reCiscoConfIntSwitchPortModeAccess   = regexp.MustCompile(RegexCiscoConfIntSwitchPortModeAccess)
	reCiscoConfIntSwitchPortModeTrunk    = regexp.MustCompile(RegexCiscoConfIntSwitchPortModeTrunk)
	reCiscoConfIntAccessVlan             = regexp.MustCompile(RegexCiscoConfIntAccessVlan)
	reCiscoConfIntVoiceVlan              = regexp.MustCompile(RegexCiscoConfIntVoiceVlan)
	reCiscoConfIntTrunkNativeVlan        = regexp.MustCompile(RegexCiscoConfIntTrunkNativeVlan)
	reCiscoConfIntTrunkAllowedVlans      = regexp.MustCompile(RegexCiscoConfIntTrunkAllowedVlans)
	reCiscoConfIntPowerInlineNever       = regexp.MustCompile(RegexCiscoConfIntPowerInlineNever)
	reCiscoConfIntNoIpV4Address          = regexp.MustCompile(RegexCiscoConfIntNoIpV4Address)
	reCiscoConfIntIpV4AddressAndMask     = regexp.MustCompile(RegexCiscoConfIntIpV4AddressAndMask)
	reCiscoConfIntPowerInlineConsumption = regexp.MustCompile(RegexCiscoConfIntPowerInlineConsumption)
	reCiscoConfIntTypeEthernet           = regexp.MustCompile(RegexCiscoConfIntTypeEthernet)
	reCiscoConfIntTypeVlan               = regexp.MustCompile(RegexCiscoConfIntTypeVlan)
)

var (
	ciscoParserClient = &ParserClient{
		Make: PARSER_CISCO,
		ConfigParser: func(s string) ConfParser {
			p := CiscoConf(s)
			return &p
		},
		ConfigVlanParser: func(s string) ConfVlanParser {
			p := CiscoConfVlan(s)
			return &p
		},
		ConfigIntParser: func(s string) ConfIntParser {
			p := CiscoConfInt(s)
			return &p
		},
	}
)

//
//
//

// CiscoConf is the full configuration for a Cisco device from the `show running-config` command output.
type CiscoConf string

// Set is used to pass a configuration to the CiscoConf type.
func (c *CiscoConf) Set(s string) {
	*c = CiscoConf(s)
}

// String is used to output the CiscoConf type as a string. It fulfills the Stringer interface.
func (c CiscoConf) String() string {
	return string(c)
}

// Hostname returns the hostname from the configuration, as a string.
func (c CiscoConf) Hostname() string {
	return GetCiscoConfMainHostname(string(c))
}

// InterfaceBlocks returns all of the "interface xxx" sections from the configuration, as a slice of string.
// Any errors in the extraction will not be reported and a nil slice will be returned.
func (c CiscoConf) InterfaceBlocks() []string {
	return GetInterfaceBlocks(string(c))
}

// InterfaceBlocksEthernet returns all of the "interface xxx" sections from the configuration,
// where the interface is an "Ethernet" type, as a slice of string.
// Any errors in the extraction will not be reported and a nil slice will be returned.
func (c CiscoConf) InterfaceBlocksEthernet() []string {
	return GetInterfaceBlocksEthernet(string(c))
}

// InterfaceBlocksEthernet returns all of the "interface xxx" sections from the configuration,
// where the interface is a "Vlan" type, as a slice of string.
// Any errors in the extraction will not be reported and a nil slice will be returned.
func (c CiscoConf) InterfaceBlocksVlan() []string {
	return GetInterfaceBlocksVlan(string(c))
}

// InterfaceBlocks returns all of the "interface xxx" sections from the configuration, as a slice of string.
// Any errors in the extraction will not be reported and a nil slice will be returned.
func (c CiscoConf) VlanBlocks() []string {
	return GetVlanBlocks(string(c))
}

// CiscoConfVlan is a single vlan configuration block for a Cisco device from the `show running-config` command output.
// For example:
//
//	vlan 101
//	  name Internet
//
// A block can be extracted from the main configuration using the `CiscoConf.VlanBlocks()` method.
type CiscoConfVlan string

// Set is used to pass a vlan configuration block to the CiscoConfVlan type.
func (c *CiscoConfVlan) Set(s string) {
	*c = CiscoConfVlan(s)
}

// String is used to output the CiscoConfVlan type as a string. It fulfills the Stringer interface.
func (c CiscoConfVlan) String() string {
	return string(c)
}

// VlanNumber returns the vlan number from the vlan configuration block, as a string.
func (c CiscoConfVlan) VlanNumber() string {
	return GetCiscoConfVlanNumber(string(c))
}

// VlanName returns the vlan name from the vlan configuration block, as a string.
func (c CiscoConfVlan) VlanName() string {
	return GetCiscoConfVlanName(string(c))
}

// CiscoConfInt is the interface configuration block for a Cisco device from the `show running-config` command output.
// For example:
//
//	interface GigabitEthernet1/0/22
//	  description [A-001]
//	  switchport access vlan 101
//	  switchport mode access
//	  power inline never
//
// A block can be extracted from the main configuration using the `CiscoConf.IntrfacBlocks()` method.
type CiscoConfInt string

// Set is used to pass an interfacee configuration block to the CiscoConfInt type.
func (c *CiscoConfInt) Set(s string) {
	*c = CiscoConfInt(s)
}

// String is used to output the CiscoConfInt type as a string. It fulfills the Stringer interface.
func (c CiscoConfInt) String() string {
	return string(c)
}

// InterfaceNameFull returns the full intrface name from the interface configuration block, as a string.
func (c CiscoConfInt) InterfaceNameFull() string {
	return GetCiscoConfIntNameFull(string(c))
}

// InterfaceNameShort returns the shortened interface name from the interface configuration block, as a string.
// For example, `GigabitEthernet1/0/22`, would be shortened to, `Gi1/0/22`.
func (c CiscoConfInt) InterfaceNameShort() string {
	return GetCiscoConfIntNameShort(string(c))
}

// Description returns the `description xxx` from the interface configuration block, as a string.
func (c CiscoConfInt) Description() string {
	return GetCiscoConfIntDescription(string(c))
}

// VlanAccess returns the `switchport access vlan xxx` number from the interface configuration block, as a string.
func (c CiscoConfInt) VlanAccess() string {
	return GetCiscoConfIntAccessVlan(string(c))
}

// VlanVoice returns the `switchport voice vlan xxx` number from the interface configuration block, as a string.
func (c CiscoConfInt) VlanVoice() string {
	return GetCiscoConfIntVoiceVlan(string(c))
}

// VlanTrunkNative returns the `switchport trunk native vlan xxx` number from the interface configuration block, as a string.
func (c CiscoConfInt) VlanTrunkNative() string {
	return GetCiscoConfIntTrunkNativeVlan(string(c))
}

// VlansTrunkAllowed returns the `switchport trunk allowed vlan xxx` list from the interface configuration block,
// as a comma separated string. For example, `switchport trunk allowed vlan 100,200,300`, would return the string, `100,200,300`.
func (c CiscoConfInt) VlansTrunkAllowed() string {
	return GetCiscoConfIntTrunkAllowedVlans(string(c))
}

// VlansTrunkAllowedSlice returns the `switchport trunk allowed vlan xxx` list from the interface configuration block,
// as a slice of strings.
func (c CiscoConfInt) VlansTrunkAllowedSlice() []string {
	return GetCiscoConfIntTrunkAllowedVlansSlice(string(c))
}

// IpV4AddressAndMask - To Be Documented
func (c CiscoConfInt) IpV4AddressAndMask() string {
	return GetCiscoConfIntIpV4AddressAndMask(string(c))
}

// IpAddressAndMask - To Be Documented
func (c CiscoConfInt) IpAddressAndMask() string {
	return GetCiscoConfIntIpV4AddressAndMask(string(c))
}

// IpV4Address - To Be Documented
func (c CiscoConfInt) IpV4Address() string {
	return GetCiscoConfIntIpV4Address(string(c))
}

// IpAddress - To Be Documented
func (c CiscoConfInt) IpAddress() string {
	return GetCiscoConfIntIpV4Address(string(c))
}

// IpV4Mask - To Be Documented
func (c CiscoConfInt) IpV4Mask() string {
	return GetCiscoConfIntIpV4Mask(string(c))
}

// IpMask - To Be Documented
func (c CiscoConfInt) IpMask() string {
	return GetCiscoConfIntIpV4Mask(string(c))
}

// PoeConsumption - To Be Documented
func (c CiscoConfInt) PoeConsumption() string {
	return GetCiscoConfIntPowerInlineConsumption(string(c))
}

// IsShutdown returns the configured interface shutdown state from the interface configuration block, as a bool.
// If the interface is shutdown/disabled, then IsShutdown will return true.
func (c CiscoConfInt) IsShutdown() bool {
	return IsCiscoConfIntShutdown(string(c))
}

// IsSwAccess returns the configured switchport access state from the interface configuration block, as a bool.
// If the interface as the config, `switchport mode access`, then IsSwAccess will return true.
func (c CiscoConfInt) IsSwAccess() bool {
	return IsCiscoConfIntSwitchPortModeAccess(string(c))
}

// IsSwTrunk returns the configured switchport trunk state from the interface configuration block, as a bool.
// If the interface as the config, `switchport mode trunk`, then IsSwTrunk will return true.
func (c CiscoConfInt) IsSwTrunk() bool {
	return IsCiscoConfIntSwitchPortModeTrunk(string(c))
}

// IsPoeDisabled tests to see if PoE is disabled on the interface. By default, PoE is enabled in auto mode,
// but there would be no configuration line for this.
//
// So it's better to test for the negative (I.e. Is PoE disabled?), as this is an explicit configuration entry.
func (c CiscoConfInt) IsPoeDisabled() bool {
	return IsCiscoConfIntPowerInlineNever(string(c))
}

// IsPoeStatic - To Be Documented
func (c CiscoConfInt) IsPoeStatic() bool {
	return IsCiscoConfIntPowerInlineConsumption(string(c))
}

// IsIpV4Int - To Be Documented
func (c CiscoConfInt) IsIpV4Int() bool {
	return IsCiscoConfIntIpV4Int(string(c))
}

// IsIpInt - To Be Documented
func (c CiscoConfInt) IsIpInt() bool {
	return IsCiscoConfIntIpV4Int(string(c))
}

// Below are the functions which support the methods above.
//
// If you simply want to call a function against a string, without all of the parser type elements, then these can be used.
// It should be noted that they haven't been tested outside of their implementation in the methods above.
// Therefore use at your own risk.

func GetCiscoConfMainHostname(s string) string {
	return regexExtractFirstMatchSubString(reCiscoConfMainHostname, s)
}

func GetCiscoConfVlanNumber(s string) string {
	return regexExtractFirstMatchSubString(reCiscoConfVlanNumber, s)
}

func GetCiscoConfVlanName(s string) string {
	return regexExtractFirstMatchSubString(reCiscoConfVlanName, s)
}

func GetCiscoConfIntNameFull(s string) string {
	return regexExtractFirstMatchSubString(reCiscoConfIntNameFull, s)
}

func GetCiscoConfIntNameShort(s string) string {
	parts := regexExtractMultipleSubString(reCiscoConfIntNameShort, s, 2)
	return strings.Join(parts, "")
}

func GetCiscoConfIntDescription(s string) string {
	return regexExtractFirstMatchSubString(reCiscoConfIntDescription, s)
}

func GetCiscoConfIntAccessVlan(s string) string {
	return regexExtractFirstMatchSubString(reCiscoConfIntAccessVlan, s)
}

func GetCiscoConfIntVoiceVlan(s string) string {
	return regexExtractFirstMatchSubString(reCiscoConfIntVoiceVlan, s)
}

func GetCiscoConfIntTrunkNativeVlan(s string) string {
	return regexExtractFirstMatchSubString(reCiscoConfIntTrunkNativeVlan, s)
}

func GetCiscoConfIntTrunkAllowedVlans(s string) string {
	return regexExtractFirstMatchSubString(reCiscoConfIntTrunkAllowedVlans, s)
}

func GetCiscoConfIntTrunkAllowedVlansSlice(s string) []string {
	vlans := regexExtractFirstMatchSubString(reCiscoConfIntTrunkAllowedVlans, s)
	return strings.Split(vlans, ",")
}

func GetCiscoConfIntIpV4AddressAndMask(s string) string {
	return regexExtractFirstMatchSubString(reCiscoConfIntIpV4AddressAndMask, s)
}

func GetCiscoConfIntIpV4Address(s string) string {
	full := GetCiscoConfIntIpV4AddressAndMask(s)
	return splitOnSpaceReturnFirstString(full)
}

func GetCiscoConfIntIpV4Mask(s string) string {
	full := GetCiscoConfIntIpV4AddressAndMask(s)
	return splitOnSpaceReturnSecondString(full)
}

func GetCiscoConfIntPowerInlineConsumption(s string) string {
	return regexExtractFirstMatchSubString(reCiscoConfIntPowerInlineConsumption, s)
}

func IsCiscoConfIntShutdown(s string) bool {
	return regexFirstMatchFalseString(reCiscoConfIntNoShutdown, reCiscoConfIntShutdown, s)
}

func IsCiscoConfIntSwitchPortModeAccess(s string) bool {
	return regexMatchString(reCiscoConfIntSwitchPortModeAccess, s)
}

func IsCiscoConfIntSwitchPortModeTrunk(s string) bool {
	return regexMatchString(reCiscoConfIntSwitchPortModeTrunk, s)
}

// IsCiscoConfIntPowerInlineNever will return true if the interface configuration contains "power inline never".
// NOTE: This indicates that PoE is disabled. I.e. `true` indicates NO PoE on the interface.
func IsCiscoConfIntPowerInlineNever(s string) bool {
	return regexMatchString(reCiscoConfIntPowerInlineNever, s)
}

func IsCiscoConfIntPowerInlineConsumption(s string) bool {
	return regexMatchString(reCiscoConfIntPowerInlineConsumption, s)
}

func IsCiscoConfIntIpV4Int(s string) bool {
	return regexMatchString(reCiscoConfIntIpV4AddressAndMask, s)
}

func GetInterfaceBlocks(s string) []string {
	b, err := splitStringBlocksByRegex(s, reCiscoConfInterface)
	if err != nil {
		return nil
	}
	return b
}

func GetInterfaceBlocksEthernet(s string) []string {
	return GetInterfaceBlocksByType(s, "ethernet")
}

func GetInterfaceBlocksVlan(s string) []string {
	return GetInterfaceBlocksByType(s, "vlan")
}

func GetInterfaceBlocksByType(s, t string) []string {
	var filteredBlocks []string
	var re *regexp.Regexp
	switch strings.ToLower(t) {
	case "ethernet", "eth":
		re = reCiscoConfIntTypeEthernet
	case "vlan", "vlans", "vl":
		re = reCiscoConfIntTypeVlan
	default:
		return nil
	}
	blocks := GetInterfaceBlocks(s)
	for i, _ := range blocks {
		if re.MatchString(blocks[i]) {
			filteredBlocks = append(filteredBlocks, blocks[i])
		}
	}
	return filteredBlocks
}

func GetVlanBlocks(s string) []string {
	b, err := splitStringBlocksByRegex(s, reCiscoConfVlan)
	if err != nil {
		return nil
	}
	return b
}

// The functions below are unexported utility functions for the Cisco parser.
//
// They are simple regex and string functions, which are needed by almost all of the other functions above.

// regexExtractFirstMatchSubString is used to return the 1st capture group match from a regex.
//
// If used with a multi-line configuration, or configuration block, this first gets split into line.
// The function will then return on the first line where the regex matches.
func regexExtractFirstMatchSubString(re *regexp.Regexp, s string) string {
	lines := strings.Split(s, "\n")
	for _, line := range lines {
		matches := re.FindStringSubmatch(line)
		if len(matches) >= 2 {
			return matches[1]
		}
	}
	return ""
}

// regexExtractMultipleSubString is used to return the multiple capture group matches from a regex.
// Where `n` is the number of capture groups to return.
//
// If used with a multi-line configuration, or configuration block, this first gets split into line.
// The function will then return on the first line where the regex matches.
func regexExtractMultipleSubString(re *regexp.Regexp, s string, n int) []string {
	lines := strings.Split(s, "\n")
	for _, line := range lines {
		matches := re.FindStringSubmatch(line)
		if len(matches) >= n+1 {
			return matches[1 : n+1]
		}
	}
	return nil
}

// regexMatchString is used as a wrapper to `regexp.MatchString`, to deal with multi-line strings.
func regexMatchString(re *regexp.Regexp, s string) bool {
	lines := strings.Split(s, "\n")
	for _, line := range lines {
		if re.MatchString(line) {
			return true
		}
	}
	return false
}

// regexFirstMatchFalseString is used to check against a negating match first.
// An example would be if we wanted to look for "shutdown" on the interface. If we see "no shutdown" then that should return false.
//
// You would call the function by providing the negative regex first, then the one to match on. For example:
//
//	regexFirstMatchFalseString(reCiscoConfIntNoShutdown, reCiscoConfIntShutdown, "interface block")
//
// This is needed because the RE2 engine in the regexp package does not support backtracking, so no option to do a lookbehind.
func regexFirstMatchFalseString(reFalse, reTrue *regexp.Regexp, s string) bool {
	lines := strings.Split(s, "\n")
	for _, line := range lines {
		if reFalse.MatchString(line) {
			return false
		} else {
			if reTrue.MatchString(line) {
				return true
			}
		}
	}
	return false
}

// splitOnSpaceReturnFirstString - To Be Documented
func splitOnSpaceReturnFirstString(s string) string {
	if s == "" {
		return ""
	}
	split := strings.Split(s, " ")
	if len(split) < 2 {
		return ""
	}
	return split[0]
}

// splitOnSpaceReturnSecondString - To Be Documented
func splitOnSpaceReturnSecondString(s string) string {
	if s == "" {
		return ""
	}
	split := strings.Split(s, " ")
	if len(split) < 2 {
		return ""
	}
	return split[1]
}
