module aws-iam-go

go 1.16

require (
	github.com/pulumi/pulumi/sdk/v3 v3.32.1
	github.com/zchase/pulumi-aws-iam/sdk v0.0.0-20220508055142-d6c8c604b61e
)

replace github.com/pulumi/pulumi-aws-iam/sdk => ../../sdk
