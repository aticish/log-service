import clickhouse_driver
import random
import ipaddress
from datetime import datetime
import time
import os

# Get environment variables
clickhouse_user = os.environ.get('CLICKHOUSE_USER')
clickhouse_password = os.environ.get('CLICKHOUSE_PASSWORD')
clickhouse_database = os.environ.get('CLICKHOUSE_DATABASE')
clickhouse_table = os.environ.get('CLICKHOUSE_TABLE')

# Clickhouse connection
client = clickhouse_driver.Client('clickhouse', user=clickhouse_user, password=clickhouse_password, port=9000, database=clickhouse_database)

# How many records will be added?
num_records = 10000000

# Random Datas
data = []
severities = ['emergency', 'alert', 'critical', 'error', 'warning', 'notice', 'info', 'debug']
actions = ['login', 'logout', 'register', 'password_change', 'update_profile', 'delete_account', 'reset_password_request', 'update_post', 'create_post', 'delete_post', 'failed_login_attempt']
agents = [
    'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/110.0.5481.178 Safari/537.36',
    'Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/110.0.5481.178 Safari/537.36',
    'Mozilla/5.0 (Linux; Android 13; Pixel 6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/110.0.5481.178 Mobile Safari/537.36',
    'Mozilla/5.0 (iPhone; CPU iPhone OS 16_3 like Mac OS X) AppleWebKit/537.36 (KHTML, like Gecko) CriOS/110.0.5481.178 Mobile/15E148 Safari/537.36',
    'Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:110.0) Gecko/20100101 Firefox/110.0',
    'Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7; rv:110.0) Gecko/20100101 Firefox/110.0',
    'Mozilla/5.0 (Macintosh; Intel Mac OS X 13_0_1) AppleWebKit/537.36 (KHTML, like Gecko) Version/16.1 Safari/537.36'
]

start_time = time.time()
for _ in range(num_records):
    user_id = random.randint(1, 9999999)
    severity = random.choice(severities)
    user_ip = str(ipaddress.IPv4Address(random.randint(0, 2**32 - 1)))
    action = random.choice(actions)
    content = ""
    agent = random.choice(agents)
    timestamp = int(time.time())

    data.append((user_id, severity, user_ip, action, content, agent, timestamp))

client.execute(
    f'INSERT INTO {clickhouse_database}.{clickhouse_table} (user_id, severity, user_ip, action, content, agent, timestamp) VALUES',
    data
)

end_time = time.time()
total_time = end_time - start_time
print(f"{num_records:,} records added in {total_time:.4f} seconds")
