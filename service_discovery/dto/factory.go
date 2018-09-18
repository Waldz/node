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
	"github.com/mysteriumnetwork/node/identity"
)

// UpdateProposal updates service proposal description with general data
func UpdateProposal(proposal *ServiceProposal, providerID identity.Identity, providerContact Contact) {
	proposal.Format = "service-proposal/v1"
	proposal.ID = 1
	proposal.ProviderID = providerID.Address
	proposal.ProviderContacts = []Contact{providerContact}
}
