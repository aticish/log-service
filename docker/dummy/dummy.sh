#!/bin/bash

# Run the initial setup command only once
if [ ! -d "/dummy/venv" ]; then
    cd "/dummy" || exit
    python -m venv venv && . venv/bin/activate && python -m pip install --upgrade pip && pip install clickhouse_driver ipaddress && python3 /dummy/create-db.py
fi

# Run the insert commands three times (insert 30 million records)
for i in {1..3}; do
    if [ ! -f "/dummy/.insert_done" ]; then
        cd "/dummy" || exit
        . venv/bin/activate && python3 /dummy/insert-data.py && sleep 1
    fi
done
touch /dummy/.insert_done
