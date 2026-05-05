# Contributing to LangOps Enterprise

Thank you for your interest in contributing to LangOps Enterprise!

## Getting Started

1. Fork the repository
2. Create your feature branch: `git checkout -b feature/amazing-feature`
3. Commit your changes: `git commit -m 'Add amazing feature'`
4. Push to the branch: `git push origin feature/amazing-feature`
5. Open a Pull Request

## Development Setup

`ash
# Backend
cd backend && go mod download

# Eval Engine
cd eval && pip install -r requirements.txt

# Frontend
cd frontend && npm install
``

## Code Standards

- Go: Follow Effective Go and gofmt
- Python: Follow PEP 8, use type hints
- TypeScript: Follow ESLint + Prettier config
- All PRs require at least one review

## Reporting Issues

Please use GitHub Issues with appropriate labels.