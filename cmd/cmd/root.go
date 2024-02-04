/*
* Copyright (c) 2019, WSO2 Inc. (http://www.wso2.org) All Rights Reserved.
*
* WSO2 Inc. licenses this file to you under the Apache License,
* Version 2.0 (the "License"); you may not use this file except
* in compliance with the License.
* You may obtain a copy of the License at
*
*    http://www.apache.org/licenses/LICENSE-2.0
*
* Unless required by applicable law or agreed to in writing,
* software distributed under the License is distributed on an
* "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
* KIND, either express or implied. See the License for the
* specific language governing permissions and limitations
* under the License.
 */

package cmd

import (
	"github.com/lithammer/dedent"
	"github.com/spf13/cobra"
	"github.com/wso2/product-mi-tooling/cmd/utils"
	"os"
	"time"
)

var cfgFile string
var verbose bool

var programName = os.Args[0]

var rootCmdShortDesc = "CLI for Micro Integrator"

var rootCmdLongDesc = dedent.Dedent(`
        ` + programName + ` is a Command Line Tool for Management of WSO2 Micro Integrator
        `)

// RootCmd represents the base command when called without any subcommands
var RootCmd = &cobra.Command{
	Use:   programName,
	Short: rootCmdShortDesc,
	Long:  rootCmdLongDesc,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the RootCmd.
func Execute() {
	if err := RootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	RootCmd.PersistentFlags().BoolVarP(&verbose, "verbose", "v", false, "Enable verbose mode")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {

	if verbose {
		utils.IsVerbose = true
		utils.EnableVerboseMode()
		t := time.Now()
		utils.Logf(utils.LogPrefixInfo+"Executed ManagementCLI (%s) on %v\n", programName, t.Format(time.RFC1123))
	} else {
		utils.IsVerbose = false
	}
}
