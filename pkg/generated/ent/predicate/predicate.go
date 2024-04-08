/*
Copyright YEAR 1block.ai.

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

package predicate

import (
	"entgo.io/ent/dialect/sql"
)

// Chat is the predicate function for chat builders.
type Chat func(*sql.Selector)

// Modelfile is the predicate function for modelfile builders.
type Modelfile func(*sql.Selector)

// Setting is the predicate function for setting builders.
type Setting func(*sql.Selector)

// User is the predicate function for user builders.
type User func(*sql.Selector)
