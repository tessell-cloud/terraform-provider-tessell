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
				Description: "ID of the snapshot",
				Computed:    true,
			},
			"name": {
				Type:        schema.TypeString,
				Description: "Name of the snapshot",
				Required:    true,
				ForceNew:    true,
			},
			"description": {
				Type:        schema.TypeString,
				Description: "Description of the snapshot",
				Optional:    true,
			},
			"snapshot_time": {
				Type:        schema.TypeString,
				Description: "Capture time of the snapshot",
				Computed:    true,
			},
			"status": {
				Type:        schema.TypeString,
				Description: "Database Backup Status",
				Computed:    true,
			},
			"size": {
				Type:        schema.TypeInt,
				Description: "Size of this snapshot (in bytes)",
				Computed:    true,
			},
			"manual": {
				Type:        schema.TypeBool,
				Description: "Specifies whether this snapshot is captured as per manual user request or per automated schedule",
				Computed:    true,
			},
			"cloud_availability": {
				Type:        schema.TypeList,
				Description: "The cloud and region information where this snapshot has been made available at",
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
							Description: "Region specific availability details for the snapshot",
							Optional:    true,
							ForceNew:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"region": {
										Type:        schema.TypeString,
										Description: "The region name",
										Required:    true,
										ForceNew:    true,
									},
									"status": {
										Type:        schema.TypeString,
										Description: "The current status of the snapshot in the respective region",
										Optional:    true,
										ForceNew:    true,
									},
								},
							},
						},
					},
				},
			},
			"availability_config": {
				Type:        schema.TypeList,
				Description: "The config information for cloud and region availability for this snapshot",
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
										Description: "The list of regions and respective availability status",
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
				Description: "The databases that are captured as part of this snapshot",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type:        schema.TypeString,
							Description: "ID of the database",
							Optional:    true,
							ForceNew:    true,
						},
						"name": {
							Type:        schema.TypeString,
							Description: "Name of the database",
							Optional:    true,
							ForceNew:    true,
						},
						"status": {
							Type:        schema.TypeString,
							Description: "Status of the database as of capture of this snapshot",
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
				DiffSuppressFunc: func(k, old, new string, d *schema.ResourceData) bool {
					if old == "0" && new == "1" {
						users := d.Get("shared_with.0.users")
						if len(users.([]interface{})) == 0 {
							return true
						}
					} else if old == "1" && new == "0" {
						users := d.GetRawState().GetAttr("shared_with").AsValueSlice()[0].GetAttr("users").AsValueSlice()
						if len(users) == 0 {
							return true
						}
					}
					return false
				},
			},
			"backup_status": {
				Type:        schema.TypeString,
				Description: "",
				Optional:    true,
				ForceNew:    true,
			},
			"availability_machine_id": {
				Type:        schema.TypeString,
				Description: "Id of the parent AvailabilityMachine, required when creating a clone",
				Required:    true,
				ForceNew:    true,
			},
			"block_until_complete": {
				Type:        schema.TypeBool,
				Description: "For any operation on this resource, block the flow until the action has completed successfully",
				Computed:    true,
			},
			"timeout": {
				Type:        schema.TypeInt,
				Description: "If block_until_complete is true, how long it should block for. (In seconds)",
				Computed:    true,
			},
		},
	}
}

func resourceDBSnapshotCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*apiClient.Client)

	var diags diag.Diagnostics
	var id string

	availabilityMachineId := d.Get("availability_machine_id").(string)

	payload := formPayloadForCreateDatabaseSnapshotRequest(d)

	response, _, err := client.CreateDatabaseSnapshotRequest(availabilityMachineId, payload)
	if err != nil {
		return diag.FromErr(err)
	}
	id = *response.ResourceId

	d.SetId(id)

	resourceDBSnapshotRead(ctx, d, meta)

	return diags
}

func resourceDBSnapshotRead(_ context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*apiClient.Client)

	var diags diag.Diagnostics

	availabilityMachineId := d.Get("availability_machine_id").(string)
	id := d.Get("id").(string)

	response, _, err := client.GetDatabaseSnapshot(availabilityMachineId, id)
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

	response, statusCode, err := client.DeleteDatabaseSnapshotRequest(availabilityMachineId, id)
	if err != nil {
		return diag.FromErr(err)
	}

	if statusCode != 200 {
		return diag.FromErr(fmt.Errorf("deletion failed for tessell_db_snapshot with id %s. Received response: %+v", id, response))
	}

	return diags
}
