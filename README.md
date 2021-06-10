# Useful stuff #

Crontab (as root)  
`30 7 * * 0,2,4,6 cp -r /home/<user>/minecraft_server/server_files/ /mnt/<mount_name>/mc_backups/$(date +"%Y-%m-%d") && rm /home/<user>/minecraft_server/server_files/logs/*.log.gz`

Mount other hdd on startup (edit `/etc/fstab` and add this)  
`/dev/disk/by-uuid/<DISK_UUID> /mnt/<mount_name>/ <etx4|exfat|other_format> defaults 0 0`

Find hdd UUID (not PARTUUID)  
`sudo blkid`

Mount the hdd now  
`sudo mount /mnt/<mount_name>`