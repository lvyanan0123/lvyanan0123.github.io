
如何自动打tag
===

编写人 *lynn*

test

将下面的文件保存为 tag,并增加可执行权限。

    vim ~/bin/tag
    chmod +x ~/bin/tag
    export PATH=~/bin:$PATH

```
#!/bin/sh
# This script will be executed after commit in placed in .git/hooks/post-commit

# 版本号类似于 MAJOR.MINOR.PATCH 遵循 Semantic Versioning 2.0.0:
# MAJOR: 主版本号，不兼容的修改，如：底层重构、框架层变动
# MINOR:次版本号，向前兼容的修改，如：增加新功能、代码优化
# PATCH: 小版本号，修复bug