package db_snapshot

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	apiClient "terraform-provider-tessell/internal/client"
)

func DataSourceDBSnapshot() *schema.Resource {
	return &schema.Resource{

		ReadContext: dataSourceDBSnapshotRead,

		Schema: map[string]*schema.Schema{
			"id": {
				Type:        schema.TypeString,
				Description: "ID of the snapshot",
				Required:    true,
			},
			"availability_machine_id": {
				Type:        schema.TypeString,
				Description: "Id of the parent AvailabilityMachine, required when creating a clone",
				Required:    true,
			},
			"name": {
				Type:        schema.TypeString,
				Description: "Name of the snapshot",
				Computed:    true,
			},
			"description": {
				Type:        schema.TypeString,
				Description: "Description of the snapshot",
				Computed:    true,
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
							Computed:    true,
						},
						"regions": {
							Type:        schema.TypeList,
							Description: "Region specific availability details for the snapshot",
							Computed:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"region": {
										Type:        schema.TypeString,
										Description: "The region name",
										Computed:    true,
									},
									"status": {
										Type:        schema.TypeString,
										Description: "The current status of the snapshot in the respective region",
										Computed:    true,
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
							Computed:    true,
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
										Computed:    true,
									},
									"regions": {
										Type:        schema.TypeList,
										Description: "The list of regions and respective availability status",
										Computed:    true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"region": {
													Type:        schema.TypeString,
													Description: "",
													Computed:    true,
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
							Required:    true,
						},
						"name": {
							Type:        schema.TypeString,
							Description: "Name of the database",
							Computed:    true,
						},
						"status": {
							Type:        schema.TypeString,
							Description: "Status of the database as of capture of this snapshot",
							Computed:    true,
						},
					},
				},
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
			"backup_status": {
				Type:        schema.TypeString,
				Description: "",
				Computed:    true,
			},
		},
	}
}

func dataSourceDBSnapshotRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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
