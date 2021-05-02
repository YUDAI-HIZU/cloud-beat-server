#!/bin/sh

set -e

sql-migrate status

exec "$@"