FROM 949191617935.dkr.ecr.cn-north-1.amazonaws.com.cn/jm-app/jm-micro:latest
ADD cert  /cert
WORKDIR /
ENTRYPOINT [ "/jm-micro" ]
