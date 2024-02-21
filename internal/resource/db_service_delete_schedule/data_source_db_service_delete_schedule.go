package db_service_delete_schedule

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	apiClient "terraform-provider-tessell/internal/client"
)

func DataSourceDBServiceDeleteSchedule() *schema.Resource {
	return &schema.Resource{

		ReadContext: dataSourceDBServiceDeleteScheduleRead,

		Schema: map[string]*schema.Schema{
			"id": {
				Type:        schema.TypeString,
				Description: "",
				Required:    true,
			},
			"delete_at": {
				Type:        schema.TypeString,
				Description: "Time at which the DB Service should be deleted at",
				Computed:    true,
			},
			"deletion_config": {
				Type:        schema.TypeList,
				Description: "If the DB Service is to be deleted, this config would be honoured if no preference is provided during deleting the service",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"retain_availability_machine": {
							Type:        schema.TypeBool,
							Description: "If specified as true, the associated Availability Machine (snapshots, sanitized-snapshots, logs) would be retained",
							Computed:    true,
						},
					},
				},
			},
		},
	}
}

func dataSourceDBServiceDeleteScheduleRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*apiClient.Client)

	var diags diag.Diagnostics

	serviceId := d.Get("service_id").(string)
	id := d.Get("id").(string)

	response, _, err := client.GetServiceDeletionScheduleTFP(serviceId, id)
	if err != nil {
		return diag.FromErr(err)
	}

	if err := setResourceData(d, response); err != nil {
		return diag.FromErr(err)
	}

	d.SetId(*response.Id)

	return diags
}
