# CI/CD

## 1. 安装配置linux

### 1.1 安装jdk8
```shell
tar -zxvf [jdk包名.tar.gz] -C /usr/local
```

### 1.2 安装maven
```shell
tar -zxvf [maven包名.tar.gz] -C /usr/local
```

### 1.3 改文件名
```shell
cd /usr/local
mv [jdk目录名] jdk
mv [maven目录名] maven
```

### 1.4 修改settings.xml，配置maven私服，配置jdk编译插件
```shell
cd maven/conf   
```
+ 找到mirror标签，更换为阿里云maven镜像
+ 找到</profiles>标签，在它上面加入 jdk1.8编译插件，改名jdk8 
+ 找到<activeProfiles>标签，复制粘贴到注释外面，保留一个activeProfile，改名jdk8

### 1.5 安装docker

### 1.6 启动docker服务，设置开机启动
```shell
systemctl start docker
systemctl enable docker
```

### 1.7 安装docker-compose

### 1.8 