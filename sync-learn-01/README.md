# sync
2020-06-16
sync 包源码学习  
sync 包提供了基本的同步原语  
golang提倡不要以共享内存的方式来通信，而是以通信的方式来共享内存
###锁是什么，什么时候需要锁呢？
锁是sync包的核心，他主要有两个方法，加锁和解锁  
单线程中是顺序执行的，但是多线程中是并发执行的为了防止错误必须加锁才能防止数据错误  


