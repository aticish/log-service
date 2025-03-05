import clickhouse_driver
import os

clickhouse_user = os.environ.get('CLICKHOUSE_USER')
clickhouse_password = os.environ.get('CLICKHOUSE_PASSWORD')
clickhouse_database = os.environ.get('CLICKHOUSE_DATABASE')
clickhouse_table = os.environ.get('CLICKHOUSE_TABLE')

# ClickHouse connection details
client = clickhouse_driver.Client('clickhouse', user=clickhouse_user, password=clickhouse_password, port=9000)

# Create database if it doesn't exist
create_db_query = f"CREATE DATABASE IF NOT EXISTS {clickhouse_database};"
client.execute(create_db_query)

# Create logs table
create_logs_table = f"""
CREATE TABLE IF NOT EXISTS {clickhouse_database}.{clickhouse_table} (
    user_id UInt64,
    user_ip String,
    severity String,
    action String,
    content String,
    agent String,
    timestamp DateTime('Asia/Istanbul'),
    INDEX idx_timestamp timestamp TYPE minmax GRANULARITY 1,
    INDEX idx_user_id user_id TYPE minmax GRANULARITY 1,
    INDEX idx_user_ip user_ip TYPE minmax GRANULARITY 1,
    INDEX idx_severity severity TYPE minmax GRANULARITY 1,
    INDEX idx_action action TYPE minmax GRANULARITY 1
) ENGINE = MergeTree()
ORDER BY (timestamp, user_id, severity, action);
"""

client.execute(create_logs_table)

print("The logs table was created successfully.")
