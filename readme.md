## 使用说明
margs 是使用go templates 来渲染输入流，并将结果输出。使用前需要了解go templates如何使用。

并且在原来的go templates上，增加了 {{eol}} {{cut slice start end}} {{join slice sep}}三个内置函数
## 安装
```
sudo cp margs /usr/bin/
```
## 快速使用
### 处理json
假设有json文件如下：
```
{
  "name": "yzh",
  "age": 18,
  "attrs": [
    "hello",
    "world",
    "1",
    "2",
    "3"
  ]
}
```
####打印name字段：
```
cat test.json|margs -j -t '{{.name}}'
# 输出：yzh
```
####循环打印attrs里面的内容
```
# {{eol}} 为内置函数，可以输出一个换行符\n
cat test.json |margs -j -t '{{range .attrs}}{{.}}{{eol}}{{end}}'
```
输出：
```
hello
world
1
2
3
```

####循环打印attrs中第2-4项
```
# {{cut slice start end}} cut为内置函数，可以取出sub_slice
cat test.json |./margs -j -t '{{range (cut .attrs 1 4)}}{{.}}{{eol}}{{end}}'
```
输出：
```
world
1
2
```
####打印attrs中第2-4项,并用“-”连接结果
```
# {{join arr seperator}} join是内置函数，类似strings.Join
cat test.json |./margs -j -t '{{join (cut .attrs 1 4) "-"}}'
```
输出：
```
world-1-2
```

### 处理输入流行
当不带-j参数，则按行读取输入流进行处理
假设有文件如下：
```
mysql-web-5d8d9894d9-hvfs9            1/1     Running     0          3h52m
redis-web-79fd749854-65gg5            1/1     Running     0          11d
kafka-web-78f68f95b4-r4wdh            2/2     Running     0          3h40m
```

#### 打印第一列
```
# -s 指定一到多个分割符，此处为对每行按空格分割，每行数据输出为一个分割后的slice
cat test.txt|margs -s " " -t '{{index . 0}}{{eol}}'
```
输出：
```
mysql-web-5d8d9894d9-hvfs9
redis-web-79fd749854-65gg5
kafka-web-78f68f95b4-r4wdh
```

#### 使用正则提取出第一个单词
```
# -r 指定一个正则表达式，使用group分割。不能与-s同时使用
cat test.txt|./margs -r '([a-zA-Z0-9]*)-.*' -t '{{index . 1}}{{eol}}'
```
输出：
```
mysql
redis
kafka
```
