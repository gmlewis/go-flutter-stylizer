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

// Package dart implements methods for manipulating Dart code.
package dart

// Options represents the configuration options for the Dart processor.
type Options struct {
	Diff  bool
	List  bool
	Write bool
}

// Client represents a Dart processor.
type Client struct {
	opts Options
}

// New returns a new Dart processor.
func New(opts Options) *Client {
	return &Client{opts: opts}
}

// StylizeFile sylizes a single Dart file using the provided options.
func (c *Client) StylizeFile(filename string) error {
	return nil
}
