1: goroutine 线程调度规则
goroutine分配和线程分配的关系
一个核，对应的线程、协程数量
多个核，默认线程分配、goroutine分配

2: goroutine切换的条件
G发生上下文切换条件
系统调用；
读写channel；
gosched主动放弃，会将G扔进全局队列；
详细调用关系

3: 全局gc在系统调用的影响
