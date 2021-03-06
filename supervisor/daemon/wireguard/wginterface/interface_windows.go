/*
 * Copyright (C) 2020 The "MysteriumNetwork/node" Authors.
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

package wginterface

import (
	"errors"
)

// New creates new WgInterface instance.
func New(requestedInterfaceName string, uid string) (*WgInterface, error) {
	return nil, errors.New("not implemented")
}

// Listen listens for WireGuard configuration changes via user space socket.
func (a *WgInterface) Listen() {

}

// Down closes device and user space api socket.
func (a *WgInterface) Down() {

}
