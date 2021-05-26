#!/bin/bash
cd $1
rm -rf "$5"

echo "开始拉取代码<br>"
git clone http://$2:$3@$4.git $5
echo "拉取代码完成<br>"
git config --global user.email "$2@TTjianbao.com"
git config --global user.name "$2"