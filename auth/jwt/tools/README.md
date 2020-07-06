## JWT 管理规则

### **1 命名规则**

- **私钥 Private Key File:** `${APP_ID}.pem`
- **公钥 Public Key File:** `${APP_ID}.pub`
- **指纹 Fingerprint File:** `${APP_ID}.fp.txt`



### 2 使用 Step 管理（推荐）

> macOS

#### 安装 Step 工具

```sh
brew install step
```

#### 生成 RSA 秘钥对与指纹

```sh
APP_ID="app-test5"
PUB_FILE=${APP_ID}.pub
PRI_FILE=${APP_ID}.pem
PASS_FILE="PASSWORD"

# 1. 生成秘钥对
step crypto keypair ${PUB_FILE} ${PRI_FILE} \
	--kty=RSA \
	--password-file=${PASS_FILE} # 密码文件
	
# 2. 根据公钥（Public Key）生成 SHA1 指纹
step crypto hash digest ${PUB_FILE} --alg=SHA1
	
```

#### 生成 JWT 签名

```sh
JWT_ALG="RS256"
JWT_AUD="com.jinmuhealth.partner.dayima.api.sys-intg"		# 端点 SRV 名称
JWT_SUB="DaYiMa"									# 客户名称
JWT_ISS="dayima"								# APP ID
JWT_PRIVATE_KEY="dayima.pem" 	# Private Key
PASS_FILE="PASSWORD"

# UNIX time: NOW
JWT_IAT=$(date +'%s')
# UNIX time: NOW+600s
JWT_EXP=$(date -r ${JWT_IAT} -v+600S +'%s')

step crypto jwt sign \
	--alg=${JWT_ALG} \
	--iss=${JWT_ISS} \
	--sub=${JWT_SUB} \
	--aud=${JWT_AUD} \
	--iat=${JWT_IAT} \
	--exp=${JWT_EXP} \
	--key=${JWT_PRIVATE_KEY} \
	--password-file=${PASS_FILE} \
	| pbcopy && pbpaste # magic line

```

#### 查看一个 JWT 签名内容

```sh
# 从剪切板里面取 JWT 值
pbpaste | step crypto jwt inspect --insecure

# ------------------
# 或者手动方式
JWT="eyJhbGciOiJSUzI1NiIsImtpZCI6IiIsInR5cCI6IkpXVCJ9.eyJhdWQiOiJjb20uamlubXVoZWFsdGgucGFydG5lci5kYXlpbWEuYXBpLnN5cy1pbnRnIiwiZXhwIjoxNTcwMTc1MDEyLCJpYXQiOjE1NzAxNzQ0MTIsImlzcyI6ImFwcC10ZXN0NSIsIm5iZiI6MTU3MDE3NDQxMiwic3ViIjoiRGFZaU1hIn0.gw-LTpaUg-XBhkYVsCYDEef3vDrEYdAACbxIoEKw7UTI9-KFDBEqNJOgGZFmYa6DGx5wKaR9kcnyAC6Z2oqO4_EzqNokljk3YgDdm6JQy58_V0MrCxwbGQ-Xjn21C0e1MDTvp9cBfiJfYJpmUQV1Kut7PJ4M2jZ8MLISI2jXOxX7EpFf-CLB1ptQrwKssUP0MxdljAXEIMzioL-nuzAMpnV15KkJO4Ij_6f9R10M-zErd9sc0o0e7PMWlcl2UI25fCtVu8NxWVCB41b5dvp3avQGFQKwzfKBfGJzfRlr_twBspT15LDLSirr87Nf_PGh8JBQhYH9GaIn8-BVTjm76A"

echo ${JWT} | step crypto jwt inspect --insecure
```

#### 验证 JWT 签名

```sh
JWT_ALG="RS256"
JWT_AUD="com.jinmuhealth.partner.dayima.api.sys-intg"		# 端点 SRV 名称
JWT_SUB="DaYiMa"									# 客户名称
JWT_ISS="app-test5"								# APP ID
JWT_PUBLIC_KEY="app-test5.pub"	 	# Public Key

JWT="eyJhbGciOiJSUzI1NiIsImtpZCI6IiIsInR5cCI6IkpXVCJ9.eyJhdWQiOiJjb20uamlubXVoZWFsdGgucGFydG5lci5kYXlpbWEuYXBpLnN5cy1pbnRnIiwiZXhwIjoxNTcwMTc1MDEyLCJpYXQiOjE1NzAxNzQ0MTIsImlzcyI6ImFwcC10ZXN0NSIsIm5iZiI6MTU3MDE3NDQxMiwic3ViIjoiRGFZaU1hIn0.gw-LTpaUg-XBhkYVsCYDEef3vDrEYdAACbxIoEKw7UTI9-KFDBEqNJOgGZFmYa6DGx5wKaR9kcnyAC6Z2oqO4_EzqNokljk3YgDdm6JQy58_V0MrCxwbGQ-Xjn21C0e1MDTvp9cBfiJfYJpmUQV1Kut7PJ4M2jZ8MLISI2jXOxX7EpFf-CLB1ptQrwKssUP0MxdljAXEIMzioL-nuzAMpnV15KkJO4Ij_6f9R10M-zErd9sc0o0e7PMWlcl2UI25fCtVu8NxWVCB41b5dvp3avQGFQKwzfKBfGJzfRlr_twBspT15LDLSirr87Nf_PGh8JBQhYH9GaIn8-BVTjm76A"

echo ${JWT} | step crypto jwt verify \
	--alg=${JWT_ALG} \
	--iss=${JWT_ISS} \
	--aud=${JWT_AUD} \
	--key=${JWT_PUBLIC_KEY}
	
# ------------------------------------
# 从剪切板里面取 JWT 值
pbpaste | step crypto jwt verify \
	--alg=${JWT_ALG} \
	--iss=${JWT_ISS} \
	--aud=${JWT_AUD} \
	--key=${JWT_KEY}
```

### 3 使用 OpenSSL 管理

> 本方式不推荐内部使用，可以用来指导外部用户操作

```sh
export APP_ID="app-test1"
export OUT_DIR="testdata"
export PASSWORD_FILE="${OUT_DIR}/PASSWORD"

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
	-pass ${PASSWORD_FILE} \
	-inr 600s

# 使用私钥进行 JWT 签名技巧 (仅 macOS)
# 将生成的 JWT Token 立刻拷贝到剪切板
go run sign/main.go \
	-key ${OUT_DIR}/${APP_ID}.pem \
	-iss ${APP_ID} \
	-pass ${PASSWORD_FILE} \
	-inr 600s | \
	pbcopy
	
# 使用公钥进行 JWT 验证
go run verify_std/main.go \
	-key ${OUT_DIR}/${APP_ID}.pub \
	-token ${JWT_TOKEN}
	
# 使用公钥进行 JWT 验证技巧 (仅 macOS)
# 从剪切板中读取已拷贝的 JWT Token
go run verify_std/main.go \
	-key ${OUT_DIR}/${APP_ID}.pub \
	-token $(pbpaste)
```

**JWT Debugger:** https://jwt.io/#debugger-io

