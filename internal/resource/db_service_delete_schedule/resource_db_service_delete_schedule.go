package db_service_delete_schedule

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	apiClient "terraform-provider-tessell/internal/client"
)

func ResourceDBServiceDeleteSchedule() *schema.Resource {
	return &schema.Resource{

		CreateContext: resourceDBServiceDeleteScheduleCreate,
		ReadContext:   resourceDBServiceDeleteScheduleRead,
		UpdateContext: resourceDBServiceDeleteScheduleUpdate,
		DeleteContext: resourceDBServiceDeleteScheduleDelete,

		Schema: map[string]*schema.Schema{
			"id": {
				Type:        schema.TypeString,
				Description: "",
				Computed:    true,
			},
			"delete_at": {
				Type:        schema.TypeString,
				Description: "Time at which the DB Service should be deleted at",
				Required:    true,
			},
			"deletion_config": {
				Type:        schema.TypeList,
				Description: "If the DB Service is to be deleted, this config would be honoured if no preference is provided during deleting the service",
				Optional:    true,
				MaxItems:    1,
				MinItems:    1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"retain_availability_machine": {
							Type:        schema.TypeBool,
							Description: "If specified as true, the associated Availability Machine (snapshots, sanitized-snapshots, logs) would be retained",
							Optional:    true,
							Default:     false,
						},
					},
				},
			},
			"service_id": {
				Type:        schema.TypeString,
				Description: "The ID of the DB Service",
				Required:    true,
				ForceNew:    true,
			},
		},
	}
}

func resourceDBServiceDeleteScheduleCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*apiClient.Client)

	var diags diag.Diagnostics
	var id string

	serviceId := d.Get("service_id").(string)

	payload := formPayloadForCreateServiceDeletionSchedule(d)

	response, _, err := client.CreateServiceDeletionSchedule(serviceId, payload)
	if err != nil {
		return diag.FromErr(err)
	}
	id = *response.Id

	d.SetId(id)

	resourceDBServiceDeleteScheduleRead(ctx, d, meta)

	return diags
}

func resourceDBServiceDeleteScheduleRead(_ context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

func resourceDBServiceDeleteScheduleUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*apiClient.Client)

	var diags diag.Diagnostics
	var pollBreakValue string
	var pollFunc func(string, string, int, int) error
	shouldPoll := false

	serviceId := d.Get("service_id").(string)
	id := d.Get("id").(string)

	payload := formPayloadForUpdateServiceDeletionScheduleTFP(d)

	_, _, err := client.UpdateServiceDeletionScheduleTFP(serviceId, id, payload)
	if err != nil {
		return diag.FromErr(err)
	}

	if shouldPoll {
		if err := pollFunc(d.Get("id").(string), pollBreakValue, d.Get("timeout").(int), 30); err != nil {
			return diag.FromErr(err)
		}
	}

	resourceDBServiceDeleteScheduleRead(ctx, d, meta)

	return diags
}

func resourceDBServiceDeleteScheduleDelete(_ context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*apiClient.Client)

	var diags diag.Diagnostics

	serviceId := d.Get("service_id").(string)
	id := d.Get("id").(string)

	response, statusCode, err := client.DeleteServiceDeletionScheduleTFP(serviceId, id)
	if err != nil {
		return diag.FromErr(err)
	}

	if statusCode != 200 {
		return diag.FromErr(fmt.Errorf("deletion failed for tessell_db_service_delete_schedule with id %s. Received response: %+v", id, response))
	}

	return diags
}
