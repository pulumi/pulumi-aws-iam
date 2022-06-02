// Copyright 2016-2022, Pulumi Corporation.
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
	"fmt"
	"reflect"

	"github.com/pulumi/pulumi-aws/sdk/v5/go/aws/iam"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

const UserIdentifier = "aws-iam:index:User"

type UserArgs struct {
	// Desired name for the IAM user.
	Name string `pulumi:"name"`

	// Desired path for the IAM user.
	Path string `pulumi:"string"`

	// When destroying this user, destroy even if it has non-Pulumi-managed IAM access keys, login
	// profile or MFA devices. Without forceDestroy a user with non-Pulumi-managed access keys
	// and login profile will fail to be destroyed.
	ForceDestroy bool `pulumi:"forceDestroy"`

	// Either a base-64 encoded PGP public key, or a keybase username in the form
	// `keybase:username`. Used to encrypt password and access key.
	PGPKey string `pulumi:"pgpKey"`

	// Whether the user should be forced to reset the generated password on first login.
	PasswordResetRequired bool `pulumi:"passwordResetRequired"`

	// The length of the generated password.
	PasswordLength int `pulumi:"passwordLength"`

	// Whether to upload a public ssh key to the IAM user.
	UploadIAMUserSSHKey bool `pulumi:"uploadIamUserSshKey"`

	// Specifies the public key encoding format to use in the response. To retrieve the
	// public key in ssh-rsa format, use SSH. To retrieve the public key in PEM format, use PEM.
	SSHKeyEncoding string `pulumi:"sshKeyEncoding"`

	// The SSH public key. The public key must be encoded in ssh-rsa format or PEM format.
	SSHPublicKey string `pulumi:"sshPublicKey"`

	// The ARN of the policy that is used to set the permissions boundary for the user.
	PermissionsBoundary string `pulumi:"permissionsBoundary"`

	// A map of tags to add.
	Tags map[string]string `pulumi:"tags"`
}

type UserInfo struct {
	// The user's name.
	Name pulumi.StringOutput `pulumi:"name"`

	// The ARN assigned by AWS for this user.
	ARN pulumi.StringOutput `pulumi:"arn"`

	// The unique ID assigned by AWS.
	UniqueID pulumi.StringOutput `pulumi:"uniqueId"`

	// The fingerprint of the PGP key used to encrypt the password.
	LoginProfileKeyFingerprint pulumi.StringOutput `pulumi:"loginProfileKeyFingerprint"`

	// The encrypted password, base64 encoded.
	LoginProfileEncryptedPassword pulumi.StringOutput `pulumi:"loginProfileEncryptedPassword"`

	// The user password.
	LoginProfilePassword pulumi.StringOutput `pulumi:"loginProfilePassword"`

	// The unique identifier for the SSH public key
	SSHKeySSHPublicKeyID pulumi.StringOutput `pulumi:"sshKeySshPublicKeyId"`

	// The unique identifier for the SSH public key.
	SSHKeyFingerprint pulumi.StringOutput `pulumi:"sshKeyFingerprint"`
}

type UserInfoOutput struct {
	*pulumi.OutputState
	UserInfo
}

func (UserInfoOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*UserInfo)(nil)).Elem()
}

type AccessKey struct {
	// The access key ID.
	ID pulumi.IDOutput `pulumi:"id"`

	// The access key secret.
	Secret pulumi.StringOutput `pulumi:"secret"`

	// The fingerprint of the PGP key used to encrypt the secret.
	KeyFingerprint pulumi.StringOutput `pulumi:"keyFingerprint"`

	// The encrypted secret, base64 encoded.
	EncryptedSecret pulumi.StringOutput `pulumi:"encryptedSecret"`

	// The secret access key converted into an SES SMTP password by applying AWS's Sigv4 conversion algorithm.
	SESSMTPPasswordV4 pulumi.StringOutput `pulumi:"sesSmtpPasswordV4"`

	// Active or Inactive. Keys are initially active, but can be made inactive by other means.
	Status pulumi.StringOutput `pulumi:"status"`
}

type AccessKeyOutput struct {
	*pulumi.OutputState
	AccessKey
}

func (AccessKeyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AccessKey)(nil)).Elem()
}

type Keybase struct {
	// Decrypt user password command.
	PasswordDecryptCommand pulumi.StringOutput `pulumi:"passwordDecryptCommand"`

	// Encrypted password
	PasswordPGPMessage pulumi.StringOutput `pulumi:"passwordPgpMessage"`

	// Decrypt access secret key command.
	SecretKeyDecryptCommand pulumi.StringOutput `pulumi:"secretKeyDecryptCommand"`

	// Encrypted access secret key.
	SecretKeyPGPMessage pulumi.StringOutput `pulumi:"secretKeyPgpMessage"`
}

type KeybaseOutput struct {
	*pulumi.OutputState
	Keybase
}

type User struct {
	pulumi.ResourceState

	// The IAM user.
	UserInfo UserInfoOutput `pulumi:"userInfo"`

	// The IAM access key.
	AccessKey AccessKeyOutput `pulumi:"accessKey"`

	// PGP key used to encrypt sensitive data for this user (if empty - secrets are not encrypted).
	PGPKey pulumi.StringOutput `pulumi:"pgpKey"`

	Keybase KeybaseOutput `pulumi:"keybase"`
}

func NewUser(ctx *pulumi.Context, name string, args *UserArgs, opts ...pulumi.ResourceOption) (*User, error) {
	if args == nil {
		args = &UserArgs{}
	}

	component := &User{}
	err := ctx.RegisterComponentResource(UserIdentifier, name, component, opts...)
	if err != nil {
		return nil, err
	}

	opts = append(opts, pulumi.Parent(component))

	userPath := "/"
	if args.Path != "" {
		userPath = args.Path
	}

	user, err := iam.NewUser(ctx, name, &iam.UserArgs{
		Name:                pulumi.String(args.Name),
		Path:                pulumi.String(userPath),
		ForceDestroy:        pulumi.BoolPtr(args.ForceDestroy),
		PermissionsBoundary: pulumi.String(args.PermissionsBoundary),
		Tags:                pulumi.ToStringMap(args.Tags),
	}, opts...)
	if err != nil {
		return nil, err
	}

	if args.PasswordLength == 0 {
		args.PasswordLength = 20
	}

	loginProfile, err := iam.NewUserLoginProfile(ctx, name, &iam.UserLoginProfileArgs{
		User:                  user.Name,
		PgpKey:                pulumi.Sprintf("%s", args.PGPKey),
		PasswordLength:        pulumi.IntPtr(args.PasswordLength),
		PasswordResetRequired: pulumi.BoolPtr(args.PasswordResetRequired),
	}, opts...)
	if err != nil {
		return nil, err
	}

	accessKeyArgs := &iam.AccessKeyArgs{
		User: user.Name,
	}
	if args.PGPKey != "" {
		accessKeyArgs.PgpKey = pulumi.Sprintf("%s", args.PGPKey)
	}

	accessKey, err := iam.NewAccessKey(ctx, name, accessKeyArgs, opts...)
	if err != nil {
		return nil, err
	}

	sshKeyEncoding := "SSH"
	if args.SSHKeyEncoding != "" {
		sshKeyEncoding = args.SSHKeyEncoding
	}

	if args.UploadIAMUserSSHKey {
		uploadedKey, err := iam.NewSshKey(ctx, name, &iam.SshKeyArgs{
			Username:  user.Name,
			Encoding:  pulumi.String(sshKeyEncoding),
			PublicKey: pulumi.String(args.SSHPublicKey),
		}, opts...)
		if err != nil {
			return nil, err
		}

		component.UserInfo.SSHKeySSHPublicKeyID = uploadedKey.SshPublicKeyId
		component.UserInfo.SSHKeyFingerprint = uploadedKey.Fingerprint
	}

	component.UserInfo.Name = user.Name
	component.UserInfo.ARN = user.Arn
	component.UserInfo.UniqueID = user.UniqueId
	component.UserInfo.LoginProfileKeyFingerprint = loginProfile.KeyFingerprint
	component.UserInfo.LoginProfileEncryptedPassword = loginProfile.EncryptedPassword
	component.UserInfo.LoginProfilePassword = loginProfile.Password

	component.AccessKey.ID = accessKey.ID()
	component.AccessKey.Secret = accessKey.Secret
	component.AccessKey.KeyFingerprint = accessKey.KeyFingerprint
	component.AccessKey.EncryptedSecret = accessKey.EncryptedSecret
	component.AccessKey.SESSMTPPasswordV4 = accessKey.SesSmtpPasswordV4
	component.AccessKey.Status = pulumi.StringOutput(accessKey.Status)

	component.PGPKey = pulumi.Sprintf("%s", args.PGPKey)

	component.Keybase.PasswordDecryptCommand = loginProfile.EncryptedPassword.ApplyT(func(encryptedPassword string) string {
		return fmt.Sprintf("echo \"%s\" | base64 --decode | keybase pgp decrypt", encryptedPassword)
	}).(pulumi.StringOutput)

	component.Keybase.PasswordPGPMessage = loginProfile.EncryptedPassword.ApplyT(func(encryptedPassword string) string {
		return fmt.Sprintf("-----BEGIN PGP MESSAGE-----\nVersion: Keybase OpenPGP v2.0.76\nComment: https://keybase.io/crypto\n%s\n-----END PGP MESSAGE-----", encryptedPassword)
	}).(pulumi.StringOutput)

	component.Keybase.SecretKeyDecryptCommand = accessKey.EncryptedSecret.ApplyT(func(encryptedSecret string) string {
		return fmt.Sprintf("echo \"%s\" | base64 --decode | keybase pgp decrypt", encryptedSecret)
	}).(pulumi.StringOutput)

	component.Keybase.SecretKeyPGPMessage = accessKey.EncryptedSecret.ApplyT(func(encryptedSecret string) string {
		return fmt.Sprintf("-----BEGIN PGP MESSAGE-----\nVersion: Keybase OpenPGP v2.0.76\nComment: https://keybase.io/crypto\n%s\n-----END PGP MESSAGE-----", encryptedSecret)
	}).(pulumi.StringOutput)

	return component, nil
}
