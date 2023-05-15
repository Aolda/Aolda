echo ---------------Install_IPFS---------------
wget https://dist.ipfs.tech/kubo/v0.20.0/kubo_v0.20.0_linux-amd64.tar.gz
tar -xvzf kubo_v0.20.0_linux-amd64.tar.gz
cd kubo
sudo bash install.sh
ipfs --version
ipfs init
sudo ipfs daemo &
echo ---------------Install_Go---------------
sudo apt-get update  
sudo apt-get -y upgrade  
wget  https://go.dev/dl/go1.19.2.linux-amd64.tar.gz 
sudo tar -xvf go1.19.2.linux-amd64.tar.gz   
sudo mv go /usr/local  
export GOROOT=/usr/local/go 
export GOPATH=$HOME/Projects/Proj1 
export PATH=$GOPATH/bin:$GOROOT/bin:$PATH 
sudo sysctl -w net.core.rmem_max=2097152 #통신 buffer 길이 늘려주기
go version
echo ---------------Done---------------




