// Code generated by smithy-go-codegen DO NOT EDIT.

package types_test

import (
	"fmt"
	"github.com/aws/aws-sdk-go-v2/service/cleanrooms/types"
)

func ExampleAnalysisRulePolicy_outputUsage() {
	var union types.AnalysisRulePolicy
	// type switches can be used to check the union value
	switch v := union.(type) {
	case *types.AnalysisRulePolicyMemberV1:
		_ = v.Value // Value is types.AnalysisRulePolicyV1

	case *types.UnknownUnionMember:
		fmt.Println("unknown tag:", v.Tag)

	default:
		fmt.Println("union is nil or unknown type")

	}
}

var _ types.AnalysisRulePolicyV1

func ExampleAnalysisRulePolicyV1_outputUsage() {
	var union types.AnalysisRulePolicyV1
	// type switches can be used to check the union value
	switch v := union.(type) {
	case *types.AnalysisRulePolicyV1MemberAggregation:
		_ = v.Value // Value is types.AnalysisRuleAggregation

	case *types.AnalysisRulePolicyV1MemberList:
		_ = v.Value // Value is types.AnalysisRuleList

	case *types.UnknownUnionMember:
		fmt.Println("unknown tag:", v.Tag)

	default:
		fmt.Println("union is nil or unknown type")

	}
}

var _ *types.AnalysisRuleAggregation
var _ *types.AnalysisRuleList

func ExampleConfiguredTableAnalysisRulePolicy_outputUsage() {
	var union types.ConfiguredTableAnalysisRulePolicy
	// type switches can be used to check the union value
	switch v := union.(type) {
	case *types.ConfiguredTableAnalysisRulePolicyMemberV1:
		_ = v.Value // Value is types.ConfiguredTableAnalysisRulePolicyV1

	case *types.UnknownUnionMember:
		fmt.Println("unknown tag:", v.Tag)

	default:
		fmt.Println("union is nil or unknown type")

	}
}

var _ types.ConfiguredTableAnalysisRulePolicyV1

func ExampleConfiguredTableAnalysisRulePolicyV1_outputUsage() {
	var union types.ConfiguredTableAnalysisRulePolicyV1
	// type switches can be used to check the union value
	switch v := union.(type) {
	case *types.ConfiguredTableAnalysisRulePolicyV1MemberAggregation:
		_ = v.Value // Value is types.AnalysisRuleAggregation

	case *types.ConfiguredTableAnalysisRulePolicyV1MemberList:
		_ = v.Value // Value is types.AnalysisRuleList

	case *types.UnknownUnionMember:
		fmt.Println("unknown tag:", v.Tag)

	default:
		fmt.Println("union is nil or unknown type")

	}
}

var _ *types.AnalysisRuleAggregation
var _ *types.AnalysisRuleList

func ExampleProtectedQueryOutput_outputUsage() {
	var union types.ProtectedQueryOutput
	// type switches can be used to check the union value
	switch v := union.(type) {
	case *types.ProtectedQueryOutputMemberS3:
		_ = v.Value // Value is types.ProtectedQueryS3Output

	case *types.UnknownUnionMember:
		fmt.Println("unknown tag:", v.Tag)

	default:
		fmt.Println("union is nil or unknown type")

	}
}

var _ *types.ProtectedQueryS3Output

func ExampleProtectedQueryOutputConfiguration_outputUsage() {
	var union types.ProtectedQueryOutputConfiguration
	// type switches can be used to check the union value
	switch v := union.(type) {
	case *types.ProtectedQueryOutputConfigurationMemberS3:
		_ = v.Value // Value is types.ProtectedQueryS3OutputConfiguration

	case *types.UnknownUnionMember:
		fmt.Println("unknown tag:", v.Tag)

	default:
		fmt.Println("union is nil or unknown type")

	}
}

var _ *types.ProtectedQueryS3OutputConfiguration

func ExampleTableReference_outputUsage() {
	var union types.TableReference
	// type switches can be used to check the union value
	switch v := union.(type) {
	case *types.TableReferenceMemberGlue:
		_ = v.Value // Value is types.GlueTableReference

	case *types.UnknownUnionMember:
		fmt.Println("unknown tag:", v.Tag)

	default:
		fmt.Println("union is nil or unknown type")

	}
}

var _ *types.GlueTableReference
