// Copyright 2016-2021, Pulumi Corporation.
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

package provider

import (
	"github.com/pkg/errors"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi/provider"
)

var resourceConstructorMap = map[string]ResourceConstructor{
	AccountIdentifier:                       createNewResourceConstructor(NewIAMAccount),
	PolicyIdentifier:                        createNewResourceConstructor(NewPolicy),
	AssumableRoleWithOIDCIdentifier:         createNewResourceConstructor(NewIAMAssumableRoleWithOIDC),
	AssumableRoleWithSAMLIdentifier:         createNewResourceConstructor(NewAssumableRoleWithSAML),
	AssumableRoleIdentifier:                 createNewResourceConstructor(NewAssumableRole),
	AssumableRolesWithSAMLIdentifier:        createNewResourceConstructor(NewAssumableRolesWithSAML),
	AssumableRolesIdentifier:                createNewResourceConstructor(NewAssumableRoles),
	EKSRoleIdentifier:                       createNewResourceConstructor(NewEKSRole),
	GroupWithAssumableRolesPolicyIdentifier: createNewResourceConstructor(NewGroupWithAssumableRolesPolicy),
	GroupWithPoliciesIdentifier:             createNewResourceConstructor(NewGroupWithPolicies),
	ReadOnlyPolicyIdentifier:                createNewResourceConstructor(NewReadOnlyPolicy),
	RoleForServiceAccountsEksIdentifier:     createNewResourceConstructor(NewRoleForServiceAccountsEks),
	UserIdentifier:                          createNewResourceConstructor(NewUser),
}

type ResourceConstructor func(ctx *pulumi.Context, name string, inputs provider.ConstructInputs, options pulumi.ResourceOption) (*provider.ConstructResult, error)

func createNewResourceConstructor[T any, P pulumi.ComponentResource](handler func(ctx *pulumi.Context, name string, inputs *T, opts ...pulumi.ResourceOption) (P, error)) ResourceConstructor {
	return func(ctx *pulumi.Context, name string, inputs provider.ConstructInputs, options pulumi.ResourceOption) (*provider.ConstructResult, error) {
		args := new(T)

		err := inputs.CopyTo(args)
		if err != nil {
			return nil, errors.Wrap(err, "setting args")
		}

		resource, err := handler(ctx, name, args, options)
		if err != nil {
			return nil, errors.Wrap(err, "creating component")
		}

		return provider.NewConstructResult(resource)
	}
}

func construct(ctx *pulumi.Context, typ, name string, inputs provider.ConstructInputs,
	options pulumi.ResourceOption) (*provider.ConstructResult, error) {
	handler, ok := resourceConstructorMap[typ]
	if !ok {
		return nil, errors.Errorf("unknown resource type %s", typ)
	}

	return handler(ctx, name, inputs, options)
}
