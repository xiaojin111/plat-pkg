**命名规则：**

- **私钥 Private Key File:** `${APP_ID}.pem`
- **公钥 Public Key File:** `${APP_ID}.pub`
- **指纹 Fingerprint File:** `${APP_ID}.fp.txt`

**命令：**

```sh
# APP ID
export APP_ID="app-test1"
# 输出的目标目录
export OUT_DIR="testdata"

# 生成私钥
openssl genrsa -out ${OUT_DIR}/${APP_ID}.pem 2048
# 根据私钥生成公钥
openssl rsa -in ${OUT_DIR}/${APP_ID}.pem -outform PEM -pubout -out ${OUT_DIR}/${APP_ID}.pub
# 根据私钥生成 SHA1 指纹
openssl rsa -in ${OUT_DIR}/${APP_ID}.pem -pubout -outform DER | \
    openssl sha1 -c | \
    awk '{print $2}'

# 生成一组 RSA 私钥、公钥、指纹
./genkey.sh ${APP_ID} ${OUT_DIR}

# 使用私钥进行 JWT 签名
go run sign.go \
	-key ${OUT_DIR}/${APP_ID}.pem \
	-iss ${APP_ID} \
	-inr 600s

# 使用私钥进行 JWT 签名技巧 (仅 macOS)
# 将生成的 JWT Token 立刻拷贝到剪切板
go run sign.go \
	-key ${OUT_DIR}/${APP_ID}.pem \
	-iss ${APP_ID} \
	-inr 600s | \
	pbcopy
	
# 使用公钥进行 JWT 验证
go run verify.go \
	-key ${OUT_DIR}/${APP_ID}.pub \
	-token ${JWT_TOKEN}
	
# 使用公钥进行 JWT 验证技巧 (仅 macOS)
# 从剪切板中读取已拷贝的 JWT Token
go run verify.go \
	-key ${OUT_DIR}/${APP_ID}.pub \
	-token $(pbpaste)
```

**JWT Debugger:** https://jwt.io/#debugger-io