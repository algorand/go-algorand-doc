#! /bin/bash

# start network
/root/node/goal network start -r /root/randpool -d /root/randpool

# UGLY HACK TODO
tail -f /dev/null