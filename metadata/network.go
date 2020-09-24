/*
 * Copyright (C) 2017 The "MysteriumNetwork/node" Authors.
 *
 * This program is free software: you can redistribute it and/or modify
 * it under the terms of the GNU General Public License as published by
 * the Free Software Foundation, either version 3 of the License, or
 * (at your option) any later version.
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU General Public License for more details.
 *
 * You should have received a copy of the GNU General Public License
 * along with this program.  If not, see <http://www.gnu.org/licenses/>.
 */

package metadata

// NetworkDefinition structure holds all parameters which describe particular network
type NetworkDefinition struct {
	MysteriumAPIAddress       string
	AccessPolicyOracleAddress string
	BrokerAddress             string
	EtherClientRPC            string
	TransactorAddress         string
	RegistryAddress           string
	HermesID                  string
	ChannelImplAddress        string
	MMNAddress                string
	MMNAPIAddress             string
}

// TestnetDefinition defines parameters for test network (currently default network)
var TestnetDefinition = NetworkDefinition{
	MysteriumAPIAddress:       "https://testnet-api.mysterium.network/v1",
	AccessPolicyOracleAddress: "https://testnet-trust.mysterium.network/api/v1/access-policies/",
	BrokerAddress:             "nats://testnet-broker.mysterium.network",
	EtherClientRPC:            "wss://goerli.infura.io/ws/v3/c2c7da73fcc84ec5885a7bb0eb3c3637",
	TransactorAddress:         "https://testnet-transactor.mysterium.network/api/v1",
	RegistryAddress:           "0x3dD81545F3149538EdCb6691A4FfEE1898Bd2ef0",
	ChannelImplAddress:        "0x3026eB9622e2C5bdC157C6b117F7f4aC2C2Db3b5",
	HermesID:                  "0x0214281cf15C1a66b51990e2E65e1f7b7C363318",
	MMNAddress:                "https://my.mysterium.network/",
	MMNAPIAddress:             "https://my.mysterium.network/api/v1",
}

// BetanetDefinition defines parameters for Betanet network (currently default network)
var BetanetDefinition = NetworkDefinition{
	MysteriumAPIAddress:       "https://betanet-api.mysterium.network/v1",
	AccessPolicyOracleAddress: "https://betanet-trust.mysterium.network/api/v1/access-policies/",
	BrokerAddress:             "nats://betanet-broker.mysterium.network",
	EtherClientRPC:            "wss://goerli.infura.io/ws/v3/c2c7da73fcc84ec5885a7bb0eb3c3637",
	TransactorAddress:         "https://betanet-transactor.mysterium.network/api/v1",
	RegistryAddress:           "0xc82Cc5B0bAe95F443e33FF053aAa70F1Eb7d312A",
	ChannelImplAddress:        "0x29a615aA7E03D8c04B24cc91B2949447D3A10bD6",
	HermesID:                  "0x42a537D649d6853C0a866470f2d084DA0f73b5E4",
	MMNAddress:                "https://betanet.mysterium.network/",
	MMNAPIAddress:             "https://betanet.mysterium.network/api/v1",
}

// LocalnetDefinition defines parameters for local network
// Expects discovery, broker and morqa services on localhost
var LocalnetDefinition = NetworkDefinition{
	MysteriumAPIAddress:       "http://localhost:8001/v1",
	AccessPolicyOracleAddress: "https://localhost:8081/api/v1/access-policies/",
	BrokerAddress:             "localhost",
	EtherClientRPC:            "http://localhost:8545",
	MMNAddress:                "http://localhost/",
	MMNAPIAddress:             "http://localhost/api/v1",
}

// DefaultNetwork defines default network values when no runtime parameters are given
var DefaultNetwork = BetanetDefinition
