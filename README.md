# Network Device Configuration Parser (confparser)

## Overview

This **Network Device Configuration Parser** library was written to fulfil the need to parse Cisco switch configuration files.

I've used it in a number of projects where I needed to grab information from these types of configuration files. It's very basic, but has saved me from having to repeatedly implement all of the regular expressions and logical multiple times.

At the moment, the library only supports Cisco IOS switch configuration files. It's been primarily tested with IOS 12.x files, but should work on any. But make sure you test and validate!

It may also work on Cisco IOS router configuration files. It's just geared towards information in Switch interfaces, as opposed to the IP interfaces, which are more common in routers.


## Disclaimer

This is the standard, "Use at your own risk," disclaimer. You are very welcome to use the library, but if you do, then that's on you. 😉

See the [**LICENSE.md**](./LICENSE.md) file for full details.


## How To Use - The Basics

First, import the library as normal:

```golang
import (
	"github.com/tinyinput/confparser"
)
```

Then the easiest was to parse a configuration file is to first create a **Client**:

```golang
package main

import (
	"github.com/tinyinput/confparser"
)

func main() {
	// Create a Client
	client, err := confparser.NewClient("cisco")
	if err != nil {
		// Handle the error
	}
}
```

Next you can pass your configuration file to the **Client**, to return a **Config Parser**:

```golang
package main

import (
	"github.com/tinyinput/confparser"
)

var (
	confFile = `Some long multi-line confguration file text`
)

func main() {
	// Create a Client
	client, err := confparser.NewClient("cisco")
	if err != nil {
		// Handle the error
	}
	// Pass your configuration file to the Client to create a Config Parser
	myConfigParser := client.ConfigParser(confFile)
}
```

Then you can call various _Methods_ on your **Config Parser** to extract elements from the configuration.

For example, if you wanted to **hostname** from the configuration, simply call that _Method_:

```golang
package main

import (
	"github.com/tinyinput/confparser"
	"fmt"
)

var (
	confFile = `Some long multi-line confguration file text`
)

func main() {
	// Create a Client
	client, err := confparser.NewClient("cisco")
	if err != nil {
		// Handle the error
	}
	// Pass your configuration file to the Client to create a Config Parser
	myConfigParser := client.ConfigParser(confFile)
	// Call the Hostname() Method to return the hostname from the configuration
	hostname := myConfigParser.Hostname()
	//
	fmt.Printf("Hostname from the Config: %s\n", hostname)
}
```

If you wanted to extract all of the **interface** configuration blocks, you could you the _Method_ `InterfaceBlocks()`. The _Methods_ can also be combined with the **Config Parser** creation, into a single line.

For example:
```golang
package main

import (
	"github.com/tinyinput/confparser"
	"fmt"
)

var (
	confFile = `Some long multi-line confguration file text`
)

func main() {
	// Create a Client
	client, err := confparser.NewClient("cisco")
	if err != nil {
		// Handle the error
	}
	// Pass your configuration file to the Client to create a Config Parser
	// and extract the interface configuration blacks
	myInterfaceConfigBlocks := client.ConfigParser(confFile).InterfaceBlocks()
	//
	fmt.Printf("Count of Interface Config Blocks from the Config: %v\n", len(myInterfaceConfigBlocks))
}
```

## How To Use - Interface Configuration

The main use-case I had for the library was to extract **interface** configurations from Cisco switch configurations.

I then needed to parse out various elements from each **interface**.

* What was the **Access Vlan Number**?
* What was the **Description**
* Was the interface a **Trunk**?
* Was the interface **Shutdown**?

The library allows me to do this in a number of ways.

First, we need the **interface** configuration blocks:
```golang
package main

import (
	"github.com/tinyinput/confparser"
)

var (
	confFile = `Some long multi-line confguration file text`
)

func main() {
	// Create a Client
	client, err := confparser.NewClient("cisco")
	if err != nil {
		// Handle the error
	}
	// Pass your configuration file to the Client to create a Config Parser
	// and extract the interface configuration blacks
	myInterfaceConfigBlocks := client.ConfigParser(confFile).InterfaceBlocks()
}
```

We can then loop over each **interface** configuration block to extract the various logical elements we need. It's up to you as to exactly _what_ you want, so below is just an example:
```golang
package main

import (
	"github.com/tinyinput/confparser"
)

var (
	confFile = `Some long multi-line confguration file text`
)

type myInterfaceType struct {
	name        string
	description string
	shutdown    bool
}

type myInterfacesType []myInterfaceType

func main() {
	// Create a Client
	client, err := confparser.NewClient("cisco")
	if err != nil {
		// Handle the error
	}
	// Pass your configuration file to the Client to create a Config Parser
	// and extract the interface configuration blacks
	myInterfaceConfigBlocks := client.ConfigParser(confFile).InterfaceBlocks()
	// Create a variable to hold our all of our interface information
	var myInterfaces myInterfacesType
	// Start a loop over all of the interface configuration blocks
	for _, block := range myInterfaceConfigBlocks {
		// Create a temporary interface struct to hold our parsed information
		var myinterface myInterfaceType
		// Create an Interface Configuration Parser, populated with the interface configuration block
		myIntParser := client.ConfigIntParser(block)
		// Use the appropriate method to extract information and store it in our temporary interface struct
		myinterface.name = myIntParser.InterfaceNameFull()
		myinterface.description = myIntParser.Description()
		myinterface.shutdown = myIntParser.IsShutdown()
		// Append the temporary interface struct to our main variable we created earlier
		myInterfaces = append(myInterfaces, myinterface)
	}
	// Start a loop over all of our parsed interfaces
	for _, intf := range myInterfaces {
		// Print the information for each interface
		fmt.Printf("Interface: %s - Description: %s - Shutdown? %v\n", intf.name, intf.description, intf.shutdown)
	}
}
```

## The Configuration Parser Interface

The following _Methods_ are defined for the `ConfParser` interface type. It is the type which can be ued to extract various elments from the main configuration.

* **Set(s string)**
	* This can be used to pass the configuration to the _Parser_
* **String()**
	* The outputs the configuration set on the _Parser_
* **Hostname()**
	* The _Hostname_ defined in the configuration
* **InterfaceBlocks()**
	* The _Interface Blocks_ from the configuration
* **InterfaceBlocksEthernet()**
	* Only the _Ethernet Interface Blocks_ from the configuration
* **InterfaceBlocksVlan()**
	* Only the _Vlan Interface Blocks_ from the configuration
* **VlanBlocks()**
	* The _Vlan Blocks_ from the configuration


## The Interface Configuration Parser Interface

(Yes... It's a confusing heading...)

The following _Methods_ are defined for the `ConfIntParser` interface type. It is the type which can be used to extract various information and logical state from individual _blocks_ of interface configuration.

* **Set(s string)**
	* This can be used to pass the interface configuration to the _Parser_
* **String()**
	* The outputs the interface configuration set on the _Parser_
* **InterfaceNameFull()**
	* The _Full_ interface name
* **InterfaceNameShort()**
	* The _Shortened_ interface name
* **Description()**
	* The interface _Description_
* **VlanAccess()**
	* The configured _Access Vlan_ number on the interface
* **VlanVoice()**
	* The configured _Voice Vlan_ number on the interface
* **VlanTrunkNative()**
	* The configured _Trunk Native Vlan_ number on the interface
* **VlansTrunkAllowed()**
	* The configured _Allowed Vlan_ numbers on the _Trunk_ for the interface
* **IpAddressAndMask()**
	* The _IP Address_ and _Subnet Mask_ configured on the interface
* **IpAddress()**
	* The _IP Address_ configured on the interface
* **IpMask()**
	* The _Subnet Mask_ configured on the interface
* **IsShutdown()**
	* Is the interface configured to be _Shutdown_?
* **IsSwAccess()**
	* Is the interface configured to be an _Access_ switch port? 
* **IsSwTrunk()**
	* Is the interface configured to be a _Trunk_ switch port?
* **IsPoeDisabled()**
	* Is _Power over Ethernet_ disabled on the interface?
	* Note: A _false_ result doesn't guarantee that _PoE_ is _Enabled_. You can only use this check to see is configuration exists to manually _Disabled_ _PoE_.
* **IsEthernet()**
	* Is the interface an _Ethernet_ interface?
* **IsEthernetFast()**
	* Is the interface a _Fast Ethernet_ interface?
* **IsEthernetGigabit()**
	* Is the interface a _Gigabit Ethernet_ interface?
* **IsEthernetTenGigabit()**
	* Is the interface a _Ten Gigabit Ethernet_ interface?
* **IsVlan()**
	* Is the interface a _Vlan_ interface?

For more details about return types and caveats, see the comments for the specific parser _Methods_ in the code.


## General Notes on Errors

One thing to note is that the _Functions_ and _Methods_ in the library will return empty strings or slices in preference to errors. The theory being, that in most cases, if a configuration item isn't found, it's not an error; it's just not in the config!


## Extras - How to Read a Configuration File

It's outside the scope of this library to discuss all of the ways in which _Golang_ can read files. But in the interests of getting people started, the following works for me:

```golang
package main

import (
	"github.com/tinyinput/confparser"
	"fmt"
	"os"
)

func main() {
	// Use the ReadFile() function from the os library
	data, err := os.ReadFile("/user/home/me/data/configFile")
	if err != nil {
		// Handle the error
	}
	// Create a Client
	client, err := confparser.NewClient("cisco")
	if err != nil {
		// Handdle the error
	}
	// Pass your configuration file to the Client, as a string, to create a Config Parser,
	// then call the `Hostname()` method
	myConfigParser := client.ConfigParser(string(data)).Hostname()
	// Print the hostname from the configuration file
	fmt.Printf("Hostname from the Config: %s\n", hostname)
}

```

This uses the _Function_ `ReadFile()` from the `os` library here: <https://pkg.go.dev/os#ReadFile>


## Extras - Adding a new Configuration "Make"

As explained above, this library was written primarily to parse Cisco switch configuration files. But at the same time, it was written so that you can add your own logic for a new type (or _Make_) of configuration file.

If you're a seasoned **Gopher**, then it should be pretty clear as to how this is done.

There are basically 3 steps:

1. Create 3 _Type_ definitions similar to: `CiscoConf`, `CiscoConfVlan` and `CiscoConfInt`
2. Create _Methods_ for each of your new _Types_ that fulfil the _Interfaces_: `ConfParser`, `ConfVlanParser` and `ConfIntParser`
3. Create a new global variable for your make of _Type_ `ParserClient`. See `ciscoParserClient` as an example
4. Update these _Functions_ with a new `case` statement: `NewClient()`, `NewConfigParser()`, `NewConfigVlanParser()` and `NewConfigIntParser()`

Then in the code which uses the updated library, you should be to call `NewClient("your custom make")` to create a **Parser Client** that supports your custom configuration file.




