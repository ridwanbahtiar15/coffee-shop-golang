FROM golang:1.21.4-alpine

# Environtmen tambahan
ENV GO_ENV=DOCKER DB_HOST=host.docker.internal

# menentukan working directory di container
WORKDIR /app

# copy project ke working directory
COPY . .

# Jalankan perintah (instalasi, build, dll) di container
# 1. Install Dependency
RUN go mod download
# 2. 
RUN go build -v -o /app/coffee-shop-golang ./cmd/main.go

# Expose port
EXPOSE 8001

# Daftarkan aplikasi
ENTRYPOINT [ "/app/coffee-shop-golang" ]