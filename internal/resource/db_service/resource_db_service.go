package db_service

import (
	"context"
	"fmt"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
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
				ForceNew:    true,
			},
			"num_of_instances": {
				Type:        schema.TypeInt,
				Description: "Number of instance (nodes) to be created for the DB Service. This is a required input for Apache Kafka. For all other engines, this input would be ignored even if specified.",
				Computed:    true,
			},
			"status": {
				Type:        schema.TypeString,
				Description: "",
				Computed:    true,
			},
			"context_info": {
				Type:        schema.TypeList,
				Description: "",
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
						"tessell_service_id": {
							Type:        schema.TypeString,
							Description: "The DB Service Id using which this DB Service clone is created",
							Optional:    true,
							ForceNew:    true,
						},
						"availability_machine_id": {
							Type:        schema.TypeString,
							Description: "The Availability Machine Id using which this DB Service clone is created",
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
							Description: "The Availaility Machine name using which this DB Service clone is created",
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
							Description: "The snapshot Id using which this DB Service clone is created",
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
							Description: "If the service was created using a maximum recoverablity from the parent service",
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
				ForceNew:    true,
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
							ForceNew:    true,
						},
						"availability_zone": {
							Type:        schema.TypeString,
							Description: "The availability-zone in which the DB Service is provisioned",
							Optional:    true,
							ForceNew:    true,
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
							ForceNew:    true,
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
							ForceNew:    true,
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
							Description: "Size in GB. This is maintained for backward compatibility and would be deprecated soon.",
							Optional:    true,
							ForceNew:    true,
							Default:     0,
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
							Description: "",
							Optional:    true,
							ForceNew:    true,
							Default:     false,
						},
						"ca_cert_id": {
							Type:        schema.TypeString,
							Description: "",
							Computed:    true,
						},
						"dns_prefix": {
							Type:        schema.TypeString,
							Description: "",
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
										Description: "",
										Optional:    true,
										ForceNew:    true,
									},
									"endpoint": {
										Type:        schema.TypeString,
										Description: "",
										Optional:    true,
										ForceNew:    true,
									},
									"master_user": {
										Type:        schema.TypeString,
										Description: "",
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
										Description: "The configured endpoint as a result of configuring the service-pricipals",
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
						"update_in_progress_info": {
							Type:        schema.TypeList,
							Description: "DB Service connectivity update-in-progress details",
							Computed:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"dns_prefix": {
										Type:        schema.TypeString,
										Description: "",
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
							Description: "",
							Required:    true,
							ForceNew:    true,
						},
					},
				},
			},
			"snapshot_configuration": {
				Type:        schema.TypeList,
				Description: "DB Service's backup configurations. If not specified, the default recommended backup configurations would be applied.",
				Optional:    true,
				MaxItems:    1,
				MinItems:    1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"auto_snapshot": {
							Type:        schema.TypeBool,
							Description: "Specify whether to capture automated snapshots for the DB Service, default true.",
							Optional:    true,
							Default:     true,
						},
						"sla": {
							Type:        schema.TypeString,
							Description: "The snapshot SLA for the DB Service. If not specified, a default SLA would be associated with the DB Service",
							Optional:    true,
						},
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
										Description: "Time value in (hh:mm) format. ex. '02:00'",
										Optional:    true,
									},
									"duration": {
										Type:        schema.TypeInt,
										Description: "The allowed duration for capturing the DB Service backup",
										Optional:    true,
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
										Description: "The parameter profile id for the database",
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
										Description: "The parameter profile id for the database",
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
										Description: "The parameter profile id for the database",
										Optional:    true,
										ForceNew:    true,
									},
									"ad_domain_id": {
										Type:        schema.TypeString,
										Description: "Active Directory Domain id",
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
										Description: "The parameter profile id for the database",
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
										Description: "The Tessell Script Id",
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
										Description: "The Tessell Script Id",
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
									clonedFromDatabaseId := d.GetRawState().GetAttr("databases").AsValueSlice()[0].GetAttr("cloned_from_info").AsValueSlice()[0].GetAttr("database_id").AsString()
									if sourceDatabaseId == clonedFromDatabaseId {
										return true
									}
								}
								return false
							},
						},
						"id": {
							Type:        schema.TypeString,
							Description: "",
							Computed:    true,
						},
						"database_name": {
							Type:        schema.TypeString,
							Description: "Database name",
							Optional:    true,
							ForceNew:    true,
						},
						"description": {
							Type:        schema.TypeString,
							Description: "Database description",
							Optional:    true,
						},
						"tessell_service_id": {
							Type:        schema.TypeString,
							Description: "Associated DB Service Id",
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
													Description: "The parameter profile id for the database",
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
													Description: "The parameter profile id for the database",
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
												"parameter_profile_id": {
													Type:        schema.TypeString,
													Description: "The parameter profile id for the database",
													Optional:    true,
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
				Computed:    true,
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
							Computed:    true,
						},
						"type": {
							Type:        schema.TypeString,
							Description: "",
							Optional:    true,
							ForceNew:    true,
						},
						"role": {
							Type:        schema.TypeString,
							Description: "DB Service Topology",
							Computed:    true,
						},
						"status": {
							Type:        schema.TypeString,
							Description: "",
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
							Computed:    true,
						},
						"availability_zone": {
							Type:        schema.TypeString,
							Description: "DB Service Instance's cloud availability zone",
							Computed:    true,
						},
						"instance_group_id": {
							Type:        schema.TypeString,
							Description: "The instance groupd Id",
							Optional:    true,
							ForceNew:    true,
						},
						"compute_type": {
							Type:        schema.TypeString,
							Description: "The compute used for creation of the Tessell Service Instance",
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
							Description: "The storage (in bytes) that has been provisioned for the DB Service instance.",
							Optional:    true,
							ForceNew:    true,
						},
						"data_volume_iops": {
							Type:        schema.TypeInt,
							Description: "",
							Optional:    true,
							ForceNew:    true,
						},
						"parameter_profile": {
							Type:        schema.TypeList,
							Description: "",
							Optional:    true,
							ForceNew:    true,
							MaxItems:    1,
							MinItems:    1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"id": {
										Type:        schema.TypeString,
										Description: "",
										Optional:    true,
										ForceNew:    true,
									},
									"name": {
										Type:        schema.TypeString,
										Description: "",
										Optional:    true,
										ForceNew:    true,
									},
									"version": {
										Type:        schema.TypeString,
										Description: "",
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
						"vpc": {
							Type:        schema.TypeString,
							Description: "The VPC used for creation of the DB Service Instance",
							Computed:    true,
						},
						"encryption_key": {
							Type:        schema.TypeString,
							Description: "The encryption key name which is used to encrypt the data at rest",
							Optional:    true,
							ForceNew:    true,
						},
						"software_image": {
							Type:        schema.TypeString,
							Description: "Software Image to be used to create the instance",
							Optional:    true,
							ForceNew:    true,
						},
						"software_image_version": {
							Type:        schema.TypeString,
							Description: "Software Image Version to be used to create the instance",
							Optional:    true,
							ForceNew:    true,
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
							Optional:    true,
							ForceNew:    true,
						},
						"last_stopped_at": {
							Type:        schema.TypeString,
							Description: "Timestamp when the Service Instance was last stopped at",
							Optional:    true,
							ForceNew:    true,
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
										Description: "",
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
										Description: "",
										Computed:    true,
									},
									"message": {
										Type:        schema.TypeString,
										Description: "",
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
										Description: "",
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
				Description: "If the service is to be deleted, this config would be honoured if no preference is provided during deleting the service",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"retain_availability_machine": {
							Type:        schema.TypeBool,
							Description: "If specified as true, the associated Availability Machine (snapshots, sanitized-snapshots, logs) would be retained",
							Optional:    true,
							ForceNew:    true,
							Default:     false,
						},
					},
				},
			},
			"deletion_schedule": {
				Type:        schema.TypeList,
				Description: "",
				Optional:    true,
				ForceNew:    true,
				MaxItems:    1,
				MinItems:    1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"delete_at": {
							Type:        schema.TypeString,
							Description: "DB Service deletion Time",
							Required:    true,
							ForceNew:    true,
						},
						"deletion_config": {
							Type:        schema.TypeList,
							Description: "If the service is to be deleted, this config would be honoured if no preference is provided during deleting the service",
							Computed:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"retain_availability_machine": {
										Type:        schema.TypeBool,
										Description: "If specified as true, the associated Availability Machine (snapshots, sanitized-snapshots, logs) would be retained",
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
				Description: "If block_until_complete is true, how long it should block for. (In seconds)",
				Optional:    true,
				Default:     3600,
			},
			"expected_status": {
				Type:        schema.TypeString,
				Description: "If provided, invoke the DB Service start/stop API",
				Optional:    true,
				Default:     "READY",
			},
		},
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
		if err := client.DBServicePollForStatus(id, "READY", d.Get("timeout").(int), 60); err != nil {
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

	response, _, err := client.GetTessellService(id)
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
	var pollBreakValue string
	var pollFunc func(string, string, int, int) error
	shouldPoll := false

	expectedStatus := d.Get("expected_status").(string)
	id := d.Get("id").(string)

	if d.HasChanges("expected_status") && expectedStatus == "READY" {
		payload := formPayloadForStartTessellService(d)

		_, _, err := client.StartTessellService(id, payload)
		if err != nil {
			return diag.FromErr(err)
		}
		shouldPoll = true
		pollBreakValue = "READY"
		pollFunc = client.DBServicePollForStatus
	} else if d.HasChanges("expected_status") && expectedStatus == "STOPPED" {
		payload := formPayloadForStopTessellService(d)

		_, _, err := client.StopTessellService(id, payload)
		if err != nil {
			return diag.FromErr(err)
		}
		shouldPoll = true
		pollBreakValue = "STOPPED"
		pollFunc = client.DBServicePollForStatus
	}

	if shouldPoll {
		if err := pollFunc(d.Get("id").(string), pollBreakValue, d.Get("timeout").(int), 30); err != nil {
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
