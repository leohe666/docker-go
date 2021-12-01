# Docker-go使用说明

这是一个通过shell结合docker，用于常规开发的工具包。方式方法来源于 https://zhuanlan.zhihu.com/p/412197352 并在此基础上进行改进优化

使用步骤：
- 进入下载目录并执行 source addpath.sh
- chmod +x go gobuild
- 新建go文件或项目
- 使用go run hello.go 运行项目
- 使用gobuild hello.go linux来打包编译,其中最后一个参数为要编译的平台包括(windows,linux,darwin)