import clickhouse_driver

# ClickHouse connection details
client = clickhouse_driver.Client('clickhouse', user='log', password='123456789', port=9000)

# Create database if it doesn't exist
create_db_query = "CREATE DATABASE IF NOT EXISTS logs;"
client.execute(create_db_query)

# Create logs table
create_logs_table = """
CREATE TABLE IF NOT EXISTS logs.logs (
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
