sudo apt install golang git
GO111MODULE=on
mkdir collectData
cd collectData
git clone https://github.com/osadsanu/piCollectData.git . 
echo "\t please run:    go to collectData directory: `cd collectData`"
echo "\t please run:    go run main.go -o serverCollect"
echo "\t please execute ./serverCollect"