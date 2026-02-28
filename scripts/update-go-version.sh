#!/bin/bash
set -e

# Read Go version from .go-version file
GO_VERSION=$(cat .go-version)

echo "🔄 Updating Go version to $GO_VERSION across all files..."

# Detect OS for sed compatibility
if [[ "$OSTYPE" == "darwin"* ]]; then
    # macOS
    SED_INPLACE="sed -i ''"
else
    # Linux/Windows Git Bash
    SED_INPLACE="sed -i"
fi

# Update go.mod
echo "📝 Updating go.mod..."
$SED_INPLACE "s/^go [0-9.]\+$/go $GO_VERSION/" go.mod

# Update all Dockerfiles
echo "🐳 Updating Dockerfiles..."
find cmd/server -name "Dockerfile" -type f -exec $SED_INPLACE "s/golang:[0-9.]\+-alpine/golang:$GO_VERSION-alpine/" {} \;

# Update docker-compose.yml if it has Go version
if grep -q "golang:" docker-compose.yml 2>/dev/null; then
    echo "🐳 Updating docker-compose.yml..."
    $SED_INPLACE "s/golang:[0-9.]\+-alpine/golang:$GO_VERSION-alpine/" docker-compose.yml
fi

# Update GitHub Actions workflows
echo "⚙️  Updating GitHub Actions workflows..."
find .github/workflows -name "*.yml" -type f -exec $SED_INPLACE "s/go-version: '[0-9.]\+'/go-version: '$GO_VERSION'/" {} \;

echo "✅ Go version updated to $GO_VERSION in all files!"
echo ""
echo "📋 Updated files:"
echo "  - go.mod"
echo "  - All Dockerfiles (10 files)"
echo "  - All GitHub Actions workflows"
echo ""
echo "🔍 Verify changes:"
echo "  git diff"
echo ""
echo "💾 Commit changes:"
echo "  git add ."
echo "  git commit -m 'chore: update Go version to $GO_VERSION'"
echo "  git push"
echo ""
echo "💡 Or just push .go-version and let GitHub Actions auto-sync:"
echo "  echo '1.25' > .go-version"
echo "  git add .go-version"
echo "  git commit -m 'chore: bump Go to 1.25'"
echo "  git push"
echo "  # GitHub Actions will automatically update all other files!"
