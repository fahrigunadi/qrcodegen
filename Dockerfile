# Stage 1: Build frontend
FROM node:20-bullseye AS frontend-builder
WORKDIR /app
COPY package*.json ./
RUN npm install
COPY . .
RUN npm run build

# Stage 2: Build Go
FROM golang:1.23-alpine AS backend-builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod tidy
COPY . .

# Copy frontend build
COPY --from=frontend-builder /app/public/build ./public/build

RUN go build -o qrcodegen .

# Stage 3: Final image
FROM alpine:latest
WORKDIR /app
COPY --from=backend-builder /app/qrcodegen .
COPY --from=backend-builder /app/public ./public
COPY --from=backend-builder /app/resources ./resources
EXPOSE 3000
CMD ["./qrcodegen"]
