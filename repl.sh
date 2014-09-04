#!/bin/bash

for i in `find .|grep -v repl.sh` ; do sed -i 's#github.com/mitchellh/packer#github.com/outscale/packer#g' $i; done

