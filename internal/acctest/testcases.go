package acctest

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-go/tfprotov6"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func (td TestData) DataSourceTest(t *testing.T, steps []resource.TestStep) {
	td.runAcceptanceTest(t, resource.TestCase{
		PreCheck:     func() { PreCheck(t) },
		CheckDestroy: nil,
		Steps:        steps,
	})
}

func (td TestData) ResourceTest(t *testing.T, steps []resource.TestStep) {
	td.runAcceptanceTest(t, resource.TestCase{
		PreCheck:     func() { PreCheck(t) },
		CheckDestroy: td.CheckDestroy(),
		Steps:        steps,
	})
}

func (td TestData) runAcceptanceTest(t *testing.T, testCase resource.TestCase) {
	testCase.ProtoV6ProviderFactories = td.providerFactories()

	resource.ParallelTest(t, testCase)
}

func (td TestData) providerFactories() map[string]func() (tfprotov6.ProviderServer, error) {
	return map[string]func() (tfprotov6.ProviderServer, error){
		// nolint: unparam
		"awscc": func() (tfprotov6.ProviderServer, error) {
			return tfsdk.NewProtocol6Server(td.provider), nil
		},
	}
}

func PreCheck(t *testing.T) {}
