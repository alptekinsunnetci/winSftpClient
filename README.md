# winSFTPClient

**winSFTPClient** is a simple Go-based SFTP client that allows you to upload files and directories to a remote server via SFTP. This tool supports connection to an SFTP server using username, password authentication, and allows recursive uploads of directories.

The code is open for modification, and you are free to customize it for your needs. Please read the instructions below to get started with the application.

## Features

- Upload files and directories recursively from your local machine to a remote server.
- Supports username/password authentication for secure connections.
- Customizable to suit your SFTP server and local path configurations.
- Debug mode for more detailed logging.

## Prerequisites

To build and run **winSFTPClient**, you need to have the following installed:

- Go (1.16 or higher)
- Git (for cloning the repository)

## Installation

1. Clone the repository:

    ```bash
    git clone https://github.com/alptekinsunnetci/winSFTPClient.git
    ```

2. Navigate into the project directory:

    ```bash
    cd winSFTPClient
    ```

3. Install required Go packages:

    ```bash
    go mod tidy
    ```

4. Build the project:

    ```bash
    go build -o winSFTPClient.exe main.go
    ```

## Usage

After building the application, you can run **winSFTPClient** from the command line.

### Command Line Arguments:

- `-u <username>`: Your SFTP server username.
- `-p <password>`: Your SFTP server password.
- `-s <host:port>`: The SFTP server's host and port.
- `-l <localDirectory>`: The local directory you want to upload.
- `-r <remoteDirectory>`: The remote directory where you want to upload files.
- `-debug`: Optional flag to enable debug logging for more detailed output.

### Example Command:

```bash
winSFTPClient.exe -u yourUser -p yourPass -s 192.168.1.1:22 -l "C:\local\path" -r "/remote/path" -debug
