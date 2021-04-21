/*
Copyright © 2021 Glenn M. Lewis

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

package dart

// Cursor represents an editor cursor and is used to advance through
// the Dart source code.
type Cursor struct {
	c         *Class
	absOffset int
	lineIndex int
	relOffset int
}

// advanceUntil searches for the provided rune, stepping through
// the Dart source code (and obeying its grammatical nuances) until found.
func (c *Cursor) advanceUntil(searchFor rune) {

}
