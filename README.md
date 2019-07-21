## pigHosts

- Download blacklist hosts from remote sites.
- Add and remove blacklist hosts, in your hosts file.
- Possibility to customize your links list from where download blacklist file.
- Possibility to esclude specific hosts.
- Only for Windows... w.i.p. Linux version.

### Use: 

```

pigHost

Usage: pigHost [load | unload | force_init] [-h | -v | -o]
 pigHost (load)
 pigHost (unload)
 pigHost (force_init)
 pigHost (--help | -h)
 pigHost (--version | -v)

Options:
 -h, --help     help online
 -o, --other    other params
 -v, --version  view version

Command:
 unload         disable and remove custom hosts
 load           load custom hosts from external urls file (if file is not declared pigHost uses /HOME_FOLDER/.pigHosts/pigHosts.urls)
 force_init     delete and create a new set of configuration files: '.pigHosts/pigHosts.excluded' and '.pigHosts/pigHosts.urls' in your user/home folder

```

### Configuration files: 

When you started first time pigHost, will be created two configuration file:

- `<USER_FOLDER>/.pigHosts/pigHosts.excluded` : list of hosts to esclude. These hosts are excluded, if present, from blacklists.
- `<USER_FOLDER>/.pigHosts/pigHosts.urls` : in this file can you add your list of urls where to download blacklists.

### Important notes

Some antivirus locks `.../etc/hosts` file. Remember to configure correctly you antivirus to do not lock this file.

- On Windows is necessary run `pigHosts` as _Administrator_.
- On Linux is necessary run `pigHosts` with elevated privileges (`sudo pighost ...`).

#### Thanks to: 
- https://firebog.net/
- https://www.squidblacklist.org/
- https://github.com/StevenBlack/hosts
- https://github.com/docopt/docopt.go
- github.com/sirupsen/logrus

