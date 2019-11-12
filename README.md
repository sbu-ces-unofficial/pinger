# pinger

A utility that uses TCP pings for monitoring and collecting information to troubleshoot network outages.

## Installation

Download the latest binary for your operating system.  It is recommended to add the binary to your
system's PATH.

The latest stable release can be found [here](https://github.com/sbu-ces-unofficial/pinger/releases/latest).

All releases can be found [here](https://github.com/sbu-ces-unofficial/pinger/releases).

Note: Unless you need the features in a pre-release, it's recommended to use the binaries tagged
under `Latest Release`.

## Usage

On Unix systems (Note in 0.0.3 release and below, sudo was required):

```bash
pinger [command]
// OR
pinger [flags]
```

On Windows systems, run as:
```bash
pinger.exe [command]
// OR
pinger.exe [flags]
```

Available commands:

- `monitor`: Track when network outages occur and their severity
- `report`: Generate a report of the network status

Available flags:

- `-h, --help`: get help on pinger
- `--version`: get the version of pinger (useful for seeing how up-to-date your copy is)

Note:
- For Unix machines, 0.0.3 release and below requires sudo privileges to make ICMP requests.  This
  is no longer necessary as newer versions make use of TCP pings.
- If the binary is not in your path and it is in the current directory, you would need to specify
  `./pinger` or `./pinger.exe`.
- All results from the `monitor` subcommand gets logged to _internet_connectivity.log_.
- All results from the `report` subcommand gets logged to _connectivity_report.txt_.

Examples:

```bash
# On Unix

# Pings google.com every minute and pings blackboard.stonybrook.edu if google.com cannot
# be reached. Log the results.
sudo pinger monitor

# Pings google.com and blackboard.stonybrook.edu and logs the result.
sudo pinger report
```

## Configuring

It is possible to configure the behavior of pinger with a _config.pflags_ file.

The basic structure of _config.pflags_ for pinger:

```
[[monitor]]

[external_urls]
// all external urls you want to ping when running the monitor subcommand

[internal_urls]
// all internals url you want to ping when running the monitor subcommand

[[report]]

[external_urls]
// all external urls you want to ping when running the report subcommand

[internal_urls]
// all internal urls you want to ping when running the report subcommand
```

Notes:

- A sample `config.pflags` have been bundled with the binary in the release.  It is written to
  follow the default behavior of `pinger` for the monitor subcommand and to follow the data
  asked in the Google Forms for the report subcommand (stonybrook.edu/mycloud has been excluded
  because it does return a valid ping result).
- Please note that comments are not supported at this time and all urls should be in quotes.
