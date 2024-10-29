FROM harbor.d-cos.com:8443/ci/alpine:202408131014
# 处理文件
#WORKDIR /app
#COPY app .
#COPY app/dist /app/
#COPY files /app/
RUN apk add -U nginx
COPY nginx.conf /etc/nginx/http.d/default.conf
COPY ssl.crt /etc/nginx/http.d/ssl.crt
COPY ssl.key /etc/nginx/http.d/ssl.key
# 自动启动
#COPY start.sh .
#RUN chmod +x ./app ./start.sh
CMD ["nginx"]