# goExamples
go的练习小例子



## slice

![切片](resources/slice.png)

如上图 ，slice由由指向数组的指针，段的长度及其容量（段的最大长度）组成。段的长度由len函数获取，cap获取再数组中切片未分配的容量。

```go
var ar = [10]int{0,1,2,3,4,5,6,7,8,9}
var a = ar[5:7] 
// reference to subarray {5,6} - len(a) is 2 and cap(a) is 5 
```
[0,1,2,3,4,5,6,7,8,9]截取时"左闭右开"所以为[5,6],而其数组指针指向5，后续只有5,6,7,8,9所以段最大长度为5


