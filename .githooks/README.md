# Git Hooks Setup

## Installation

To enable these hooks locally, run:

```bash
git config core.hooksPath .githooks
chmod +x .githooks/pre-commit
```

## Available Hooks

- `pre-commit`: Checks for hardcoded passwords and secrets before allowing commits
