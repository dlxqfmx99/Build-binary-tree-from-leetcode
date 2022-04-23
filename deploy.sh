#!/bin/bash



echo -e "\033[0;32mDeploying updates to GitHub...\033[0m"



git add .
msg="rebuilding site `date`"
if [ $# -eq 1 ]

  then msg="$1"

fi

git commit -m "$msg"

# push source to github

git push -f origin main:main