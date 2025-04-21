#!/bin/bash
cd $FABRIC_CFG_PATH
# cryptogen generate --config crypto-config.yaml --output keyfiles
configtxgen -profile OrdererGenesis -outputBlock genesis.block -channelID systemchannel

configtxgen -printOrg dealer-auto-com > JoinRequest_dealer-auto-com.json
configtxgen -printOrg mvd-auto-com > JoinRequest_mvd-auto-com.json
configtxgen -printOrg owner-auto-com > JoinRequest_owner-auto-com.json
