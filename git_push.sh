#!/bin/bash
cd $1
pwd
echo "开始push<br>"

git add -A
git commit  -m "$2"
git push origin master
git tag -a $3 -m "$2"
git push --tags
echo "push 完成"
git remote rm origin
