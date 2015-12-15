package dhcp

// Options defined in RFC1533
// Options updated by RFC2132

type option byte

const (
	// RFC1497 Vendor Extensions (RFC1533 section-3)
	pad                    option = 0
	end                    option = 255
	subnetMask             option = 1
	timeOffset             option = 2
	router                 option = 3
	timeServer             option = 4
	nameServer             option = 5
	domainNameServer       option = 6
	logServer              option = 7
	cookieServer           option = 8
	lprServer              option = 9  // Line Printer Servers
	impressServer          option = 10 // Imagen Impress Servers
	resourceLocationServer option = 11
	hostName               option = 12 // Client Host Name
	bootFileSize           option = 13
	meritDumpFile          option = 14
	domainName             option = 15
	swapServer             option = 16
	rootPath               option = 17
	extensionsPath         option = 18

	// IP Layer Parameters per Host (section-4)
	ipForwarding              option = 19 // Enable/Disable
	nonLocalSourceRouting     option = 20 // Enable/Disable
	policyFilter              option = 21
	maxDatagramReassemblySize option = 22
	defaultIPTimeToLive       option = 23
	pathMTUAgingTimeout       option = 24
	pathMTUPlateauTable       option = 25

	// IP Layer Parameters per Interface (section-5)
	interfaceMTU              option = 26
	allSubnetsAreLocal        option = 27
	broadcastAddress          option = 28
	performMaskDiscovery      option = 29
	maskSupplier              option = 30
	performRouterDiscovery    option = 31
	routerSolicitationAddress option = 32
	staticRoute               option = 33

	// Link Layer Parameters per Interface (section-6)
	trailerEncapsulation  option = 34
	arpCacheTimeout       option = 35
	ethernetEncapsulation option = 36

	// TCP Parameters (section-7)
	tcpDefaultTTL        option = 37
	tcpKeepaliveInterval option = 38
	tcpKeepaliveGarbage  option = 39

	// Application and Service Parameters (section-8)
	networkInformationServiceDomain            option = 40
	networkInformationServers                  option = 41
	networkTimeProtocolServers                 option = 42
	vendorSpecificInformation                  option = 43
	netBIOSOverTCPIPNameServer                 option = 44
	netBIOSOverTCPIPDatagramDistributionServer option = 45
	netBIOSOverTCPIPNodeType                   option = 46
	netBIOSOverTCPIPScope                      option = 47
	xWindowSystemFontServer                    option = 48
	xWindowSystemDisplayManager                option = 49

	//  DHCP Extensions (section-9)
	requestedIPAddress     option = 50
	ipAddressLeaseTime     option = 51
	overloadOption         option = 52
	dhcpMessageType        option = 53 // DHCP Message Type
	serverIdentifier       option = 54
	parameterRequestList   option = 55
	message                option = 56
	maximumDHCPMessageSize option = 57
	renewalTimeValue       option = 58 // T1
	rebindingTimeValue     option = 59 // T2
	classIdentifier        option = 60
	clientIdentifier       option = 61

	// Application and Service Parameters updated by rfc2132 (section-8)
	networkInformationServicePlusDomain  option = 64
	networkInformationServicePlusServers option = 65
	mobileIPHomeAgent                    option = 68
	simpleMailTransportProtocol          option = 69
	postOfficeProtocolServer             option = 70
	networkNewsTransportProtocol         option = 71
	defaultWorldWideWebServer            option = 72
	defaultFingerServer                  option = 73
	defaultInternetRelayChatServer       option = 74
	streetTalkServer                     option = 75
	streetTalkDirectoryAssistance        option = 76
)

type Options struct {
	cookie [4]byte
	_      map[option][]byte
}
