### 检查清单✅

> 以下适用的项目中标注 `x` 打钩；不适用的留空。


- [ ] ⚠️确认当前准备将PR提交合并到哪个**目标（develop/master/release-*）分支（左边）**。
- [ ] ⚠️确认当前正准备提交一个 **feature/hotfix 分支（右边）**。🚫不要提交你自己的 **master/develop** 分支。
- [ ] 👓确保PR变更范围足够小，以方便 Review。
- [ ] 📝确保所有的 Public 公开 API 都已经标注了文档注释。
- [ ] 📝更新 README 或其它文档，如果必须的话。
- [ ] 🔰确保PR中**不包含**破坏安全性的信息（Token、秘钥、密码等），用于自动化测试的除外。
- [ ] ⚠️确认已经**从 upstream 更新合并进来必要的 commits** 到本 PR；如果有冲突，则冲突必须解决。
- [ ] ⚠️检查每一个 **commit 的 message** 内容都符合风格规范。
- [ ] ⚠️检查代码没有 **Lint** 失败。**Lint 的 warning 视为 error**。
- [ ] ⚠️检查代码已经通过**单元测试**。
- [ ] ⚠️检查代码能够**正常编译**；🚫不应当有任何**编译 warning 或 error**。
- [ ] 🎈考虑将与本次提交的PR所解决的问题相关的测试都包含了进来，且测试通过。
- [ ] ⚠️避免 CI 执行失败。



### PR 基本概况

1. 标注 `x` 勾选 PR 的主要类型，单选：

  - [ ] Bugfix 修复Bug
  - [ ] Feature 功能
  - [ ] Code style update (formatting, renaming) 代码样式更新（格式化、重命名等）
  - [ ] Refactoring (no functional changes, no api changes) 重构（无功能变化、无API变化）
  - [ ] Build related changes 构建相关的变更
  - [ ] Documentation content changes 文档内容变化
  - [ ] Other 其它 (在后面描述): ???

2. 本 PR 是否引入了 Break Change 破坏性变更：

  - [ ] Yes
  - [ ] No
  - [ ] 不确定，请后面描述




### PR描述📝

请在此详细描述你的PR涉及的内容……

……



**🔗相关 GitHub/Jira Issue 或链接：**

- foo/bar
- ……



### 其它备注📝

请在此描述其它需要 Reviewer 注意的相关信息……

……

