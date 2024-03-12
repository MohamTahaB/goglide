
VERSION="v1.0.0"

BINARY_NAME="goglide"

# Build the project
echo "Building $BINARY_NAME..."
go build -o "$BINARY_NAME"

# Check if the build was successful
if [ $? -eq 0 ]; then
    echo "Build successful."

    # Tag the release
    echo "Creating Git tag $VERSION..."
    git tag -a "$VERSION" -m "Release $VERSION"
    git push origin "$VERSION"

    # Check if RELEASE.md exists
    if [ -f RELEASE.md ]; then
        RELEASE_NOTES="RELEASE.md"
        echo "RELEASE.md found."
    else
        echo "Release aborted: no RELEASE.md found in the root dir."
        exit 1
    fi

    # Create the release on GitHub
    echo "Creating GitHub release..."
    gh release create "$VERSION" -t "$VERSION" -n "$(cat "$RELEASE_NOTES")" "$BINARY_NAME"
else
    echo "Build failed."
    exit 1
fi
