# Go Version Management

## Overview

This project uses a centralized Go version management system to keep all files in sync.

## Single Source of Truth

**File**: `.go-version`
```
1.24
```

This file controls the Go version used in:
- ✅ `go.mod`
- ✅ All Dockerfiles (10 files)
- ✅ All GitHub Actions workflows
- ✅ Docker Compose

---

## How to Update Go Version

### Method 1: Automatic (Recommended) ✅

```bash
# 1. Update the version file
echo "1.25" > .go-version

# 2. Run the update script
./scripts/update-go-version.sh

# 3. Commit changes
git add .
git commit -m "chore: update Go version to 1.25"
git push
```

### Method 2: Manual

```bash
# Edit .go-version
echo "1.25" > .go-version

# Manually update:
# - go.mod (line: go 1.25)
# - All Dockerfiles (FROM golang:1.25-alpine)
# - All workflows (go-version: '1.25')
```

---

## Automatic Verification

### GitHub Action: `check-go-version.yml`

Runs on every PR and push to verify all files use the same Go version.

**What it checks:**
- ✅ `go.mod` matches `.go-version`
- ✅ All Dockerfiles use correct version
- ✅ All workflows use correct version

**If mismatch found:**
- ❌ Build fails
- 💡 Shows how to fix

---

## Files Managed

### 1. `.go-version` (Source of Truth)
```
1.24
```

### 2. `go.mod`
```go
go 1.24
```

### 3. Dockerfiles (10 files)
```dockerfile
FROM golang:1.24-alpine AS builder
```

**Locations:**
- `cmd/server/gateway/Dockerfile`
- `cmd/server/risk-agent/Dockerfile`
- `cmd/server/trading-agent/Dockerfile`
- `cmd/server/quantum-agent/Dockerfile`
- `cmd/server/compliance-agent/Dockerfile`
- `cmd/server/consensus-agent/Dockerfile`
- `cmd/server/blockchain-agent/Dockerfile`
- `cmd/server/digital-twin-agent/Dockerfile`
- `cmd/server/optimization-agent/Dockerfile`
- `cmd/server/regulation-agent/Dockerfile`

### 4. GitHub Actions Workflows
```yaml
- name: Setup Go
  uses: actions/setup-go@v4
  with:
    go-version: '1.24'
```

**Locations:**
- `.github/workflows/dev-deploy.yml`
- `.github/workflows/main-deploy.yml`
- `.github/workflows/pr-checks.yml`
- `.github/workflows/security-scan.yml`

---

## Benefits

### ✅ Consistency
- All files always use the same Go version
- No version mismatches
- No build failures due to version conflicts

### ✅ Easy Updates
- Change one file (`.go-version`)
- Run one script
- Everything updates automatically

### ✅ Automatic Verification
- CI/CD checks for mismatches
- Fails fast if versions don't match
- Prevents deployment of inconsistent code

### ✅ Documentation
- `.go-version` file clearly shows current version
- Easy for new developers to understand
- Self-documenting system

---

## Troubleshooting

### Build Fails: "go.mod requires go >= X.XX"

**Cause**: Dockerfile uses older Go version than go.mod

**Fix**:
```bash
./scripts/update-go-version.sh
```

### CI Fails: "Go versions are inconsistent"

**Cause**: Files have different Go versions

**Fix**:
```bash
# Check current version
cat .go-version

# Update all files
./scripts/update-go-version.sh

# Commit and push
git add .
git commit -m "fix: sync Go versions"
git push
```

### Script Doesn't Work on Windows

**Fix**: Use Git Bash or WSL
```bash
# In Git Bash or WSL
bash ./scripts/update-go-version.sh
```

---

## When to Update Go Version

### Minor Updates (1.24.0 → 1.24.1)
- Security patches
- Bug fixes
- Update immediately

### Major Updates (1.24 → 1.25)
- New features
- Breaking changes possible
- Test thoroughly before updating

### Process:
1. Check Go release notes
2. Update `.go-version`
3. Run `./scripts/update-go-version.sh`
4. Run tests: `make test`
5. Build Docker images: `make docker-build`
6. If all pass, commit and push

---

## Example: Updating to Go 1.25

```bash
# 1. Update version file
echo "1.25" > .go-version

# 2. Run update script
./scripts/update-go-version.sh

# Output:
# 🔄 Updating Go version to 1.25 across all files...
# 📝 Updating go.mod...
# 🐳 Updating Dockerfiles...
# ⚙️  Updating GitHub Actions workflows...
# ✅ Go version updated to 1.25 in all files!

# 3. Verify changes
git diff

# 4. Test locally
make test
make docker-build

# 5. Commit
git add .
git commit -m "chore: update Go version to 1.25"
git push
```

---

## Maintenance

### Regular Tasks
- **Monthly**: Check for new Go releases
- **Quarterly**: Update to latest stable version
- **Yearly**: Review and update dependencies

### Monitoring
- Watch Go release announcements
- Subscribe to Go security mailing list
- Check Dependabot alerts

---

**Last Updated**: February 28, 2026  
**Current Go Version**: 1.24  
**Next Review**: May 2026
