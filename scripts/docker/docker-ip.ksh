docker inspect --format '{{ .NetworkSettings.IPAddress }}' "$@"
