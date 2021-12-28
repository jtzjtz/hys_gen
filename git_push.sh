#!/bin/bash
cd $1
pwd
echo "开始push $1<br>"

git add -A
git commit  -m "$2"
git tag -a $3 -m "$2"

git push https://$4 master:main --tags
echo "push 完成"
#git push  http://github.com/jtzjtz/hys_gen.git master:master --tags
