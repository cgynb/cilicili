# cilicili

## 普通部署

> os: ubuntu

### 配置redis
```shell
sudo apt install redis-server
```
### 配置mysql
用的是腾讯云的mysql，所以就没有配置
如果想在本机跑，可以再`config.toml`中进行配置

### 运行项目
1. 创建项目文件夹
```shell
cd
mkdir cilicili
```
2. 将二进制文件main，配置文件config.toml传到服务器上，放入项目文件夹(cilicili)中
3. 跑起来
```shell
chmod +x main
./main > l.log
```

## docker 部署
之后看下 :-)
