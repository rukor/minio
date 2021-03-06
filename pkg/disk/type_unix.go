// +build freebsd

/*
 * Minio Cloud Storage, (C) 2015 Minio, Inc.
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

package disk

import (
	"strconv"
	"syscall"
)

// fsType2StringMap - list of filesystems supported by donut on linux
var fsType2StringMap = map[string]string{
	"35": "UFS",
}

// getFSType returns the filesystem type of the underlying mounted filesystem
func getFSType(path string) (string, error) {
	s := syscall.Statfs_t{}
	err := syscall.Statfs(path, &s)
	if err != nil {
		return "", err
	}
	fsTypeHex := strconv.FormatInt(int64(s.Type), 16)
	fsTypeString, ok := fsType2StringMap[fsTypeHex]
	if ok == false {
		return "UNKNOWN", nil
	}
	return fsTypeString, nil
}
