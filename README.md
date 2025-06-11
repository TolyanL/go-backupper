# Backupper 🔄

Backupper is my tool designed to automate the backup process of my applications. It supports executing remote jobs, copying generated backup files to a local storage directory, and compressing them efficiently using Zstandard (zstd). 💾⚡

## Features ✨

*   **📝 Configurable Tasks:** Define multiple backup tasks with specific remote hosts, users, and job details.  
*   **🌐 Remote Job Execution:** Integrate with various remote backup strategies (e.g., `docker exec` for database dumps within containers, simple file copies).  
*   **💽 Local Storage:** Automatically copy backup files from remote servers to a specified local directory.  
*   **🗜️ Efficient Compression:** Utilizes Zstandard (`.zst`) for fast and effective compression of backup archives.  
*   **📊 Flexible Logging:** Configurable logging output (JSON or text) and verbosity levels.  
*   **🧹 Temporary File Cleanup:** Automatically removes uncompressed local copies after compression.  

## Getting Started 🚀

### Prerequisites ✅

*   Go (version 1.18 or higher recommended)  
*   SSH client installed on the machine running Backupper (for remote operations) 🔑  
*   zstd installed on the machine running Backupper  

### Build the Application 🔨

To build the Backupper executable, navigate to the project's root directory and run:  

```bash
make build
```

Or you can build manually

```bash
go build -o backupper cmd/backupper/main.go
```

## Usage Configuration ⚙️

Backupper relies on a `config.yaml` file to define its behavior and backup tasks. By default, it looks for `./config.yaml`. You can specify a different path using the `--config` flag.  

## Example config.yaml: 📄

```yaml
backupper:
  store_dir: ./store/                 # Path to store backups. This will be converted to an absolute path.
  logger:
    enabled: true                     # Enable or disable logging
    dir: ./logs                       # Directory to store log files
    format: json                      # json | text - format of logs
    level: info                       # debug | info | warn | error - verbosity level of logs
  tasks:
    - name: database_backup        # Descriptive name for the task
      user: root                      # SSH user for connecting to the remote server
      address: 192.168.1.100          # IP address or hostname of the remote server
      job:
        # Example for backing up a PostgreSQL database inside a Docker container
        container_name: container_db # Name of the Docker container
        use_command:                  # Optional: custom command to execute BEFORE the backup (e.g., dump specific table)
        work_dir: /var/backups        # Remote directory where the backup file will be generated
      postgresql:                     # Optional: Database specific settings if backing up PostgreSQL
        host: localhost               # PostgreSQL host (relative to the remote server)
        port: 5432                    # PostgreSQL port
        user: dbuser                  # PostgreSQL username
        password: _db_password     # PostgreSQL password
        database: app_db           # PostgreSQL database name
    - name: another_server_files      # Another example task
      user: backup_user               # SSH user for connecting to the remote server
      address: another.server      # IP address or hostname of the remote server
      job:
        # Example for backing up general files/directories
        container_name: container_media  # Name of the Docker container
        work_dir: /home/backup_user/data # The directory on the remote server to be backed up
        media_path: ./media/             # Path to the media directory
        # media_path and work_dir must specified, it will be copied.
        # No container_name or postgresql section implies a direct file/directory backup.
```

Additional info (all parameters) can be found in `config.example.yaml` ℹ️  

## Running Backupper ▶️

Once you have built the executable and configured your `config.yaml` (or used the example config), you can run Backupper:  
```bash
./backupper
```

If your `config.yaml` is not in the default location (`./config.yaml`):  
```bash
./backupper --config /path/to/your/config.yaml
```

Backupper will iterate through all defined tasks in your config.yaml, execute them, copy files, compress them, and log the results. 🔄📂  

## Output and Logging 📋  

Backupper provides informative output to log files (if enabled).  

Console Output Example:  
```log
👋 Hello from backupper's logger  
✅ in 5.34s | Completed task 'database_backup': 📦 archive name 'my_postgres_db_20231027103000.sql.zst' size: 10.5 MB ➡️  2.1 MB  
✅ in 2.12s | Completed task 'another_server_files': 📦 archive name 'data_archive_20231027103015.tar.zst' size: 500.0 KB ➡️  120.5 KB  
```  

## Zstandard (ZSTD) Compression 🏗️  
All backup files copied to the local `store_dir` are compressed using the Zstandard algorithm, resulting in files with a `.zst` extension (e.g., `db_backup.sql.zst`). Zstandard is known for its high compression ratios and extremely fast decompression speeds, making it ideal for backup archives. ⚡📉  

After a file is copied locally and successfully compressed, the original uncompressed local copy is automatically removed to save disk space. 🧹🗑️