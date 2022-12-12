package provider

import (
	"context"

	"terraform-provider-tessell/internal/client"
	"terraform-provider-tessell/internal/resource/availability_machine"
	"terraform-provider-tessell/internal/resource/dataflix"
	"terraform-provider-tessell/internal/resource/dataflix_catalog"
	"terraform-provider-tessell/internal/resource/db_service"
	"terraform-provider-tessell/internal/resource/db_snapshot"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func init() {
	// Set descriptions to support markdown syntax, this will be used in document generation
	// and the language server.
	schema.DescriptionKind = schema.StringMarkdown

	// Customize the content of descriptions when output. For example you can add defaults on
	// to the exported descriptions if present.
	// schema.SchemaDescriptionBuilder = func(s *schema.Schema) string {
	// 	desc := s.Description
	// 	if s.Default != nil {
	// 		desc += fmt.Sprintf(" Defaults to `%v`.", s.Default)
	// 	}
	// 	return strings.TrimSpace(desc)
	// }
}

func New(terraformVersion string) func() *schema.Provider {
	return func() *schema.Provider {
		provider := &schema.Provider{
			Schema: map[string]*schema.Schema{
				"api_address": {
					Type:         schema.TypeString,
					Required:     true,
					ValidateFunc: validation.IsURLWithHTTPS,
					DefaultFunc:  schema.EnvDefaultFunc("TESSELL_API_ADDRESS", nil),
				},
				"api_key": {
					Type:         schema.TypeString,
					Required:     true,
					Sensitive:    true,
					ValidateFunc: validation.StringLenBetween(32, 40),
					DefaultFunc:  schema.EnvDefaultFunc("TESSELL_API_KEY", nil),
				},
				"tenant_id": {
					Type:         schema.TypeString,
					Required:     true,
					ValidateFunc: validation.IsUUID,
					DefaultFunc:  schema.EnvDefaultFunc("TESSELL_TENANT_ID", nil),
				},
			},
			DataSourcesMap: map[string]*schema.Resource{
				"tessell_db_service":            db_service.DataSourceDBService(),
				"tessell_db_services":           db_service.DataSourceDBServices(),
				"tessell_db_snapshot":           db_snapshot.DataSourceDBSnapshot(),
				"tessell_dataflix":              dataflix.DataSourceDataflix(),
				"tessell_dataflixes":            dataflix.DataSourceDataflixes(),
				"tessell_availability_machine":  availability_machine.DataSourceAvailabilityMachine(),
				"tessell_availability_machines": availability_machine.DataSourceAvailabilityMachines(),
				"tessell_dataflix_catalog":      dataflix_catalog.DataSourceDataflixCatalog(),
			},
			ResourcesMap: map[string]*schema.Resource{
				"tessell_db_service":  db_service.ResourceDBService(),
				"tessell_db_snapshot": db_snapshot.ResourceDBSnapshot(),
			},
		}

		provider.ConfigureContextFunc = configure(terraformVersion, provider)

		return provider
	}
}

func configure(terraformVersion string, provider *schema.Provider) func(context.Context, *schema.ResourceData) (interface{}, diag.Diagnostics) {
	return func(_ context.Context, d *schema.ResourceData) (interface{}, diag.Diagnostics) {
		apiAddress := d.Get("api_address").(string)
		apiKey := d.Get("api_key").(string)
		tenantId := d.Get("tenant_id").(string)

		var diags diag.Diagnostics

		c, err := client.NewClient(&apiAddress, &apiKey, &tenantId)
		if err != nil {
			return nil, diag.FromErr(err)
		}
		return c, diags
	}
}
