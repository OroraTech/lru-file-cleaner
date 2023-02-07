# lru-file-cleaner
Delete files and directories in least recently used order until size quota is reached. Only files and directories at the top-level of the specified directory are considered. For directories, the cumulative size of all contents is considered. 

## Usage

```bash
lru-file-cleaner -dirpath ./foo -quota 1000
```

This will delete files and directories from the directory `./foo`, in least recently used order, until the quota of 1000 bytes is reached.