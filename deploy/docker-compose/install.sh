#!/bin/bash

command_exists() {
	command -v "$1" 2>&1
}

space_left() {
    dir="$1"
    while [ ! -d "$dir" ]; do
        dir=`dirname "$dir"`;
    done
    echo `df -h "$dir" --output='avail' | tail -n 1`
}

start_docker() {
    systemctl start docker && systemctl enable docker
}

confirm() {
    echo -e -n "\033[34m[扬帆测试平台] $* \033[1;36m(Y/n)\033[0m"
    read -n 1 -s opt

    [[ "$opt" == $'\n' ]] || echo

    case "$opt" in
        'y' | 'Y' ) return 0;;
        'n' | 'N' ) return 1;;
        *) confirm "$1";;
    esac
}

info() {
    echo -e "\033[37m[扬帆测试平台] $*\033[0m"
}

warning() {
    echo -e "\033[33m[扬帆测试平台] $*\033[0m"
}

abort() {
    echo -e "\033[31m[扬帆测试平台] $*\033[0m"
    exit 1
}

trap 'onexit' INT
onexit() {
    echo
    abort "用户手动结束安装"
}


yangfan_path='/data/yangfan'
ENV_TAG='v1.3.1'
export ENV_TAG

if [ -z "$BASH" ]; then
    abort "请用 bash 执行本脚本"
fi

if [ ! -t 0 ]; then
    abort "STDIN 不是标准的输入设备"
fi

if [ "$#" -ne "0" ]; then
    abort "当前脚本无需任何参数"
fi

if [ "$EUID" -ne "0" ]; then
    abort "请以 root 权限运行"
fi
info "脚本调用方式确认正常"

if [ -z `command_exists docker` ]; then
    warning "缺少 Docker 环境"
    if confirm "是否需要自动安装 Docker"; then
        curl -sSLk https://get.docker.com/ | bash
        if [ $? -ne "0" ]; then
            abort "Docker 安装失败"
        fi
        info "Docker 安装完成"
    else
        abort "中止安装"
    fi
fi
info "发现 Docker 环境: '`command -v docker`'"

start_docker
docker version > /dev/null 2>&1
if [ $? -ne "0" ]; then
    abort "Docker 服务工作异常"
fi
info "Docker 工作状态正常"

compose_command="docker compose"
if $compose_command version; then
    info "发现 Docker Compose Plugin"
else
    warning "未发现 Docker Compose Plugin"
    compose_command="docker-compose"
    if [ -z `command_exists "docker-compose"` ]; then
        warning "未发现 docker-compose 组件"
        if confirm "是否需要自动安装 Docker Compose Plugin"; then
            curl -sSLk https://get.docker.com/ | bash
            if [ $? -ne "0" ]; then
                abort "Docker Compose Plugin 安装失败"
            fi
            info "Docker Compose Plugin 安装完成"
            compose_command="docker compose"
        else
            abort "中止安装"
        fi
    else
        info "发现 docker-compose 组件: '`command -v docker-compose`'"
    fi
fi
current_dir=$(pwd)
env_file="${current_dir}/.env"
while true; do

    if [ -e ".env" ]; then
        echo ".env 文件存在"

        # 读取 .env 文件
        source "$env_file"

        # 判断 YANGFAN_DIR 变量是否存在并非空
        if [ -n "$YANGFAN_DIR" ]; then
            info "YANGFAN_DIR 存在并已设置为: $YANGFAN_DIR"
            yangfan_path=$YANGFAN_DIR
            info "安装目录确认完成: '$yangfan_path'"
        else
            echo -e -n "\033[34m[扬帆测试平台] 安装目录 (留空则为 '$yangfan_path'): \033[0m"
            read input_path
            [[ -z "$input_path" ]] && input_path=$yangfan_path

            if [[ ! $input_path == /* ]]; then
                warning "'$input_path' 不是合法的绝对路径"
                continue
            fi
            yangfan_path=$input_path
        fi
    else
        echo -e -n "\033[34m[扬帆测试平台] 安装目录 (留空则为 '$yangfan_path'): \033[0m"
        read input_path
        [[ -z "$input_path" ]] && input_path=$yangfan_path

        if [[ ! $input_path == /* ]]; then
            warning "'$input_path' 不是合法的绝对路径"
            continue
        fi
        yangfan_path=$input_path
    fi

    if confirm "目录 '$yangfan_path' 当前剩余存储空间为 `space_left \"$yangfan_path\"` , 至少需要 5G, 是否确定"; then
        break
    fi
done

info "再次确认安装目录确认完成: '$yangfan_path'"
mkdir -p "$yangfan_path"
if [ $? -ne "0" ]; then
    abort "创建安装目录 '$yangfan_path' 失败"
fi
info "创建安装目录 '$yangfan_path' 成功"
cd "$yangfan_path"

mkdir -p ./config
mkdir -p ./mysql/data
mkdir -p ./data/grafana/provisioning/dashboards
mkdir -p ./data/grafana/provisioning/datasources
mkdir -p ./data/grafana/provisioning/yangfan-json



# 下载 compose.yaml 脚本
wget "http://docs.yangfan.gd.cn/install/compose.yaml" --no-check-certificate -O compose.yaml
if [ $? -ne "0" ]; then
    abort "下载 compose.yaml 脚本失败"
else
    info "下载 compose.yaml 脚本成功"
fi

# 下载 docker.config.yaml 文件
config_file="./config/docker.config.yaml"
if [ ! -f "$config_file" ]; then
    wget "http://docs.yangfan.gd.cn/install/docker.config.yaml" --no-check-certificate -O "$config_file"

    if [ $? -ne 0 ]; then
        abort "下载 docker.config.yaml 文件失败"
    else
        info "下载 docker.config.yaml 文件成功"
    fi
else
    info "docker.config.yaml 文件已存在，跳过下载"
fi


# 下载 my.conf 文件
config_file="./config/my.conf"

if [ ! -f "$config_file" ]; then
    wget "http://docs.yangfan.gd.cn/install/my.conf" --no-check-certificate -O "$config_file"
    if [ $? -ne 0 ]; then
        abort "下载 my.conf 文件失败"
    else
        info "下载 my.conf 文件成功"
    fi
else
    info "my.conf 文件已存在，跳过下载"
fi


# 下载 prometheus.yml 文件
wget "http://docs.yangfan.gd.cn/install/prometheus.yml" --no-check-certificate -O ./config/prometheus.yml
if [ $? -ne "0" ]; then
    abort "下载 prometheus.yml 文件失败"
else
    info "下载 prometheus.yml 文件成功"
fi

# 下载 yangfan.sql 文件
wget "http://docs.yangfan.gd.cn/install/yangfan.sql" --no-check-certificate -O ./mysql/yangfan.sql
if [ $? -ne "0" ]; then
    abort "下载 yangfan.sql 文件失败"
else
    info "下载 yangfan.sql 文件成功"
fi

# 下载 grafana.ini 文件
wget "http://docs.yangfan.gd.cn/install/grafana.ini" --no-check-certificate -O ./data/grafana/provisioning/grafana.ini
if [ $? -ne "0" ]; then
    abort "下载 grafana.ini 文件失败"
else
    info "下载 grafana.ini 文件成功"
fi

# 下载 yangfan.grafana.dashboard.yaml 文件
wget "http://docs.yangfan.gd.cn/install/yangfan.grafana.dashboard.yaml" --no-check-certificate -O ./data/grafana/provisioning/dashboards/yangfan.grafana.dashboard.yaml
if [ $? -ne "0" ]; then
    abort "下载 yangfan.grafana.dashboard.yaml 文件失败"
else
    info "下载 yangfan.grafana.dashboard.yaml 文件成功"
fi

# 下载 yangfan.grafana.prometheus.yaml 文件
wget "http://docs.yangfan.gd.cn/install/yangfan.grafana.prometheus.yaml" --no-check-certificate -O ./data/grafana/provisioning/datasources/yangfan.grafana.prometheus.yaml
if [ $? -ne "0" ]; then
    abort "下载 yangfan.grafana.prometheus.yaml 文件失败"
else
    info "下载 yangfan.grafana.prometheus.yaml 文件成功"
fi

# 下载 yangfan-for-distributed-load-testing.json 文件
wget "http://docs.yangfan.gd.cn/install/yangfan-for-distributed-load-testing.json" --no-check-certificate -O ./data/grafana/provisioning/yangfan-json/yangfan-for-distributed-load-testing.json
if [ $? -ne "0" ]; then
    abort "下载 yangfan-for-distributed-load-testing.json 文件失败"
else
    info "下载 yangfan-for-distributed-load-testing.json 文件成功"
fi

# 下载 yangfan-for-node-status.json 文件
wget "http://docs.yangfan.gd.cn/install/yangfan-for-node-status.json" --no-check-certificate -O ./data/grafana/provisioning/yangfan-json/yangfan-for-node-status.json
if [ $? -ne "0" ]; then
    abort "下载 yangfan-for-node-status.json 文件失败"
else
    info "下载 yangfan-for-node-status.json 文件成功"
fi

touch "$env_file"
if [ $? -ne "0" ]; then
    abort "创建 .env 脚本失败"
fi
info "创建 .env 脚本成功"

echo "YANGFAN_DIR=$yangfan_path" > "$env_file"
echo "YANGFAN_TAG=$ENV_TAG" > "$env_file"


info "即将开始下载 Docker 镜像"
$compose_command pull
if [ $? -ne "0" ]; then
    abort "镜像下载失败"
fi

info "即将运行 Docker 镜像"
$compose_command up -d

if [ $? -ne "0" ]; then
    abort "启动 Docker 容器失败"
fi

info "当前安装版本：$ENV_TAG"


