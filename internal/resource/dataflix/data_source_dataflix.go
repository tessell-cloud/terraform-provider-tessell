package dataflix

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	apiClient "terraform-provider-tessell/internal/client"
)

func DataSourceDataflix() *schema.Resource {
	return &schema.Resource{

		ReadContext: dataSourceDataflixRead,

		Schema: map[string]*schema.Schema{
			"availability_machine_id": {
				Type:        schema.TypeString,
				Description: "ID of the Availability Machine",
				Required:    true,
			},
			"tessell_service_id": {
				Type:        schema.TypeString,
				Description: "ID of the associated DB Service",
				Computed:    true,
			},
			"service_name": {
				Type:        schema.TypeString,
				Description: "Name of the associated DB Service",
				Computed:    true,
			},
			"engine_type": {
				Type:        schema.TypeString,
				Description: "",
				Computed:    true,
			},
			"cloud_availability": {
				Type:        schema.TypeList,
				Description: "The cloud and region information where the data is available for access",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"cloud": {
							Type:        schema.TypeString,
							Description: "",
							Computed:    true,
						},
						"regions": {
							Type:        schema.TypeList,
							Description: "The regions details",
							Computed:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"region": {
										Type:        schema.TypeString,
										Description: "The cloud region name",
										Computed:    true,
									},
									"availability_zones": {
										Type:        schema.TypeList,
										Description: "",
										Computed:    true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
								},
							},
						},
					},
				},
			},
			"owner": {
				Type:        schema.TypeString,
				Description: "Owner of the Availability Machine",
				Computed:    true,
			},
			"tsm": {
				Type:        schema.TypeBool,
				Description: "Specify whether the associated DB Service is created using TSM compute type",
				Computed:    true,
			},
			"shared_with": {
				Type:        schema.TypeList,
				Description: "Tessell Entity ACL Sharing Summary Info",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"users": {
							Type:        schema.TypeList,
							Description: "",
							Computed:    true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
					},
				},
			},
			"storage_provider_type": {
				Type:        schema.TypeString,
				Description: "",
				Computed:    true,
			},
		},
	}
}

func dataSourceDataflixRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*apiClient.Client)

	var diags diag.Diagnostics

	availabilityMachineId := d.Get("availability_machine_id").(string)

	response, _, err := client.GetDataflixByAmId(availabilityMachineId)
	if err != nil {
		return diag.FromErr(err)
	}

	if err := setResourceData(d, response); err != nil {
		return diag.FromErr(err)
	}

	d.SetId(*response.AvailabilityMachineId)

	return diags
}
