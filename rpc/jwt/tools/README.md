**命名规则：**

- **私钥 Private Key File:** `${APP_ID}.pem`
- **公钥 Public Key File:** `${APP_ID}.pub`
- **指纹 Fingerprint File:** `${APP_ID}.fp.txt`

**命令：**

```sh
export APP_ID="app-test1"
export OUT_DIR="testdata"

# 1. 生成私钥
openssl genrsa -out ${OUT_DIR}/${APP_ID}.pem 2048

# 2. 根据私钥生成公钥
openssl rsa -in ${OUT_DIR}/${APP_ID}.pem -outform PEM -pubout -out ${OUT_DIR}/${APP_ID}.pub

# 3. 根据私钥（Private Key）生成 SHA1 指纹
openssl rsa -in ${OUT_DIR}/${APP_ID}.pem -pubout -outform DER | \
	openssl sha1 -c | \
  awk '{print $2}'

# 4. 根据公钥（Public Key）生成 SHA1 指纹
# 注意：本步骤输出结果应当与第3步输出结果完全一致
openssl rsa -in ${OUT_DIR}/${APP_ID}.pub -pubin -pubout -outform DER | \
	openssl sha1 -c | \
  awk '{print $2}'


# 快速生成一组 RSA 私钥、公钥、指纹
./genkey.sh ${APP_ID} ${OUT_DIR}

# 使用私钥进行 JWT 签名
go run sign/main.go \
	-key ${OUT_DIR}/${APP_ID}.pem \
	-iss ${APP_ID} \
	-inr 600s

# 使用私钥进行 JWT 签名技巧 (仅 macOS)
# 将生成的 JWT Token 立刻拷贝到剪切板
go run sign/main.go \
	-key ${OUT_DIR}/${APP_ID}.pem \
	-iss ${APP_ID} \
	-inr 600s | \
	pbcopy
	
# 使用公钥进行 JWT 验证
go run verify/main.go \
	-key ${OUT_DIR}/${APP_ID}.pub \
	-token ${JWT_TOKEN}
	
# 使用公钥进行 JWT 验证技巧 (仅 macOS)
# 从剪切板中读取已拷贝的 JWT Token
go run verify/main.go \
	-key ${OUT_DIR}/${APP_ID}.pub \
	-token $(pbpaste)
```

**JWT Debugger:** https://jwt.io/#debugger-io
