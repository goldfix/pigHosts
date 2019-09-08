[![GitHub](https://img.shields.io/github/license/goldfix/pigHosts)](https://github.com/goldfix/pigHosts/blob/master/LICENSE)
[![GitHub release (latest SemVer)](https://img.shields.io/github/v/release/goldfix/pigHosts?sort=semver)](https://github.com/goldfix/pigHosts/releases)
[![Travis (.com) branch](https://img.shields.io/travis/com/goldfix/pigHosts/master?label=master%20build)](https://travis-ci.com/goldfix/pigHosts)

## pigHosts

- Latest version: [![GitHub release (latest SemVer)](https://img.shields.io/github/v/release/goldfix/pigHosts?sort=semver)](https://github.com/goldfix/pigHosts/releases)
- Download blocklist hosts from remote sites.
- Add and remove blocklist hosts, from your hosts file.
- Possibility to customize your download blocklist links.
- Possibility to esclude specific hosts.
- Windows & Linux version.

### Use:

```

pigHost

Usage:
 pigHost (load | unload | force_init) [--debug] | (--version) | (--check-update)
 pigHost (--help | -h)

Options:
 -h, --help       help online
 -v, --version    view version
 --check-update   check if there is a new version
 --debug          view debug info

Command:
 load           load custom hosts from external urls declared in the file: '<USER_FOLDER>/.pigHosts/pigHosts.urls'
 unload         disable and remove custom hosts
 force_init     delete and create a new set of configuration files: '<USER_FOLDER>/.pigHosts/pigHosts.excluded' and '<USER_FOLDER>/.pigHosts/pigHosts.urls'

```

### Configuration files:

When you started first time pigHost, will be created two configuration file:

- `<USER_FOLDER>/.pigHosts/pigHosts.excluded` : list of hosts to esclude. These hosts are excluded, if present, from blocklist.
- `<USER_FOLDER>/.pigHosts/pigHosts.urls` : in this file you can add your list of urls where to download blocklist .

### Important notes

Some antivirus locks `.../etc/hosts` file. Remember to configure the antivirus correctly so not to block this file.

- On Windows is necessary run `pigHosts` as _Administrator_.
- On Linux is necessary run `pigHosts` with elevated privileges (`sudo pighost ...`).

### Build from source

- Download [go-task](https://github.com/go-task/task/releases).
- Set your `GOPATH` environment variable.
- In the `./src` folder, execute command: `task build-mod-vendor`.

#### Thanks to:

Blocklist Collection sites:

- https://firebog.net/
- https://www.squidblacklist.org/
- https://github.com/StevenBlack/hosts
- https://someonewhocares.org/

Tools and libraries:

- https://github.com/docopt/docopt.go
- https://github.com/sirupsen/logrus
- https://github.com/go-task/task
- https://github.com/goreleaser/goreleaser
- https://github.com/briandowns/spinner
