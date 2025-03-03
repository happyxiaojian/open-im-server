// Copyright © 2023 OpenIM. All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"github.com/openimsdk/open-im-server/v3/internal/msggateway"
	config2 "github.com/openimsdk/open-im-server/v3/pkg/common/config"

	//"github.com/openimsdk/open-im-server/internal/msggateway".
	"github.com/spf13/cobra"

	"github.com/OpenIMSDK/protocol/constant"
)

type MsgGatewayCmd struct {
	*RootCmd
}

func NewMsgGatewayCmd() *MsgGatewayCmd {
	ret := &MsgGatewayCmd{NewRootCmd("msgGateway")}
	ret.SetRootCmdPt(ret)
	return ret
}

func (m *MsgGatewayCmd) AddWsPortFlag() {
	m.Command.Flags().IntP(constant.FlagWsPort, "w", 0, "ws server listen port")
}

func (m *MsgGatewayCmd) getWsPortFlag(cmd *cobra.Command) int {
	port, _ := cmd.Flags().GetInt(constant.FlagWsPort)
	if port == 0 {
		port = m.PortFromConfig(constant.FlagWsPort)
	}
	return port
}

func (m *MsgGatewayCmd) addRunE() {
	m.Command.RunE = func(cmd *cobra.Command, args []string) error {
		return msggateway.RunWsAndServer(m.getPortFlag(cmd), m.getWsPortFlag(cmd), m.getPrometheusPortFlag(cmd))
	}
}

func (m *MsgGatewayCmd) Exec() error {
	m.addRunE()
	return m.Execute()
}
func (m *MsgGatewayCmd) GetPortFromConfig(portType string) int {
	if portType == constant.FlagWsPort {
		return config2.Config.LongConnSvr.OpenImWsPort[0]
	} else if portType == constant.FlagPort {
		return config2.Config.LongConnSvr.OpenImMessageGatewayPort[0]
	} else if portType == constant.FlagPrometheusPort {
		return 0
	} else {
		return 0
	}
}
