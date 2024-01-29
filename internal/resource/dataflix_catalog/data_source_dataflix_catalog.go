package dataflix_catalog

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	apiClient "terraform-provider-tessell/internal/client"
)

func DataSourceDataflixCatalog() *schema.Resource {
	return &schema.Resource{

		ReadContext: dataSourceDataflixCatalogRead,

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
				Description: "Database Engine Type",
				Computed:    true,
			},
			"time_zone": {
				Type:        schema.TypeString,
				Description: "Timezone applicable for timestamps that are returned in this response",
				Computed:    true,
			},
			"owner": {
				Type:        schema.TypeString,
				Description: "Owner of the Availability Machine",
				Computed:    true,
			},
			"pitr_catalog": {
				Type:        schema.TypeList,
				Description: "Catalog information for the point-in-time recoverability",
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
							Description: "",
							Computed:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"region": {
										Type:        schema.TypeString,
										Description: "Region name",
										Computed:    true,
									},
									"time_ranges": {
										Type:        schema.TypeList,
										Description: "",
										Computed:    true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"from_time": {
													Type:        schema.TypeString,
													Description: "Recoverability start timestamp",
													Computed:    true,
												},
												"to_time": {
													Type:        schema.TypeString,
													Description: "Recoverability end timestamp",
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
											},
										},
									},
								},
							},
						},
					},
				},
			},
			"snapshot_catalog": {
				Type:        schema.TypeList,
				Description: "Catalog information for the available snapshots",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type:        schema.TypeString,
							Description: "ID of the snapshot",
							Required:    true,
						},
						"name": {
							Type:        schema.TypeString,
							Description: "Name of the snapshot",
							Computed:    true,
						},
						"description": {
							Type:        schema.TypeString,
							Description: "Description for the snapshot",
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
							Description: "Specifies whether the backup is captured as per manual user request or as per the automated schedule",
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
				},
			},
		},
	}
}

func dataSourceDataflixCatalogRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*apiClient.Client)

	var diags diag.Diagnostics

	availabilityMachineId := d.Get("availability_machine_id").(string)

	response, _, err := client.GetDataflixCatalog(availabilityMachineId)
	if err != nil {
		return diag.FromErr(err)
	}

	if err := setResourceData(d, response); err != nil {
		return diag.FromErr(err)
	}

	d.SetId(*response.AvailabilityMachineId)

	return diags
}
