# DevOps指南

采用Jenkins Pipeline作为自动化CI/CD的主要工具，
集成Harbor，SonarQube，golangci-lint，eslint，Checkmarx等工具。
使用Github Webhooks触发自动化流程，Github App与Github Action同时做PR的代码检查。
放弃K8s部署方案。

## 1. 准备工作

### 1.1 VM配置要求
+ CPU：2核
+ 内存：8G
+ 硬盘空间：30G
+ 操作系统：Centos7.9
```shell
# RSA密钥生成命令，在.ssh文件夹里
ssh-keygen -t rsa -b 4096
```

### 1.2 必备工具 - jdk，node，go
```shell
# 安装包位置
cd ~
mkdir pkg
cd pkg

# 解压安装包
# jdk8u333
tar -zxvf jdk-8u333-linux-aarch64.tar.gz -C /home/opc/tools
# node-v16
tar xf node-v16.17.0-linux-x64.tar.xz -C /home/opc/tools
# go1.17
tar -zxvf go1.17.5.linux-amd64.tar.gz -C /home/opc/tools

# 修改文件名 （go不用改）
cd /home/opc/tools
mv jdk1.8.0_333 jdk
mv node-v16.17.0-linux-x64 node

# 修改环境变量
sudo vi /etc/profile
# 找到export，在PATH变量上面一行加入
export NODE_HOME=/home/opc/tools/node
export GO_HOME=/home/opc/tools/go
# 整合PATH行
export PATH=$PATH:$NODE_HOME/bin:$GO_HOME/bin
# 载入环境变量
source /etc/profile
```

### 1.3 必备工具 - docker，docker-compose
```shell
# docker
# 下载安装 要等好久
curl -fsSL https://get.docker.com | bash -s docker --mirror Aliyun
# 启动 重启用restart
sudo systemctl start docker
# 开机启动
sudo systemctl enable docker
# 检查安装结果
docker --version

# docker-compose
# 下载安装
sudo curl -L https://github.com/docker/compose/releases/download/1.23.2/docker-compose-`uname -s`-`uname -m` -o /usr/bin/docker-compose
# 增加执行权限
sudo chmod +x /usr/bin/docker-compose
# 检查安装结果
docker-compose --version
```

### 1.4 必备工具 - Jenkins
```shell
# Jenkins
# 拉取镜像（所用版本为下面）
sudo docker pull jenkins/jenkins
sudo docker pull jenkins/jenkins:2.319.3-lts
# 创建docker-compose.yaml路径
cd
mkdir docker
cd docker
mkdir jenkins_docker
cd jenkins_docker
# 拷贝配置好的docker-compose.yaml（配置见下面yaml）
# 启动Jenkins
cd /home/ops/docker/jenkins_docker
sudo docker-compose up -d
# 赋予宿主机jenkins data目录权限
sudo chmod -R 777 data
# 重启Jenkins容器
sudo docker-compose restart
# 将node挂载入jenkins容器
cd /home/opc/tools
cp -r node /home/opc/docker/jenkins_docker/data
```

添加daemon.json，为了后面Harbor登录追加仓库地址信息
```shell
sudo vi /etc/docker/daemon.json
```
```json
{
  "insecure-registries": ["168.138.30.224:9002"]
}
```
重启docker
```shell
sudo systemctl daemon-reload
sudo systemctl restart docker
```

Jenkins配置文件docker-compose.yaml
```yaml
version: '3.1'
services:
  jenkins:
    image: 'jenkins/jenkins:2.319.3-lts'
    container_name: jenkins
    privileged: true
    restart: always
    user: root
    ports:
      - '9000:8080'
      - '50000:50000'
    volumes:
      - './data:/var/jenkins_home'
      - '/usr/bin/docker:/usr/bin/docker'
      - '/var/run/docker.sock:/var/run/docker.sock'
      - '/etc/docker/daemon.json:/etc/docker/daemon.json'
```


## 2. 配置Jenkins

### 2.1 初始化Jenkins
```shell
# 查看jenkins日志
sudo docker logs -f jenkins
```
找到日志里的密码，浏览器中访问jenkins，输入密码。<br/>
密码在`Please use the following password to proceed to installation:`后面。两坨星星中间夹着。

选择手动安装插件（右边按钮）。默认选项即可，直接点击安装。<br/>
如果失败，进入面板后找到`Mange Jenkins`→`Manage Plugins`→`Available`，手动安装<br/>
如果总是不成功，就更改为国内的插件源。

设置用户名密码。根据提示操作，直到进入到Welcome界面。

### 2.2 安装重要插件
`Manage Jenkins` → `Manage Plugins` → `Available` → 

`Search` `Git Parameter` → 勾选 → 

`Search` `Publish Over SSH` → 勾选 →

`Search` `NodeJS` → 勾选 → 

点击`Install without restart`

### 2.3 挂载jdk
```shell
mv /home/opc/tools/jdk /home/opc/docker/jenkins_docker/data
```
回到浏览器jenkins的`Dashboard` → `Global Tool Configuration` → `JDK` → `Add JDK`，取消自动安装，按下图填写。

![挂载jdk.png](devops/挂载jdk.png)

### 2.4 挂载nodejs
```shell
cp -r /home/opc/tools/node /home/opc/docker/jenkins_docker/data
```
回到浏览器jenkins的`Dashboard` → `Global Tool Configuration` → `NodeJS` → `NodeJS`，取消自动安装，按下图填写。

![挂载nodejs.png](devops/挂载nodejs.png)

### 2.5 连接代码仓库
`Dashboard` → `Configure System`，找到`Publish over SSH`，点击`Add`

![连接代码仓库.png](devops/连接代码仓库.png)

这个username是远程服务器的username，密码和私钥也是连接那个服务器用的。

`Test Configuration`测试是否跑通。如果提示路径问题，注意mkdir。显示SUCCESS则保存。


## 3. 流水线Pipeline

### 3.1 创建流水线
`New Item` → 输入名称 → `Pipeline` → OK

![流水线脚本.png](devops/流水线脚本.png)

`Pipeline Syntax`可进入生成指定功能脚本的页面，比如checkout功能的脚本

点击`Generate Pipeline Script`，复制粘贴代码到Jenkinsfile中。

回到Pipeline配置区域 → 选择`Pipeline script from SCM` → 选择`Git`

![配置流水线拉取分支.png](devops/配置流水线拉取分支.png)
![配置流水线脚本地址.png](devops/配置流水线脚本地址.png)

> 注意：Script Path第一个字符不要"/"
> 注意：按上图中填写Branch Specifier，才会按照tag构建，如果填写的是*/main，即便选择按tag构建，也只会拉取main分支的代码进行构建。
> 不知为何，when { buildingTag() } 语法不起作用。可能是只在多分支流水线中才有用。
> 目前CI/CD的流水线是分离的。

### 3.2 拉取代码
来到项目的Configure，按图中点击

![参数化构建.png](devops/参数化构建.png)
![配置代码仓库信息.png](devops/配置代码仓库信息.png)

转到`Pipeline Syntax`生成语法

### 3.3 静态代码分析
SonarQube 8.9.9
```shell
# 拉取镜像
sudo docker pull postgres
sudo docker pull sonarqube:8.9.9-community
# 创建docker-compose.yaml并运行
cd /home/opc/docker
mkdir sonarqube_docker
cd sonarqube_docker
# 粘贴入此路径，然后执行命令
sudo docker-compose up -d
# 查看日志
sudo docker logs -f sonarqube
# 有两个ERROR，扩增虚拟内存
sudo vi /etc/sysctl.conf
# 在注释后面增加如下，并保存
vm.max_map_count=262144
# 执行生效
sudo sysctl -p
# 重启容器
sudo docker-compose up -d
# 查看日志，还有错也没办法
sudo docker logs -f sonarqube
```

```yaml
version: "3.1"
services:
  sonar_db:
    image: postgres
    container_name: sonar_db
    privileged: true
    restart: always
    user: root
    ports:
      - 5432:5432
    networks:
      - sonarnet
    environment:
      POSTGRES_USER: sonar
      POSTGRES_PASSWORD: sonar
  sonarqube:
    image: sonarqube:8.9.9-community
    container_name: sonarqube
    privileged: true
    restart: always
    user: root
    depends_on:
      - sonar_db
    ports:
      - 9001:9000
    networks:
      - sonarnet
    environment:
      SONAR_JDBC_URL: jdbc:postgresql://sonar_db:5432/sonar
      SONAR_JDBC_USERNAME: sonar
      SONAR_JDBC_PASSWORD: sonar
    volumes:
      - '/home/opc/tools/node/bin/node:/usr/bin/node'
      - '/home/opc/tools/node/bin/npm:/usr/bin/npm'
      - '/home/opc/tools/node/bin/npx:/usr/bin/npx'
      - '/usr/bin/docker:/usr/bin/docker'
      - '/var/run/docker.sock:/var/run/docker.sock'
networks:
  sonarnet:
    driver: bridge
```

在浏览器输入ip和预设端口号，sonarqube初始账户和密码都是`admin`

按提示操作，到能看到面板为止。

下载 `sonar-scanner-cli-4.6.0.2311-linux.zip`
> https://binaries.sonarsource.com/?prefix=Distribution/sonar-scanner-cli/

```shell
# 安装解压zip工具
sudo yum -y install unzip
# 解压
cd /home/opc/pkg
unzip sonar-scanner-cli-4.6.0.2311-linux.zip
mv sonar-scanner-4.6.0.2311-linux/ sonar-scanner
# 移动sonar-scanner到jenkins挂载卷目录
mv ~/pkg/sonar-scanner /home/opc/docker/jenkins_docker/data
# 配置sonar-scanner
cd /home/opc/docker/jenkins_docker/data/sonar-scanner/conf
vi sonar-scanner.properties
```
撤销注释并修改 `sonar.host.url=http://宿主机ip:9001`<br/>
撤销UTF-8注释

给sonar-scanner设置token，来到浏览器sonar面板右上角

![sonar-scanner设置token.png](devops/sonar-scanner设置token.png)

![sonar-scanner生成token.png](devops/sonar-scanner生成token.png)

```shell
# 使用扫描命令
# 运行一次流水线就会产生workspace文件夹
# cd /home/opc/docker/jenkins_docker/data/workspace
# cd main  
/home/opc/docker/jenkins_docker/data/sonar-scanner/bin/sonar-scanner \
    -Dsonar.sources=./ \
    -Dsonar.projectname=${JOB_NAME} \
    -Dsonar.login=8ed4baf9ba0204179095df34597c503c4b857b8a \
    -Dsonar.projectKey=${JOB_NAME} \
    -Dsonar.nodejs.executable=/usr/bin/node \
    -Dsonar.inclusions=src/**/*.js,src/**/*.jsx \
    -Dsonar.coverage.exclusions=node_modules/**/*,server/build/**/*,config/**/*,scripts/**/*,public/**/*,src/config/**/* \
    -Dsonar.qualitygate.wait=true
```

成功会看到

![sonar-scanner成功.png](devops/sonar-scanner成功.png)

<h3>SonarQube加入流水线</h3>

jenkins安装SonarQube Scanner插件

`Dashboard` → `Configure System`，找到`SonarQube servers`，点击`Add`，注意添加token

![img.png](devops/配置sonarqube-server.png)

![增加sonar的token.png](devops/增加sonar的token.png)

`Dashboard` → `Global Tool Configuration`，找到`SonarQube Scanner`，点击`Add`

![添加sonar-scanner全局工具.png](devops/添加sonar-scanner全局工具.png)

修改流水线对应位置命令，注意命令bin是在容器内部的地址
```shell
/home/opc/docker/jenkins_docker/data/sonar-scanner/bin/sonar-scanner \
    -Dsonar.sources=./ \
    -Dsonar.projectname=${JOB_NAME} \
    -Dsonar.login=8ed4baf9ba0204179095df34597c503c4b857b8a \
    -Dsonar.projectKey=${JOB_NAME} \
    -Dsonar.nodejs.executable=/usr/bin/node \
    -Dsonar.inclusions=src/**/*.js,src/**/*.jsx \
    -Dsonar.coverage.exclusions=node_modules/**/*,server/build/**/*,config/**/*,scripts/**/*,public/**/*,src/config/**/* \
    -Dsonar.qualitygate.wait=true
```

<h3>设置代码质量标准</h3>
![代码质量标准.png](devops/代码质量标准.png)

### 3.5 安全漏洞扫描
checkmarx, Tenable.io CS, Harness

### 3.6 构建镜像
镜像仓库Harbor-v2.4.3
```shell
# 安装Harbor
tar -zxvf harbor-offline-installer-v2.4.3.tgz -C /home/opc/tools
cd /home/opc/tools/harbor
# 配置harbor
cp harbor.yml.tmpl harbor.yml
vi harbor.yml
```
![配置harbor.png](devops/配置harbor.png)

初始帐号密码是：admin Harbor12345
```shell
# 创建harbor容器
cd /home/opc/tools/harbor
sudo ./install.sh
```

创建一个新项目

![Harbor新项目.png](devops/Harbor新项目.png)

在Jenkinsfile定义变量
```
environment {
    harborUsername = "admin"
    harborPassword = "Harbor12345"
    harborAddress = "10.0.0.167:9002"
    harborRepo = "repo"
}
```

在pipeline的stage里加入
```shell
docker login -u ${harborUsername} --password-stdin ${harborPassword} ${harborAddress}
docker tag ${JOB_NAME}:${tag} ${harborAddress}/${harborRepo}/${JOB_NAME}:${tag}
docker push ${harborAddress}/${harborRepo}/${JOB_NAME}:${tag}
```

编写deploy.sh放入/usr/bin

```shell
sudo chmod +x deploy.sh
```

使用pipeline代码生成，选择`ssh publisher`，输入要执行的shell代码，生成并复制到pipeline

```shell
sudo deploy.sh $harborAddress $harborRepo $JOB_NAME $tag $container_port $host_port
```
将生成后的shell脚本和上面相同的语句的单引号，改为双引号。

至此CD流程结束。之后如果运行时需要package的一些配置，可以增加stage。

<h3>设置pipeline结束后的操作</h3>
failed用error收尾

如果选择每个JOB有不同的保留个数，可以在config里选择`discard a build`，然后生成`buildDiscard`的语法。
并粘贴进pipeline，与environment并列。

为所有Job设置结束后保留N条流水线运行记录

`Dashboard` → `Configure System`，找到`Global Build Discarders`，点击`Add`

![配置保留任务.png](devops/配置保留任务.png)

## 4. 问题

> 注意对于一次性设置，例如export，每一个stage都要设定
> npm报错建议回退到npm@6

### 4.1 代理原因卡死，不报错
报错 `npm ERR! network This is a problem related to network connectivity.`
```shell
${npmHome}/bin/npm config set registry https://registry.npmjs.org/
```

### 4.2 内存原因卡死，查看monitor，未卡死时渐满，卡死时monitor调不出
```shell
# RAM被限制为512MB
${npmHome}/bin/npm install --max_old_space_size=512
```

### 4.3 连接SSH server，test connection不过
可能是RSA标题不对

使用下面标题，不要含有`OPENSSH`
```shell
-----BEGIN RSA PRIVATE KEY-----
```

生成命令是
```shell
ssh-keygen -t rsa -b 2048
```

### 4.4 Harbor报错HTTPS
报错`Error response from daemon: Get "https://xx.xx.xx.xx:xxxx/v2/": http: server gave HTTP response to HTTPS client`

检查pipeline的Harbor的address变量是否和daemon.json中的一致。

### 4.5 部署脚本运行不成功
```shell
sudo deploy.sh $harborAddress $harborRepo $JOB_NAME $tag $container_port $host_port
```
将生成后的shell脚本和上面相同的语句的单引号，改为双引号。

### 4.6 创建容器的脚本执行报错127
Exec exit status not zero. Status [127]

脚本不存在，或者脚本的依赖不存在，无法执行脚本。

+ 检查生成语句中.sh执行语句的引号是否为双引号
+ 检查生成语句中.sh执行语句是否需要sudo

### 4.7 在CI的多分支流水线中，无法发现最新版本
添加`Discover tags`

![ci发现标签.png](devops/ci发现标签.png)

### 4.8 sonar报错key不合法
`ERROR: Malformed key for Project: 'xxx/main'. Allowed characters are alphanumeric, '-', '_', '.' and ':', with at least one non-digit.`

多分支流水线会在JOB前面加入一个文件夹路径，来标记多分支流水线，所以会生成"/"，比如 xxx/main

修改pipeline里，sonar的stage里，projectkey字段。


## 5. PR触发CI

### 5.1 安装插件
`GitHub Pull Request Builder`

### 5.2 配置JOB
`Dashboard` → `Configure System`，找到`GitHub Pull Request Builder`

![配置githubPR插件.png](devops/配置githubPR插件.png)

### 5.3 配置webhook
github setting里，github pr builder插件的access token要加上hook相关的权限。

![webhook配置1.png](devops/webhook配置1.png)

secret在jenkins的user config里面生成API token

下面选上`pull request`

成功后显示`√`，见下图

![webhook配置2.png](devops/webhook配置2.png)

### 5.4 配置action
在项目中添加`.github/workflow/xxx.yaml`文件

这种方法是直接利用Github的Action来检查PR的代码

建议使用`GitHub App`方法

~~如果想要使用jenkins流水线检查代码，需要使用Github API~~

> ~~https://medium.com/@shandredreyer/github-using-jenkins-webhooks-and-status-to-prevent-a-pull-request-from-being-merged-443b43f36099~~

~~或者使用插件`CloudBees`~~

> ~~https://docs.cloudbees.com/docs/cloudbees-ci-kb/latest/client-and-managed-masters/github-customize-status-checks-for-pull-request~~

或者使用插件`GitHub App`

> https://www.youtube.com/watch?v=aDmeeVDrp0o&t=1100s

### 5.5 配置branch规则
在repo的settings里，选择branch，在规则里增加`Require status checks to pass before merging`，找到配置好的jenkins的job，
添加为Required

成功后会在PR界面显示

![githubChecks.png](devops/githubChecks.png)

至此，全套DevOps流程结束。
