# Firestore document configuration for Chapter 15 Cooking Battle App
# This creates sample recipe data in Firestore

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

# Local variable to define recipe documents
locals {
  recipes = [
    # Italian recipes
    {
      document_id = "recipe_001"
      fields = jsonencode({
        "name"       = { "stringValue" = "Classic Tomato Pasta" }
        "category"   = { "stringValue" = "Italian" }
        "difficulty" = { "stringValue" = "easy" }
        "prepTime"   = { "integerValue" = "15" }
        "cookTime"   = { "integerValue" = "20" }
        "ingredients" = {
          "arrayValue" = {
            "values" = [
              {
                "mapValue" = {
                  "fields" = {
                    "name"   = { "stringValue" = "pasta" }
                    "amount" = { "integerValue" = "200" }
                    "unit"   = { "stringValue" = "g" }
                  }
                }
              },
              {
                "mapValue" = {
                  "fields" = {
                    "name"   = { "stringValue" = "tomato" }
                    "amount" = { "integerValue" = "3" }
                    "unit"   = { "stringValue" = "pieces" }
                  }
                }
              },
              {
                "mapValue" = {
                  "fields" = {
                    "name"   = { "stringValue" = "garlic" }
                    "amount" = { "integerValue" = "2" }
                    "unit"   = { "stringValue" = "cloves" }
                  }
                }
              },
              {
                "mapValue" = {
                  "fields" = {
                    "name"   = { "stringValue" = "olive_oil" }
                    "amount" = { "integerValue" = "30" }
                    "unit"   = { "stringValue" = "ml" }
                  }
                }
              }
            ]
          }
        }
        "nutrition" = {
          "mapValue" = {
            "fields" = {
              "calories" = { "integerValue" = "450" }
              "protein"  = { "integerValue" = "15" }
              "carbs"    = { "integerValue" = "65" }
              "fat"      = { "integerValue" = "12" }
            }
          }
        }
        "tags" = {
          "arrayValue" = {
            "values" = [
              { "stringValue" = "vegetarian" },
              { "stringValue" = "quick" },
              { "stringValue" = "budget-friendly" }
            ]
          }
        }
      })
    },

    # Japanese recipes
    {
      document_id = "recipe_002"
      fields = jsonencode({
        "name"       = { "stringValue" = "Grilled Chicken Teriyaki" }
        "category"   = { "stringValue" = "Japanese" }
        "difficulty" = { "stringValue" = "medium" }
        "prepTime"   = { "integerValue" = "30" }
        "cookTime"   = { "integerValue" = "25" }
        "ingredients" = {
          "arrayValue" = {
            "values" = [
              {
                "mapValue" = {
                  "fields" = {
                    "name"   = { "stringValue" = "chicken" }
                    "amount" = { "integerValue" = "500" }
                    "unit"   = { "stringValue" = "g" }
                  }
                }
              },
              {
                "mapValue" = {
                  "fields" = {
                    "name"   = { "stringValue" = "soy_sauce" }
                    "amount" = { "integerValue" = "50" }
                    "unit"   = { "stringValue" = "ml" }
                  }
                }
              },
              {
                "mapValue" = {
                  "fields" = {
                    "name"   = { "stringValue" = "mirin" }
                    "amount" = { "integerValue" = "30" }
                    "unit"   = { "stringValue" = "ml" }
                  }
                }
              },
              {
                "mapValue" = {
                  "fields" = {
                    "name"   = { "stringValue" = "sugar" }
                    "amount" = { "integerValue" = "15" }
                    "unit"   = { "stringValue" = "g" }
                  }
                }
              }
            ]
          }
        }
        "nutrition" = {
          "mapValue" = {
            "fields" = {
              "calories" = { "integerValue" = "380" }
              "protein"  = { "integerValue" = "42" }
              "carbs"    = { "integerValue" = "18" }
              "fat"      = { "integerValue" = "15" }
            }
          }
        }
        "tags" = {
          "arrayValue" = {
            "values" = [
              { "stringValue" = "protein-rich" },
              { "stringValue" = "japanese" },
              { "stringValue" = "grilled" }
            ]
          }
        }
      })
    },

    # Quick & Easy recipes
    {
      document_id = "recipe_003"
      fields = jsonencode({
        "name"       = { "stringValue" = "Egg Fried Rice" }
        "category"   = { "stringValue" = "Asian" }
        "difficulty" = { "stringValue" = "easy" }
        "prepTime"   = { "integerValue" = "10" }
        "cookTime"   = { "integerValue" = "15" }
        "ingredients" = {
          "arrayValue" = {
            "values" = [
              {
                "mapValue" = {
                  "fields" = {
                    "name"   = { "stringValue" = "rice" }
                    "amount" = { "integerValue" = "300" }
                    "unit"   = { "stringValue" = "g" }
                  }
                }
              },
              {
                "mapValue" = {
                  "fields" = {
                    "name"   = { "stringValue" = "egg" }
                    "amount" = { "integerValue" = "3" }
                    "unit"   = { "stringValue" = "pieces" }
                  }
                }
              },
              {
                "mapValue" = {
                  "fields" = {
                    "name"   = { "stringValue" = "green_onion" }
                    "amount" = { "integerValue" = "2" }
                    "unit"   = { "stringValue" = "stalks" }
                  }
                }
              },
              {
                "mapValue" = {
                  "fields" = {
                    "name"   = { "stringValue" = "soy_sauce" }
                    "amount" = { "integerValue" = "20" }
                    "unit"   = { "stringValue" = "ml" }
                  }
                }
              }
            ]
          }
        }
        "nutrition" = {
          "mapValue" = {
            "fields" = {
              "calories" = { "integerValue" = "420" }
              "protein"  = { "integerValue" = "18" }
              "carbs"    = { "integerValue" = "52" }
              "fat"      = { "integerValue" = "14" }
            }
          }
        }
        "tags" = {
          "arrayValue" = {
            "values" = [
              { "stringValue" = "quick" },
              { "stringValue" = "leftover-friendly" },
              { "stringValue" = "one-pan" }
            ]
          }
        }
      })
    },

    # Healthy options
    {
      document_id = "recipe_004"
      fields = jsonencode({
        "name"       = { "stringValue" = "Greek Salad Bowl" }
        "category"   = { "stringValue" = "Mediterranean" }
        "difficulty" = { "stringValue" = "easy" }
        "prepTime"   = { "integerValue" = "15" }
        "cookTime"   = { "integerValue" = "0" }
        "ingredients" = {
          "arrayValue" = {
            "values" = [
              {
                "mapValue" = {
                  "fields" = {
                    "name"   = { "stringValue" = "cucumber" }
                    "amount" = { "integerValue" = "1" }
                    "unit"   = { "stringValue" = "piece" }
                  }
                }
              },
              {
                "mapValue" = {
                  "fields" = {
                    "name"   = { "stringValue" = "tomato" }
                    "amount" = { "integerValue" = "2" }
                    "unit"   = { "stringValue" = "pieces" }
                  }
                }
              },
              {
                "mapValue" = {
                  "fields" = {
                    "name"   = { "stringValue" = "feta_cheese" }
                    "amount" = { "integerValue" = "100" }
                    "unit"   = { "stringValue" = "g" }
                  }
                }
              },
              {
                "mapValue" = {
                  "fields" = {
                    "name"   = { "stringValue" = "olives" }
                    "amount" = { "integerValue" = "50" }
                    "unit"   = { "stringValue" = "g" }
                  }
                }
              }
            ]
          }
        }
        "nutrition" = {
          "mapValue" = {
            "fields" = {
              "calories" = { "integerValue" = "250" }
              "protein"  = { "integerValue" = "12" }
              "carbs"    = { "integerValue" = "15" }
              "fat"      = { "integerValue" = "18" }
            }
          }
        }
        "tags" = {
          "arrayValue" = {
            "values" = [
              { "stringValue" = "healthy" },
              { "stringValue" = "vegetarian" },
              { "stringValue" = "no-cook" }
            ]
          }
        }
      })
    }
  ]

  # Ingredient stock data (for checkIngredientStock tool)
  ingredient_stock = [
    {
      document_id = "stock_001"
      fields = jsonencode({
        "ingredient" = { "stringValue" = "pasta" }
        "quantity"   = { "integerValue" = "1000" }
        "unit"       = { "stringValue" = "g" }
        "available"  = { "booleanValue" = true }
      })
    },
    {
      document_id = "stock_002"
      fields = jsonencode({
        "ingredient" = { "stringValue" = "tomato" }
        "quantity"   = { "integerValue" = "10" }
        "unit"       = { "stringValue" = "pieces" }
        "available"  = { "booleanValue" = true }
      })
    },
    {
      document_id = "stock_003"
      fields = jsonencode({
        "ingredient" = { "stringValue" = "chicken" }
        "quantity"   = { "integerValue" = "2000" }
        "unit"       = { "stringValue" = "g" }
        "available"  = { "booleanValue" = true }
      })
    }
  ]
}

# Create recipe documents in Firestore
resource "google_firestore_document" "recipes" {
  for_each    = { for recipe in local.recipes : recipe.document_id => recipe }
  project     = var.project_id
  collection  = "recipes"
  document_id = each.value.document_id
  fields      = each.value.fields
}

# Create ingredient stock documents in Firestore
resource "google_firestore_document" "ingredient_stock" {
  for_each    = { for stock in local.ingredient_stock : stock.document_id => stock }
  project     = var.project_id
  collection  = "ingredient_stock"
  document_id = each.value.document_id
  fields      = each.value.fields
}

# Output the created document IDs
output "recipe_ids" {
  value = [for doc in google_firestore_document.recipes : doc.document_id]
  description = "List of created recipe document IDs"
}

output "stock_ids" {
  value = [for doc in google_firestore_document.ingredient_stock : doc.document_id]
  description = "List of created ingredient stock document IDs"
}