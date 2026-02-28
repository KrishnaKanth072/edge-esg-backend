# Contributing to EDGE ESG Backend

Thank you for your interest in contributing to EDGE ESG Backend! This document provides guidelines for contributing to the project.

## Development Setup

1. **Prerequisites**
   - Go 1.21 or higher
   - Docker & Docker Compose
   - Make
   - Git

2. **Clone and Setup**
   ```bash
   git clone https://github.com/KrishnaKanth072/edge-esg-backend.git
   cd edge-esg-backend
   make up
   ```

3. **Run Tests**
   ```bash
   make test
   ```

## Branch Strategy

- `main` - Production-ready code
- `dev` - Development branch
- `feature/*` - New features
- `bugfix/*` - Bug fixes
- `hotfix/*` - Urgent production fixes

## Workflow

1. **Create a Feature Branch**
   ```bash
   git checkout dev
   git pull origin dev
   git checkout -b feature/your-feature-name
   ```

2. **Make Changes**
   - Write clean, documented code
   - Follow Go best practices
   - Add tests for new features
   - Update documentation

3. **Test Your Changes**
   ```bash
   make test
   make lint
   ```

4. **Commit**
   ```bash
   git add .
   git commit -m "feat: add new feature description"
   ```

   Use conventional commits:
   - `feat:` - New feature
   - `fix:` - Bug fix
   - `docs:` - Documentation changes
   - `chore:` - Maintenance tasks
   - `refactor:` - Code refactoring
   - `test:` - Test updates

5. **Push and Create PR**
   ```bash
   git push origin feature/your-feature-name
   ```
   Then create a Pull Request to `dev` branch on GitHub.

## Code Standards

- Follow Go formatting: `gofmt -s -w .`
- Pass linting: `golangci-lint run`
- Maintain test coverage above 70%
- Document exported functions and types
- Use meaningful variable names

## RBI Compliance

When contributing, ensure:
- Data encryption for sensitive fields
- Audit trail for all operations
- Rate limiting on API endpoints
- Input validation and sanitization
- No hardcoded credentials

## Questions?

Open an issue or contact the maintainers.
