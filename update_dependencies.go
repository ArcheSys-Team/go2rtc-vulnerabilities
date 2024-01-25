#!/bin/bash

# Go module path (replace with your module path)
MODULE_PATH="github.com/ArcheSys-Team/go2rtc-vulnerabilities"

# Get the list of outdated dependencies
OUTDATED_DEPS=$(go list -u -m -json all | jq -c 'select(.Update != null)')

# Check if there are outdated dependencies
if [ -n "$OUTDATED_DEPS" ]; then
    echo "Outdated Dependencies found:"
    echo "$OUTDATED_DEPS"

    # Loop through each outdated dependency and update it
    echo "Updating dependencies..."
    while read -r line; do
        package=$(echo "$line" | jq -r '.Path')
        current_version=$(echo "$line" | jq -r '.Version')
        latest_version=$(go list -u -m "$package" | grep " => " | awk '{ print $3 }')

        if [ "$latest_version" != "$current_version" ]; then
            go get -u "$package"
            echo "Updated $package from $current_version to $latest_version"
        else
            echo "No update needed for $package"
        fi
    done <<< "$OUTDATED_DEPS"

    # Print a message to remind users to commit changes
    echo "Please review and commit the changes to go.mod and go.sum files."
else
    echo "No outdated dependencies found."
fi