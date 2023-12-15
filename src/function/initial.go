package function

/*
1. init函数：初始化函数，可以用来做一些初始化操作。每一个源文件都可以包含一个init函数，该函数会在main函数执行前，被Go框架调用
2. 全局变量定义，init函数，main函数，它们三个的执行顺序是：
		Step1： 全局变量的定义
		Step2： init函数的调用
		Step3： 执行main函数
3. 多个源文件都有init函数，执行顺序：先执行非main中的init，最后执行main中的init函数
*/
