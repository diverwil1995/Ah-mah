# syntax=docker/dockerfile:1
# 從Docker Hub下載名為 golang的官方映像檔 版本號 1.19
FROM golang:1.19
# 指定容器內的工作目錄為 "/app"
WORKDIR /app
# 將本地目錄的內容複製到容器目錄
COPY . .
# 執行命令 "go mod download" 依據 go.mod 逐一下載依賴套件
RUN go mod download
# 執行編譯命令 "go build" 將 .go文件編譯並生成可執行文件
RUN CGO_ENABLED=0 GOOS=linux go build
# 指定容器運行時開放 8080 Port
EXPOSE 8080
# 運行 main 這個編譯過的可執行文件
CMD [ "/main" ]