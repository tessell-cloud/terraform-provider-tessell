package db_service_start_stop_schedule

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	apiClient "terraform-provider-tessell/internal/client"
)

func ResourceDBServiceStartStopSchedule() *schema.Resource {
	return &schema.Resource{

		CreateContext: resourceDBServiceStartStopScheduleCreate,
		ReadContext:   resourceDBServiceStartStopScheduleRead,
		UpdateContext: resourceDBServiceStartStopScheduleUpdate,
		DeleteContext: resourceDBServiceStartStopScheduleDelete,

		Schema: map[string]*schema.Schema{
			"id": {
				Type:        schema.TypeString,
				Description: "The ID of the schedule",
				Computed:    true,
			},
			"name": {
				Type:        schema.TypeString,
				Description: "Name of the schedule",
				Computed:    true,
			},
			"description": {
				Type:        schema.TypeString,
				Description: "",
				Computed:    true,
			},
			"service_id": {
				Type:        schema.TypeString,
				Description: "The ID of the DB Service",
				Required:    true,
				ForceNew:    true,
			},
			"status": {
				Type:        schema.TypeString,
				Description: "StartStopScheduleStatus",
				Computed:    true,
			},
			"schedule_info": {
				Type:        schema.TypeList,
				Description: "Describes the start/stop schedule of tessell service",
				Required:    true,
				MaxItems:    1,
				MinItems:    1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"one_time": {
							Type:        schema.TypeList,
							Description: "One time start/stop schedule details for the DB Service",
							Optional:    true,
							MaxItems:    1,
							MinItems:    1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"db_service_start_at": {
										Type:        schema.TypeString,
										Description: "Time at which the DB Service should be started at",
										Optional:    true,
									},
									"db_service_stop_at": {
										Type:        schema.TypeString,
										Description: "Time at which the DB Service should be stopped at",
										Optional:    true,
									},
								},
							},
						},
						"recurring": {
							Type:        schema.TypeList,
							Description: "Recurring start/stop schedule details for the DB Service tessell service",
							Optional:    true,
							MaxItems:    1,
							MinItems:    1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"schedule_start_date": {
										Type:        schema.TypeString,
										Description: "Date from which the given recurring schedule would be applicable from",
										Optional:    true,
									},
									"db_service_start_at": {
										Type:        schema.TypeString,
										Description: "Time at which the DB Service should be started at",
										Optional:    true,
									},
									"db_service_stop_at": {
										Type:        schema.TypeString,
										Description: "Time at which the DB Service should be stopped at",
										Optional:    true,
									},
									"schedule_expiry": {
										Type:        schema.TypeList,
										Description: "Schedule expiry details for recurring start/stop schedule for the DB Service",
										Optional:    true,
										MaxItems:    1,
										MinItems:    1,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"on": {
													Type:        schema.TypeString,
													Description: "Date after which the schedule would not be applicable",
													Optional:    true,
												},
												"after_occurrences": {
													Type:        schema.TypeInt,
													Description: "Number of occurrences which the schedule would not be applicable",
													Optional:    true,
												},
												"never": {
													Type:        schema.TypeBool,
													Description: "If set to True, the schedule will be applicable forever",
													Optional:    true,
												},
											},
										},
									},
									"daily_schedule": {
										Type:        schema.TypeBool,
										Description: "Whether the given schedule is a daily schedule i.e. a schedule which is executed daily",
										Optional:    true,
									},
									"weekly_schedule": {
										Type:        schema.TypeList,
										Description: "Weekly recurring start/stop schedule details for the DB Service",
										Optional:    true,
										MaxItems:    1,
										MinItems:    1,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"days": {
													Type:        schema.TypeList,
													Description: "Days of the week on which the recurring start/stop schedule would be applicable for the DB Service",
													Optional:    true,
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
					},
				},
			},
			"metadata": {
				Type:        schema.TypeList,
				Description: "",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"schedule_counter": {
							Type:        schema.TypeInt,
							Description: "",
							Computed:    true,
						},
					},
				},
			},
			"date_created": {
				Type:        schema.TypeString,
				Description: "Timestamp when the schedule was created",
				Computed:    true,
			},
			"date_modified": {
				Type:        schema.TypeString,
				Description: "Timestamp when the schedule was last modified",
				Computed:    true,
			},
			"last_run": {
				Type:        schema.TypeString,
				Description: "The date-time at which this schedule was last executed",
				Computed:    true,
			},
		},
	}
}

func resourceDBServiceStartStopScheduleCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*apiClient.Client)

	var diags diag.Diagnostics
	var id string

	serviceId := d.Get("service_id").(string)

	payload := formPayloadForCreateServiceStartStopSchedule(d)

	response, _, err := client.CreateServiceStartStopSchedule(serviceId, payload)
	if err != nil {
		return diag.FromErr(err)
	}
	id = *response.Id

	d.SetId(id)

	resourceDBServiceStartStopScheduleRead(ctx, d, meta)

	return diags
}

func resourceDBServiceStartStopScheduleRead(_ context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*apiClient.Client)

	var diags diag.Diagnostics

	serviceId := d.Get("service_id").(string)
	id := d.Get("id").(string)

	response, _, err := client.GetServiceStartStopSchedule(serviceId, id)
	if err != nil {
		return diag.FromErr(err)
	}

	if err := setResourceData(d, response); err != nil {
		return diag.FromErr(err)
	}

	d.SetId(*response.Id)

	return diags
}

func resourceDBServiceStartStopScheduleUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*apiClient.Client)

	var diags diag.Diagnostics
	var pollBreakValue string
	var pollFunc func(string, string, int, int) error
	shouldPoll := false

	serviceId := d.Get("service_id").(string)
	id := d.Get("id").(string)

	payload := formPayloadForUpdateServiceStartStopSchedule(d)

	_, _, err := client.UpdateServiceStartStopSchedule(serviceId, id, payload)
	if err != nil {
		return diag.FromErr(err)
	}

	if shouldPoll {
		if err := pollFunc(d.Get("id").(string), pollBreakValue, d.Get("timeout").(int), 30); err != nil {
			return diag.FromErr(err)
		}
	}

	resourceDBServiceStartStopScheduleRead(ctx, d, meta)

	return diags
}

func resourceDBServiceStartStopScheduleDelete(_ context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*apiClient.Client)

	var diags diag.Diagnostics

	serviceId := d.Get("service_id").(string)
	id := d.Get("id").(string)

	response, statusCode, err := client.DeleteServiceStartStopSchedule(serviceId, id)
	if err != nil {
		return diag.FromErr(err)
	}

	if statusCode != 200 {
		return diag.FromErr(fmt.Errorf("deletion failed for tessell_db_service_start_stop_schedule with id %s. Received response: %+v", id, response))
	}

	return diags
}
