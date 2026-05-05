# Build stage for Go backend
FROM golang:1.22-bookworm AS backend-builder
WORKDIR /app
COPY backend/go.mod backend/go.sum ./
RUN go mod download
COPY backend/ ./
RUN CGO_ENABLED=0 GOOS=linux go build -o /langops-server ./cmd/server

# Build stage for frontend
FROM node:20-bookworm AS frontend-builder
WORKDIR /app
COPY frontend/package.json frontend/package-lock.json* ./
RUN npm install
COPY frontend/ ./
RUN npm run build

# Runtime stage
FROM debian:bookworm-slim
RUN apt-get update && apt-get install -y ca-certificates && rm -rf /var/lib/apt/lists/*
WORKDIR /app
COPY --from=backend-builder /langops-server ./langops-server
COPY --from=frontend-builder /app/dist ./frontend/dist
COPY configs/ ./configs/
EXPOSE 8080
CMD ["./langops-server"]