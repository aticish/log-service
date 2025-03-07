# A simple yet powerful logging microservice

The Log Microservice is a lightweight and scalable logging solution designed to efficiently record and query system events. Built with Docker, Golang, and ClickHouse, this service offers high-performance log processing while maintaining a minimal and straightforward structure.

Despite its simplicity, it provides flexible query support and a scalable architecture, making it suitable for systems handling large volumes of data. With easy setup and low resource consumption, it can be seamlessly integrated into projects of any size.

### Requirements

- Docker

### Installation

1- First, clone this project into your working directory and open it in the terminal:

```bash
git clone git@github.com:aticish/log-service.git log-service && cd log-service
```

2- Rename the `env.local` file to `.env` and configure its contents according to your setup. The `LOGAPI_TOKEN` inside this file will serve as your bearer token and will be used in requests.

3- Run the following command in the root directory of the project to start:

```bash
docker-compose up
```

This command will set up the project, create the required database and table for ClickHouse, and insert 30 million dummy log records. If you do not want to insert these records, simply remove the Python container from the `docker-compose.yml` file.

**Note:** If you remove the Python container, you will need to manually create the database and tables yourself. You can look `docker/dummy/create-db.py` file for database structure...

4- Add the following lines to your hosts file:

```bash
127.0.0.1   traefik.logapi.test  # from .env file
127.0.0.1   logapi.test  # from .env file
```

5- Enjoy 🚀

### Usage

- All requests must be of type POST and should be made to the following endpoint:

```bash
http://domain.com/api/v1/
```

Requests made to other routes or the root directory will return a `404 Not Found` response.

- Requests must include the following header:

```bash
Authentication: Bearer LOGAPI_TOKEN
```

If this header is missing or invalid, the server will return a `401 Unauthorized` error.

#### Example Request for Reading Logs

To read logs, you can use the following request example:

```json
{
    "method": "read",
    "data": {
        "user_ip": "*.*.23.150",
        "user_id": "*",
        "action": "*",
        "severity": "*",
        "start_date": "2023-09-18 10:11:11",
        "ending_date": "2023-09-19 10:11:11",
        "page": 1,
        "limit": 1000,
        "order": "timestamp",
        "direction": "DESC"
    }
}
```

- **method**: Must be either `write` or `read`.
- **data**: Contains filtering options:
    - **user_ip**: Defines IP filtering rules. Supported formats:
        - `"*"`: Includes logs from all IP addresses.
        - `"*.*.*.123"`: Retrieves logs where the IP address ends in `123`
        - `"127.0.0.1,127.0.0.2"`: Retrieves logs for specific IPs (`127.0.0.1` and `127.0.0.2`).
        - `"123"`: Retrieves logs where the IP address contains `123`
        - `"127.0.0.1"`: Retrieves logs only for `127.0.0.1`
    - **user_id**: Filters logs by user ID.
        - `"*"`: Includes logs from all users
        - `"123"`: Retrieves logs where the user id is `123``
        - `"16,26,33,9573"`: Retrieves logs where the user ids are `16`, `26`, `33` and `9573`
    - **action**: Defines action filtering rules.
        - `"*"`: all actions
        - `"login"`: only `login` logs
        - `"register,login"`: `register` and `login` logs
    - **severity**: Defines severity filtering rules: Supported severities are, `emergency`, `alert`, `critical`, `error`, `warning`, `notice`, `info`, `debug`
        - `"*"`: all severities
        - `"emergency"`: only `emergency` logs
        - `"emergency,alert"`: `emergency` and `alert` logs
    - **start_date**: Defines the start time for filtering logs. (format: `YYYY-MM-DD HH:MM:SS`). If left blank, it filters starting from the logs 1 month ago.
    - **ending_date**: Defines the end time for filtering logs (format: `YYYY-MM-DD HH:MM:SS`). If left blank, tomorrow's date (today + 1 day) will be used as basis.
    - **page**: The page number for pagination.
    - **limit**: The number of results per page. Must be between 1 and 10000. Default value is 1000.
    - **order**: Sorting field. Supported values are `timestamp`, `user_id`, `user_ip`, `severity`, `action`. Default is `timestamp`.
    - **direction**: Sorting direction (`"ASC"` for ascending, `"DESC"` for descending`).

#### Example Request for Writing Logs

To log an event, use the following request example:

```json
{
    "method": "write",
    "data": {
        "user_ip": "11.33.0.1",
        "user_id": "1",
        "action": "login",
        "severity": "info",
        "agent": "mozilla/5.0 (macintosh; intel mac os x 10.15; rv:136.0) gecko/20100101 firefox/136.0",
        "content": {}
    }
}
```

- **method:** "write" – creates a new log entry.
- **data:** Contains log details:
    - **user_ip:** The IP address of the user initiating the action.
    - **user_id:** The ID of the user performing the action.
    - **action:** The action performed (e.g., "login").
    - **severity:** The log severity level (e.g., "info", "warning", "error").
    - **agent:** The user agent string from the request, capturing browser and device information.
    - **content:** An object that can store additional data related to the log entry.
