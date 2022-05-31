package utils

import (
	"github.com/pulumi/pulumi-aws/sdk/v5/go/aws/iam"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func sanitizePolicyDocumentPrincipals(principals []iam.GetPolicyDocumentStatementPrincipal) []iam.GetPolicyDocumentStatementPrincipal {
	var sanitizedPrincipals []iam.GetPolicyDocumentStatementPrincipal
	for _, principal := range principals {
		if len(principal.Identifiers) > 0 {
			sanitizedPrincipals = append(sanitizedPrincipals, principal)
		}
	}
	return sanitizedPrincipals
}

func GetIAMPolicyDocument(ctx *pulumi.Context, args *iam.GetPolicyDocumentArgs) (*iam.GetPolicyDocumentResult, error) {
	var policyStatements []iam.GetPolicyDocumentStatement
	for _, statement := range args.Statements {
		sanitizedPrincipals := sanitizePolicyDocumentPrincipals(statement.Principals)
		statement.Principals = sanitizedPrincipals
		policyStatements = append(policyStatements, statement)
	}

	return iam.GetPolicyDocument(ctx, &iam.GetPolicyDocumentArgs{
		Statements: policyStatements,
	})
}
