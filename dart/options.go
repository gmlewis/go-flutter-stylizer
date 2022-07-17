/*
Copyright © 2022 Glenn M. Lewis

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

import "log"

var defaultMemberOrdering = []string{
	"public-constructor",
	"named-constructors",
	"public-static-variables",
	"public-instance-variables",
	"public-override-variables",
	"private-static-variables",
	"private-instance-variables",
	"public-override-methods",
	"public-other-methods",
	"private-other-methods", // optional!
	"build-method",
}

var allOptions = []string{
	"public-constructor",
	"named-constructors",
	"public-static-variables",
	"public-static-properties",
	"public-instance-variables",
	"public-instance-properties",
	"public-override-variables",
	"public-override-properties",
	"private-static-variables",
	"private-static-properties",
	"private-instance-variables",
	"private-instance-properties",
	"public-override-methods",
	"public-other-methods",
	"private-other-methods",
	"operators",
	"build-method",
}

func (o Options) validate() Options {
	if !validateMemberOrdering(o.MemberOrdering) {
		o.MemberOrdering = defaultMemberOrdering
	}

	var hasProperties bool
	for _, v := range o.MemberOrdering {
		switch v {
		case "public-static-properties":
			o.hasPublicStaticProperties = true
			hasProperties = true
		case "public-instance-properties":
			o.hasPublicInstanceProperties = true
			hasProperties = true
		case "public-override-properties":
			o.hasPublicOverrideProperties = true
			hasProperties = true
		case "private-static-properties":
			o.hasPrivateStaticProperties = true
			hasProperties = true
		case "private-instance-properties":
			o.hasPrivateInstanceProperties = true
			hasProperties = true
		case "private-other-methods":
			o.hasPrivateOtherMethods = true
		case "operators":
			o.hasOperators = true
		}
	}

	if o.GroupAndSortGetterMethods && hasProperties {
		log.Printf("WARNING: newer *-properties ordering overrides groupAndSortGetterMethods.")
		o.GroupAndSortGetterMethods = false
	}

	return o
}

func validateMemberOrdering(memberOrdering []string) bool {
	if got, want := len(memberOrdering), len(defaultMemberOrdering); got < want-1 {
		log.Printf("flutterStylizer.memberOrdering must have at least %v values, but found %v. Ignoring and using defaults.", want-1, got)
		return false
	}

	lookup := map[string]bool{}
	for _, s := range allOptions {
		lookup[s] = true
	}

	seen := map[string]bool{}
	for _, el := range memberOrdering {
		if !lookup[el] {
			log.Printf("Unknown member %q in flutterStylizer.memberOrdering. Ignoring and using defaults.", el)
			return false
		}
		if seen[el] {
			log.Printf("Duplicate member %q in flutterStylizer.memberOrdering. Ignoring and using defaults.", el)
			return false
		}
		seen[el] = true
	}

	return true
}
