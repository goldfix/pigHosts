package pighosts

var PigHostsUrls = ""
var PigHostsExcluded = ""

const numHostPerLine = 9

const nonRoutable = "0.0.0.0"
const localHost = "127.0.0.1"

var filterSpecificHostDefault = []string{
	"127.0.0.1 localhost",
	"127.0.0.1 localhost.localdomain",
	"127.0.0.1 local",
	"255.255.255.255 broadcasthost",
	"::1 localhost",
	"::1 ip6-localhost",
	"::1 ip6-loopback",
	"fe80::1%lo0 localhost",
	"ff00::0 ip6-localnet",
	"ff00::0 ip6-mcastprefix",
	"ff02::1 ip6-allnodes",
	"ff02::2 ip6-allrouters",
	"ff02::3 ip6-allhosts",
	"0.0.0.0 0.0.0.0",
	"localhost",
	"localhost.localdomain",
	"local",
	"broadcasthost",
	"ip6-localhost",
	"ip6-loopback",
	"ip6-localnet",
	"ip6-mcastprefix",
	"ip6-allnodes",
	"ip6-allrouters",
	"ip6-allhosts",
	"0.0.0.0",
}

var defaultHostsUrlsDefault = []string{
	"https://raw.githubusercontent.com/StevenBlack/hosts/master/hosts",
	"https://www.squidblacklist.org/downloads/dg-ads.acl",
	"https://www.squidblacklist.org/downloads/dg-malicious.acl",
}

var defaultHostsUrlsTmp = []string{}
var filterSpecificHostTmp = []string{}
