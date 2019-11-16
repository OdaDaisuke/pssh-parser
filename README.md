# PSSH parser written in go

This repository provides PSSH parser.

In DRM context, PSSH is used in MPEG-DASH content protectection.

## Installation

```shell
go get github.com/OdaDaisuke/pssh-parser/cmd/pssh
```

## Usage

Specify PSSH raw data or base64 encoded format.

### Raw data

If you want to use base64 encoded value, use `-e` option.

```shell
pssh -s AAAAxnBzc2gBAAAA7e+LqXnWSs6jyCfc1R0h7QAAAAINw+xPdoNUi4HnPGTlguE2FEe37S9mVyu9EwbOfPNhDQAAAIISEBRHt+0vZlcrvRMGznzzYQ0SEFrGoR6qL17Vv2aMQByBNMoSEG7hNRbI51h7rp9+zT6Zom4SEPnsEqYaJl1Hj4MzTjp40scSEA3D7E92g1SLgec8ZOWC4TYaDXdpZGV2aW5lX3Rlc3QiEXVuaWZpZWQtc3RyZWFtaW5nSOPclZsG -e
```

If you want to pass data by file, use `-f` option like this.

```shell
pssh -f ./pssh_data.txt
```
