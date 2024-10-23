# Bước 1: Thiết lập hình ảnh cơ sở
FROM node:latest AS builder

# Bước 2: Thiết lập thư mục làm việc
WORKDIR /app

# Bước 3: Sao chép package.json và package-lock.json (nếu có)
COPY package*.json ./

# Bước 4: Cài đặt dependencies
RUN npm install

# Bước 5: Sao chép mã nguồn vào container
COPY . .

# Bước 6: Biên dịch ứng dụng
RUN npm run build

# Bước 7: Chạy ứng dụng
CMD ["npm", "start"]
