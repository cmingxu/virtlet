// +build linux

/*
Copyright 2018 Mirantis

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

package utils

import "syscall"

type mounter struct{}

var _ Mounter = &mounter{}

// NewMounter creates linux mounter struct
func NewMounter() Mounter {
	return &mounter{}
}

func (mounter *mounter) Mount(source string, target string, fstype string, flags uintptr) error {
	return syscall.Mount(source, target, fstype, flags, "")
}

func (mounter *mounter) Unmount(target string, flags int) error {
	return syscall.Unmount(target, flags)
}
