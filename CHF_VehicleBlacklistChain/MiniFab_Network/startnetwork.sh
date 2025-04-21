
#!/bin/sh

echo "Starting MiniFab Network"
minifab netup -s couchdb -e true -i 2.4.8 -o owner.auto.com

sleep 5

echo "Create the Channel"
minifab create -c autochannel

sleep 2

echo "Join the Channel"
minifab join -c autochannel

sleep 2

echo "anchor update"
minifab anchorupdate