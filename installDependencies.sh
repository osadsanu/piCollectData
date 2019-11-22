sudo apt install golang git
GO111MODULE=on
mkdir collectData
cd collectData
git clone https://github.com/osadsanu/piCollectData.git . 
echo "please run:    go run main.go -o serverCollect"
echo "please execute ./serverCollect"