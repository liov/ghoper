tar -xzvf openresty-VERSION.tar.gz
cd openresty-VERSION/
./configure
make
sudo make install DESTDIR=/whereto make install