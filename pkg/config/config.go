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

package config

import (
	"os"
	"strings"

	"github.com/spf13/pflag"

	"github.com/spf13/viper"
)

var defaultConfig = map[string]interface{}{
	"server": map[string]interface{}{
		"address": "127.0.0.1",
		"port":    "8080",
	},
	"fabric": map[string]interface{}{
		"configpath": "./sdk_config.yaml",
		"user":       "admin",
	},
}

var name = ""

// ReadConfig reads configuration from file or env
func ReadConfig(envPrefix string) error {
	var err error

	for key, value := range defaultConfig {
		viper.SetDefault(key, value)
	}

	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	viper.SetEnvPrefix(envPrefix)
	viper.AutomaticEnv()

	flagset := pflag.NewFlagSet("Fabric SDK Server", pflag.ExitOnError)
	addPFlags(defaultConfig, flagset)
	flagset.Parse(os.Args)
	viper.BindPFlags(flagset)

	return err
}

func addPFlags(value interface{}, flagset *pflag.FlagSet) {
	switch value.(type) {
	case map[string]interface{}:
		for n, v := range value.(map[string]interface{}) {
			name = name + "." + n
			addPFlags(v, flagset)
		}
	default:
		flagset.String(name[1:], "", "")
		removeLast(&name)
		return
	}
	removeLast(&name)
}

func removeLast(input *string) {
	items := strings.Split(*input, ".")
	if len(items) > 0 {
		items = items[:len(items)-1]
	}
	*input = strings.Join(items, ".")
}
