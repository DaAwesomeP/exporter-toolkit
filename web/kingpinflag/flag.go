// Copyright 2020 The Prometheus Authors
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
package kingpinflag

import (
	"gopkg.in/alecthomas/kingpin.v2"
)

type FlagStruct struct {
    WebListenAddresses *[]string
    WebSystemdSocket *bool
    WebConfigFile *string
}

// AddFlags adds the flags used by this package to the Kingpin application.
// To use the default Kingpin application, call AddFlags(kingpin.CommandLine)
func AddFlags(a *kingpin.Application) *FlagStruct {
	flags := FlagStruct{
		WebListenAddresses: a.Flag(
			"web.listen-address",
			"Addresses on which to expose metrics and web interface.",
		).Default(":9100").Strings(),
		WebSystemdSocket: kingpin.Flag(
			"web.systemd-socket",
			"Use systemd socket activation listeners instead of port listeners.",
		).Bool(),
		WebConfigFile: a.Flag(
			"web.config.file",
			"[EXPERIMENTAL] Path to configuration file that can enable TLS or authentication.",
		).Default("").String(),
	}
	return &flags
}
