/*
 * Copyright (C) 2019 The "MysteriumNetwork/node" Authors.
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

package release

import (
	log "github.com/cihub/seelog"
	cenv "github.com/mysteriumnetwork/go-ci/env"
	"github.com/mysteriumnetwork/go-ci/shell"
	"github.com/mysteriumnetwork/node/ci/env"
)

func releaseDebianPPA(repository, version, buildNumber string) error {
	err := shell.NewCmd("mkdir -p ~/.ssh").Run()
	if err != nil {
		return err
	}
	err = shell.NewCmd("chmod 0700 ~/.ssh").Run()
	if err != nil {
		return err
	}
	err = shell.NewCmdf("cp -f %s ~/.ssh/id_rsa", cenv.Str(env.LaunchpadSSHKey)).Run()
	if err != nil {
		return err
	}
	err = shell.NewCmd("chmod 600 ~/.ssh/id_rsa").Run()
	if err != nil {
		return err
	}
	err = shell.NewCmdf("bin/release_ppa %s %s %s %s", repository, version, buildNumber, "xenial").Run()
	if err != nil {
		return err
	}
	err = shell.NewCmdf("bin/release_ppa %s %s %s %s", repository, version, buildNumber, "bionic").Run()
	if err != nil {
		return err
	}
	return nil
}

// ReleaseDebianPPASnapshot releases to node-dev PPA
func ReleaseDebianPPASnapshot() error {
	err := cenv.EnsureEnvVars(
		cenv.SnapshotBuild,
		cenv.BuildVersion,
		cenv.BuildNumber,
		env.LaunchpadSSHKey,
	)
	if err != nil {
		return err
	}
	// TODO uncomment after testing
	//if !cenv.Bool(cenv.SnapshotBuild) {
	//	log.Info("not a snapshot build, skipping ReleaseDebianPPASnapshot action...")
	//	return nil
	//}
	return releaseDebianPPA("node-dev", cenv.Str(cenv.BuildVersion), cenv.Str(cenv.BuildNumber))
}

// ReleaseDebianPPAPreRelease releases to node-pre PPA (which is then manually promoted to node PPA)
func ReleaseDebianPPAPreRelease() error {
	err := cenv.EnsureEnvVars(
		cenv.TagBuild,
		cenv.BuildVersion,
		cenv.BuildNumber,
		env.LaunchpadSSHKey,
	)
	if err != nil {
		return err
	}
	if !cenv.Bool(cenv.TagBuild) {
		log.Info("not a tag build, skipping ReleaseDebianPPAPreRelease action...")
		return nil
	}
	return releaseDebianPPA("node-pre", cenv.Str(cenv.BuildVersion), cenv.Str(cenv.BuildNumber))
}
