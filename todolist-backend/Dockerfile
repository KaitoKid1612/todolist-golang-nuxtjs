# Bước 1: Chọn hình ảnh cơ sở
FROM golang:1.23 AS builder

# Bước 2: Thiết lập thư mục làm việc
WORKDIR /app

# Bước 3: Sao chép go.mod và go.sum và cài đặt các dependency
COPY go.mod go.sum ./
RUN go mod download

# Bước 4: Sao chép mã nguồn vào container
COPY . .

# Bước 5: Biên dịch ứng dụng
RUN go build -o main .

# Bước 6: Chạy ứng dụng
CMD ["./main"]
