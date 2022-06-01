module aws-iam-go

go 1.16

require (
	github.com/pulumi/pulumi-aws-iam v0.0.2-alpha6 // indirect
	github.com/pulumi/pulumi-aws-iam/sdk v0.0.0-00010101000000-000000000000
	github.com/pulumi/pulumi/sdk/v3 v3.32.1
	gopkg.in/yaml.v3 v3.0.1 // indirect
)

replace github.com/pulumi/pulumi-aws-iam/sdk => ../../sdk
