# pinger

A utility to help determine when network outage occurs and how much of the network is done.

## Installation

Download the latest binary for your operating system.  It is recommended to add the binary to your
system's PATH.

The latest stable release can be found [here](https://github.com/sbu-ces-unofficial/pinger/releases/latest).

All releases can be found [here](https://github.com/sbu-ces-unofficial/pinger/releases).

Note: Unless you need the features in a pre-release, it's recommended to use the binaries tagged
under `Latest Release`.

## Usage

On Unix systems, run as (note that at the moment, `sudo` is required):

```bash
sudo pinger [flags]
```

On Windows systems, run as:
```bash
pinger.exe [flags]
```

Available flags:

- `-h, --help`: get help on pinger
- `--version`: get the version of pinger (useful for seeing how up-to-date your copy is)

Note:
- Administrator permission is required on Unix systems to make ICMP requests
  - As a result, logs are also created as root and can only be deleted by root
- If the binary is not in your path and it is in the current directory, you would need to specify
  `./pinger` or `./pinger.exe`.
- All results get logged to _internet_connectivity.log_.

Examples:

```bash
# Pings google.com every minute and pings blackboard.stonybrook.edu if google.com cannot
# be reached. Log the results.
sudo pinger
```
