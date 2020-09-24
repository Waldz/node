/*
 * Copyright (C) 2018 The "MysteriumNetwork/node" Authors.
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

package node

import (
	"path"

	"github.com/mysteriumnetwork/node/config"
	"github.com/mysteriumnetwork/node/core/port"
	"github.com/mysteriumnetwork/node/logconfig"
	openvpn_core "github.com/mysteriumnetwork/node/services/openvpn/core"
	"github.com/mysteriumnetwork/node/services/wireguard/resources"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

// Openvpn interface is abstraction over real openvpn options to unblock mobile development
// will disappear as soon as go-openvpn will unify common factory for openvpn creation
type Openvpn interface {
	Check() error
	BinaryPath() string
}

// TODO this struct will disappear when we unify go-openvpn embedded lib and external process based session creation/handling
type wrapper struct {
	nodeOptions openvpn_core.NodeOptions
}

func (w wrapper) Check() error {
	return w.nodeOptions.Check()
}

func (w wrapper) BinaryPath() string {
	return w.nodeOptions.BinaryPath
}

var _ Openvpn = wrapper{}

// Options describes options which are required to start Node
type Options struct {
	Directories OptionsDirectory

	TequilapiAddress string
	TequilapiPort    int
	TequilapiEnabled bool
	BindAddress      string
	UI               OptionsUI
	FeedbackURL      string

	Keystore OptionsKeystore

	logconfig.LogOptions
	OptionsNetwork
	Discovery  OptionsDiscovery
	Quality    OptionsQuality
	Location   OptionsLocation
	Transactor OptionsTransactor
	Hermes     OptionsHermes

	Openvpn  Openvpn
	Firewall OptionsFirewall

	Payments OptionsPayments

	Consumer bool

	P2PPorts *port.Range
}

// GetOptions retrieves node options from the app configuration.
func GetOptions() *Options {
	network := OptionsNetwork{
		Testnet:               config.GetBool(config.FlagTestnet),
		Localnet:              config.GetBool(config.FlagLocalnet),
		Betanet:               config.GetBool(config.FlagBetanet),
		ExperimentNATPunching: config.GetBool(config.FlagNATPunching),
		MysteriumAPIAddress:   config.GetString(config.FlagAPIAddress),
		BrokerAddress:         config.GetString(config.FlagBrokerAddress),
		EtherClientRPC:        config.GetString(config.FlagEtherRPC),
	}
	return &Options{
		Directories:      *GetOptionsDirectory(&network),
		TequilapiAddress: config.GetString(config.FlagTequilapiAddress),
		TequilapiPort:    config.GetInt(config.FlagTequilapiPort),
		TequilapiEnabled: true,
		BindAddress:      config.GetString(config.FlagBindAddress),
		UI: OptionsUI{
			UIEnabled:     config.GetBool(config.FlagUIEnable),
			UIBindAddress: config.GetString(config.FlagUIAddress),
			UIPort:        config.GetInt(config.FlagUIPort),
		},
		FeedbackURL: config.GetString(config.FlagFeedbackURL),
		Keystore: OptionsKeystore{
			UseLightweight: config.GetBool(config.FlagKeystoreLightweight),
		},
		LogOptions:     *GetLogOptions(),
		OptionsNetwork: network,
		Discovery:      *GetDiscoveryOptions(),
		Quality: OptionsQuality{
			Type:    QualityType(config.GetString(config.FlagQualityType)),
			Address: config.GetString(config.FlagQualityAddress),
		},
		Location: OptionsLocation{
			IPDetectorURL: config.GetString(config.FlagIPDetectorURL),
			Type:          LocationType(config.GetString(config.FlagLocationType)),
			Address:       config.GetString(config.FlagLocationAddress),
			Country:       config.GetString(config.FlagLocationCountry),
			City:          config.GetString(config.FlagLocationCity),
			NodeType:      config.GetString(config.FlagLocationNodeType),
		},
		Transactor: OptionsTransactor{
			TransactorEndpointAddress:       config.GetString(config.FlagTransactorAddress),
			RegistryAddress:                 config.GetString(config.FlagTransactorRegistryAddress),
			ChannelImplementation:           config.GetString(config.FlagTransactorChannelImplementation),
			ProviderMaxRegistrationAttempts: config.GetInt(config.FlagTransactorProviderMaxRegistrationAttempts),
			ProviderRegistrationRetryDelay:  config.GetDuration(config.FlagTransactorProviderRegistrationRetryDelay),
			ProviderRegistrationStake:       config.GetBigInt(config.FlagTransactorProviderRegistrationStake),
		},
		Payments: OptionsPayments{
			MaxAllowedPaymentPercentile:    config.GetInt(config.FlagPaymentsMaxHermesFee),
			BCTimeout:                      config.GetDuration(config.FlagPaymentsBCTimeout),
			HermesPromiseSettlingThreshold: config.GetFloat64(config.FlagPaymentsHermesPromiseSettleThreshold),
			SettlementTimeout:              config.GetDuration(config.FlagPaymentsHermesPromiseSettleTimeout),
			MystSCAddress:                  config.GetString(config.FlagPaymentsMystSCAddress),
			ConsumerUpperGBPriceBound:      config.GetBigInt(config.FlagPaymentsConsumerPricePerGBUpperBound),
			ConsumerLowerGBPriceBound:      config.GetBigInt(config.FlagPaymentsConsumerPricePerGBLowerBound),
			ConsumerUpperMinutePriceBound:  config.GetBigInt(config.FlagPaymentsConsumerPricePerMinuteUpperBound),
			ConsumerLowerMinutePriceBound:  config.GetBigInt(config.FlagPaymentsConsumerPricePerMinuteLowerBound),
			ConsumerDataLeewayMegabytes:    config.GetUInt64(config.FlagPaymentsConsumerDataLeewayMegabytes),
			ProviderInvoiceFrequency:       config.GetDuration(config.FlagPaymentsProviderInvoiceFrequency),
			MaxUnpaidInvoiceValue:          config.GetBigInt(config.FlagPaymentsMaxUnpaidInvoiceValue),
		},
		Hermes: OptionsHermes{
			HermesID: config.GetString(config.FlagHermesID),
		},
		Openvpn: wrapper{nodeOptions: openvpn_core.NodeOptions{
			BinaryPath: config.GetString(config.FlagOpenvpnBinary),
		}},
		Firewall: OptionsFirewall{
			BlockAlways: config.GetBool(config.FlagFirewallKillSwitch),
		},
		P2PPorts: getP2PListenPorts(),
		Consumer: config.GetBool(config.FlagConsumer),
	}
}

// GetLogOptions retrieves logger options from the app configuration.
func GetLogOptions() *logconfig.LogOptions {
	filepath := ""
	if logDir := config.GetString(config.FlagLogDir); logDir != "" {
		filepath = path.Join(logDir, "mysterium-node")
	}
	level, err := zerolog.ParseLevel(config.GetString(config.FlagLogLevel))
	if err != nil {
		log.Error().Err(err).Msg("Failed to parse logging level")
		level = zerolog.DebugLevel
	}
	return &logconfig.LogOptions{
		LogLevel: level,
		LogHTTP:  config.GetBool(config.FlagLogHTTP),
		Filepath: filepath,
	}
}

// GetDiscoveryOptions retrieves discovery options from the app configuration.
func GetDiscoveryOptions() *OptionsDiscovery {
	typeValues := config.GetStringSlice(config.FlagDiscoveryType)
	types := make([]DiscoveryType, len(typeValues))
	for i, typeValue := range typeValues {
		types[i] = DiscoveryType(typeValue)
	}

	return &OptionsDiscovery{
		Types:         types,
		PingInterval:  config.GetDuration(config.FlagDiscoveryPingInterval),
		FetchEnabled:  true,
		FetchInterval: config.GetDuration(config.FlagDiscoveryFetchInterval),
	}
}

// OptionsKeystore stores the keystore configuration
type OptionsKeystore struct {
	UseLightweight bool
}

func getP2PListenPorts() *port.Range {
	p2pPortRange, err := port.ParseRange(config.GetString(config.FlagP2PListenPorts))
	if err != nil {
		log.Warn().Err(err).Msg("Failed to parse p2p listen port range, using default value")
		p2pPortRange = port.UnspecifiedRange()
	}
	if p2pPortRange.Capacity() > resources.MaxConnections {
		log.Warn().Msgf("Specified p2p port range exceeds maximum number of connections allowed for the platform (%d), "+
			"using default value", resources.MaxConnections)
		p2pPortRange = port.UnspecifiedRange()
	}
	return p2pPortRange
}
