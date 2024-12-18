#!/bin/bash

# Define old and new repository references
OLD_REPO="github.com/devalvamseezeeve/iopn-distrubution"
NEW_REPO="github.com/devalvamseezeeve/iopn-distrubution"

# Log replacement activity
echo "Updating references from '$OLD_REPO' to '$NEW_REPO'..."

# Replace in go.mod
if [ -f "go.mod" ]; then
  echo "Updating go.mod..."
  sed -i "s|$OLD_REPO|$NEW_REPO|g" go.mod
fi

# Replace in Makefile
if [ -f "Makefile" ]; then
  echo "Updating Makefile..."
  sed -i "s|$OLD_REPO|$NEW_REPO|g" Makefile
fi

# Replace in all Go source files
echo "Updating Go source files..."
find . -type f -name "*.go" -exec sed -i "s|$OLD_REPO|$NEW_REPO|g" {} +

# Replace in shell and script files
echo "Updating shell and script files..."
find . -type f \( -name "*.sh" -o -name "*.sh.template" \) -exec sed -i "s|$OLD_REPO|$NEW_REPO|g" {} +

# Replace in documentation and other text files
echo "Updating markdown and documentation files..."
find . -type f \( -name "*.md" -o -name "*.txt" -o -name "*.yaml" -o -name "*.yml" \) -exec sed -i "s|$OLD_REPO|$NEW_REPO|g" {} +

# Replace in protobuf files
echo "Updating protobuf files..."
find . -type f \( -name "*.proto" -o -name "*.pb.go" \) -exec sed -i "s|$OLD_REPO|$NEW_REPO|g" {} +

# Confirmation
echo "Replacement completed successfully!"

