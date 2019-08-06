wget https://ftp.postgresql.org/pub/source/v11.3/postgresql-11.3.tar.gz
tar -xzvf
apt install libreadline-dev
./configure --prefix=/usr/local/postgresql

useradd postgres

sudo passwd postgres

su postgres

mkdir /home/postgre/data

cd /home/jyb/

chown -R postgres.postgres postgre

./initdb -E UTF-8 -D /home/postgre/data --locale=en_US.UTF-8 -U postgres -W

./pg_ctl -D /home/postgre/data start

vim /home/postgre/data/pg_hba.conf

host    all     all     0.0.0.0/0                       md5

vim /home/postgre/data/postgresql.conf

listen_addresses = '*'

./pg_ctl -D /home/postgre/data/ reload

./psql

ALTER USER postgres WITH PASSWORD '123456';



sudo  passwd -d postgres

sudo -u postgres passwd
