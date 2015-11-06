ifconfig $1 2>/dev/null | grep "inet addr" | awk -F: '{print $2}' | awk '{print $1}'
