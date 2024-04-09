#!/bin/bash
set -e

exec dumb-init -- llmos-dashboard "${@}"
