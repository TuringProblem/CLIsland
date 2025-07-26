# Setup Guide

This guide will help you set up the complete CI/CD pipeline for your Go project.

## Prerequisites

1. **GitHub Repository**: Your project should be hosted on GitHub
2. **Go 1.21+**: Ensure you're using Go 1.21 or later
3. **Codecov Account**: Set up at [codecov.io](https://codecov.io)

## Step 1: GitHub Repository Setup

### 1.1 Enable GitHub Actions
- Go to your repository settings
- Navigate to "Actions" → "General"
- Ensure "Allow all actions and reusable workflows" is selected

### 1.2 Add Repository Secrets
Go to Settings → Secrets and variables → Actions and add:

```
CODECOV_TOKEN=your_codecov_token_here
```

To get your Codecov token:
1. Go to [codecov.io](https://codecov.io)
2. Add your repository
3. Copy the token from your repository settings

## Step 2: Codecov Setup

### 2.1 Add Repository to Codecov
1. Visit [codecov.io](https://codecov.io)
2. Sign in with GitHub
3. Add your repository
4. Copy the token and add it to GitHub secrets

### 2.2 Configure Codecov
The `.codecov.yml` file is already configured with:
- 80% coverage target
- Coverage thresholds
- Ignored files (tests, stubs, etc.)

## Step 3: Local Development Setup

### 3.1 Install Development Tools
```bash
# Install all required tools
make dev-setup
```

This installs:
- `golangci-lint` - Code linting
- `goimports` - Import formatting
- `gosec` - Security scanning

### 3.2 Verify Setup
```bash
# Run all CI checks locally
make ci

# Check available targets
make help
```

## Step 4: GitHub Actions Workflows

### 4.1 CI Workflow (`.github/workflows/ci.yml`)
This workflow runs on every push and PR:

**Jobs:**
- **Test**: Runs tests with Go 1.21 and 1.22, uploads coverage to Codecov
- **Lint**: Runs golangci-lint with comprehensive rules
- **Build**: Builds for multiple platforms (Linux, macOS, Windows)
- **Security**: Runs gosec security scanner

### 4.2 Release Workflow (`.github/workflows/release.yml`)
This workflow runs when you push a tag:

**Features:**
- Builds for all platforms
- Creates GitHub release
- Generates checksums
- Auto-generates release notes

### 4.3 QLTY.dev Workflow (`.github/workflows/qulty.yml`)
Optional workflow for code quality analysis:

**Setup:**
1. Install the QLTY.dev GitHub app
2. Uncomment the API key line if you have one
3. Add `QLTY_API_KEY` to repository secrets

## Step 5: Making Releases

### 5.1 Create a Release
```bash
# Tag your release
git tag v1.0.0

# Push the tag
git push origin v1.0.0
```

The release workflow will automatically:
- Build for all platforms
- Create a GitHub release
- Upload artifacts
- Generate release notes

### 5.2 Local Release Testing
```bash
# Test release build locally
make release
```

## Step 6: Badges and Status

### 6.1 Add Badges to README
The README already includes these badges:

```markdown
[![Go Version](https://img.shields.io/badge/go-1.21+-blue.svg)](https://golang.org)
[![Build Status](https://github.com/TuringProblem/CLIsland/workflows/CI/badge.svg)](https://github.com/TuringProblem/CLIsland/actions)
[![Test Coverage](https://codecov.io/gh/TuringProblem/CLIsland/branch/main/graph/badge.svg)](https://codecov.io/gh/TuringProblem/CLIsland)
[![Go Report Card](https://goreportcard.com/badge/github.com/TuringProblem/CLIsland)](https://goreportcard.com/report/github.com/TuringProblem/CLIsland)
[![GoDoc](https://godoc.org/github.com/TuringProblem/CLIsland?status.svg)](https://godoc.org/github.com/TuringProblem/CLIsland)
[![License](https://img.shields.io/badge/license-MIT-blue.svg)](LICENSE)
[![Release](https://img.shields.io/github/v/release/TuringProblem/CLIsland)](https://github.com/TuringProblem/CLIsland/releases)
```

### 6.2 Update Badge URLs
Replace `TuringProblem/CLIsland` with your repository name in the badge URLs.

## Step 7: Configuration Files

### 7.1 golangci-lint (`.golangci.yml`)
Configured with:
- 50+ linters enabled
- Custom rules for your project
- Exclusions for test files and stubs
- 5-minute timeout

### 7.2 Codecov (`.codecov.yml`)
Configured with:
- 80% coverage target
- Coverage thresholds
- Ignored files
- Branch detection

### 7.3 Dependabot (`.github/dependabot.yml`)
Automated dependency updates:
- Weekly Go module updates
- Weekly GitHub Actions updates
- Auto-assignment and labeling

## Step 8: Troubleshooting

### 8.1 Common Issues

**Build Failures:**
```bash
# Check Go version
go version

# Clean and rebuild
make clean
make build
```

**Lint Failures:**
```bash
# Run linter locally
make lint

# Auto-fix what can be fixed
make lint-fix
```

**Test Failures:**
```bash
# Run tests locally
make test-unit

# Run with coverage
make test-coverage
```

### 8.2 GitHub Actions Issues

**Workflow Not Running:**
- Check repository settings → Actions → General
- Ensure workflows are enabled
- Check branch protection rules

**Secret Issues:**
- Verify `CODECOV_TOKEN` is set correctly
- Check token permissions in Codecov

### 8.3 Codecov Issues

**Coverage Not Uploading:**
- Verify token is correct
- Check workflow logs for upload errors
- Ensure tests are generating coverage files

## Step 9: Advanced Configuration

### 9.1 Custom Linting Rules
Edit `.golangci.yml` to:
- Add/remove linters
- Adjust thresholds
- Add custom exclusions

### 9.2 Coverage Thresholds
Edit `.codecov.yml` to:
- Change coverage targets
- Adjust thresholds
- Modify ignored files

### 9.3 Build Matrix
Edit `.github/workflows/ci.yml` to:
- Add more Go versions
- Add more platforms
- Customize build flags

## Step 10: Monitoring

### 10.1 GitHub Actions
- Monitor workflow runs in Actions tab
- Set up notifications for failures
- Review build artifacts

### 10.2 Codecov
- Monitor coverage trends
- Review coverage reports
- Set up coverage alerts

### 10.3 Dependabot
- Review dependency updates
- Monitor security advisories
- Configure update schedules

## Conclusion

Your project now has:
- ✅ Automated testing on multiple Go versions
- ✅ Comprehensive linting with golangci-lint
- ✅ Security scanning with gosec
- ✅ Code coverage tracking with Codecov
- ✅ Automated releases with multi-platform builds
- ✅ Dependency updates with Dependabot
- ✅ Professional badges and documentation

The CI/CD pipeline will help maintain code quality and automate the release process! 