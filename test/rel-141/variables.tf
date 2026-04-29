variable "api_address" {
  description = "Tessell API endpoint"
  type        = string
  default     = "https://api.terraformrelease140001.tsl-terls.cloud"
}

variable "tenant_id" {
  description = "Tessell tenant ID"
  type        = string
  default     = "fbab7165-97a3-42c7-8be5-abcf31cb62f7"
}

variable "api_key" {
  description = "Tessell API key — pass via TF_VAR_api_key or -var flag, never hardcode"
  type        = string
  sensitive   = true
}
