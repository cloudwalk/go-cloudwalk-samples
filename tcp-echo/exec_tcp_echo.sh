#!/bin/bash

upstate=`nmap localhost | grep 800 | cut -c1-3`

if [ "$upstate" != "800" ]; 
then
    `./tcp-echo &`
    `^C`
fi
exit