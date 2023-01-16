go version
export WORKSPACEPATH=$(cd $(dirname $0);pwd)
export CHEETAHPATH=/home/cheeath
export SERVERPATH=$WORKSPACEPATH/server
export WEBPATH=$WORKSPACEPATH/web
echo $CHEETAHPATH
echo $SERVERPATH
echo $WEBPATH
cd $SERVERPATH
export GO111MODULE=on
export GOPROXY=https://goproxy.io

# go mod tidy 根据实际情况，如果没有更新库可以不用执行
go mod tidy
go build -o cheetah main.go
# 删除cheetah进程，首次部署可以不用停止
{
	killall cheetah
} || {
  echo '没有cheetah进程'
}

{
  rm -rf $CHEETAHPATH/cheetah
}||{
  echo 'cheetah可执行文件不存在'
}

{
  mkdir $CHEETAHPATH
} || {
  echo '$CHEETAHPATH 已存在'
}

{
  mkdir $CHEETAHPATH/resource
} || {
  echo '$CHEETAHPATH/resource 已存在'
}
cp $SERVERPATH/cheetah $CHEETAHPATH/cheetah
cp $WORKSPACEPATH/config.production.yaml $CHEETAHPATH/config.yaml
cp -r $SERVERPATH/resource/* $CHEETAHPATH/resource/
cd $CHEETAHPATH
# nohup ./cheetah &
chmod u+x cheetah
nohup ./cheetah > cheetah.log 2>&1 &
sleep 10






