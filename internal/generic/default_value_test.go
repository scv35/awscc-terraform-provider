package generic

import (
	"context"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-provider-awscc/internal/tfresource"
)

func TestDefaultStringValue(t *testing.T) {
	t.Parallel()

	type testCase struct {
		plannedValue  types.String
		currentValue  types.String
		defaultValue  types.String
		expectedValue types.String
		expectError   bool
	}
	tests := map[string]testCase{
		"non-default non-Null string": {
			plannedValue:  types.StringValue("gamma"),
			currentValue:  types.StringValue("beta"),
			defaultValue:  types.StringValue("alpha"),
			expectedValue: types.StringValue("gamma"),
		},
		"non-default non-Null string, current Null": {
			plannedValue:  types.StringValue("gamma"),
			currentValue:  types.StringNull(),
			defaultValue:  types.StringValue("alpha"),
			expectedValue: types.StringValue("gamma"),
		},
		"non-default Null string, current Null": {
			plannedValue:  types.StringNull(),
			currentValue:  types.StringValue("beta"),
			defaultValue:  types.StringValue("alpha"),
			expectedValue: types.StringNull(),
		},
		"default string": {
			plannedValue:  types.StringNull(),
			currentValue:  types.StringValue("alpha"),
			defaultValue:  types.StringValue("alpha"),
			expectedValue: types.StringValue("alpha"),
		},
		"default string on create": {
			plannedValue:  types.StringNull(),
			currentValue:  types.StringNull(),
			defaultValue:  types.StringValue("alpha"),
			expectedValue: types.StringNull(),
		},
	}

	for name, test := range tests {
		name, test := name, test
		t.Run(name, func(t *testing.T) {
			ctx := context.TODO()
			request := planmodifier.StringRequest{
				PlanValue:  test.plannedValue,
				Path:       path.Root("test"),
				StateValue: test.currentValue,
			}
			response := planmodifier.StringResponse{}
			StringDefaultValue(test.defaultValue).PlanModifyString(ctx, request, &response)

			if !response.Diagnostics.HasError() && test.expectError {
				t.Fatal("expected error, got no error")
			}

			if response.Diagnostics.HasError() && !test.expectError {
				t.Fatalf("got unexpected error: %s", tfresource.DiagsError(response.Diagnostics))
			}

			if diff := cmp.Diff(response.PlanValue, test.expectedValue); diff != "" {
				t.Errorf("unexpected diff (+wanted, -got): %s", diff)
			}
		})
	}
}
