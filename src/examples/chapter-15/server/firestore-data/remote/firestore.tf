# Firestore document configuration for Recipe Quest App
# Simple schema for educational Tool demonstration

terraform {
  required_version = ">= 1.0"
  required_providers {
    google = {
      source  = "hashicorp/google"
      version = ">= 5.0"
    }
  }
}

provider "google" {
  project = var.project_id
  region  = var.region
}

# Local variable to define ingredient compatibility combinations
locals {
  ingredient_combinations = [
    # Perfect pairings (score 9-10)
    {
      document_id = "combo_001"
      fields = jsonencode({
        "ingredients" = {
          "arrayValue" = {
            "values" = [
              { "stringValue" = "chicken" },
              { "stringValue" = "rice" }
            ]
          }
        }
        "compatibility_score" = { "integerValue" = "9" }
        "flavor_profile"      = { "stringValue" = "savory" }
        "cuisine_style"       = { "stringValue" = "asian" }
        "tips"               = { "stringValue" = "Perfect for comfort food and one-pot dishes" }
        "difficulty_bonus"    = { "integerValue" = "1" }
      })
    },

    {
      document_id = "combo_002"  
      fields = jsonencode({
        "ingredients" = {
          "arrayValue" = {
            "values" = [
              { "stringValue" = "salmon" },
              { "stringValue" = "lemon" }
            ]
          }
        }
        "compatibility_score" = { "integerValue" = "10" }
        "flavor_profile"      = { "stringValue" = "fresh" }
        "cuisine_style"       = { "stringValue" = "mediterranean" }
        "tips"               = { "stringValue" = "Classic pairing with bright citrus notes" }
        "difficulty_bonus"    = { "integerValue" = "0" }
      })
    },

    {
      document_id = "combo_003"
      fields = jsonencode({
        "ingredients" = {
          "arrayValue" = {
            "values" = [
              { "stringValue" = "pasta" },
              { "stringValue" = "garlic" }
            ]
          }
        }
        "compatibility_score" = { "integerValue" = "8" }
        "flavor_profile"      = { "stringValue" = "aromatic" }
        "cuisine_style"       = { "stringValue" = "italian" }
        "tips"               = { "stringValue" = "Foundation of many Italian classics" }
        "difficulty_bonus"    = { "integerValue" = "0" }
      })
    },

    # Good combinations (score 7-8)
    {
      document_id = "combo_004"
      fields = jsonencode({
        "ingredients" = {
          "arrayValue" = {
            "values" = [
              { "stringValue" = "beef" },
              { "stringValue" = "potatoes" }
            ]
          }
        }
        "compatibility_score" = { "integerValue" = "8" }
        "flavor_profile"      = { "stringValue" = "hearty" }
        "cuisine_style"       = { "stringValue" = "western" }
        "tips"               = { "stringValue" = "Great for stews and roasts" }
        "difficulty_bonus"    = { "integerValue" = "1" }
      })
    },

    {
      document_id = "combo_005"
      fields = jsonencode({
        "ingredients" = {
          "arrayValue" = {
            "values" = [
              { "stringValue" = "shrimp" },
              { "stringValue" = "avocado" }
            ]
          }
        }
        "compatibility_score" = { "integerValue" = "7" }
        "flavor_profile"      = { "stringValue" = "light" }
        "cuisine_style"       = { "stringValue" = "modern" }
        "tips"               = { "stringValue" = "Perfect for fresh, healthy dishes" }
        "difficulty_bonus"    = { "integerValue" = "0" }
      })
    },

    # Creative combinations (score 6-7)
    {
      document_id = "combo_006"
      fields = jsonencode({
        "ingredients" = {
          "arrayValue" = {
            "values" = [
              { "stringValue" = "tofu" },
              { "stringValue" = "ginger" }
            ]
          }
        }
        "compatibility_score" = { "integerValue" = "7" }
        "flavor_profile"      = { "stringValue" = "umami" }
        "cuisine_style"       = { "stringValue" = "asian" }
        "tips"               = { "stringValue" = "Excellent for stir-fries and broths" }
        "difficulty_bonus"    = { "integerValue" = "1" }
      })
    },

    {
      document_id = "combo_007"
      fields = jsonencode({
        "ingredients" = {
          "arrayValue" = {
            "values" = [
              { "stringValue" = "quinoa" },
              { "stringValue" = "lime" }
            ]
          }
        }
        "compatibility_score" = { "integerValue" = "6" }
        "flavor_profile"      = { "stringValue" = "zesty" }
        "cuisine_style"       = { "stringValue" = "south_american" }
        "tips"               = { "stringValue" = "Modern healthy pairing with bright flavors" }
        "difficulty_bonus"    = { "integerValue" = "0" }
      })
    },

    # Adventurous combinations (score 5-6)
    {
      document_id = "combo_008"
      fields = jsonencode({
        "ingredients" = {
          "arrayValue" = {
            "values" = [
              { "stringValue" = "pork" },
              { "stringValue" = "miso" }
            ]
          }
        }
        "compatibility_score" = { "integerValue" = "8" }
        "flavor_profile"      = { "stringValue" = "rich" }
        "cuisine_style"       = { "stringValue" = "japanese" }
        "tips"               = { "stringValue" = "Deep umami flavors for sophisticated dishes" }
        "difficulty_bonus"    = { "integerValue" = "2" }
      })
    }
  ]
}

# Create ingredient combination documents in Firestore
resource "google_firestore_document" "ingredient_combinations" {
  for_each    = { for combo in local.ingredient_combinations : combo.document_id => combo }
  project     = var.project_id
  collection  = "ingredient_combinations"
  document_id = each.value.document_id
  fields      = each.value.fields
}

# Output the created document IDs
output "combination_ids" {
  value = [for doc in google_firestore_document.ingredient_combinations : doc.document_id]
  description = "List of created ingredient combination document IDs"
}