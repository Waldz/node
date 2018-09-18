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

package dto

import (
	"testing"

	"github.com/mysteriumnetwork/node/identity"
	"github.com/stretchr/testify/assert"
)

var (
	providerID      = identity.FromAddress("123456")
	providerContact = Contact{
		Type: "type1",
	}
)

func Test_NewServiceProposal(t *testing.T) {
	proposal := ServiceProposal{ID: 123, ProviderID: "123"}
	UpdateProposal(&proposal, providerID, providerContact)

	assert.Exactly(
		t,
		ServiceProposal{
			ID:               1,
			Format:           "service-proposal/v1",
			ProviderID:       providerID.Address,
			ProviderContacts: []Contact{providerContact},
		},
		proposal,
	)
}
