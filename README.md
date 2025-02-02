# mic-project-layout
自用 micro project layout

## 安装Gonew：

```shell
  go install golang.org/x/tools/cmd/gonew@latest
```

## 克隆模板

```shell
  gonew github.com/qmute/mic-project-layout your.domain/myProg
```

## 测试 

> 个人使用gonew时，更改go mod名称时偶尔会有点小问题

```shell
  cd myProg
  make wire
  or 
  make run
 ```

## 修改 mic-project-layout

在项目中搜索`mic-project-layout`，替换为你的项目名。