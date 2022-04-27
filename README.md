LumenX is the blockchain built using the [Cosmos SDK](https://github.com/cosmos/cosmos-sdk). LumenX will be interact with other sovereign blockchains using a protocol called [IBC](https://github.com/cosmos/ics/tree/master/ibc) that enables Inter-Blockchain Communication.

# LumenX network

## Mainnet Full Node Quick Start
With each version of the LumenX, the chain is restarted from a new Genesis state. We are currently on lumenx.

Get mainnet config [here](https://github.com/metaprotocol-ai/lumenx/tree/master/config)

```
- Hardware requirements
CPU: 4 core
RAM: 16 GB
DISK Storage: SSD 1,000 GB

- Software requirements
OS: Ubuntu Server 20.04
Go version: Go 1.17+
```

### Build from code

This assumes that you're running Linux and have installed [Go 1.17+](https://golang.org/dl/).  This guide helps you:

Build, Install, and Name your Node
```bash
# Clone LumenX from the latest release found here: https://github.com/metaprotocol-ai/lumenx/releases
git clone https://github.com/metaprotocol-ai/lumenx
# Enter the folder LumenX was cloned into
cd lumenx
# Compile and install LumenX
make build
# Check LumenX version
build/lumenxd version
```

### To join mainnet follow this steps

#### Genesis & Seeds
Download [genesis.json](https://raw.githubusercontent.com/metaprotocol-ai/lumenx/master/config/genesis.json)
```
wget -O $HOME/.lumenx/config/genesis.json https://raw.githubusercontent.com/metaprotocol-ai/lumenx/master/config/genesis.json
```
Download [config.toml](https://raw.githubusercontent.com/metaprotocol-ai/lumenx/master/config/config.toml) with predefined seeds and persistent peers
```
wget -O $HOME/.lumenx/config/config.toml https://raw.githubusercontent.com/metaprotocol-ai/lumenx/master/config/config.toml
```

Alternatively enter persistent peers to config.toml provided [here](https://github.com/metaprotocol-ai/lumenx/tree/master/config)

1) Open ~/.lumenx/config/config.toml with text editor. Alternatively you can use cli editor, like nano ``` nano ~/.lumenx/config/config.toml ```
2) Scroll down to persistant peers in `config.toml`, and add the persistant peers as a comma-separated list

#### Setting Up a New Node
You can edit this moniker, in the ~/.lumenx/config/config.toml file:
```bash
# A custom human readable name for this node
moniker = "<your_custom_moniker>"
```

You can edit the ~/.lumenx/config/app.toml file in order to enable the anti spam mechanism and reject incoming transactions with less than the minimum gas prices:
```
# This is a TOML config file.
# For more information, see https://github.com/toml-lang/toml
##### main base config options #####
# The minimum gas prices a validator is willing to accept for processing a
# transaction. A transaction's fees must meet the minimum of any denomination
# specified in this config (e.g. 10ulumen).
minimum-gas-prices = ""
```
Your full node has been initialized!

#### Run a full node
```
# Start LumenX
build/lumenxd start
# to run process in background run
screen -dmSL lumenxd build/lumenxd start
# Check your node's status with LumenX cli
build/lumenxd status
```
If you want to run lumenxd as a permanent background service, see [here](https://github.com/metaprotocol-ai/lumenx/tree/master/run-a-node-as-a-background-service)

### Create a key
Add new
``` bash
build/lumenxd keys add <key_name>
```

Or import via mnemonic
```bash
build/lumenxd keys add <key_name> -i
```

As a result, you got
```bash
- name: <key_name>
  type: local
  address: <key_address>
  pubkey: <key_pubkey>
  mnemonic: ""
  threshold: 0
  pubkeys: []
**Important** write this mnemonic phrase in a safe place.
It is the only way to recover your account if you ever forget your password.
<key_mnemonic>
```

### To become a validator follow this steps
Before setting up your validator node, make sure you've already gone through the [Full Node Setup](https://github.com/metaprotocol-ai/lumenx#to-join-mainnet-follow-this-steps)

#### What is a Validator?
[Validators](https://docs.cosmos.network/v0.44/modules/staking/01_state.html#validator) are responsible for committing new blocks to the blockchain through voting. A validator's stake is slashed if they become unavailable or sign blocks at the same height.
Please read about [Sentry Node Architecture](https://hub.cosmos.network/main/validators/security.html#sentry-nodes-ddos-protection) to protect your node from DDOS attacks and to ensure high-availability.

#### Create Your Validator

Your `lumenvalconspub` can be used to create a new validator by staking tokens. You can find your validator pubkey by running:

```bash
build/lumenxd tendermint show-validator
```

To create your validator, just use the following command:

Check if your key(address) has enough balance:

```bash
build/lumenxd query bank balances <key address>
```

For test nodes, `chain-id` is `LumenX`.\
You need transction fee `2000000ulumen` to make your transaction for creating validator.\
Don't use more `ulumen` than you have!

```bash
build/lumenxd tx staking create-validator \
  --amount=10000000ulumen \
  --pubkey=$(build/lumenxd tendermint show-validator) \
  --moniker=<choose a moniker> \
  --chain-id=<chain_id> \
  --commission-rate="0.10" \
  --commission-max-rate="0.20" \
  --commission-max-change-rate="0.01" \
  --min-self-delegation="1" \
  --from=<key_name> \
  --fees=2000000ulumen
```

* NOTE: If you have troubles with \'\\\' symbol, run the command in a single line like `build/lumenxd tx staking create-validator --amount=1000000ulumen --pubkey=$(build/lumenxd tendermint show-validator) ...`

When specifying commission parameters, the `commission-max-change-rate` is used to measure % _point_ change over the `commission-rate`. E.g. 1% to 2% is a 100% rate increase, but only 1 percentage point.

`Min-self-delegation` is a strictly positive integer that represents the minimum amount of self-delegated voting power your validator must always have. A `min-self-delegation` of 1 means your validator will never have a self-delegation lower than `1000000ulumen`

You can check that you are in the validator set by using a third party explorer or using cli tool
```bash
build/lumenxd query staking validators --chain-id=<chain_id>
```

* Note: You can edit the params after, by running command `build/lumenxd tx staking edit-validator ... —from <key_name> --chain-id=<chain_id> --fees=2ulumen` with the necessary options

## How to init chain

Add key to your keyring
```build/lumenxd keys add <key name>```

Initialize genesis and config files.
```build/lumenxd init <moniker> --chain-id <chain id>```

Replace all denoms `stake` to `ulumen` in `genesis.json`

Add genesis account
```build/lumenxd add-genesis-account <key name> 200000000000ulumen``` - 200000lumen

Create genesis transaction
```build/lumenxd gentx <key name> 100000000000ulumen --chain-id <chain id>``` - create CreateValidator transaction

Collect all of gentxs
```build/lumenxd collect-gentxs```

Run network
```build/lumenxd start```


## Run a node as a background service

Create some necessary files
```
sudo mkdir -p /var/log/lumenxd
sudo touch /var/log/lumenxd/digitaloceand.log
sudo touch /var/log/lumenxd/lumenxd_error.log
sudo touch /etc/systemd/system/lumenxd.service
```

Edit systemd service file for myblockchaind.
```sudo nano /etc/systemd/system/lumenxd.service```

Add following configuration to it:
```
Description=lumenxd daemon
After=network-online.target
[Service]
User=ubuntu
ExecStart=/home/ubuntu/lumenx/build/lumenxd start --home=/home/ubuntu/.lumenx
WorkingDirectory=/home/ubuntu/go/bin
StandardOutput=file:/var/log/lumenxd/digitaloceand.log
StandardError=file:/var/log/lumenxd/digitaloceand_error.log
Restart=always
RestartSec=3
LimitNOFILE=4096
[Install]
WantedBy=multi-user.target
```

Enabled it to run all the time even after it reboots.
```
sudo systemctl enable lumenxd.service
#Start process
sudo systemctl start lumenxd.service
#Stop process
sudo systemctl stop lumenxd.service
#View logs
sudo journalctl -u lumenxd -f
```
