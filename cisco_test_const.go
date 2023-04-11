package confparser

const (
	testCiscoConfVlan = `vlan 101
 name INTERNET`
	testCiscoConfVlanNumberResult = `101`
	testCiscoConfVlanNameResult   = `INTERNET`
)

const (
	testCiscoIntAccessNoPoe = `interface GigabitEthernet1/0/1
 description STANDARD ACCESS INTERFACE
 switchport access vlan 101
 switchport mode access
 power inline never`
	testCiscoIntAccessNoPoeResultInterfaceNameFull    = `GigabitEthernet1/0/1`
	testCiscoIntAccessNoPoeResultInterfaceNameShort   = `Gi1/0/1`
	testCiscoIntAccessNoPoeResultDescription          = `STANDARD ACCESS INTERFACE`
	testCiscoIntAccessNoPoeResultAccessVlan           = `101`
	testCiscoIntAccessNoPoeResultVoiceVlan            = ``
	testCiscoIntAccessNoPoeResultTrunkNativeVlan      = ``
	testCiscoIntAccessNoPoeResultTrunkAllowedVlans    = ``
	testCiscoIntAccessNoPoeResultIsShutdown           = false
    testCiscoIntAccessNoPoeResultIsSwAccess           = true
    testCiscoIntAccessNoPoeResultIsSwTrunk            = false
	testCiscoIntAccessNoPoeResultIsPoeDisabled        = true
    testCiscoIntAccessNoPoeResultIsEthernet           = true
    testCiscoIntAccessNoPoeResultIsEthernetFast       = false
    testCiscoIntAccessNoPoeResultIsEthernetGigabit    = true
    testCiscoIntAccessNoPoeResultIsEthernetTenGigabit = false
    testCiscoIntAccessNoPoeResultIsVlan               = false
)

const (
	testCiscoIntAccessVoiceVlan = `interface GigabitEthernet5/0/48
 description VOICE ENABLED INTERFACE
 switchport access vlan 2
 switchport mode access
 switchport voice vlan 901
 srr-queue bandwidth share 1 30 35 5
 priority-queue out
 mls qos trust cos
 auto qos trust
 spanning-tree portfast`
	testCiscoIntAccessVoiceVlanResultInterfaceNameFull  = `GigabitEthernet5/0/48`
	testCiscoIntAccessVoiceVlanResultInterfaceNameShort = `Gi5/0/48`
	testCiscoIntAccessVoiceVlanResultDescription        = `VOICE ENABLED INTERFACE`
	testCiscoIntAccessVoiceVlanResultAccessVlan         = `2`
	testCiscoIntAccessVoiceVlanResultVoiceVlan          = `901`
	testCiscoIntAccessVoiceVlanResultTrunkNativeVlan    = ``
	testCiscoIntAccessVoiceVlanResultTrunkAllowedVlans  = ``
	testCiscoIntAccessVoiceVlanResultIsShutdown         = false
	testCiscoIntAccessVoiceVlanIsPoeDisabled            = false
)

const (
	testCiscoIntTrunkShutdown = `interface GigabitEthernet2/0/23
 description STANDARD TRUNK INTERFACE
 switchport trunk native vlan 1101
 switchport trunk allowed vlan 1101,1109,1117,1121,1122
 switchport mode trunk
 shutdown`
	testCiscoIntTrunkShutdownResultInterfaceNameFull  = `GigabitEthernet2/0/23`
	testCiscoIntTrunkShutdownResultInterfaceNameShort = `Gi2/0/23`
	testCiscoIntTrunkShutdownResultDescription        = `STANDARD TRUNK INTERFACE`
	testCiscoIntTrunkShutdownResultAccessVlan         = ``
	testCiscoIntTrunkShutdownResultVoiceVlan          = ``
	testCiscoIntTrunkShutdownResultTrunkNativeVlan    = `1101`
	testCiscoIntTrunkShutdownResultTrunkAllowedVlans  = `1101,1109,1117,1121,1122`
	testCiscoIntTrunkShutdownResultIsShutdown         = true
    testCiscoIntTrunkShutdownResultIsSwAccess         = false
    testCiscoIntTrunkShutdownResultIsSwTrunk          = true
	testCiscoIntTrunkShutdownResultIsPoeDisabled      = false
)

const (
    testCiscoIntVlan = `interface Vlan1
 no ip address
 shutdown`
    testCiscoIntVlanResultIsShutdown           = true
    testCiscoIntVlanResultIsEthernet           = false
    testCiscoIntVlanResultIsEthernetFast       = false
    testCiscoIntVlanResultIsEthernetGigabit    = false
    testCiscoIntVlanResultIsEthernetTenGigabit = false
    testCiscoIntVlanResultIsVlan               = true
)

const (
    testCiscoIntEthernetFast = `interface FastEthernet0
 no ip address
 shutdown`
    testCiscoIntEthernetFastResultIsShutdown           = true
    testCiscoIntEthernetFastResultIsEthernet           = true
    testCiscoIntEthernetFastResultIsEthernetFast       = true
    testCiscoIntEthernetFastResultIsEthernetGigabit    = false
    testCiscoIntEthernetFastResultIsEthernetTenGigabit = false
    testCiscoIntEthernetFastResultIsVlan               = false
)

const (
    testCiscoIntEthernetGigabit = `interface GigabitEthernet1/0/17
 description [A-017:OF.1.001]
 shutdown`
    testCiscoIntEthernetGigabitResultIsShutdown           = true
    testCiscoIntEthernetGigabitResultIsEthernet           = true
    testCiscoIntEthernetGigabitResultIsEthernetFast       = false
    testCiscoIntEthernetGigabitResultIsEthernetGigabit    = true
    testCiscoIntEthernetGigabitResultIsEthernetTenGigabit = false
    testCiscoIntEthernetGigabitResultIsVlan               = false
)

const (
    testCiscoIntEthernetTenGigabit = `interface TenGigabitEthernet1/0/1
 description [DP:OF.2.999]-CORE-SW01
 switchport trunk native vlan 10
 switchport mode trunk
 srr-queue bandwidth share 1 30 35 5
 priority-queue out 
 mls qos trust cos
 auto qos trust 
 ip dhcp snooping trust`
    testCiscoIntEthernetTenGigabitResultIsShutdown           = true
    testCiscoIntEthernetTenGigabitResultIsEthernet           = true
    testCiscoIntEthernetTenGigabitResultIsEthernetFast       = false
    testCiscoIntEthernetTenGigabitResultIsEthernetGigabit    = false
    testCiscoIntEthernetTenGigabitResultIsEthernetTenGigabit = true
    testCiscoIntEthernetTenGigabitResultIsVlan               = false
)

const (
	testCiscoConf = `version 15.0
no service pad
service tcp-keepalives-in
service tcp-keepalives-out
service timestamps debug datetime msec localtime show-timezone
service timestamps log datetime msec localtime show-timezone
service password-encryption
service sequence-numbers
!
hostname TEST-SWITCH-01
!
boot-start-marker
boot-end-marker
!
logging rate-limit all 100 except errors
no logging console
no logging monitor
enable secret 5 $1$UWtU$seXH/uSt3l8361btwo59S/
!
username ciscoroot secret 5 $1$oqnX$L8EAakyd2gTX6mQfMIicn.
aaa new-model
!
!
aaa group server tacacs+ ACS
 server 10.10.10.1
 server 10.10.10.2
!
aaa authentication login default group ACS local-case
aaa authentication enable default group ACS enable
aaa authorization console
aaa authorization config-commands
aaa authorization exec default group ACS local 
aaa authorization commands 1 default group ACS local 
aaa authorization commands 15 default group ACS local 
!
!
!
!
!
!
aaa session-id common
clock timezone GMT 0 0
clock summer-time BST recurring last Sun Mar 1:00 last Sun Oct 2:00
switch 1 provision ws-c2960x-24fpd-l
switch 2 provision ws-c2960x-24fpd-l
no ip source-route
!
!
ip dhcp snooping vlan 1-4094
no ip dhcp snooping information option
ip dhcp snooping
no ip domain-lookup
ip domain-name net.example.com
login on-failure log
login on-success log
ipv6 mld snooping
vtp mode transparent
udld enable

!
mls qos map cos-dscp 0 8 16 24 32 46 48 56
mls qos srr-queue output cos-map queue 1 threshold 3 4 5
mls qos srr-queue output cos-map queue 2 threshold 1 2
mls qos srr-queue output cos-map queue 2 threshold 2 3
mls qos srr-queue output cos-map queue 2 threshold 3 6 7
mls qos srr-queue output cos-map queue 3 threshold 3 0
mls qos srr-queue output cos-map queue 4 threshold 3 1
mls qos srr-queue output dscp-map queue 1 threshold 3 32 33 40 41 42 43 44 45
mls qos srr-queue output dscp-map queue 1 threshold 3 46 47
mls qos srr-queue output dscp-map queue 2 threshold 1 16 17 18 19 20 21 22 23
mls qos srr-queue output dscp-map queue 2 threshold 1 26 27 28 29 30 31 34 35
mls qos srr-queue output dscp-map queue 2 threshold 1 36 37 38 39
mls qos srr-queue output dscp-map queue 2 threshold 2 24
mls qos srr-queue output dscp-map queue 2 threshold 3 48 49 50 51 52 53 54 55
mls qos srr-queue output dscp-map queue 2 threshold 3 56 57 58 59 60 61 62 63
mls qos srr-queue output dscp-map queue 3 threshold 3 0 1 2 3 4 5 6 7
mls qos srr-queue output dscp-map queue 4 threshold 1 8 9 11 13 15
mls qos srr-queue output dscp-map queue 4 threshold 2 10 12 14
mls qos queue-set output 1 threshold 1 100 100 50 200
mls qos queue-set output 1 threshold 2 125 125 100 400
mls qos queue-set output 1 threshold 3 100 100 100 400
mls qos queue-set output 1 threshold 4 60 150 50 200
mls qos queue-set output 1 buffers 15 25 40 20
mls qos
no setup express
!
!
archive
 log config
  logging enable
  logging size 200
  notify syslog contenttype plaintext
  hidekeys
!
spanning-tree mode rapid-pvst
spanning-tree portfast default
spanning-tree portfast bpduguard default
spanning-tree extend system-id
!
!
!
!
!
!
vlan internal allocation policy ascending
!
vlan 2
 name NULL
!
vlan 10
 name Trunk_Native
!
vlan 101
 name Office
!
vlan 201
 name Internet
!
vlan 301
 name Management
!
vlan 901
 name Phones
!
vlan 1001
 name INTERNET
!
ip ssh version 2
lldp timer 15
lldp run
!
!
!
!
!
interface FastEthernet0
 no ip address
 shutdown
!
interface GigabitEthernet1/0/1
 description [A-001:OF.1.001]
 switchport access vlan 101
 switchport mode access
!
interface GigabitEthernet1/0/2
 description [A-002:OF.1.001]
 switchport access vlan 101
 switchport mode access
!
interface GigabitEthernet1/0/3
 description [A-003:OF.1.001]
 switchport access vlan 101
 switchport mode access
!
interface GigabitEthernet1/0/4
 description [A-004:OF.1.001]
 shutdown
!
interface GigabitEthernet1/0/5
 description [A-005:OF.1.001]
 shutdown
!
interface GigabitEthernet1/0/6
 description [A-006:OF.1.001]
 shutdown
!
interface GigabitEthernet1/0/7
 description [A-007:OF.1.001]
 shutdown
!
interface GigabitEthernet1/0/8
 description [A-008:OF.1.001]
 shutdown
!
interface GigabitEthernet1/0/9
 description [A-009:OF.1.001]
 shutdown
!
interface GigabitEthernet1/0/10
 description [A-010:OF.1.001]
 switchport access vlan 101
 switchport mode access
!
interface GigabitEthernet1/0/11
 description [A-011:OF.1.001]-WIFI-AP001
 switchport trunk native vlan 2101
 switchport trunk allowed vlan 2101,2109,2117,2121,2122
 switchport mode trunk
!
interface GigabitEthernet1/0/12
 description [A-0012:OF.1.001]
 shutdown
!
interface GigabitEthernet1/0/13
 description [A-013:OF.1.001]
 shutdown
!
interface GigabitEthernet1/0/14
 description [A-0014:OF.1.001]
 switchport access vlan 2
 switchport mode access
 switchport voice vlan 917
 srr-queue bandwidth share 1 30 35 5
 priority-queue out 
 mls qos trust cos
 auto qos trust 
!
interface GigabitEthernet1/0/15
 description [A-015:OF.1.001]
 shutdown
!
interface GigabitEthernet1/0/16
 description [A-016:OF.1.001]
 shutdown
!
interface GigabitEthernet1/0/17
 description [A-017:OF.1.001]
 shutdown
!
interface GigabitEthernet1/0/18
 description [A-018:OF.1.001]
 shutdown
!
interface GigabitEthernet1/0/19
 description [A-019:OF.1.001]
 shutdown
!
interface GigabitEthernet1/0/20
 description [A-020:OF.1.001]
 shutdown
!
interface GigabitEthernet1/0/21
 description [A-021:OF.1.001]
 switchport trunk native vlan 2101
 switchport trunk allowed vlan 2101,2109,2117,2121,2122
 switchport mode trunk
!
interface GigabitEthernet1/0/22
 description [A-022:OF.1.001]
 switchport access vlan 137
 switchport mode access
 power inline never
!
interface GigabitEthernet1/0/23
 description [A-023:OF.1.001]
 switchport access vlan 137
 switchport mode access
 power inline never
!
interface GigabitEthernet1/0/24
 description [A-024:OF.1.001]
 switchport access vlan 137
 switchport mode access
 power inline never
!
interface GigabitEthernet1/0/25
 shutdown
!
interface GigabitEthernet1/0/26
 shutdown
!
interface TenGigabitEthernet1/0/1
 description [DP:OF.2.999]-CORE-SW01
 switchport trunk native vlan 10
 switchport mode trunk
 srr-queue bandwidth share 1 30 35 5
 priority-queue out 
 mls qos trust cos
 auto qos trust 
 ip dhcp snooping trust
!
interface TenGigabitEthernet1/0/2
 shutdown
!
interface GigabitEthernet2/0/1
 description [A-024:OF.1.001]
 switchport access vlan 2101
 switchport mode access
 switchport voice vlan 917
 shutdown
 srr-queue bandwidth share 1 30 35 5
 priority-queue out 
 mls qos trust cos
 auto qos trust 
!
interface GigabitEthernet2/0/2
 description [A-025:OF.1.001]
 switchport access vlan 2101
 switchport mode access
 switchport voice vlan 917
 shutdown
 srr-queue bandwidth share 1 30 35 5
 priority-queue out 
 mls qos trust cos
 auto qos trust 
!
interface GigabitEthernet2/0/3
 description [A-026:OF.1.001]
 switchport access vlan 2101
 switchport mode access
 switchport voice vlan 917
 shutdown
 srr-queue bandwidth share 1 30 35 5
 priority-queue out 
 mls qos trust cos
 auto qos trust 
!
interface GigabitEthernet2/0/4
 description [A-027:OF.1.001]
 switchport access vlan 2101
 switchport mode access
 switchport voice vlan 917
 shutdown
 srr-queue bandwidth share 1 30 35 5
 priority-queue out 
 mls qos trust cos
 auto qos trust 
!
interface GigabitEthernet2/0/5
 description [A-028:OF.1.001]
 switchport access vlan 2101
 switchport mode access
 switchport voice vlan 917
 shutdown
 srr-queue bandwidth share 1 30 35 5
 priority-queue out 
 mls qos trust cos
 auto qos trust 
!
interface GigabitEthernet2/0/6
 description [A-029:OF.1.001]
 switchport access vlan 2101
 switchport mode access
 switchport voice vlan 917
 shutdown
 srr-queue bandwidth share 1 30 35 5
 priority-queue out 
 mls qos trust cos
 auto qos trust 
!
interface GigabitEthernet2/0/7
 description [A-030:OF.1.001]
 switchport access vlan 2101
 switchport mode access
 switchport voice vlan 917
 shutdown
 srr-queue bandwidth share 1 30 35 5
 priority-queue out 
 mls qos trust cos
 auto qos trust 
!
interface GigabitEthernet2/0/8
 description [A-031:OF.1.001]
 switchport access vlan 2110
 switchport mode access
 shutdown
 power inline consumption 9000
!
interface GigabitEthernet2/0/9
 description [A-032:OF.1.001]
 switchport access vlan 2101
 switchport mode access
 switchport voice vlan 917
 shutdown
 srr-queue bandwidth share 1 30 35 5
 priority-queue out 
 mls qos trust cos
 auto qos trust 
!
interface GigabitEthernet2/0/10
 description [A-033:OF.1.001]
 switchport access vlan 2101
 switchport mode access
 switchport voice vlan 917
 shutdown
 srr-queue bandwidth share 1 30 35 5
 priority-queue out 
 mls qos trust cos
 auto qos trust 
!
interface GigabitEthernet2/0/11
 description [A-034:OF.1.001]
 switchport trunk native vlan 2101
 switchport trunk allowed vlan 2101,2109,2117,2121,2122
 switchport mode trunk
 shutdown
!
interface GigabitEthernet2/0/12
 shutdown
!
interface GigabitEthernet2/0/13
 shutdown
!
interface GigabitEthernet2/0/14
 shutdown
!
interface GigabitEthernet2/0/15
 shutdown
!
interface GigabitEthernet2/0/16
 shutdown
!
interface GigabitEthernet2/0/17
 shutdown
!
interface GigabitEthernet2/0/18
 shutdown
!
interface GigabitEthernet2/0/19
 shutdown
!
interface GigabitEthernet2/0/20
 shutdown
!
interface GigabitEthernet2/0/21
 shutdown
!
interface GigabitEthernet2/0/22
 shutdown
!
interface GigabitEthernet2/0/23
 shutdown
!
interface GigabitEthernet2/0/24
 shutdown
!
interface GigabitEthernet2/0/25
 shutdown
!
interface GigabitEthernet2/0/26
 shutdown
!
interface TenGigabitEthernet2/0/1
 description [DP:OF.2.999]-CORE-SW02
 switchport trunk native vlan 10
 switchport mode trunk
 srr-queue bandwidth share 1 30 35 5
 priority-queue out 
 mls qos trust cos
 auto qos trust 
 ip dhcp snooping trust
!
interface TenGigabitEthernet2/0/2
 shutdown
!
interface Vlan1
 no ip address
 shutdown
!
interface Vlan301
 ip address 10.10.1.1 255.255.255.0
!
ip default-gateway 10.10.1.254
no ip http server
no ip http secure-server
!
!
ip access-list standard snmp-access
 permit 10.10.254.1
 permit 10.10.254.2
 permit 10.100.100.0 0.0.0.15
ip access-list standard vty-access
 permit 10.10.254.1
 permit 10.10.254.2
 permit 10.100.100.0 0.0.0.15
logging history size 100
logging history errors
logging facility syslog
logging source-interface Vlan301
logging host 10.10.254.2
cdp timer 20
!
snmp-server community aeltcsnmpro RO snmp-access
tacacs-server host 10.10.10.1 key 7 095E4F0A0A111610130F17
tacacs-server host 10.10.10.2 key 7 15568A0F173E2A27293026
tacacs-server directed-request
!
!
!
alias exec sis show interface status
alias exec sir show ip route
alias exec siib show ip interface brief
alias exec mac show mac address-table
alias exec cdp show cdp neighbors
alias exec sri show run interface
alias exec sid show interface description
!
line con 0
line vty 0 4
 access-class vty-access in
 length 0
 transport input ssh
line vty 5 15
 access-class vty-access in
 length 0
 transport input ssh
!
ntp server 10.10.10.1 prefer
ntp server 10.10.10.2
end`
	testCiscoConfResultHostname            = `TEST-SWITCH-01`
	testCiscoConfResultInterfaceBlocksList = `interface FastEthernet0
 no ip address
 shutdown
!
interface GigabitEthernet1/0/1
 description [A-001:OF.1.001]
 switchport access vlan 101
 switchport mode access
!
interface GigabitEthernet1/0/2
 description [A-002:OF.1.001]
 switchport access vlan 101
 switchport mode access
!
interface GigabitEthernet1/0/3
 description [A-003:OF.1.001]
 switchport access vlan 101
 switchport mode access
!
interface GigabitEthernet1/0/4
 description [A-004:OF.1.001]
 shutdown
!
interface GigabitEthernet1/0/5
 description [A-005:OF.1.001]
 shutdown
!
interface GigabitEthernet1/0/6
 description [A-006:OF.1.001]
 shutdown
!
interface GigabitEthernet1/0/7
 description [A-007:OF.1.001]
 shutdown
!
interface GigabitEthernet1/0/8
 description [A-008:OF.1.001]
 shutdown
!
interface GigabitEthernet1/0/9
 description [A-009:OF.1.001]
 shutdown
!
interface GigabitEthernet1/0/10
 description [A-010:OF.1.001]
 switchport access vlan 101
 switchport mode access
!
interface GigabitEthernet1/0/11
 description [A-011:OF.1.001]-WIFI-AP001
 switchport trunk native vlan 2101
 switchport trunk allowed vlan 2101,2109,2117,2121,2122
 switchport mode trunk
!
interface GigabitEthernet1/0/12
 description [A-0012:OF.1.001]
 shutdown
!
interface GigabitEthernet1/0/13
 description [A-013:OF.1.001]
 shutdown
!
interface GigabitEthernet1/0/14
 description [A-0014:OF.1.001]
 switchport access vlan 2
 switchport mode access
 switchport voice vlan 917
 srr-queue bandwidth share 1 30 35 5
 priority-queue out 
 mls qos trust cos
 auto qos trust 
!
interface GigabitEthernet1/0/15
 description [A-015:OF.1.001]
 shutdown
!
interface GigabitEthernet1/0/16
 description [A-016:OF.1.001]
 shutdown
!
interface GigabitEthernet1/0/17
 description [A-017:OF.1.001]
 shutdown
!
interface GigabitEthernet1/0/18
 description [A-018:OF.1.001]
 shutdown
!
interface GigabitEthernet1/0/19
 description [A-019:OF.1.001]
 shutdown
!
interface GigabitEthernet1/0/20
 description [A-020:OF.1.001]
 shutdown
!
interface GigabitEthernet1/0/21
 description [A-021:OF.1.001]
 switchport trunk native vlan 2101
 switchport trunk allowed vlan 2101,2109,2117,2121,2122
 switchport mode trunk
!
interface GigabitEthernet1/0/22
 description [A-022:OF.1.001]
 switchport access vlan 137
 switchport mode access
 power inline never
!
interface GigabitEthernet1/0/23
 description [A-023:OF.1.001]
 switchport access vlan 137
 switchport mode access
 power inline never
!
interface GigabitEthernet1/0/24
 description [A-024:OF.1.001]
 switchport access vlan 137
 switchport mode access
 power inline never
!
interface GigabitEthernet1/0/25
 shutdown
!
interface GigabitEthernet1/0/26
 shutdown
!
interface TenGigabitEthernet1/0/1
 description [DP:OF.2.999]-CORE-SW01
 switchport trunk native vlan 10
 switchport mode trunk
 srr-queue bandwidth share 1 30 35 5
 priority-queue out 
 mls qos trust cos
 auto qos trust 
 ip dhcp snooping trust
!
interface TenGigabitEthernet1/0/2
 shutdown
!
interface GigabitEthernet2/0/1
 description [A-024:OF.1.001]
 switchport access vlan 2101
 switchport mode access
 switchport voice vlan 917
 shutdown
 srr-queue bandwidth share 1 30 35 5
 priority-queue out 
 mls qos trust cos
 auto qos trust 
!
interface GigabitEthernet2/0/2
 description [A-025:OF.1.001]
 switchport access vlan 2101
 switchport mode access
 switchport voice vlan 917
 shutdown
 srr-queue bandwidth share 1 30 35 5
 priority-queue out 
 mls qos trust cos
 auto qos trust 
!
interface GigabitEthernet2/0/3
 description [A-026:OF.1.001]
 switchport access vlan 2101
 switchport mode access
 switchport voice vlan 917
 shutdown
 srr-queue bandwidth share 1 30 35 5
 priority-queue out 
 mls qos trust cos
 auto qos trust 
!
interface GigabitEthernet2/0/4
 description [A-027:OF.1.001]
 switchport access vlan 2101
 switchport mode access
 switchport voice vlan 917
 shutdown
 srr-queue bandwidth share 1 30 35 5
 priority-queue out 
 mls qos trust cos
 auto qos trust 
!
interface GigabitEthernet2/0/5
 description [A-028:OF.1.001]
 switchport access vlan 2101
 switchport mode access
 switchport voice vlan 917
 shutdown
 srr-queue bandwidth share 1 30 35 5
 priority-queue out 
 mls qos trust cos
 auto qos trust 
!
interface GigabitEthernet2/0/6
 description [A-029:OF.1.001]
 switchport access vlan 2101
 switchport mode access
 switchport voice vlan 917
 shutdown
 srr-queue bandwidth share 1 30 35 5
 priority-queue out 
 mls qos trust cos
 auto qos trust 
!
interface GigabitEthernet2/0/7
 description [A-030:OF.1.001]
 switchport access vlan 2101
 switchport mode access
 switchport voice vlan 917
 shutdown
 srr-queue bandwidth share 1 30 35 5
 priority-queue out 
 mls qos trust cos
 auto qos trust 
!
interface GigabitEthernet2/0/8
 description [A-031:OF.1.001]
 switchport access vlan 2110
 switchport mode access
 shutdown
 power inline consumption 9000
!
interface GigabitEthernet2/0/9
 description [A-032:OF.1.001]
 switchport access vlan 2101
 switchport mode access
 switchport voice vlan 917
 shutdown
 srr-queue bandwidth share 1 30 35 5
 priority-queue out 
 mls qos trust cos
 auto qos trust 
!
interface GigabitEthernet2/0/10
 description [A-033:OF.1.001]
 switchport access vlan 2101
 switchport mode access
 switchport voice vlan 917
 shutdown
 srr-queue bandwidth share 1 30 35 5
 priority-queue out 
 mls qos trust cos
 auto qos trust 
!
interface GigabitEthernet2/0/11
 description [A-034:OF.1.001]
 switchport trunk native vlan 2101
 switchport trunk allowed vlan 2101,2109,2117,2121,2122
 switchport mode trunk
 shutdown
!
interface GigabitEthernet2/0/12
 shutdown
!
interface GigabitEthernet2/0/13
 shutdown
!
interface GigabitEthernet2/0/14
 shutdown
!
interface GigabitEthernet2/0/15
 shutdown
!
interface GigabitEthernet2/0/16
 shutdown
!
interface GigabitEthernet2/0/17
 shutdown
!
interface GigabitEthernet2/0/18
 shutdown
!
interface GigabitEthernet2/0/19
 shutdown
!
interface GigabitEthernet2/0/20
 shutdown
!
interface GigabitEthernet2/0/21
 shutdown
!
interface GigabitEthernet2/0/22
 shutdown
!
interface GigabitEthernet2/0/23
 shutdown
!
interface GigabitEthernet2/0/24
 shutdown
!
interface GigabitEthernet2/0/25
 shutdown
!
interface GigabitEthernet2/0/26
 shutdown
!
interface TenGigabitEthernet2/0/1
 description [DP:OF.2.999]-CORE-SW02
 switchport trunk native vlan 10
 switchport mode trunk
 srr-queue bandwidth share 1 30 35 5
 priority-queue out 
 mls qos trust cos
 auto qos trust 
 ip dhcp snooping trust
!
interface TenGigabitEthernet2/0/2
 shutdown
!
interface Vlan1
 no ip address
 shutdown
!
interface Vlan301
 ip address 10.10.1.1 255.255.255.0`
	testCiscoConfResultInterfaceEthernetBlocksList = `interface FastEthernet0
 no ip address
 shutdown
!
interface GigabitEthernet1/0/1
 description [A-001:OF.1.001]
 switchport access vlan 101
 switchport mode access
!
interface GigabitEthernet1/0/2
 description [A-002:OF.1.001]
 switchport access vlan 101
 switchport mode access
!
interface GigabitEthernet1/0/3
 description [A-003:OF.1.001]
 switchport access vlan 101
 switchport mode access
!
interface GigabitEthernet1/0/4
 description [A-004:OF.1.001]
 shutdown
!
interface GigabitEthernet1/0/5
 description [A-005:OF.1.001]
 shutdown
!
interface GigabitEthernet1/0/6
 description [A-006:OF.1.001]
 shutdown
!
interface GigabitEthernet1/0/7
 description [A-007:OF.1.001]
 shutdown
!
interface GigabitEthernet1/0/8
 description [A-008:OF.1.001]
 shutdown
!
interface GigabitEthernet1/0/9
 description [A-009:OF.1.001]
 shutdown
!
interface GigabitEthernet1/0/10
 description [A-010:OF.1.001]
 switchport access vlan 101
 switchport mode access
!
interface GigabitEthernet1/0/11
 description [A-011:OF.1.001]-WIFI-AP001
 switchport trunk native vlan 2101
 switchport trunk allowed vlan 2101,2109,2117,2121,2122
 switchport mode trunk
!
interface GigabitEthernet1/0/12
 description [A-0012:OF.1.001]
 shutdown
!
interface GigabitEthernet1/0/13
 description [A-013:OF.1.001]
 shutdown
!
interface GigabitEthernet1/0/14
 description [A-0014:OF.1.001]
 switchport access vlan 2
 switchport mode access
 switchport voice vlan 917
 srr-queue bandwidth share 1 30 35 5
 priority-queue out 
 mls qos trust cos
 auto qos trust 
!
interface GigabitEthernet1/0/15
 description [A-015:OF.1.001]
 shutdown
!
interface GigabitEthernet1/0/16
 description [A-016:OF.1.001]
 shutdown
!
interface GigabitEthernet1/0/17
 description [A-017:OF.1.001]
 shutdown
!
interface GigabitEthernet1/0/18
 description [A-018:OF.1.001]
 shutdown
!
interface GigabitEthernet1/0/19
 description [A-019:OF.1.001]
 shutdown
!
interface GigabitEthernet1/0/20
 description [A-020:OF.1.001]
 shutdown
!
interface GigabitEthernet1/0/21
 description [A-021:OF.1.001]
 switchport trunk native vlan 2101
 switchport trunk allowed vlan 2101,2109,2117,2121,2122
 switchport mode trunk
!
interface GigabitEthernet1/0/22
 description [A-022:OF.1.001]
 switchport access vlan 137
 switchport mode access
 power inline never
!
interface GigabitEthernet1/0/23
 description [A-023:OF.1.001]
 switchport access vlan 137
 switchport mode access
 power inline never
!
interface GigabitEthernet1/0/24
 description [A-024:OF.1.001]
 switchport access vlan 137
 switchport mode access
 power inline never
!
interface GigabitEthernet1/0/25
 shutdown
!
interface GigabitEthernet1/0/26
 shutdown
!
interface TenGigabitEthernet1/0/1
 description [DP:OF.2.999]-CORE-SW01
 switchport trunk native vlan 10
 switchport mode trunk
 srr-queue bandwidth share 1 30 35 5
 priority-queue out 
 mls qos trust cos
 auto qos trust 
 ip dhcp snooping trust
!
interface TenGigabitEthernet1/0/2
 shutdown
!
interface GigabitEthernet2/0/1
 description [A-024:OF.1.001]
 switchport access vlan 2101
 switchport mode access
 switchport voice vlan 917
 shutdown
 srr-queue bandwidth share 1 30 35 5
 priority-queue out 
 mls qos trust cos
 auto qos trust 
!
interface GigabitEthernet2/0/2
 description [A-025:OF.1.001]
 switchport access vlan 2101
 switchport mode access
 switchport voice vlan 917
 shutdown
 srr-queue bandwidth share 1 30 35 5
 priority-queue out 
 mls qos trust cos
 auto qos trust 
!
interface GigabitEthernet2/0/3
 description [A-026:OF.1.001]
 switchport access vlan 2101
 switchport mode access
 switchport voice vlan 917
 shutdown
 srr-queue bandwidth share 1 30 35 5
 priority-queue out 
 mls qos trust cos
 auto qos trust 
!
interface GigabitEthernet2/0/4
 description [A-027:OF.1.001]
 switchport access vlan 2101
 switchport mode access
 switchport voice vlan 917
 shutdown
 srr-queue bandwidth share 1 30 35 5
 priority-queue out 
 mls qos trust cos
 auto qos trust 
!
interface GigabitEthernet2/0/5
 description [A-028:OF.1.001]
 switchport access vlan 2101
 switchport mode access
 switchport voice vlan 917
 shutdown
 srr-queue bandwidth share 1 30 35 5
 priority-queue out 
 mls qos trust cos
 auto qos trust 
!
interface GigabitEthernet2/0/6
 description [A-029:OF.1.001]
 switchport access vlan 2101
 switchport mode access
 switchport voice vlan 917
 shutdown
 srr-queue bandwidth share 1 30 35 5
 priority-queue out 
 mls qos trust cos
 auto qos trust 
!
interface GigabitEthernet2/0/7
 description [A-030:OF.1.001]
 switchport access vlan 2101
 switchport mode access
 switchport voice vlan 917
 shutdown
 srr-queue bandwidth share 1 30 35 5
 priority-queue out 
 mls qos trust cos
 auto qos trust 
!
interface GigabitEthernet2/0/8
 description [A-031:OF.1.001]
 switchport access vlan 2110
 switchport mode access
 shutdown
 power inline consumption 9000
!
interface GigabitEthernet2/0/9
 description [A-032:OF.1.001]
 switchport access vlan 2101
 switchport mode access
 switchport voice vlan 917
 shutdown
 srr-queue bandwidth share 1 30 35 5
 priority-queue out 
 mls qos trust cos
 auto qos trust 
!
interface GigabitEthernet2/0/10
 description [A-033:OF.1.001]
 switchport access vlan 2101
 switchport mode access
 switchport voice vlan 917
 shutdown
 srr-queue bandwidth share 1 30 35 5
 priority-queue out 
 mls qos trust cos
 auto qos trust 
!
interface GigabitEthernet2/0/11
 description [A-034:OF.1.001]
 switchport trunk native vlan 2101
 switchport trunk allowed vlan 2101,2109,2117,2121,2122
 switchport mode trunk
 shutdown
!
interface GigabitEthernet2/0/12
 shutdown
!
interface GigabitEthernet2/0/13
 shutdown
!
interface GigabitEthernet2/0/14
 shutdown
!
interface GigabitEthernet2/0/15
 shutdown
!
interface GigabitEthernet2/0/16
 shutdown
!
interface GigabitEthernet2/0/17
 shutdown
!
interface GigabitEthernet2/0/18
 shutdown
!
interface GigabitEthernet2/0/19
 shutdown
!
interface GigabitEthernet2/0/20
 shutdown
!
interface GigabitEthernet2/0/21
 shutdown
!
interface GigabitEthernet2/0/22
 shutdown
!
interface GigabitEthernet2/0/23
 shutdown
!
interface GigabitEthernet2/0/24
 shutdown
!
interface GigabitEthernet2/0/25
 shutdown
!
interface GigabitEthernet2/0/26
 shutdown
!
interface TenGigabitEthernet2/0/1
 description [DP:OF.2.999]-CORE-SW02
 switchport trunk native vlan 10
 switchport mode trunk
 srr-queue bandwidth share 1 30 35 5
 priority-queue out 
 mls qos trust cos
 auto qos trust 
 ip dhcp snooping trust
!
interface TenGigabitEthernet2/0/2
 shutdown`
	testCiscoConfResultInterfaceVlanBlocksList = `interface Vlan1
 no ip address
 shutdown
!
interface Vlan301
 ip address 10.10.1.1 255.255.255.0`
	testCiscoConfResultVlanBlocksList = `vlan 2
 name NULL
!
vlan 10
 name Trunk_Native
!
vlan 101
 name Office
!
vlan 201
 name Internet
!
vlan 301
 name Management
!
vlan 901
 name Phones
!
vlan 1001
 name INTERNET`
)
