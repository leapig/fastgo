#!/bin/sh

# 获取本地 IP
local_ip=$(ifconfig | grep "inet " | grep -v 127.0.0.1 | awk '{print $2}' | cut -d '/' -f1 | head -n 1)
echo "本地 IP: $local_ip"

# 临时文件
temp_file="temp.toml"

# 读取 frpc.toml 文件，将 localIP 替换为获取到的本地 IP
while IFS= read -r line; do
    if [[ $line =~ ^localIP.* ]]; then
        echo "localIP = \"$local_ip\"" >> $temp_file
    else
        echo "$line" >> $temp_file
    fi
done < frpc.toml

# 将更新后的临时文件替换原有的 frpc.toml 文件
mv $temp_file frpc.toml

# 停止并删除 frpc 容器
docker stop frpc
docker rm frpc

# 启动新的 frpc 容器，加载更新后的 frpc.toml 文件
docker run -d --name frpc -v $PWD/frpc.toml:/etc/frp/frpc.toml snowdreamtech/frpc:latest
