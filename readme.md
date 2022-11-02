# backend-app

### Setup Postgressql LInux Debian

```sh
sudo apt update
sudo apt install postgresql
sudo apt install postgresql-client

# Setting up new password
sudo passwd postgres

user:~$ sudo -i -u postgres
postgres@user:~$ psql
postgres=# ALTER USER postgres PASSWORD 'mynewpassword';
```
Make it open to public.

```sh
sudo nano /etc/postgresql/9.3/main/postgresql.conf
# edit: listen_addresses = '*'
# save
sudo nano /etc/postgresql/9.3/main/pg_hba.conf
# add: host    all             all             192.168.7.0/24          md5
# save
sudo systemctl restart postgresql
# Check from another device using nmap
nmap <ip-adress> -p5432
```
Make new database
```sh
user:~$ sudo -i -u postgres
postgres@user:~$ psql
postgres=# create database go_movies;
postgres=# \l
```
Import go_movies.sql database.
```sh
postgres@user:~$ psql -d go_movies -f go_movies.sql
```
