/*
Copyright 2016 The Kubernetes Authors All rights reserved.

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

package main

import (
	"io"
	"testing"

	"github.com/spf13/cobra"
)

func TestGetCmd(t *testing.T) {
	tests := []releaseCase{
		{
			name:     "get with a release",
			resp:     releaseMock(&releaseOptions{name: "thomas-guide"}),
			args:     []string{"thomas-guide"},
			expected: "REVISION: 1\nRELEASED: (.*)\nCHART: foo-0.1.0-beta.1\nUSER-SUPPLIED VALUES:\nname: \"value\"\nCOMPUTED VALUES:\nname: value\n\nHOOKS:\n---\n# pre-install-hook\n" + mockHookTemplate + "\nMANIFEST:",
		},
		{
			name: "get requires release name arg",
			err:  true,
		},
	}

	cmd := func(c *fakeReleaseClient, out io.Writer) *cobra.Command {
		return newGetCmd(c, out)
	}
	runReleaseCases(t, tests, cmd)
}
