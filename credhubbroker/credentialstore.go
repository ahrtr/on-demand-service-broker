// Copyright (C) 2015-Present Pivotal Software, Inc. All rights reserved.

// This program and the accompanying materials are made available under
// the terms of the under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at

// http://www.apache.org/licenses/LICENSE-2.0

// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package credhubbroker

import "code.cloudfoundry.org/credhub-cli/credhub/permissions"

//go:generate counterfeiter -o fakes/credentialstore.go . CredentialStore
type CredentialStore interface {
	Set(key string, value interface{}) error
	Delete(key string) error
	AddPermissions(credentialName string, perms []permissions.Permission) ([]permissions.Permission, error)
}
