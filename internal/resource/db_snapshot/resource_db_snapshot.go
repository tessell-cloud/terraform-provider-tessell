package db_snapshot

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	apiClient "terraform-provider-tessell/internal/client"
)

func ResourceDBSnapshot() *schema.Resource {
	return &schema.Resource{

		CreateContext: resourceDBSnapshotCreate,
		ReadContext:   resourceDBSnapshotRead,
		UpdateContext: resourceDBSnapshotUpdate,
		DeleteContext: resourceDBSnapshotDelete,

		Schema: map[string]*schema.Schema{
			"id": {
				Type:        schema.TypeString,
				Description: "DB Service snapshot Id",
				Computed:    true,
			},
			"name": {
				Type:        schema.TypeString,
				Description: "DB Service snapshot name",
				Optional:    true,
				ForceNew:    true,
			},
			"description": {
				Type:        schema.TypeString,
				Description: "Description for the snapshot",
				Optional:    true,
				ForceNew:    true,
			},
			"snapshot_time": {
				Type:        schema.TypeString,
				Description: "DB Service snapshot capture time",
				Computed:    true,
			},
			"status": {
				Type:        schema.TypeString,
				Description: "Database Backup Status",
				Computed:    true,
			},
			"size": {
				Type:        schema.TypeInt,
				Description: "Database Backup size in bytes",
				Computed:    true,
			},
			"manual": {
				Type:        schema.TypeBool,
				Description: "Specifies whether the backup is captured manually",
				Computed:    true,
			},
			"cloud_availability": {
				Type:        schema.TypeList,
				Description: "",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"cloud": {
							Type:        schema.TypeString,
							Description: "",
							Required:    true,
							ForceNew:    true,
						},
						"regions": {
							Type:        schema.TypeList,
							Description: "The regions details",
							Optional:    true,
							ForceNew:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"region": {
										Type:        schema.TypeString,
										Description: "The cloud region name",
										Required:    true,
										ForceNew:    true,
									},
									"availability_zones": {
										Type:        schema.TypeList,
										Description: "",
										Optional:    true,
										ForceNew:    true,
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
			"availability_config": {
				Type:        schema.TypeList,
				Description: "",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"availability_configured_manually": {
							Type:        schema.TypeBool,
							Description: "",
							Optional:    true,
							ForceNew:    true,
							Default:     false,
						},
						"dap_id": {
							Type:        schema.TypeString,
							Description: "",
							Computed:    true,
						},
						"cloud_availability_config": {
							Type:        schema.TypeList,
							Description: "",
							Computed:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"cloud": {
										Type:        schema.TypeString,
										Description: "",
										Required:    true,
										ForceNew:    true,
									},
									"regions": {
										Type:        schema.TypeList,
										Description: "The list of regions and respective avaoilability status",
										Optional:    true,
										ForceNew:    true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"region": {
													Type:        schema.TypeString,
													Description: "",
													Required:    true,
													ForceNew:    true,
												},
												"status": {
													Type:        schema.TypeString,
													Description: "Database Backup Status",
													Computed:    true,
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
			"databases": {
				Type:        schema.TypeList,
				Description: "The databases that are captured as part of the snapshot",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type:        schema.TypeString,
							Description: "Databases Id",
							Optional:    true,
							ForceNew:    true,
						},
						"name": {
							Type:        schema.TypeString,
							Description: "Databases name",
							Optional:    true,
							ForceNew:    true,
						},
						"status": {
							Type:        schema.TypeString,
							Description: "Databases status",
							Optional:    true,
							ForceNew:    true,
						},
					},
				},
			},
			"shared_with": {
				Type:        schema.TypeList,
				Description: "Tessell Entity ACL Sharing Summary Info",
				Optional:    true,
				ForceNew:    true,
				MaxItems:    1,
				MinItems:    1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"users": {
							Type:        schema.TypeList,
							Description: "",
							Optional:    true,
							ForceNew:    true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
					},
				},
			},
			"availability_machine_id": {
				Type:        schema.TypeString,
				Description: "Id of the parent AvailabilityMachine, required when creating a clone",
				Optional:    true,
				ForceNew:    true,
			},
			"block_until_complete": {
				Type:        schema.TypeBool,
				Description: "For any operation on this resource, block the flow until the action has completed successfully",
				Optional:    true,
				Default:     true,
			},
			"timeout": {
				Type:        schema.TypeInt,
				Description: "If block_until_complete is true, how long it should block for. (In seconds)",
				Optional:    true,
				Default:     1200,
			},
		},
	}
}

func resourceDBSnapshotCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*apiClient.Client)

	var diags diag.Diagnostics
	var id string

	availabilityMachineId := d.Get("availability_machine_id").(string)

	payload := formPayloadForCreateTessellServiceBackupRequest(d)

	response, _, err := client.CreateTessellServiceBackupRequest(availabilityMachineId, payload)
	if err != nil {
		return diag.FromErr(err)
	}
	id = *response.ResourceId

	d.SetId(id)

	if d.Get("block_until_complete").(bool) {
		//if err := client.WaitTillReady(resourceId, d.Get("timeout").(int)); err != nil {
		if err := client.DBSnapshotPollForStatus(availabilityMachineId, id, "AVAILABLE", d.Get("timeout").(int), 60); err != nil {
			return diag.FromErr(err)
		}
	}

	resourceDBSnapshotRead(ctx, d, meta)

	return diags
}

func resourceDBSnapshotRead(_ context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*apiClient.Client)

	var diags diag.Diagnostics

	availabilityMachineId := d.Get("availability_machine_id").(string)
	id := d.Get("id").(string)

	response, _, err := client.GetBackup(availabilityMachineId, id)
	if err != nil {
		return diag.FromErr(err)
	}

	if err := setResourceData(d, response); err != nil {
		return diag.FromErr(err)
	}

	d.SetId(*response.Id)

	return diags
}
func resourceDBSnapshotUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	return diags
}

func resourceDBSnapshotDelete(_ context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*apiClient.Client)

	var diags diag.Diagnostics

	availabilityMachineId := d.Get("availability_machine_id").(string)
	id := d.Get("id").(string)

	response, statusCode, err := client.DeleteBackupRequest(availabilityMachineId, id)
	if err != nil {
		return diag.FromErr(err)
	}

	if statusCode != 200 {
		return diag.FromErr(fmt.Errorf("deletion failed for tessell_db_snapshot with resourceId %s. Received response: %+v", id, response))
	}

	//err = client.WaitTillDeleted(databaseDeletionResponse.TaskId, d.Get("timeout").(int), "Database Deletion")
	err = client.DBSnapshotPollForStatusCode(availabilityMachineId, id, 404, d.Get("timeout").(int), 30)
	if err != nil {
		return diag.FromErr(err)
	}

	return diags
}
