#!/bin/bash

# Script to update Genkit version in all src/examples Go modules
# Usage: ./scripts/update-genkit-version.sh <new_version>
# Example: ./scripts/update-genkit-version.sh v0.7.0

set -e

# Colors for output
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
BLUE='\033[0;34m'
NC='\033[0m' # No Color

# Function to print colored output
print_info() {
    echo -e "${BLUE}[INFO]${NC} $1"
}

print_success() {
    echo -e "${GREEN}[SUCCESS]${NC} $1"
}

print_warning() {
    echo -e "${YELLOW}[WARNING]${NC} $1"
}

print_error() {
    echo -e "${RED}[ERROR]${NC} $1"
}

# Check if version parameter is provided
if [ $# -eq 0 ]; then
    print_error "Usage: $0 <new_version>"
    print_error "Example: $0 v0.7.0"
    exit 1
fi

NEW_VERSION="$1"

# Validate version format (should start with v and contain dots)
if [[ ! "$NEW_VERSION" =~ ^v[0-9]+\.[0-9]+\.[0-9]+.*$ ]]; then
    print_warning "Version '$NEW_VERSION' doesn't follow semantic versioning format (vX.Y.Z)"
    read -p "Continue anyway? (y/N): " -n 1 -r
    echo
    if [[ ! $REPLY =~ ^[Yy]$ ]]; then
        exit 1
    fi
fi

print_info "Updating Genkit version to $NEW_VERSION in all src/examples Go modules..."

# Get the repository root directory
REPO_ROOT="$(cd "$(dirname "${BASH_SOURCE[0]}")/.." && pwd)"
EXAMPLES_DIR="$REPO_ROOT/src/examples"

# Check if examples directory exists
if [ ! -d "$EXAMPLES_DIR" ]; then
    print_error "Examples directory not found: $EXAMPLES_DIR"
    exit 1
fi

# Find all go.mod files in src/examples
GO_MOD_FILES=($(find "$EXAMPLES_DIR" -name "go.mod" -type f))

if [ ${#GO_MOD_FILES[@]} -eq 0 ]; then
    print_error "No go.mod files found in $EXAMPLES_DIR"
    exit 1
fi

print_info "Found ${#GO_MOD_FILES[@]} go.mod files to update"

# Track success and failure counts
SUCCESS_COUNT=0
FAILURE_COUNT=0
FAILED_MODULES=()

# Process each go.mod file
for GO_MOD_FILE in "${GO_MOD_FILES[@]}"; do
    MODULE_DIR="$(dirname "$GO_MOD_FILE")"
    # Create relative path manually (compatible with macOS)
    RELATIVE_PATH="${MODULE_DIR#$REPO_ROOT/}"
    
    print_info "Processing: $RELATIVE_PATH"
    
    # Check if the go.mod file contains genkit dependency
    if ! grep -q "github.com/firebase/genkit/go" "$GO_MOD_FILE"; then
        print_warning "  No Genkit dependency found in $RELATIVE_PATH/go.mod, skipping..."
        continue
    fi
    
    # Change to module directory
    cd "$MODULE_DIR"
    
    # Update the Genkit version
    print_info "  Updating Genkit version..."
    if go get "github.com/firebase/genkit/go@$NEW_VERSION"; then
        print_success "  ✓ Updated Genkit dependency"
    else
        print_error "  ✗ Failed to update Genkit dependency"
        FAILURE_COUNT=$((FAILURE_COUNT + 1))
        FAILED_MODULES+=("$RELATIVE_PATH (dependency update)")
        continue
    fi
    
    # Run go mod tidy
    print_info "  Running go mod tidy..."
    if go mod tidy; then
        print_success "  ✓ go mod tidy completed"
    else
        print_error "  ✗ go mod tidy failed"
        FAILURE_COUNT=$((FAILURE_COUNT + 1))
        FAILED_MODULES+=("$RELATIVE_PATH (go mod tidy)")
        continue
    fi
    
    # Check if there's a main.go file or if we can build the module
    BUILD_SUCCESS=false
    
    if [ -f "main.go" ]; then
        print_info "  Building main.go..."
        if go build -o /tmp/genkit-test-build main.go; then
            print_success "  ✓ Build successful (main.go)"
            BUILD_SUCCESS=true
            rm -f /tmp/genkit-test-build
        else
            print_error "  ✗ Build failed (main.go)"
        fi
    elif [ -f "*.go" ] || ls *.go >/dev/null 2>&1; then
        print_info "  Building module..."
        if go build -o /tmp/genkit-test-build .; then
            print_success "  ✓ Build successful"
            BUILD_SUCCESS=true
            rm -f /tmp/genkit-test-build
        else
            print_error "  ✗ Build failed"
        fi
    else
        print_info "  Running go mod verify (no buildable Go files found)..."
        if go mod verify; then
            print_success "  ✓ Module verification successful"
            BUILD_SUCCESS=true
        else
            print_error "  ✗ Module verification failed"
        fi
    fi
    
    if [ "$BUILD_SUCCESS" = true ]; then
        SUCCESS_COUNT=$((SUCCESS_COUNT + 1))
        print_success "  ✓ Module $RELATIVE_PATH updated successfully"
    else
        FAILURE_COUNT=$((FAILURE_COUNT + 1))
        FAILED_MODULES+=("$RELATIVE_PATH (build/verify)")
    fi
    
    echo # Empty line for readability
done

# Return to repository root
cd "$REPO_ROOT"

# Print summary
echo "=================================================="
print_info "UPDATE SUMMARY"
echo "=================================================="
print_info "Target Genkit version: $NEW_VERSION"
print_info "Total modules processed: $((SUCCESS_COUNT + FAILURE_COUNT))"
print_success "Successful updates: $SUCCESS_COUNT"

if [ $FAILURE_COUNT -gt 0 ]; then
    print_error "Failed updates: $FAILURE_COUNT"
    echo
    print_error "Failed modules:"
    for failed_module in "${FAILED_MODULES[@]}"; do
        print_error "  - $failed_module"
    done
    echo
    print_warning "Please review the failed modules manually."
    exit 1
else
    print_success "All modules updated successfully!"
    echo
    print_info "Next steps:"
    print_info "1. Test your applications to ensure they work with the new Genkit version"
    print_info "2. Update any version references in documentation"
    print_info "3. Commit the changes to your repository"
fi
