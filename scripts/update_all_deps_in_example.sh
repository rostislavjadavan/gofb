#!/bin/bash

SCRIPT_DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"
EXAMPLES_DIR="$SCRIPT_DIR/../examples"

if [ ! -d "$EXAMPLES_DIR" ]; then
    echo "Examples directory not found: $EXAMPLES_DIR"
    exit 1
fi

for dir in "$EXAMPLES_DIR"/*; do
    if [ -d "$dir" ] && [ -f "$dir/go.mod" ]; then
        echo "Updating dependencies in $dir"
        cd "$dir" && go get -u && cd -
        cd ..
    else
        echo "Skipping $dir - no go.mod found"
    fi
done

echo "Dependency update complete!"
