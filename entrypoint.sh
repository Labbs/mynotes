#!/bin/sh
set -e

# Lance les migrations
./app migration -c /config/config.yaml

# Lance le serveur
exec ./app server -c /config/config.yaml