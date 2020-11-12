    Go 1.11 introduced Modules, enabling an alternative workflow.
This new approach will gradually become the default mode, deprecating the use of GOPATH.
    Modules aim to solve problems related to dependency management, version selection and reproducible builds;
they also enable users to run Go code outside of GOPATH.

mod使用实例
1 新建项目目录  mkdir learn-go-with-tests
2 cd learn-go-with-tests
3 go mod init github.com/haimingwey/leanr-go-with-tests
