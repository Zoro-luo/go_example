Docker -- 入门到实践  https://www.bookstack.cn/books/docker_practice-v1.1.0
安装：
	(1) Ubuntu/CentOS 安装 Docker CE:
		//在测试或开发环境中 Docker 官方为了简化安装流程，提供了一套便捷的安装脚本，Ubuntu 系统上可以使用这套脚本安装：
		$ curl -fsSL get.docker.com -o get-docker.sh
		$ sudo sh get-docker.sh --mirror Aliyun
		//启动 Docker CE
		sudo systemctl enable docker
		sudo systemctl start docker
		//建立 docker 用户组
		$ sudo groupadd docker
		//将当前用户加入 docker 组：
		$ sudo usermod -aG docker $USER
		//退出当前终端并重新登录，进行如下测试,测试 Docker 是否安装正确
		$ docker run hello-world
使用镜像：
	(1) 获取镜像：
		//仓库名：如之前所说，这里的仓库名是两段式名称，即 <用户名>/<软件名>。对于 Docker Hub，如果不给出用户名，则默认为 library，也就是官方镜像。
		docker pull ubuntu:18.04
		//运行
		/* --rm：这个参数是说容器退出后随之将其删除。默认情况下，为了排障需求，退出的容器并不会立即删除，除非手动 docker rm。
		 * 我们这里只是随便执行个命令，看看结果，不需要排障和保留结果，因此使用 --rm 可以避免浪费空间。
		 * ubuntu:18.04：这是指用 ubuntu:18.04 镜像为基础来启动容器。
		 * bash：放在镜像名后的是 命令，这里我们希望有个交互式 Shell，因此用的是 bash。
		 */
		$ docker run -it --rm \				
    	ubuntu:18.04 \
   		bash
   		// 执行 cat /etc/os-release，返回容器内是 Ubuntu 18.04.1 LTS 系统
   		cat /etc/os-release
   	(2) 退出容器：	通过 exit 退出这个容器
   	(3) 列出镜像:
   		docker image ls
   		//虚悬镜像
   		<none>               <none>              00285df0df87        5 days ago          342 MB
   		//显示虚悬镜像：
   		docker image ls -f dangling=true
   		//一般来说，虚悬镜像已经失去了存在的价值，是可以随意删除的，可以用下面的命令删除。
   		$ docker image prune
   		//列出部分镜像
   		docker image ls ubuntu
   		//显示在 mongo:3.2 之后建立的镜像
   		$ docker image ls -f since=mongo:3.2
   		//以特定格式显示
   		$ docker image ls -q
   	(4) 删除本地镜像:
   		docker image ls
   		$ docker image rm 501 			//短 ID 删除镜像
   		$ docker image rm centos 		//镜像名 删除镜像
