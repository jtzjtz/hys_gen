#!/bin/bash
cd $1
pwd
echo "开始push<br>"

git add -A
git commit  -m "$2"
git tag -a $3 -m "$2"

git push https://$4 master:main --tags
echo "push 完成"
