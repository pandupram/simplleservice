# Gunakan image Golang sebagai base image
FROM golang:1.22-alpine

# Buat direktori kerja di dalam container
WORKDIR /app

# Copy semua file ke dalam container
COPY . .

# Unduh semua dependensi
RUN go mod download

# Build aplikasi
RUN go build -o main .

# Ekspos port yang digunakan oleh aplikasi
EXPOSE 8080

# Jalankan aplikasi
CMD ["./main"]
