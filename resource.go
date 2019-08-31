package pighosts

import "time"

var pigHostsUrls = ""
var pigHostsExcluded = ""

const numHostPerLine = 9

const nonRoutable = "0.0.0.0"
const localHost = "127.0.0.1"
const localHost_ipv6 = "::1"

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
	"fe00::0",
	"ff00::0",
	"ff02::1",
	"ff02::2",
}

var defaultHostsUrlsDefault = []string{
	"# https://raw.githubusercontent.com/StevenBlack/hosts/master/hosts",
	"# https://www.squidblacklist.org/downloads/dg-ads.acl",
	"# https://www.squidblacklist.org/downloads/dg-malicious.acl",
	"https://someonewhocares.org/hosts/ipv6/hosts",
}

var defaultHostsUrlsTmp = []string{}
var filterSpecificHostTmp = []string{}

const headerHostFile = "###--pigHost_START------------------------------------"
const footerHostFile = "###--pigHosts_END-------------------------------------"

const hostFile = "/Windows/System32/drivers/etc/hosts"
const hostFileNew = "/tmp/pigHostBak/host.new"
const hostFileEmpty = "/tmp/pigHostBak/host.empty"

var hostFileBak = "/tmp/pigHostBak/host_" + time.Now().Format("20060201T1504") + ".bak"

const manifest = `<?xml version="1.0" encoding="UTF-8" standalone="yes"?>
<assembly xmlns="urn:schemas-microsoft-com:asm.v1" manifestVersion="1.0">
<assemblyIdentity
    version="1.0.0.0"
    name="pigHosts.exe"
    type="win32"
/>
<description>pigHosts.exe</description>
<trustInfo xmlns="urn:schemas-microsoft-com:asm.v3">
    <security>
        <requestedPrivileges>
            <requestedExecutionLevel level="requireAdministrator" uiAccess="false"/>
        </requestedPrivileges>
    </security>
</trustInfo>
</assembly>`
