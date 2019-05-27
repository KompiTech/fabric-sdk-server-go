/*
  Copyright 2019 KompiTech GmbH

  Licensed under the Apache License, Version 2.0 (the "License");
  you may not use this file except in compliance with the License.
  You may obtain a copy of the License at

      http://www.apache.org/licenses/LICENSE-2.0

  Unless required by applicable law or agreed to in writing, software
  distributed under the License is distributed on an "AS IS" BASIS,
  WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
  See the License for the specific language governing permissions and
  limitations under the License.
*/

package fabric

import (
	"github.com/hyperledger/fabric-sdk-go/pkg/client/channel"
	"github.com/hyperledger/fabric-sdk-go/pkg/core/config"
	"github.com/hyperledger/fabric-sdk-go/pkg/fabsdk"
)

// NewChannelClient returns client for specified channel
func NewChannelClient(channelName, configPath, user string) (*channel.Client, func(), error) {
	sdk, err := fabsdk.New(config.FromFile(configPath))
	if err != nil {
		return nil, nil, err
	}

	channelContext := sdk.ChannelContext(channelName, fabsdk.WithUser(user))

	channelClient, err := channel.New(channelContext)
	if err != nil {
		sdk.Close()
		return nil, nil, err
	}
	return channelClient, sdk.Close, nil
}
