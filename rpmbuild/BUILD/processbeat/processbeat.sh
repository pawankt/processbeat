#!/bin/bash

# Script to run processbeat in foreground with the same path settings that
# the init script / systemd unit file would do.

/usr/share/processbeat/bin/processbeat \
  -path.home /usr/share/processbeat \
  -path.config /etc/processbeat \
  -path.data /var/lib/processbeat \
  -path.logs /var/log/processbeat \
  $@
