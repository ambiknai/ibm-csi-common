/**
 * Copyright 2021 IBM Corp.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

// Package mountmanager ...
package mountmanager

import (
	mount "k8s.io/mount-utils"
	exec "k8s.io/utils/exec"
	testExec "k8s.io/utils/exec/testing"
)

// NewNodeMounter ...
func NewFakeNodeMounter() Mounter {
	fakesafemounter := NewFakeSafeMounter()
	return &NodeMounter{fakesafemounter}
}

// NewFakeSafeMounter ...
func NewFakeSafeMounter() *mount.SafeFormatAndMount {
	fakeMounter := &mount.FakeMounter{MountPoints: []mount.MountPoint{{
		Device: "valid-devicePath",
		Path:   "valid-vol-path",
		Type:   "ext4",
		Opts:   []string{"defaults"},
		Freq:   1,
		Pass:   2,
	}},
	}

	var fakeExec exec.Interface = &testExec.FakeExec{
		CommandScript: []testExec.FakeCommandAction{
			func(cmd string, args ...string) exec.Cmd {
				return nil
			},
		},
	}

	return &mount.SafeFormatAndMount{
		Interface: fakeMounter,
		Exec:      fakeExec,
	}
}
