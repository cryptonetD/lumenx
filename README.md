LumenX is the blockchain built using the [Cosmos SDK](https://github.com/cosmos/cosmos-sdk). LumenX will be interact with other sovereign blockchains using a protocol called [IBC](https://github.com/cosmos/ics/tree/master/ibc) that enables Inter-Blockchain Communication.

# LumenX network

## Mainnet Full Node Quick Start
With each version of the LumenX, the chain is restarted from a new Genesis state.

Get mainnet config [here](https://github.com/metaprotocol-ai/lumenx/tree/master/config)

```
- Hardware requirements
CPU: 4 core or higher
RAM: 16 GB (32 GB recommended)
DISK Storage: SSD(NVME) 2 TB (minimum 1 TB)

- Software requirements
OS: Ubuntu Server 20.04
Go version: Go 1.17+
```

### Build from code

These instructions assume you are running Linux and have installed [Go 1.17+](https://golang.org/dl/).  This guide helps you:

Build, Install, and Name your Node
```bash
# Clone LumenX from the latest release found here: https://github.com/metaprotocol-ai/lumenx/releases
git clone https://github.com/metaprotocol-ai/lumenx
# Enter the folder LumenX was cloned into
cd lumenx
# Compile and install LumenX
make install
# Check LumenX version
lumenxd version
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
#### Moniker characters

Monikers can only contain ASCII characters; using Unicode characters will render your node unreachable by other peers in the network.

#### Update minimum gas prices

1. Open `~/.lumenx/config/app.toml`.

2. Modify `minimum-gas-prices` and set the minimum price of gas a validator will accept to validate a transaction and to prevent spam.

**Example**:

```toml
# The minimum gas prices a validator is willing to accept for processing a
# transaction. A transaction's fees must meet the minimum of any denomination
# specified in this config (e.g. 0.25token1;0.0001token2).
minimum-gas-prices = "0.0025ulumen"
```
or just download [app.toml](https://raw.githubusercontent.com/metaprotocol-ai/lumenx/master/config/app.toml)
```
wget -O $HOME/.lumenx/config/app.toml https://raw.githubusercontent.com/metaprotocol-ai/lumenx/master/config/app.toml
```

Your full node has been initialized!

#### Run a full node
```
# Start LumenX
lumenxd start
# Check your node's status with LumenX cli
lumenxd status
```
If you want to run lumenxd as a permanent background service, see [here](https://github.com/metaprotocol-ai/lumenx#run-a-node-as-a-background-service)

Your node is now syncing. This process will take a long time. Make sure you've set it up on a stable connection so you can leave while it syncs.

```
Sync start times

Nodes take at least an hour to start syncing. This wait is normal.
Before troubleshooting a sync, please wait an hour for the sync to start.
```

### Create a key
Add new
``` bash
lumenxd keys add <key_name>
```

Or import via mnemonic
```bash
lumenxd keys add <key_name> -i
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
lumenxd tendermint show-validator
```

To create your validator, just use the following command:

Check if your key(address) has enough balance:

```bash
lumenxd query bank balances <key address>
```

For the main node, `chain-id` is `LumenX`.\
You need transction fee `2000000ulumen`(2lumen) to make your transaction for creating validator.\
Don't use more `ulumen` than you have!

```bash
lumenxd tx staking create-validator \
  --amount=10000000ulumen \
  --pubkey=$(lumenxd tendermint show-validator) \
  --moniker=<choose a moniker> \
  --chain-id=<chain_id> \
  --commission-rate="0.10" \
  --commission-max-rate="0.20" \
  --commission-max-change-rate="0.01" \
  --min-self-delegation="1" \
  --from=<key_name> \
  --fees=2000000ulumen
```

* NOTE: If you have troubles with \'\\\' symbol, run the command in a single line like `lumenxd tx staking create-validator --amount=1000000ulumen --pubkey=$(lumenxd tendermint show-validator) ...`

When specifying commission parameters, the `commission-max-change-rate` is used to measure % _point_ change over the `commission-rate`. E.g. 1% to 2% is a 100% rate increase, but only 1 percentage point.

`Min-self-delegation` is a strictly positive integer that represents the minimum amount of self-delegated voting power your validator must always have. A `min-self-delegation` of 1 means your validator will never have a self-delegation lower than `1000000ulumen`

You can check that you are in the validator set by using a third party explorer or using cli tool
```bash
lumenxd query staking validators --chain-id=<chain_id>
```

* Note: You can edit the params after, by running command `lumenxd tx staking edit-validator ... —from <key_name> --chain-id=<chain_id> --fees=2ulumen` with the necessary options


#### Edit Validator Description
You can edit your validator's public description. This info is to identify your validator, and will be relied on by delegators to decide which validators to stake to. Make sure to provide input for every flag below. If a flag is not included in the command the field will default to empty (--moniker defaults to the machine name) if the field has never been set or remain the same if it has been set in the past.

```bash
lumenxd tx staking edit-validator
  --moniker=<choose a moniker> \
  --website=<your website url> \
  --identity=<keybase identity> \
  --details=<description> \
  --chain-id=<chain_id> \
  --from=<key_name> \
  --commission-rate=<commission rate>
```

#### Unique Flags

| Name, shorthand     | type   | Required | Default  | Description                                                         |
| --------------------| -----  | -------- | -------- | ------------------------------------------------------------------- |
| --commission-rate   | string | float    | 0.0      | Commission rate percentage |
| --moniker           | string | false    | ""       | Validator name |
| --identity          | string | false    | ""       | Optional identity signature (ex. UPort or Keybase) |
| --website           | string | false    | ""       | Optional website  |
| --details           | string | false    | ""       | Optional details |


#### Examples
```bash
lumenxd tx staking edit-validator --commission-rate 0.10 --from <key_name>
```

#### Put a thumbnail
Create a [Keybase Account](https://keybase.io/) follow the Keybase instructions to set up a PGP key, and upload a profile picture.
And link your Keybase profile to your validator.
```bash
lumenxd tx staking edit-validator --identity="6A0D65E29A4CBC8E" --from <key_name>
```

#### Unjail a validator
The unjail command allows users to unjail a validator previously jailed for downtime.\
You need transction fee `500ulumen` to make your transaction.
```bash
lumenxd tx slashing unjail --chain-id=LumenX --fees=500ulumen --from <key_name>
```

## How to init chain

Add key to your keyring
```lumenxd keys add <key_name>```

Initialize genesis and config files.
```lumenxd init <moniker> --chain-id <chain_id>```

Replace all denoms `stake` to `ulumen` in `genesis.json`

Add genesis account
```lumenxd add-genesis-account <key_name> 200000000000ulumen``` - 200000lumen

Create genesis transaction
```lumenxd gentx <key_name> 100000000000ulumen --chain-id <chain_id>``` - create CreateValidator transaction

Collect all of gentxs
```lumenxd collect-gentxs```

Run network
```lumenxd start```


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
ExecStart=/home/ubuntu/go/bin/lumenxd start --home=/home/ubuntu/.lumenx
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
#enable service
sudo systemctl enable lumenxd.service
#Start process
sudo systemctl start lumenxd.service
#Stop process
sudo systemctl stop lumenxd.service
#View logs
sudo journalctl -u lumenxd -f
```
