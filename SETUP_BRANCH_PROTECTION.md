# Branch Protection Setup Guide

## 🔒 Setup Branch Protection Rules

Follow these steps to require PR approval and passing builds before merging:

### Step 1: Go to Repository Settings

1. Open: https://github.com/KrishnaKanth072/edge-esg-backend
2. Click **"Settings"** tab (top right)
3. Click **"Branches"** (left sidebar)

---

### Step 2: Add Branch Protection Rule

1. Click **"Add branch protection rule"**
2. In **"Branch name pattern"**, type: `main`

---

### Step 3: Configure Protection Rules

Check these boxes:

#### ✅ Require a pull request before merging
- Check this box
- Then check: **"Require approvals"**
  - Set **"Required number of approvals"** to: `1`
- Check: **"Dismiss stale pull request approvals when new commits are pushed"**
- Check: **"Require review from Code Owners"** (optional)

#### ✅ Require status checks to pass before merging
- Check this box
- Check: **"Require branches to be up to date before merging"**
- In the search box, add these status checks:
  - `lint`
  - `test`
  - `build`
  - `security`
  - `compliance-check`

#### ✅ Require conversation resolution before merging
- Check this box

#### ✅ Do not allow bypassing the above settings
- Check this box

#### ❌ Allow force pushes
- Leave UNCHECKED

#### ❌ Allow deletions
- Leave UNCHECKED

---

### Step 4: Save Changes

1. Scroll to bottom
2. Click **"Create"** or **"Save changes"**

---

## 🎯 What This Does:

### Before Protection:
- ❌ Anyone can push directly to main
- ❌ No approval needed
- ❌ Can merge even if tests fail

### After Protection:
- ✅ **Must create Pull Request** (no direct push to main)
- ✅ **All CI/CD checks must pass** (lint, test, build, security)
- ✅ **You must approve** before merge
- ✅ **Conversations must be resolved**
- ✅ **Branch must be up-to-date**

---

## 📋 Workflow After Setup:

### For You (Admin):

1. Create feature branch:
```bash
git checkout -b feature/new-feature
# Make changes
git add .
git commit -m "feat: add new feature"
git push origin feature/new-feature
```

2. Create Pull Request on GitHub
3. Wait for CI/CD to pass (all checks green ✅)
4. Review your own PR
5. Click **"Approve"**
6. Click **"Merge pull request"**

### For Collaborators:

1. Create feature branch
2. Push changes
3. Create Pull Request
4. Wait for CI/CD to pass
5. **Wait for YOUR approval** ⏳
6. After you approve → They can merge

---

## 🔐 Additional Security (Optional):

### Require Signed Commits

1. In branch protection settings
2. Check: **"Require signed commits"**
3. This ensures all commits are verified

### Restrict Who Can Push

1. Check: **"Restrict who can push to matching branches"**
2. Add only trusted users

### Require Deployments to Succeed

1. Check: **"Require deployments to succeed before merging"**
2. Select environments: `staging`, `production`

---

## 🧪 Test the Protection:

### Try Direct Push (Should Fail):
```bash
git checkout main
echo "test" >> test.txt
git add test.txt
git commit -m "test: direct push"
git push origin main
```

**Expected:** ❌ Error: "protected branch hook declined"

### Correct Way (Should Work):
```bash
git checkout -b test/branch-protection
echo "test" >> test.txt
git add test.txt
git commit -m "test: via PR"
git push origin test/branch-protection
# Then create PR on GitHub
```

**Expected:** ✅ PR created, waiting for checks and approval

---

## 📊 Status Checks Required:

These checks must pass before merge:

1. **lint** - Code quality (golangci-lint)
2. **test** - Unit tests
3. **build** - Build verification
4. **security** - Security scan (gosec)
5. **compliance-check** - RBI compliance checks

---

## 🎯 Quick Setup Checklist:

- [ ] Go to Settings → Branches
- [ ] Add rule for `main` branch
- [ ] ✅ Require pull request
- [ ] ✅ Require 1 approval
- [ ] ✅ Require status checks (lint, test, build, security)
- [ ] ✅ Require conversation resolution
- [ ] ✅ Do not allow bypassing
- [ ] ❌ Disable force pushes
- [ ] ❌ Disable deletions
- [ ] Click "Create"

---

## 🎉 Done!

Your `main` branch is now protected!

**Only you can approve and merge PRs after all checks pass!** 🔒

---

**Setup URL:** https://github.com/KrishnaKanth072/edge-esg-backend/settings/branches
