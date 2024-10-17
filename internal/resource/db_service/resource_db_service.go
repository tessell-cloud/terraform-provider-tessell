package db_service

import (
	"context"
	"fmt"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/customdiff"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	apiClient "terraform-provider-tessell/internal/client"
)

func ResourceDBService() *schema.Resource {
	return &schema.Resource{

		CreateContext: resourceDBServiceCreate,
		ReadContext:   resourceDBServiceRead,
		UpdateContext: resourceDBServiceUpdate,
		DeleteContext: resourceDBServiceDelete,

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(60 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			"id": {
				Type:        schema.TypeString,
				Description: "Tessell generated UUID for the DB Service. This is the unique identifier for the DB Service.",
				Computed:    true,
			},
			"availability_machine_id": {
				Type:        schema.TypeString,
				Description: "Associated Availability Machine Id",
				Computed:    true,
			},
			"snapshot_id": {
				Type:        schema.TypeString,
				Description: "Tessell service snapshot Id, using which the clone is to be created",
				Optional:    true,
				ForceNew:    true,
			},
			"pitr": {
				Type:        schema.TypeString,
				Description: "PITR Timestamp, using which the clone is to be created",
				Optional:    true,
				ForceNew:    true,
			},
			"name": {
				Type:        schema.TypeString,
				Description: "Name of the DB Service",
				Required:    true,
			},
			"description": {
				Type:        schema.TypeString,
				Description: "DB Service's description",
				Optional:    true,
			},
			"tenant_id": {
				Type:        schema.TypeString,
				Description: "The tenant-id for the DB Service",
				Computed:    true,
			},
			"subscription": {
				Type:        schema.TypeString,
				Description: "Tessell Subscription in which the DB Service is to be created",
				Required:    true,
				ForceNew:    true,
			},
			"engine_type": {
				Type:        schema.TypeString,
				Description: "",
				Required:    true,
				ForceNew:    true,
			},
			"topology": {
				Type:        schema.TypeString,
				Description: "",
				Required:    true,
			},
			"num_of_instances": {
				Type:        schema.TypeInt,
				Description: "Number of instance (nodes) to be created for the DB Service. This is a required input for Apache Kafka. For all other engines, this input would be ignored even if specified.",
				Computed:    true,
			},
			"status": {
				Type:        schema.TypeString,
				Description: "The current status of the DB Service",
				Computed:    true,
			},
			"context_info": {
				Type:        schema.TypeList,
				Description: "Provide more context of DB Service state",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"sub_status": {
							Type:        schema.TypeString,
							Description: "",
							Optional:    true,
							ForceNew:    true,
						},
						"description": {
							Type:        schema.TypeString,
							Description: "",
							Optional:    true,
							ForceNew:    true,
						},
					},
				},
			},
			"license_type": {
				Type:        schema.TypeString,
				Description: "DB Service License Type",
				Computed:    true,
			},
			"edition": {
				Type:        schema.TypeString,
				Description: "",
				Optional:    true,
				ForceNew:    true,
			},
			"software_image": {
				Type:        schema.TypeString,
				Description: "Software Image to be used to create the DB Service",
				Required:    true,
				ForceNew:    true,
			},
			"software_image_version": {
				Type:        schema.TypeString,
				Description: "Software Image Version to be used to create the DB Service",
				Required:    true,
				ForceNew:    true,
			},
			"software_image_version_family": {
				Type:        schema.TypeString,
				Description: "Software Image Family DB Service belongs to",
				Computed:    true,
			},
			"auto_minor_version_update": {
				Type:        schema.TypeBool,
				Description: "Specify whether to automatically update minor version for DB Service",
				Optional:    true,
				Default:     true,
			},
			"enable_deletion_protection": {
				Type:        schema.TypeBool,
				Description: "Specify whether to enable deletion protection for the DB Service",
				Optional:    true,
				Default:     true,
			},
			"enable_stop_protection": {
				Type:        schema.TypeBool,
				Description: "This field specifies whether to enable stop protection for the DB Service. If this is enabled, the stop for the DB Service would be disallowed until this setting is disabled.",
				Optional:    true,
				Default:     false,
			},
			"owner": {
				Type:        schema.TypeString,
				Description: "DB Service owner email address",
				Computed:    true,
			},
			"logged_in_user_role": {
				Type:        schema.TypeString,
				Description: "Access role for the currently logged in user",
				Computed:    true,
			},
			"date_created": {
				Type:        schema.TypeString,
				Description: "Timestamp when the DB Service was created at",
				Computed:    true,
			},
			"started_at": {
				Type:        schema.TypeString,
				Description: "Timestamp when the DB Service was last started at",
				Computed:    true,
			},
			"stopped_at": {
				Type:        schema.TypeString,
				Description: "Timestamp when the DB Service was last stopped at",
				Computed:    true,
			},
			"cloned_from_info": {
				Type:        schema.TypeList,
				Description: "If the DB Service is created as a clone from some other DB Service, this section describes the parent DB Service and cloning details",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"clone_type": {
							Type:        schema.TypeString,
							Description: "",
							Optional:    true,
						},
						"content_type": {
							Type:        schema.TypeString,
							Description: "",
							Optional:    true,
						},
						"tessell_service_id": {
							Type:        schema.TypeString,
							Description: "The DB Service ID using which this DB Service clone is created",
							Optional:    true,
							ForceNew:    true,
						},
						"availability_machine_id": {
							Type:        schema.TypeString,
							Description: "The Availability Machine ID using which this DB Service clone is created",
							Optional:    true,
							ForceNew:    true,
						},
						"tessell_service": {
							Type:        schema.TypeString,
							Description: "The DB Service name using which this DB Service clone is created",
							Optional:    true,
							ForceNew:    true,
						},
						"availability_machine": {
							Type:        schema.TypeString,
							Description: "The Availability Machine name using which this DB Service clone is created",
							Optional:    true,
							ForceNew:    true,
						},
						"snapshot_name": {
							Type:        schema.TypeString,
							Description: "The snapshot using which this DB Service clone is created",
							Optional:    true,
							ForceNew:    true,
						},
						"snapshot_id": {
							Type:        schema.TypeString,
							Description: "The snapshot ID using which this DB Service clone is created",
							Optional:    true,
							ForceNew:    true,
						},
						"snapshot_time": {
							Type:        schema.TypeString,
							Description: "DB Service snapshot capture time",
							Optional:    true,
							ForceNew:    true,
						},
						"pitr_time": {
							Type:        schema.TypeString,
							Description: "If the database was created using a Point-In-Time mechanism, it specifies the timestamp in UTC",
							Optional:    true,
							ForceNew:    true,
						},
						"maximum_recoverability": {
							Type:        schema.TypeBool,
							Description: "If the service was created using a maximum recoverability from the parent service",
							Optional:    true,
							ForceNew:    true,
							Default:     false,
						},
					},
				},
			},
			"infrastructure": {
				Type:        schema.TypeList,
				Description: "This field contains DB Service's infrastructure related information, like, where the service is hosted - cloud, region; what compute shape, or network is is configured with.",
				Required:    true,
				MaxItems:    1,
				MinItems:    1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"cloud": {
							Type:        schema.TypeString,
							Description: "The cloud-type in which the DB Service is provisioned (ex. aws, azure)",
							Optional:    true,
							ForceNew:    true,
						},
						"region": {
							Type:        schema.TypeString,
							Description: "The region in which the DB Service provisioned",
							Optional:    true,
							Computed:    true,
						},
						"availability_zone": {
							Type:        schema.TypeString,
							Description: "The availability-zone in which the DB Service is provisioned",
							Optional:    true,
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
						"vpc": {
							Type:        schema.TypeString,
							Description: "The VPC to be used for provisioning the DB Service",
							Optional:    true,
							Computed:    true,
						},
						"private_subnet": {
							Type:        schema.TypeString,
							Description: "The private subnet to be used for provisioning the compute resource",
							Optional:    true,
						},
						"enable_encryption": {
							Type:        schema.TypeBool,
							Description: "",
							Optional:    true,
							ForceNew:    true,
							Default:     false,
							DiffSuppressFunc: func(k, old, new string, d *schema.ResourceData) bool {
								if old != "" {
									clonedFromDatabaseId := d.GetRawState().GetAttr("databases").AsValueSlice()[0].GetAttr("cloned_from_info").AsValueSlice()[0].GetAttr("database_id").AsString()
									return clonedFromDatabaseId != ""
								}
								return false
							},
						},
						"encryption_key": {
							Type:        schema.TypeString,
							Description: "The encryption key name which is used to encrypt the data at rest",
							Optional:    true,
							ForceNew:    true,
							DiffSuppressFunc: func(k, old, new string, d *schema.ResourceData) bool {
								if old != "" {
									encryptionKey := d.Get(k)
									clonedFromDatabaseId := d.GetRawState().GetAttr("databases").AsValueSlice()[0].GetAttr("cloned_from_info").AsValueSlice()[0].GetAttr("database_id").AsString()
									return old == encryptionKey && new == "" && clonedFromDatabaseId != ""
								}
								return false
							},
						},
						"compute_type": {
							Type:        schema.TypeString,
							Description: "The compute-type to be used for provisioning the DB Service",
							Optional:    true,
							Computed:    true,
						},
						"aws_infra_config": {
							Type:        schema.TypeList,
							Description: "",
							Optional:    true,
							ForceNew:    true,
							MaxItems:    1,
							MinItems:    1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"aws_cpu_options": {
										Type:        schema.TypeList,
										Description: "",
										Optional:    true,
										ForceNew:    true,
										MaxItems:    1,
										MinItems:    1,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"vcpus": {
													Type:        schema.TypeInt,
													Description: "Number of vcpus for aws cpu options",
													Optional:    true,
													ForceNew:    true,
												},
											},
										},
									},
								},
							},
						},
						"storage": {
							Type:        schema.TypeInt,
							Description: "The storage (in bytes) that has been provisioned for the DB Service",
							Computed:    true,
						},
						"additional_storage": {
							Type:        schema.TypeInt,
							Description: "Storage in bytes that is over and above the storage included with compute. This is maintained for backward compatibility and would be deprecated soon.",
							Optional:    true,
							ForceNew:    true,
							Default:     0,
						},
						"enable_compute_sharing": {
							Type:        schema.TypeBool,
							Description: "Specify if the computes should be shared across DB Services",
							Optional:    true,
							ForceNew:    true,
							Default:     false,
						},
						"timezone": {
							Type:        schema.TypeString,
							Description: "The timezone detail",
							Optional:    true,
							ForceNew:    true,
							Default:     "UTC",
						},
						"multi_disk": {
							Type:        schema.TypeBool,
							Description: "Specify whether the DB service uses multiple data disks",
							Computed:    true,
						},
						"iops": {
							Type:        schema.TypeInt,
							Description: "IOPS requested for the DB Service",
							Optional:    true,
							Computed:    true,
							DiffSuppressFunc: func(k, old, new string, d *schema.ResourceData) bool {
								if old != "0" && new == "0" && !d.GetRawState().IsNull() {
									return true
								}
								return false
							},
						},
						"throughput": {
							Type:        schema.TypeInt,
							Description: "throughput requested for the DB Service",
							Optional:    true,
							Computed:    true,
							DiffSuppressFunc: func(k, old, new string, d *schema.ResourceData) bool {
								if old != "0" && new == "0" && !d.GetRawState().IsNull() {
									return true
								}
								return false
							},
						},
						"compute_name_prefix": {
							Type:        schema.TypeString,
							Description: "If not specified, it will be autogenerated",
							Optional:    true,
							ForceNew:    true,
						},
						"computes": {
							Type:        schema.TypeList,
							Description: "",
							Optional:    true,
							Computed:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"name": {
										Type:        schema.TypeString,
										Description: "",
										Optional:    true,
									},
									"instance_group_name": {
										Type:        schema.TypeString,
										Description: "",
										Optional:    true,
									},
									"region": {
										Type:        schema.TypeString,
										Description: "The region in which the compute is to be provisioned",
										Optional:    true,
										ForceNew:    true,
									},
									"availability_zone": {
										Type:        schema.TypeString,
										Description: "The availability-zone in which the compute is to be provisioned",
										Optional:    true,
										ForceNew:    true,
									},
									"role": {
										Type:        schema.TypeString,
										Description: "",
										Optional:    true,
										ForceNew:    true,
									},
									"vpc": {
										Type:        schema.TypeString,
										Description: "The VPC to be used for provisioning the compute resource",
										Optional:    true,
										ForceNew:    true,
									},
									"private_subnet": {
										Type:        schema.TypeString,
										Description: "The private subnet to be used for provisioning the compute resource",
										Optional:    true,
									},
									"compute_type": {
										Type:        schema.TypeString,
										Description: "The compute-type to be used for provisioning the compute resource",
										Optional:    true,
										ForceNew:    true,
									},
									"compute_id": {
										Type:        schema.TypeString,
										Description: "Specify the compute resource if it has to be shared",
										Optional:    true,
										ForceNew:    true,
									},
									"timezone": {
										Type:        schema.TypeString,
										Description: "The timezone detail",
										Optional:    true,
										ForceNew:    true,
										Default:     "UTC",
									},
									"storage_config": {
										Type:        schema.TypeList,
										Description: "The storage details to be provisioned.",
										Optional:    true,
										ForceNew:    true,
										MaxItems:    1,
										MinItems:    1,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"provider": {
													Type:        schema.TypeString,
													Description: "",
													Required:    true,
													ForceNew:    true,
												},
												"fsx_net_app_config": {
													Type:        schema.TypeList,
													Description: "The FSx NetApp details to be provisioned",
													Optional:    true,
													ForceNew:    true,
													MaxItems:    1,
													MinItems:    1,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"file_system_id": {
																Type:        schema.TypeString,
																Description: "File System Id of the FSx NetApp registered with Tessell",
																Required:    true,
																ForceNew:    true,
															},
															"svm_id": {
																Type:        schema.TypeString,
																Description: "Storage Virtual Machine Id of the FSx NetApp registered with Tessell",
																Required:    true,
																ForceNew:    true,
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
						"storage_provider": {
							Type:        schema.TypeString,
							Description: "",
							Computed:    true,
						},
					},
				},
			},
			"refresh_info": {
				Type:        schema.TypeList,
				Description: "Service refresh details",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"content_type": {
							Type:        schema.TypeString,
							Description: "",
							Computed:    true,
						},
						"snapshot_name": {
							Type:        schema.TypeString,
							Description: "",
							Computed:    true,
						},
						"snapshot_time": {
							Type:        schema.TypeString,
							Description: "Time at which snapshot is created.",
							Computed:    true,
						},
						"pitr": {
							Type:        schema.TypeString,
							Description: "",
							Computed:    true,
						},
						"script_info": {
							Type:        schema.TypeList,
							Description: "",
							Computed:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"pre_script_info": {
										Type:        schema.TypeList,
										Description: "",
										Computed:    true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"script_id": {
													Type:        schema.TypeString,
													Description: "The Tessell Script ID",
													Computed:    true,
												},
												"script_version": {
													Type:        schema.TypeString,
													Description: "The Tessell Script version",
													Computed:    true,
												},
											},
										},
									},
									"post_script_info": {
										Type:        schema.TypeList,
										Description: "",
										Computed:    true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"script_id": {
													Type:        schema.TypeString,
													Description: "The Tessell Script ID",
													Computed:    true,
												},
												"script_version": {
													Type:        schema.TypeString,
													Description: "The Tessell Script version",
													Computed:    true,
												},
											},
										},
									},
								},
							},
						},
						"schedule_id": {
							Type:        schema.TypeString,
							Description: "If refreshed using schedule then schedule id, else null",
							Computed:    true,
						},
						"last_successful_refresh_time": {
							Type:        schema.TypeString,
							Description: "Time at which refresh would be successfully completed.",
							Computed:    true,
						},
					},
				},
			},
			"service_connectivity": {
				Type:        schema.TypeList,
				Description: "DB Service's connectivity information",
				Required:    true,
				MaxItems:    1,
				MinItems:    1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"enable_ssl": {
							Type:        schema.TypeBool,
							Description: "Specify whether to enable SSL to the DB Service, default false",
							Optional:    true,
							ForceNew:    true,
							Default:     false,
						},
						"ca_cert_id": {
							Type:        schema.TypeString,
							Description: "The CA certificate ID associated with the DB Service",
							Computed:    true,
						},
						"dns_prefix": {
							Type:        schema.TypeString,
							Description: "DNS Prefix associated with the DB Service",
							Optional:    true,
							DiffSuppressFunc: func(k, old, new string, d *schema.ResourceData) bool {
								dnsPrefix := d.Get(k)
								if old != "" && new == "" && !d.GetRawState().IsNull() {
									dnsPrefixInState := d.GetRawState().GetAttr("service_connectivity").AsValueSlice()[0].GetAttr("dns_prefix").AsString()
									if dnsPrefix == dnsPrefixInState {
										return true
									}
								}
								return false
							},
						},
						"service_port": {
							Type:        schema.TypeInt,
							Description: "The connection port for the DB Service",
							Optional:    true,
							ForceNew:    true,
						},
						"enable_public_access": {
							Type:        schema.TypeBool,
							Description: "Specify whether to enable public access to the DB Service, default false",
							Optional:    true,
						},
						"allowed_ip_addresses": {
							Type:        schema.TypeList,
							Description: "The list of allowed ipv4 addresses that can connect to the DB Service",
							Optional:    true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"connect_strings": {
							Type:        schema.TypeList,
							Description: "The list of connect strings for the DB Service",
							Computed:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"type": {
										Type:        schema.TypeString,
										Description: "",
										Optional:    true,
										ForceNew:    true,
									},
									"usage_type": {
										Type:        schema.TypeString,
										Description: "",
										Optional:    true,
										ForceNew:    true,
									},
									"connect_descriptor": {
										Type:        schema.TypeString,
										Description: "The connection description for the DB Service",
										Optional:    true,
										ForceNew:    true,
									},
									"endpoint": {
										Type:        schema.TypeString,
										Description: "The connection end point for the DB Service",
										Optional:    true,
										ForceNew:    true,
									},
									"master_user": {
										Type:        schema.TypeString,
										Description: "The master user name for the DB Service",
										Optional:    true,
										ForceNew:    true,
									},
									"service_port": {
										Type:        schema.TypeInt,
										Description: "The connection port for the DB Service",
										Optional:    true,
										ForceNew:    true,
									},
								},
							},
						},
						"private_link": {
							Type:        schema.TypeList,
							Description: "The interface endpoint or Gateway Load Balancer endpoint to connect to your DB service.",
							Computed:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"status": {
										Type:        schema.TypeString,
										Description: "",
										Optional:    true,
										ForceNew:    true,
									},
									"service_principals": {
										Type:        schema.TypeList,
										Description: "The list of AWS account principals that are currently enabled",
										Computed:    true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"endpoint_service_name": {
										Type:        schema.TypeString,
										Description: "The configured endpoint as a result of configuring the service-principals",
										Computed:    true,
									},
									"client_azure_subscription_ids": {
										Type:        schema.TypeList,
										Description: "The list of Azure subscription Ids",
										Optional:    true,
										ForceNew:    true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"private_link_service_alias": {
										Type:        schema.TypeString,
										Description: "The Azure private link service alias",
										Optional:    true,
										ForceNew:    true,
									},
								},
							},
						},
						"computes_connectivity": {
							Type:        schema.TypeList,
							Description: "The Genie endpoint to connect to your DB service.",
							Computed:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"compute_resource_id": {
										Type:        schema.TypeString,
										Description: "",
										Computed:    true,
									},
									"port_access_config": {
										Type:        schema.TypeList,
										Description: "",
										Computed:    true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"port": {
													Type:        schema.TypeInt,
													Description: "The connection port for the DB Service",
													Computed:    true,
												},
												"enable_public_access": {
													Type:        schema.TypeBool,
													Description: "Enable public access to database (true/false)",
													Computed:    true,
												},
												"allowed_ip_addresses": {
													Type:        schema.TypeList,
													Description: "Set allowed IP address if enablePublicAccess is true.",
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
						"update_in_progress_info": {
							Type:        schema.TypeList,
							Description: "DB Service connectivity update-in-progress details",
							Computed:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"dns_prefix": {
										Type:        schema.TypeString,
										Description: "The DNS prefix associated with the DB Service",
										Optional:    true,
										ForceNew:    true,
									},
									"enable_public_access": {
										Type:        schema.TypeBool,
										Description: "Specify whether to enable public access to the DB Service, default false",
										Optional:    true,
										ForceNew:    true,
									},
									"allowed_ip_addresses": {
										Type:        schema.TypeList,
										Description: "The list of allowed ipv4 addresses that can connect to the DB Service",
										Optional:    true,
										ForceNew:    true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"private_link": {
										Type:        schema.TypeList,
										Description: "The interface endpoint or Gateway Load Balancer endpoint to connect to your DB service.",
										Optional:    true,
										ForceNew:    true,
										MaxItems:    1,
										MinItems:    1,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"service_principals": {
													Type:        schema.TypeList,
													Description: "The list of AWS account principals that are currently enabled",
													Optional:    true,
													ForceNew:    true,
													Elem: &schema.Schema{
														Type: schema.TypeString,
													},
												},
												"client_azure_subscription_ids": {
													Type:        schema.TypeList,
													Description: "The list of Azure subscription Ids",
													Optional:    true,
													ForceNew:    true,
													Elem: &schema.Schema{
														Type: schema.TypeString,
													},
												},
											},
										},
									},
									"computes_connectivity": {
										Type:        schema.TypeList,
										Description: "The Genie endpoint to connect to your DB service.",
										Computed:    true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"compute_resource_id": {
													Type:        schema.TypeString,
													Description: "",
													Computed:    true,
												},
												"port_access_config": {
													Type:        schema.TypeList,
													Description: "",
													Computed:    true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"port": {
																Type:        schema.TypeInt,
																Description: "The connection port for the DB Service",
																Computed:    true,
															},
															"enable_public_access": {
																Type:        schema.TypeBool,
																Description: "Enable public access to database (true/false)",
																Computed:    true,
															},
															"allowed_ip_addresses": {
																Type:        schema.TypeList,
																Description: "Set allowed IP address if enablePublicAccess is true.",
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
			"tessell_genie_status": {
				Type:        schema.TypeString,
				Description: "DB Service's Genie status",
				Computed:    true,
			},
			"creds": {
				Type:        schema.TypeList,
				Description: "DB Service's credential details",
				Required:    true,
				ForceNew:    true,
				MaxItems:    1,
				MinItems:    1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"master_user": {
							Type:        schema.TypeString,
							Description: "DB Service's master username",
							Required:    true,
						},
						"master_password": {
							Type:        schema.TypeString,
							Description: "DB Service's master password",
							Required:    true,
						},
					},
				},
			},
			"maintenance_window": {
				Type:        schema.TypeList,
				Description: "This field details the DB Service maintenance related details.",
				Optional:    true,
				ForceNew:    true,
				MaxItems:    1,
				MinItems:    1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"day": {
							Type:        schema.TypeString,
							Description: "",
							Required:    true,
							ForceNew:    true,
						},
						"time": {
							Type:        schema.TypeString,
							Description: "Time value in (hh:mm) format. ex. '02:00'",
							Required:    true,
							ForceNew:    true,
						},
						"duration": {
							Type:        schema.TypeInt,
							Description: "The duration during which the maintenance window will be allowed to trigger",
							Required:    true,
							ForceNew:    true,
						},
					},
				},
			},
			"snapshot_configuration": {
				Type:        schema.TypeList,
				Description: "DB Service's snapshot retention configurations. If not specified, the default recommended retention configurations would be applied.",
				Optional:    true,
				MaxItems:    1,
				MinItems:    1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"snapshot_window": {
							Type:        schema.TypeList,
							Description: "",
							Optional:    true,
							MaxItems:    1,
							MinItems:    1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"time": {
										Type:        schema.TypeString,
										Description: "Time value in (hh:mm) format. ex. '02:00'. Deprecated, please use backupStartTime in schedule.",
										Optional:    true,
									},
								},
							},
						},
						"sla": {
							Type:        schema.TypeString,
							Description: "The snapshot SLA for the DB Service. If not specified, a default SLA would be associated with the DB Service",
							Optional:    true,
						},
						"schedule": {
							Type:        schema.TypeList,
							Description: "Schedule Information",
							Optional:    true,
							ForceNew:    true,
							MaxItems:    1,
							MinItems:    1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"backup_start_time": {
										Type:        schema.TypeList,
										Description: "Clock time format value in hour and minute.",
										Optional:    true,
										ForceNew:    true,
										MaxItems:    1,
										MinItems:    1,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"hour": {
													Type:        schema.TypeInt,
													Description: "",
													Optional:    true,
													ForceNew:    true,
													Default:     1,
												},
												"minute": {
													Type:        schema.TypeInt,
													Description: "",
													Optional:    true,
													ForceNew:    true,
													Default:     0,
												},
											},
										},
									},
									"daily_schedule": {
										Type:        schema.TypeList,
										Description: "",
										Optional:    true,
										ForceNew:    true,
										MaxItems:    1,
										MinItems:    1,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"backups_per_day": {
													Type:        schema.TypeInt,
													Description: "The number of backups to be captured per day.",
													Optional:    true,
													ForceNew:    true,
												},
											},
										},
									},
									"weekly_schedule": {
										Type:        schema.TypeList,
										Description: "",
										Optional:    true,
										ForceNew:    true,
										MaxItems:    1,
										MinItems:    1,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"days": {
													Type:        schema.TypeList,
													Description: "Days in a week to retain weekly backups for",
													Optional:    true,
													ForceNew:    true,
													Elem: &schema.Schema{
														Type: schema.TypeString,
													},
												},
											},
										},
									},
									"monthly_schedule": {
										Type:        schema.TypeList,
										Description: "Definition for taking month specific schedule.",
										Optional:    true,
										ForceNew:    true,
										MaxItems:    1,
										MinItems:    1,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"common_schedule": {
													Type:        schema.TypeList,
													Description: "",
													Optional:    true,
													ForceNew:    true,
													MaxItems:    1,
													MinItems:    1,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"dates": {
																Type:        schema.TypeList,
																Description: "Dates in a month to retain monthly backups",
																Optional:    true,
																ForceNew:    true,
																Elem: &schema.Schema{
																	Type: schema.TypeInt,
																},
															},
															"last_day_of_month": {
																Type:        schema.TypeBool,
																Description: "",
																Optional:    true,
																ForceNew:    true,
																Default:     false,
															},
														},
													},
												},
											},
										},
									},
									"yearly_schedule": {
										Type:        schema.TypeList,
										Description: "",
										Optional:    true,
										ForceNew:    true,
										MaxItems:    1,
										MinItems:    1,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"common_schedule": {
													Type:        schema.TypeList,
													Description: "",
													Optional:    true,
													ForceNew:    true,
													MaxItems:    1,
													MinItems:    1,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"dates": {
																Type:        schema.TypeList,
																Description: "Dates in a month to retain monthly backups",
																Optional:    true,
																ForceNew:    true,
																Elem: &schema.Schema{
																	Type: schema.TypeInt,
																},
															},
															"last_day_of_month": {
																Type:        schema.TypeBool,
																Description: "",
																Optional:    true,
																ForceNew:    true,
																Default:     false,
															},
															"months": {
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
												"month_specific_schedule": {
													Type:        schema.TypeList,
													Description: "",
													Optional:    true,
													ForceNew:    true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"month": {
																Type:        schema.TypeString,
																Description: "Name of a month",
																Required:    true,
																ForceNew:    true,
															},
															"dates": {
																Type:        schema.TypeList,
																Description: "",
																Required:    true,
																ForceNew:    true,
																Elem: &schema.Schema{
																	Type: schema.TypeInt,
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
						"full_backup_schedule": {
							Type:        schema.TypeList,
							Description: "The schedule at which full backups would be triggered",
							Optional:    true,
							ForceNew:    true,
							MaxItems:    1,
							MinItems:    1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"weekly_schedule": {
										Type:        schema.TypeList,
										Description: "",
										Optional:    true,
										ForceNew:    true,
										MaxItems:    1,
										MinItems:    1,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"days": {
													Type:        schema.TypeList,
													Description: "Days in a week to retain weekly backups for",
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
					},
				},
			},
			"engine_configuration": {
				Type:        schema.TypeList,
				Description: "This field details the DB Service engine configuration details like - parameter profile, or options profile (if applicable) are used to configure the DB Service.",
				Required:    true,
				ForceNew:    true,
				MaxItems:    1,
				MinItems:    1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"oracle_config": {
							Type:        schema.TypeList,
							Description: "",
							Optional:    true,
							ForceNew:    true,
							MaxItems:    1,
							MinItems:    1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"multi_tenant": {
										Type:        schema.TypeBool,
										Description: "Specify whether the DB Service is multi-tenant.",
										Optional:    true,
										ForceNew:    true,
										Default:     false,
									},
									"parameter_profile_id": {
										Type:        schema.TypeString,
										Description: "The parameter profile id for the database",
										Optional:    true,
										ForceNew:    true,
									},
									"options_profile": {
										Type:        schema.TypeString,
										Description: "The options profile for the database",
										Optional:    true,
										ForceNew:    true,
									},
									"sid": {
										Type:        schema.TypeString,
										Description: "SID for oracle database",
										Optional:    true,
										Computed:    true,
									},
									"character_set": {
										Type:        schema.TypeString,
										Description: "The character-set for the database",
										Optional:    true,
										ForceNew:    true,
									},
									"national_character_set": {
										Type:        schema.TypeString,
										Description: "The national-character-set for the database",
										Optional:    true,
										ForceNew:    true,
									},
									"enable_archive_mode": {
										Type:        schema.TypeBool,
										Description: "To explicitly enable archive mode, when PITR is disabled",
										Optional:    true,
										ForceNew:    true,
										Default:     true,
									},
								},
							},
						},
						"postgresql_config": {
							Type:        schema.TypeList,
							Description: "",
							Optional:    true,
							ForceNew:    true,
							MaxItems:    1,
							MinItems:    1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"parameter_profile_id": {
										Type:        schema.TypeString,
										Description: "The parameter profile ID for the database",
										Optional:    true,
										ForceNew:    true,
									},
									"ad_domain_id": {
										Type:        schema.TypeString,
										Description: "Active Directory Domain ID",
										Optional:    true,
										ForceNew:    true,
									},
									"proxy_port": {
										Type:        schema.TypeInt,
										Description: "",
										Optional:    true,
										ForceNew:    true,
									},
								},
							},
						},
						"mysql_config": {
							Type:        schema.TypeList,
							Description: "",
							Optional:    true,
							ForceNew:    true,
							MaxItems:    1,
							MinItems:    1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"parameter_profile_id": {
										Type:        schema.TypeString,
										Description: "The parameter profile ID for the database",
										Optional:    true,
										ForceNew:    true,
									},
									"ad_domain_id": {
										Type:        schema.TypeString,
										Description: "Active Directory Domain ID",
										Optional:    true,
										ForceNew:    true,
									},
								},
							},
						},
						"sql_server_config": {
							Type:        schema.TypeList,
							Description: "",
							Optional:    true,
							ForceNew:    true,
							MaxItems:    1,
							MinItems:    1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"parameter_profile_id": {
										Type:        schema.TypeString,
										Description: "The parameter profile ID for the database",
										Optional:    true,
										ForceNew:    true,
									},
									"ad_domain_id": {
										Type:        schema.TypeString,
										Description: "Active Directory Domain ID",
										Optional:    true,
										ForceNew:    true,
									},
								},
							},
						},
						"apache_kafka_config": {
							Type:        schema.TypeList,
							Description: "",
							Optional:    true,
							ForceNew:    true,
							MaxItems:    1,
							MinItems:    1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"parameter_profile_id": {
										Type:        schema.TypeString,
										Description: "The parameter profile id for the database",
										Optional:    true,
										ForceNew:    true,
									},
								},
							},
						},
						"mongodb_config": {
							Type:        schema.TypeList,
							Description: "",
							Optional:    true,
							ForceNew:    true,
							MaxItems:    1,
							MinItems:    1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"cluster_name": {
										Type:        schema.TypeString,
										Description: "The MongoDB Cluster name",
										Optional:    true,
										ForceNew:    true,
									},
									"parameter_profile_id": {
										Type:        schema.TypeString,
										Description: "The parameter profile ID for the database",
										Optional:    true,
										ForceNew:    true,
									},
								},
							},
						},
						"milvus_config": {
							Type:        schema.TypeList,
							Description: "",
							Optional:    true,
							ForceNew:    true,
							MaxItems:    1,
							MinItems:    1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"parameter_profile_id": {
										Type:        schema.TypeString,
										Description: "The parameter profile ID for the database",
										Optional:    true,
										ForceNew:    true,
									},
								},
							},
						},
						"pre_script_info": {
							Type:        schema.TypeList,
							Description: "",
							Optional:    true,
							ForceNew:    true,
							MaxItems:    1,
							MinItems:    1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"script_id": {
										Type:        schema.TypeString,
										Description: "The Tessell Script ID",
										Optional:    true,
										ForceNew:    true,
									},
									"script_version": {
										Type:        schema.TypeString,
										Description: "The Tessell Script version",
										Optional:    true,
										ForceNew:    true,
									},
								},
							},
						},
						"post_script_info": {
							Type:        schema.TypeList,
							Description: "",
							Optional:    true,
							ForceNew:    true,
							MaxItems:    1,
							MinItems:    1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"script_id": {
										Type:        schema.TypeString,
										Description: "The Tessell Script ID",
										Optional:    true,
										ForceNew:    true,
									},
									"script_version": {
										Type:        schema.TypeString,
										Description: "The Tessell Script version",
										Optional:    true,
										ForceNew:    true,
									},
								},
							},
						},
					},
				},
			},
			"databases": {
				Type:        schema.TypeList,
				Description: "Databases that are part of this DB Service",
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"source_database_id": {
							Type:        schema.TypeString,
							Description: "Required while creating a clone. It specifies the Id of the source database from which the clone is being created.",
							Optional:    true,
							ForceNew:    true,
							DiffSuppressFunc: func(k, old, new string, d *schema.ResourceData) bool {
								sourceDatabaseId := d.Get(k)
								if old == "" && new == sourceDatabaseId && !d.GetRawState().IsNull() {
									clonedFromInfoNested := d.GetRawState().GetAttr("databases").AsValueSlice()[0].GetAttr("cloned_from_info").AsValueSlice()
									if len(clonedFromInfoNested) == 0 {
										return true
									}
									clonedFromDatabaseId := clonedFromInfoNested[0].GetAttr("database_id").AsString()
									if sourceDatabaseId == clonedFromDatabaseId {
										return true
									}
								}
								return false
							},
						},
						"id": {
							Type:        schema.TypeString,
							Description: "Tessell generated UUID for the database",
							Computed:    true,
						},
						"database_name": {
							Type:        schema.TypeString,
							Description: "Database name",
							Optional:    true,
							Computed:    true, // in case of oracle
						},
						"description": {
							Type:        schema.TypeString,
							Description: "Database description",
							Optional:    true,
						},
						"tessell_service_id": {
							Type:        schema.TypeString,
							Description: "Associated DB Service ID",
							Computed:    true,
						},
						"engine_type": {
							Type:        schema.TypeString,
							Description: "Database Engine Type",
							Computed:    true,
						},
						"status": {
							Type:        schema.TypeString,
							Description: "Database status",
							Computed:    true,
						},
						"date_created": {
							Type:        schema.TypeString,
							Description: "Timestamp when the entity was created",
							Computed:    true,
						},
						"tessell_created": {
							Type:        schema.TypeBool,
							Description: "Database created from Tessell platform",
							Optional:    true,
							Computed:    true,
						},
						"cloned_from_info": {
							Type:        schema.TypeList,
							Description: "If a database is created as a clone from some other DB Service's database, this section describes the original database details",
							Computed:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"database_id": {
										Type:        schema.TypeString,
										Description: "The original database Id using which this database clone is created",
										Computed:    true,
									},
								},
							},
						},
						"database_configuration": {
							Type:        schema.TypeList,
							Description: "",
							Optional:    true,
							ForceNew:    true,
							MaxItems:    1,
							MinItems:    1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"oracle_config": {
										Type:        schema.TypeList,
										Description: "",
										Optional:    true,
										ForceNew:    true,
										MaxItems:    1,
										MinItems:    1,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"parameter_profile_id": {
													Type:        schema.TypeString,
													Description: "The parameter profile id for the database",
													Optional:    true,
													ForceNew:    true,
												},
												"options_profile": {
													Type:        schema.TypeString,
													Description: "The options profile for the database",
													Optional:    true,
													ForceNew:    true,
												},
												"username": {
													Type:        schema.TypeString,
													Description: "Username for the oracle database",
													Optional:    true,
													ForceNew:    true,
												},
											},
										},
									},
									"postgresql_config": {
										Type:        schema.TypeList,
										Description: "",
										Optional:    true,
										ForceNew:    true,
										MaxItems:    1,
										MinItems:    1,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"parameter_profile_id": {
													Type:        schema.TypeString,
													Description: "The parameter profile ID for the database",
													Optional:    true,
													ForceNew:    true,
												},
											},
										},
									},
									"mysql_config": {
										Type:        schema.TypeList,
										Description: "",
										Optional:    true,
										ForceNew:    true,
										MaxItems:    1,
										MinItems:    1,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"parameter_profile_id": {
													Type:        schema.TypeString,
													Description: "The parameter profile ID for the database",
													Optional:    true,
													ForceNew:    true,
												},
											},
										},
									},
									"sql_server_config": {
										Type:        schema.TypeList,
										Description: "",
										Optional:    true,
										ForceNew:    true,
										MaxItems:    1,
										MinItems:    1,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"parameter_profile_id": {
													Type:        schema.TypeString,
													Description: "The parameter profile ID for the database",
													Optional:    true,
													ForceNew:    true,
												},
											},
										},
									},
									"mongodb_config": {
										Type:        schema.TypeList,
										Description: "",
										Optional:    true,
										ForceNew:    true,
										MaxItems:    1,
										MinItems:    1,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"parameter_profile_id": {
													Type:        schema.TypeString,
													Description: "The parameter profile ID for the database",
													Optional:    true,
													ForceNew:    true,
												},
											},
										},
									},
									"milvus_config": {
										Type:        schema.TypeList,
										Description: "",
										Optional:    true,
										ForceNew:    true,
										MaxItems:    1,
										MinItems:    1,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"parameter_profile_id": {
													Type:        schema.TypeString,
													Description: "The parameter profile ID for the database",
													Optional:    true,
													ForceNew:    true,
												},
											},
										},
									},
								},
							},
						},
						"connect_string": {
							Type:        schema.TypeList,
							Description: "",
							Computed:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"connect_descriptor": {
										Type:        schema.TypeString,
										Description: "",
										Computed:    true,
									},
									"master_user": {
										Type:        schema.TypeString,
										Description: "",
										Computed:    true,
									},
									"endpoint": {
										Type:        schema.TypeString,
										Description: "",
										Computed:    true,
									},
									"service_port": {
										Type:        schema.TypeString,
										Description: "",
										Computed:    true,
									},
								},
							},
						},
					},
				},
			},
			"integrations_config": {
				Type:        schema.TypeList,
				Description: "Integrations to be enabled for the DB Service",
				Optional:    true,
				MaxItems:    1,
				MinItems:    1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"integrations": {
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
						integrations := d.Get("integrations_config.0.integrations")
						if len(integrations.([]interface{})) == 0 {
							return true
						}
					} else if old == "1" && new == "0" {
						integrationsConfig := d.GetRawState().GetAttr("integrations_config").AsValueSlice()[0]
						integrations := integrationsConfig.GetAttr("integrations").AsValueSlice()
						if len(integrations) == 0 {
							return true
						}
					}
					return false
				},
			},
			"tags": {
				Type:        schema.TypeList,
				Description: "The tags to be associated with the DB Service",
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:        schema.TypeString,
							Description: "Case sensitive, tag name",
							Optional:    true,
							ForceNew:    true,
						},
						"value": {
							Type:        schema.TypeString,
							Description: "Case sensitive, tag value",
							Optional:    true,
							ForceNew:    true,
						},
					},
				},
				DiffSuppressFunc: func(k, old, new string, d *schema.ResourceData) bool {
					if old == "0" && new == "1" {
						tags := d.Get("tags.0")
						if len(tags.(map[string]interface{})) == 0 {
							return true
						}
					} else if old == "1" && new == "0" {
						tags := d.GetRawState().GetAttr("tags").AsValueSlice()[0].AsValueMap()
						if len(tags) == 0 {
							return true
						}
					}
					return false
				},
			},
			"instances": {
				Type:        schema.TypeList,
				Description: "Instances associated with this DB Service",
				Optional:    true,
				Computed:    true, // TODO: remove this once instances computes are removed from infra
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type:        schema.TypeString,
							Description: "Tessell generated UUID for the DB Service Instance",
							Computed:    true,
						},
						"name": {
							Type:        schema.TypeString,
							Description: "Name of the DB Service Instance",
							Required:    true,
						},
						"instance_group_name": {
							Type:        schema.TypeString,
							Description: "Name of the instance group",
							Required:    true,
						},
						"type": {
							Type:        schema.TypeString,
							Description: "DB Service instance type",
							Computed:    true,
							Optional:    true,
						},
						"role": {
							Type:        schema.TypeString,
							Description: "DB Service instance role",
							Required:    true,
						},
						"status": {
							Type:        schema.TypeString,
							Description: "DB Service instance status",
							Computed:    true,
						},
						"tessell_service_id": {
							Type:        schema.TypeString,
							Description: "DB Service Instance's associated DB Service id",
							Computed:    true,
						},
						"cloud": {
							Type:        schema.TypeString,
							Description: "DB Service Instance's cloud type",
							Computed:    true,
						},
						"region": {
							Type:        schema.TypeString,
							Description: "DB Service Instance's cloud region",
							Required:    true,
						},
						"availability_zone": {
							Type:        schema.TypeString,
							Description: "DB Service Instance's cloud availability zone",
							Optional:    true,
							Computed:    true,
						},
						"instance_group_id": {
							Type:        schema.TypeString,
							Description: "The instance groupd Id",
							Computed:    true,
						},
						"compute_type": {
							Type:        schema.TypeString,
							Description: "The compute used for creation of the Tessell Service Instance",
							Required:    true,
						},
						"aws_infra_config": {
							Type:        schema.TypeList,
							Description: "",
							Optional:    true,
							Computed:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"aws_cpu_options": {
										Type:        schema.TypeList,
										Description: "",
										Optional:    true,
										Computed:    true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"vcpus": {
													Type:        schema.TypeInt,
													Description: "Number of vcpus for aws cpu options",
													Optional:    true,
													Computed:    true,
												},
											},
										},
									},
								},
							},
						},
						"compute_id": {
							Type:        schema.TypeString,
							Description: "The associated compute identifier",
							Optional:    true,
							Computed:    true,
						},
						"compute_name": {
							Type:        schema.TypeString,
							Description: "The associated compute name",
							Computed:    true,
						},
						"storage": {
							Type:        schema.TypeInt,
							Description: "The storage (in bytes) that has been provisioned for the DB Service instance.",
							Computed:    true,
						},
						"data_volume_iops": {
							Type:        schema.TypeInt,
							Description: "",
							Optional:    true,
							Computed:    true,
						},
						"throughput": {
							Type:        schema.TypeInt,
							Description: "Throughput requested for this DB Service instance",
							Optional:    true,
							Computed:    true,
						},
						"enable_perf_insights": {
							Type:        schema.TypeBool,
							Description: "",
							Optional:    true,
						},
						"parameter_profile": {
							Type:        schema.TypeList,
							Description: "",
							Computed:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"id": {
										Type:        schema.TypeString,
										Description: "Tessell generated UUID for the the parameter profile",
										Optional:    true,
										ForceNew:    true,
									},
									"name": {
										Type:        schema.TypeString,
										Description: "The name used to identify the parameter profile",
										Optional:    true,
										ForceNew:    true,
									},
									"version": {
										Type:        schema.TypeString,
										Description: "The version of the parameter profile associated with the instance",
										Optional:    true,
										ForceNew:    true,
									},
									"status": {
										Type:        schema.TypeString,
										Description: "",
										Optional:    true,
										ForceNew:    true,
									},
								},
							},
						},
						"monitoring_config": {
							Type:        schema.TypeList,
							Description: "",
							Computed:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"perf_insights": {
										Type:        schema.TypeList,
										Description: "",
										Optional:    true,
										ForceNew:    true,
										MaxItems:    1,
										MinItems:    1,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"perf_insights_enabled": {
													Type:        schema.TypeBool,
													Description: "",
													Optional:    true,
													ForceNew:    true,
												},
												"monitoring_deployment_id": {
													Type:        schema.TypeString,
													Description: "",
													Optional:    true,
													ForceNew:    true,
												},
											},
										},
									},
								},
							},
						},
						"vpc": {
							Type:        schema.TypeString,
							Description: "The VPC used for creation of the DB Service Instance",
							Required:    true,
						},
						"public_subnet": {
							Type:        schema.TypeString,
							Description: "The public subnet used for creation of the DB Service Instance",
							Computed:    true,
						},
						"private_subnet": {
							Type:        schema.TypeString,
							Description: "The private subnet used for creation of the DB Service Instance",
							Optional:    true,
						},
						"encryption_key": {
							Type:        schema.TypeString,
							Description: "The encryption key name which is used to encrypt the data at rest",
							Optional:    true,
						},
						"software_image": {
							Type:        schema.TypeString,
							Description: "Software Image to be used to create the instance",
							Computed:    true,
						},
						"software_image_version": {
							Type:        schema.TypeString,
							Description: "Software Image Version to be used to create the instance",
							Computed:    true,
						},
						"date_created": {
							Type:        schema.TypeString,
							Description: "Timestamp when the entity was created",
							Computed:    true,
						},
						"connect_string": {
							Type:        schema.TypeList,
							Description: "",
							Computed:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"connect_descriptor": {
										Type:        schema.TypeString,
										Description: "",
										Computed:    true,
									},
									"master_user": {
										Type:        schema.TypeString,
										Description: "",
										Computed:    true,
									},
									"endpoint": {
										Type:        schema.TypeString,
										Description: "",
										Computed:    true,
									},
									"service_port": {
										Type:        schema.TypeString,
										Description: "",
										Computed:    true,
									},
								},
							},
						},
						"updates_in_progress": {
							Type:        schema.TypeList,
							Description: "The updates that are in progress for this resource",
							Computed:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"update_type": {
										Type:        schema.TypeString,
										Description: "Type of the update",
										Computed:    true,
									},
									"reference_id": {
										Type:        schema.TypeString,
										Description: "The reference-id of the update request",
										Computed:    true,
									},
									"submitted_at": {
										Type:        schema.TypeString,
										Description: "Timestamp when the resource update was requested",
										Computed:    true,
									},
									"update_info": {
										Type:        schema.TypeMap,
										Description: "The specific details for a Tessell resource that are being updated",
										Computed:    true,
									},
								},
							},
						},
						"last_started_at": {
							Type:        schema.TypeString,
							Description: "Timestamp when the service instance was last started at",
							Computed:    true,
						},
						"last_stopped_at": {
							Type:        schema.TypeString,
							Description: "Timestamp when the Service Instance was last stopped at",
							Computed:    true,
						},
						"sync_mode": {
							Type:        schema.TypeString,
							Description: "",
							Optional:    true,
						},
						"engine_configuration": {
							Type:        schema.TypeList,
							Description: "This field details the DB Service Instance engine configuration details like - access mode",
							Optional:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"oracle_config": {
										Type:        schema.TypeList,
										Description: "",
										Optional:    true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"access_mode": {
													Type:        schema.TypeString,
													Description: "",
													Optional:    true,
												},
											},
										},
									},
								},
							},
						},
						"storage_config": {
							Type:        schema.TypeList,
							Description: "",
							Optional:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"provider": {
										Type:        schema.TypeString,
										Description: "",
										Optional:    true,
									},
									"fsx_net_app_config": {
										Type:        schema.TypeList,
										Description: "",
										Optional:    true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"file_system_name": {
													Type:        schema.TypeString,
													Description: "",
													Optional:    true,
												},
												"svm_name": {
													Type:        schema.TypeString,
													Description: "",
													Optional:    true,
												},
												"volume_name": {
													Type:        schema.TypeString,
													Description: "",
													Optional:    true,
												},
												"file_system_id": {
													Type:        schema.TypeString,
													Description: "File System Id of the FSx NetApp registered with Tessell",
													Optional:    true,
												},
												"svm_id": {
													Type:        schema.TypeString,
													Description: "Storage Virtual Machine Id of the FSx NetApp registered with Tessell",
													Optional:    true,
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
			"shared_with": {
				Type:        schema.TypeList,
				Description: "Tessell Entity ACL Sharing Info",
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
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"email_id": {
										Type:        schema.TypeString,
										Description: "",
										Optional:    true,
										ForceNew:    true,
									},
									"role": {
										Type:        schema.TypeString,
										Description: "",
										Optional:    true,
										ForceNew:    true,
									},
								},
							},
						},
					},
				},
			},
			"upcoming_scheduled_actions": {
				Type:        schema.TypeList,
				Description: "",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"start_stop": {
							Type:        schema.TypeList,
							Description: "",
							Computed:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"action": {
										Type:        schema.TypeString,
										Description: "Action which can be either start/stop",
										Computed:    true,
									},
									"at": {
										Type:        schema.TypeString,
										Description: "The time at which the specified action is to be performed",
										Computed:    true,
									},
								},
							},
						},
						"patch": {
							Type:        schema.TypeList,
							Description: "",
							Computed:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"at": {
										Type:        schema.TypeString,
										Description: "The updated time at which the specified action is to be performed",
										Computed:    true,
									},
									"message": {
										Type:        schema.TypeString,
										Description: "Details for the update",
										Optional:    true,
										ForceNew:    true,
									},
								},
							},
						},
						"delete": {
							Type:        schema.TypeList,
							Description: "",
							Computed:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"at": {
										Type:        schema.TypeString,
										Description: "The scheduled time for the action to be deleted",
										Computed:    true,
									},
								},
							},
						},
					},
				},
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
				DiffSuppressFunc: func(k, old, new string, d *schema.ResourceData) bool {
					if old == "0" && new == "1" {
						deletionConfig := d.Get("deletion_config")
						if len(deletionConfig.([]interface{})) == 0 {
							return true
						}
					} else if old == "1" && new == "0" {
						deletionConfigFromState := d.GetRawState().GetAttr("deletion_config").AsValueSlice()[0]
						retainAvailabilityMachineIsNullInState := deletionConfigFromState.GetAttr("retain_availability_machine").False()
						if retainAvailabilityMachineIsNullInState {
							return true
						}
					}
					return false
				},
			},
			"deletion_schedule": {
				Type:        schema.TypeList,
				Description: "Details of the deletion schedule on a DB Service",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type:        schema.TypeString,
							Description: "",
							Computed:    true,
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
				},
			},
			"updates_in_progress": {
				Type:        schema.TypeList,
				Description: "The updates that are in progress for this resource",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"update_type": {
							Type:        schema.TypeString,
							Description: "Type of the update",
							Computed:    true,
						},
						"reference_id": {
							Type:        schema.TypeString,
							Description: "The reference-id of the update request",
							Computed:    true,
						},
						"submitted_at": {
							Type:        schema.TypeString,
							Description: "Timestamp when the resource update was requested",
							Computed:    true,
						},
						"update_info": {
							Type:        schema.TypeMap,
							Description: "The specific details for a Tessell resource that are being updated",
							Computed:    true,
						},
					},
				},
			},
			"parent_availability_machine_id": {
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
				Description: "Timeout for terraform polling, when block_until_complete is true (default: true). (In seconds)",
				Optional:    true,
				Default:     7200,
			},
			"expected_status": {
				Type:        schema.TypeString,
				Description: "If provided, invoke the DB Service start/stop API",
				Optional:    true,
				Default:     "READY",
			},
		},
		CustomizeDiff: customdiff.All(
			customdiff.ValidateChange("instances", func(ctx context.Context, oldInstances, newInstances, meta interface{}) error {
				instances := newInstances.([]interface{})
				names := make(map[string]bool)
				primaryCount := 0
				for _, instanceRaw := range instances {
					instance := instanceRaw.(map[string]interface{})
					name := instance["name"].(string)
					role := instance["role"].(string)
					if names[name] {
						return fmt.Errorf("duplicate instance name: %s", name)
					}
					names[name] = true
					if role == "primary" {
						primaryCount++
						if primaryCount > 1 {
							return fmt.Errorf("more than one instance has 'primary' role")
						}
					}
				}
				if primaryCount == 0 && len(instances) != 0 {
					return fmt.Errorf("no instance is marked as 'primary'")
				}
				return nil
			}),
		),
	}
}

func resourceDBServiceCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*apiClient.Client)

	var diags diag.Diagnostics
	var id string

	parentAvailabilityMachineId := d.Get("parent_availability_machine_id").(string)
	snapshotId := d.Get("snapshot_id").(string)
	pitr := d.Get("pitr").(string)

	if snapshotId != "" || pitr != "" {
		payload := formPayloadForCloneTessellService(d)

		response, _, err := client.CloneTessellService(parentAvailabilityMachineId, payload)
		if err != nil {
			return diag.FromErr(err)
		}
		id = *response.ResourceId
	} else {
		payload := formPayloadForProvisionTessellService(d)

		response, _, err := client.ProvisionTessellService(payload)
		if err != nil {
			return diag.FromErr(err)
		}
		id = *response.ResourceId
	}

	d.SetId(id)

	if d.Get("block_until_complete").(bool) {
		if err := client.DBServicePollForStatus(d, id, "READY", d.Get("timeout").(int), 60); err != nil {
			return diag.FromErr(err)
		}
	}

	resourceDBServiceRead(ctx, d, meta)

	return diags
}

func resourceDBServiceRead(_ context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*apiClient.Client)

	var diags diag.Diagnostics

	id := d.Get("id").(string)

	response, _, err := client.GetTessellService(id, d)
	if err != nil {
		return diag.FromErr(err)
	}

	if err := setResourceData(d, response); err != nil {
		return diag.FromErr(err)
	}

	d.SetId(*response.Id)

	return diags
}

func resourceDBServiceUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*apiClient.Client)

	var diags diag.Diagnostics

	expectedStatus := d.Get("expected_status").(string)
	status := d.GetRawState().GetAttr("status").AsString()
	id := d.Get("id").(string)

	if (d.HasChanges("expected_status") || status == "STOPPED") && expectedStatus == "READY" {
		payload := formPayloadForStartTessellService(d)

		_, _, err := client.StartTessellService(id, payload)
		if err != nil {
			d.Set("expected_status", "")
			return diag.FromErr(err)
		}

		if err := client.DBServicePollForStatus(d, d.Get("id").(string), "READY", d.Get("timeout").(int), 30); err != nil {
			return diag.FromErr(err)
		}
	}

	if d.HasChanges("name") ||
		d.HasChanges("description") ||
		d.HasChanges("enable_deletion_protection") ||
		d.HasChanges("enable_stop_protection") ||
		d.HasChanges("auto_minor_version_update") {
		payload := formPayloadForUpdateTessellService(d)

		_, _, err := client.UpdateTessellService(id, payload)
		if err != nil {
			return diag.FromErr(err)
		}
	}

	if d.HasChanges("creds.0.master_password") {
		payload := formPayloadForUpdateTessellServiceCredentials(d)

		taskSummary, _, err := client.UpdateTessellServiceCredentials(id, payload)
		if err != nil {
			return diag.FromErr(err)
		}

		if err := client.DBServicePollForUpdateInProgress(d.Get("id").(string), *taskSummary.TaskId, d.Get("timeout").(int), 30); err != nil {
			return diag.FromErr(err)
		}
	}

	// Updates in instance
	if d.HasChanges("instances") {
		// Add Instance
		tessellServiceResponse, _, err := client.GetTessellService(id, d)
		if err != nil {
			return diag.FromErr(err)
		}
		newTFInstances := getNewTFInstances(d, tessellServiceResponse.Instances)
		if newTFInstances != nil && len(*newTFInstances) != 0 {
			for _, newTfInstance := range *newTFInstances {
				instanceAddPayload := formPayloadForAddTessellServiceInstances(d, &newTfInstance)
				if instanceAddPayload != nil {
					_, _, err := client.AddTessellServiceInstances(id, instanceAddPayload)
					if err != nil {
						return diag.FromErr(err)
					}
					// poll for instance addition
					if err := client.DBServicePollForInstanceAddition(d.Get("id").(string), *newTfInstance.Name, d.Get("timeout").(int), 30); err != nil {
						return diag.FromErr(err)
					}
				}
			}
		}

		// Switchover Instance
		tessellServiceResponse, _, err = client.GetTessellService(id, d)
		if err != nil {
			return diag.FromErr(err)
		}
		switchoverPayload := formPayloadForSwitchoverTessellService(d, tessellServiceResponse.Instances)
		if switchoverPayload != nil {
			_, _, err := client.SwitchoverTessellService(id, switchoverPayload)
			if err != nil {
				return diag.FromErr(err)
			}
			// poll for instance switchover
			if err := client.DBServicePollForInstanceSwitchover(d.Get("id").(string), *switchoverPayload.SwitchToInstanceId, d.Get("timeout").(int), 30); err != nil {
				return diag.FromErr(err)
			}
		}

		// Remove Instance
		tessellServiceResponse, _, err = client.GetTessellService(id, d)
		if err != nil {
			return diag.FromErr(err)
		}
		deleteInstancePayload := formPayloadForDeleteTessellServiceInstances(d, tessellServiceResponse.Instances)
		if deleteInstancePayload != nil && len(*deleteInstancePayload.InstanceIds) != 0 {
			_, _, err := client.DeleteTessellServiceInstances(id, deleteInstancePayload)
			if err != nil {
				return diag.FromErr(err)
			}
			// poll for instance removal
			for _, instanceId := range *deleteInstancePayload.InstanceIds {
				if err := client.DBServicePollForInstanceDeletion(d.Get("id").(string), instanceId, d.Get("timeout").(int), 30); err != nil {
					return diag.FromErr(err)
				}
			}
		}
	}

	if (d.HasChanges("expected_status") || status == "READY") && expectedStatus == "STOPPED" {
		payload := formPayloadForStopTessellService(d)

		_, _, err := client.StopTessellService(id, payload)
		if err != nil {
			d.Set("expected_status", "")
			return diag.FromErr(err)
		}

		if err := client.DBServicePollForStatus(d, d.Get("id").(string), "STOPPED", d.Get("timeout").(int), 30); err != nil {
			return diag.FromErr(err)
		}
	}

	resourceDBServiceRead(ctx, d, meta)

	return diags
}

func resourceDBServiceDelete(_ context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*apiClient.Client)

	var diags diag.Diagnostics

	id := d.Get("id").(string)

	payload := formPayloadForDeleteTessellService(d)

	response, statusCode, err := client.DeleteTessellService(id, payload)
	if err != nil {
		return diag.FromErr(err)
	}

	if statusCode != 200 {
		return diag.FromErr(fmt.Errorf("deletion failed for tessell_db_service with id %s. Received response: %+v", id, response))
	}

	if d.Get("block_until_complete").(bool) {
		if err := client.DBServicePollForStatusCode(id, 404, d.Get("timeout").(int), 30); err != nil {
			return diag.FromErr(err)
		}
	}

	return diags
}
