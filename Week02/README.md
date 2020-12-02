# 作业

我们在数据库操作的时候，比如 dao 层中当遇到一个 sql.ErrNoRows 的时候，是否应该 Wrap 这个 error，抛给上层。为什么，应该怎么做请写出代码？

# 思路

应该Wrap这个error，因为当前业务的后续查询和逻辑处理有可能受此结果的影响，所以应当交给service层消化此error
