#! /bin/bash

# start kmd
/app/goal kmd start -d /app/nodes/randpool/Primary

# start network
/app/goal network start -r /app/nodes/randpool

# UGLY HACK TODO
tail -f /dev/null