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

APP_NAME=yangfan

# go mod tidy 根据实际情况，如果没有更新库可以不用执行
go mod tidy
go build -o yangfan main.go
# 删除cheetah进程，首次部署可以不用停止

is_exist(){
  echo "获取pid"
  pid=`ps -ef|grep $APP_NAME|grep -v grep|awk '{print $2}'`
  echo "yangfan pid is ${pid}"
  #如果不存在返回1，存在返回0     
  if [ -z "${pid}" ]; then
   return 1
  else
    return 0
  fi
}

stop(){
  echo "停止脚本"
  is_exist
  if [ $? -eq "0" ]; then
	{
	  echo "正在停止"
	  kill -9 $pid
	}||{
	  echo '停止出现异常'
	}
  else
    echo "${APP_NAME} is not running"
  fi  
  return 0
}
copy(){
	echo "复制文件"
	cp $SERVERPATH/yangfan $CHEETAHPATH/yangfan
	cp $WORKSPACEPATH/config.production.yaml $CHEETAHPATH/config.yaml
	cp -r $SERVERPATH/resource/* $CHEETAHPATH/resource/
	cd $CHEETAHPATH
}

set_path(){
	{
	  echo "执行 rm -rf ${CHEETAHPATH}/yangfan"
	  rm -rf $CHEETAHPATH/yangfan
	}||{
	  echo 'cheetah可执行文件不存在'
	}

	{
	  echo "执行 mkdir ${CHEETAHPATH}"
	  mkdir $CHEETAHPATH
	} || {
	  echo '$CHEETAHPATH 已存在'
	}

	{
	  mkdir $CHEETAHPATH/resource
	} || {
	  echo '$CHEETAHPATH/resource 已存在'
	}
}
run_server(){
	echo "运行"
	chmod u+x yangfan
	nohup ./yangfan > yangfan.log 2>&1 &
	sleep 10
}

start(){
	set_path
	copy
	run_server
}



restart(){
	stop
	echo "准备运行"
	start
}
restart




