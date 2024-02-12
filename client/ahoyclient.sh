#!/bin/bash
alias sail="./ahoyclient"

version="1.00"
# option: -d  /[DIRECTORY REQUESTED]/
sudo systemctl start httpd
sudo systemctl status httpd




# option: -help
Help(){
echo "Ahoy! version:$version"
echo "options:"









}


while getopts ":h" option; do
   case $option in
      h) # display Help
         Help
         exit;;
   esac
done
