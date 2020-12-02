## 题目

我们在数据库操作的时候，比如 dao 层中当遇到一个 sql.ErrNoRows 的时候，是否应该 Wrap 这个 error，抛给上层。为什么，应该怎么做请写出代码？

## 解答

答：应该wrap。

为什么：

除了调用同一个包中的函数出错了，直接返回，不用wrap（因为wrap同一个包的函数会导致双倍堆栈信息）。

这三种情况都应该wrap

- 调用自身的基础库
- 调用官方的标准库
- 调用github的第三方库

[代码](https://github.com/wljgithub/Go-000/blob/main/Week02/main.go)