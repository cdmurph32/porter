/*
 * (c) 2016-2018 Adobe. All rights reserved.
 * This file is licensed to you under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License. You may obtain a copy
 * of the License at http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software distributed under
 * the License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR REPRESENTATIONS
 * OF ANY KIND, either express or implied. See the License for the specific language
 * governing permissions and limitations under the License.
 */
package build

import (
	"os/exec"

	"github.com/adobe-platform/porter/constants"
	"github.com/phylake/go-cli"
)

type CleanCmd struct{}

func (recv *CleanCmd) Name() string {
	return "clean"
}

func (recv *CleanCmd) ShortHelp() string {
	return "Remove porter temporary files"
}

func (recv *CleanCmd) LongHelp() string {
	return ""
}

func (recv *CleanCmd) SubCommands() []cli.Command {
	return nil
}

func (recv *CleanCmd) Execute(args []string) bool {
	err := exec.Command("rm", "-fr", constants.TempDir).Run()
	if err != nil {
		panic(err)
	}
	return true
}
