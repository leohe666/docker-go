#!/bin/bash

# 获取当前脚本目录
project_path=$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )
# 将脚本目录添加到用户shell配置文件中
echo 'PATH='$project_path':$PATH' >> ~/.bashrc
curl -sSfL https://raw.githubusercontent.com/cosmtrek/air/master/install.sh | sh -s -- -b $PWD/packages/bin