// Copyright 2015 Google Inc. All Rights Reserved.
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

package fs

import "github.com/jacobsa/gcloud/syncutil"

type DirInode struct {
	/////////////////////////
	// Constant data
	/////////////////////////

	// The name of the GCS object backing the inode. Special case: the empty
	// string means this is the root inode.
	//
	// INVARIANT: name == "" || name[len(name)-1] == '/'
	name string

	/////////////////////////
	// Mutable state
	/////////////////////////

	// A mutex that must be held when calling certain methods. See documentation
	// for each method.
	Mu syncutil.InvariantMutex
}
