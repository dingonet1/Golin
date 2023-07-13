
# 使用说明
> 此版本初衷是只做CLI版本,压缩体积小、跨平台可用，随着功能越多、参数越多使用人记不住参数，增加了基于Python的简易Gui版本，可满足的等保模式的增加资产、运行不同模式，运行Cli指令，增加了基于Web的多主机、单主机模式以及模拟定级功能。注：完整使用功能建议使用Cli版本。（如弱口令扫描、端口扫描、目录扫描、子域名扫描、syslog...）

## web功能预览
![web](images/web.gif)

## 暴力破解功能预览
![crack](images/crack.gif)

## 端口扫描功能预览
![port](images/port.jpg)

## 子域名扫描功能预览
![domain功能](images/domain.gif)

## 目录扫描功能预览
![dirsearch功能](images/dirsearch.gif)

## Windows模式预览
![Windows](images/windows.png)


## 自动化测评现阶段支持类型
| 序号 |      类型       | 是否支持 |     备注     |
|:--:|:-------------:|:----:|:----------:|
| 1  |    Centos     |  √   | SSH远程或本地运行 |
| 2  |    Windows    |  √   |    本地运行    |
| 3  |     Redis     |  √   |  远程或本地运行   |
| 4  |  PostgreSQL   |  √   |  远程或本地运行   |
| 5  |    Oracle     |  √   |  远程或本地运行   |
| 6  |     MSSQL     |  √   |  远程或本地运行   |
| 7  |      H3C      |  √   |   SSH远程    |
| 8  |      华为       |  √   |   SSH远程    |
| 9  |      AIX      |  √   | 可自定义命令未内置  |
| 10 |    Ubuntu     |  √   | 可自定义命令未内置  |
| 11 |    MongoDB    |      |            |
| 12 | Elasticsearch |      |            |

## 弱口令现阶段支持类型
| 序号 |      类型       | 是否支持 |       备注        |
|:--:|:-------------:|:----:|:---------------:|
| 1  |      SSH      |  √   |     支持协议识别      |
| 2  |      RDP      |  x   | 已移除,建议此需求使用Goby |
| 3  |      FTP      |  √   |     支持协议识别      |
| 4  |     MySQL     |  √   |                 |
| 5  |  PostgreSQL   |  √   |                 |
| 6  |     Redis     |  √   |                 |
| 7  |     MSSQL     |  √   |                 |
| 8  |      SMB      |  √   |    超时参数暂时不可用    |
| 9  |    Telnet     |  √   |                 |
| 10 |    Tomcat     |  √   |                 |
| 11 |     xlsx      |  √   |    仅允许xlsx格式    |
| 12 |     SNMP      |      |                 |
| 13 |    MangoDB    |      |                 |
| 14 | Elasticsearch |      |                 |

## 端口扫描现阶段支持功能
| 序号 |    功能     | 是否支持 |                       备注                        |
|:--:|:---------:|:----:|:-----------------------------------------------:|
| 1  |    多线程    |  √   |               默认为30并发,可通过-c指定并发数                |
| 2  |   指定端口    |  √   |                 格式支持1,2,3,2-20                  |
| 3  |   指定IP    |  √   |   格式支持192.168.1.1,192.168.1.1/24,192.168.1-10   |
| 4  |   排除端口    |  √   |                   格式支持:1,2,3                    |
| 5  | 扫描前探测主机存活 |  √   |            基于ping,可通过--noping跳过探测存活             |
| 6  |  打乱主机顺序   |  √   |              默认不打乱,可通过--random进行打乱              |
| 7  |   协议识别    |  √   | 目前支持常见协议:ssh、redis、https、https、MySQL、pgsql、ftp等 |
| 8  |   超时时间    |  √   |                  默认3秒,可通过-t指定                   |
| 9  |   识别web   |  √   |               目前支持识别server、title                |
| 10 |   结果保存    |  √   |         默认不保存,可通过--save保存到portscan.xlsx         |
| 11 | 主机操作系统识别  |  √   |                      基于ttl                      |



## web目录扫描现阶段支持功能
| 序号 |      功能       | 是否支持 |       备注        |
|:--:|:-------------:|:----:|:---------------:|
| 1  |      多线程      |  √   |     默认为30并发     |
| 2  |    自定义状态码     |  √   |     默认为200      |
| 3  |     代理模式      |  √   |  http/s、socks   |
| 4  |    返回title    |  √   |                 |
| 5  |    超时等待时常     |  √   |      默认为3秒      |
| 6  |     循环等待      |  √   |     默认为无限制      |
| 7  |     内置url     |      |                 |
| 8  | 自定义User-Agent |  √   |                 |
| 9  |      重传       |      |                 |
| 10 |      爬虫       |      |                 |
| 11 |     结果保存      |  √   | 保存到dirScan.json |

## 子域名扫描现阶段支持功能
| 序号 |     功能     | 是否支持 |              备注              |
|:--:|:----------:|:----:|:----------------------------:|
| 1  |    多线程     |  √   |           默认为30并发            |
| 2  |   内置常用域名   |  √   |             100+             |
| 3  |   指定字典文件   |  √   |        可基于字典文件实现多层扫描         |
| 4  |   fofa调用   |  √   | 需在环境变量设置 FOFA_EMAIL、FOFA_KEY |
| 5  | RapidDNS调用 |  √   |                              |


# 常用启动参数
```
golin web (通过web方式启动,仅支持等保功能)
golin port -i 192.168.1.1/24 (扫描c段端口)
golin dirsearch -u https://test.com -f 字典.txt --code 200,406 (扫描状态码为200以及300的web目录)
golin crack ssh -i 192.168.1.1 (基于内置鉴别信息进行扫描弱口令)
golin domain -u baidu.com --api (扫描子域名,并且调用fofa、RapidDNS的API)
golin update (检查是否可更新)
golin windows (运行采集当前主机的安全配置生成报告,需管理员身份运行)
golin [linux、mysql、oracle、sqlserver、redis...] (按照3级等保要求核查各项安全配置)
```

# 免责声明
本工具仅面向合法授权的企业安全建设行为，如您需要测试本工具的可用性，请自行搭建靶机环境。

在使用本工具进行检测时，您应确保该行为符合当地的法律法规，并且已经取得了足够的授权。请勿对非授权目标进行扫描。

如您在使用本工具的过程中存在任何非法行为，您需自行承担相应一切后果。