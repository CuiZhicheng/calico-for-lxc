#!/bin/bash
name=$1
network=$2
pid=$(sudo lxc-info -n $name | grep 'PID' | cut -d ':' -f 2 |  tr -d '[[:space:]]')

sudo CNI_PATH=/opt/bin /opt/bin/cnitool add $network /var/run/netns/$pid
