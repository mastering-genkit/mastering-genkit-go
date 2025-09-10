# Terraform variables for Chapter 15 Cooking Battle App

variable "project_id" {
  description = "Google Cloud Project ID with Firebase enabled"
  type        = string
}

variable "region" {
  description = "Default region for resources"
  type        = string
  default     = "asia-northeast1"
}

variable "environment" {
  description = "Environment name (dev, staging, prod)"
  type        = string
  default     = "dev"
}