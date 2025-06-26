# Go URL Shortener - BETA
![overview](assest/program.gif)

A simple URL shortening service written in Go. It reads URL mappings from a CSV file and provides short URLs that redirect to the original URLs.

---

## Features

- Shortens long URLs into short URLs of the form `http://l.sh/<id>`
- Prevents duplicate entries for the same long URL
- Stores URL data in a CSV file (`data.csv`)
- Redirects HTTP requests to the original URLs based on short URL path
- Simple command-line interface to add new URLs
- Validates URLs before shortening
- Automatically creates `data.csv` with example data if missing
- Runs as a root process to bind on port 80

---

## Installation & Setup

1. Make sure you have [Go](https://golang.org/dl/) installed.

2. Clone the repository and navigate to the project folder.

3. Create or verify `data.csv` exists (the program creates it automatically if missing).

4. Run the URL shortener CLI to add URLs:
    ``` bash
    sudo go run client.go
    ```
5. Run the HTTP server as root (to bind port 80):
    ``` bash
    sudo go run server.go
    ```

---

## Usage

### Add a URL

Run `client.go` and enter the long URL when prompted. You will receive a short URL like `http://l.sh/abc1`.

### Redirect

Open a browser and go to `http://l.sh/<short-id>`. The server will redirect to the original URL.

---

## CSV Format

The CSV file (`data.csv`) has the following columns:

| no | shorten_url         | id   | long_url                  | date        |
|----|---------------------|------|---------------------------|-------------|
| 1  | http://l.sh/test123 | test | https://example.com       | 1728850133  |

- **no**: incremental number
- **shorten_url**: full short URL
- **id**: unique short identifier (used in URL path)
- **long_url**: original URL
- **date**: Unix timestamp of when added

---
### Note

To make the URL shortener work locally, you may need to add the following entry to your system’s hosts file:

- On Windows: `C:\Windows\System32\drivers\etc\hosts`  
- On Linux/macOS: `/etc/hosts`

Add this line to the hosts file (use a plain text editor with administrator/root privileges):

> 0.0.0.0 l.sh

This will redirect requests for `l.sh` to your local machine, allowing you to test the shortener service internally.  
**Note:** This only works internally on your machine and does not affect external DNS resolution.

