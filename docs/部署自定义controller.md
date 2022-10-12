1.构建镜像
使用dockfile文件
运行命令docker build -t kubebuilder-demo:v1 .
使用docker images查看是否构建成功
你也可以用docker push kubebuilder-demo:v1将本地镜像推送到镜像库中
2.部署
kubectl apply -f config/samples提交







