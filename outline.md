# 实验环境

https://sourceforge.net/projects/mingw-w64/files/

MinGW-W64 GCC-8.1.0

i686-win32-dwarf

文件名：i686-8.1.0-release-win32-dwarf-rt_v6-rev0.7z

解压到任意文件夹，例如：C:\mingw

C++编译器文件：C:\mingw\bin\g++.exe

设置环境变量：

```
C:\>setx path "%path%;C:\mingw\bin"

成功: 指定的值已得到保存。

C:\>

```

重新打开命令提示符，输入：g++ --version

出现版本信息，表示实验环境搭建成功。

```
C:\>g++ --version
g++ (i686-win32-dwarf-rev0, Built by MinGW-W64 project) 8.1.0
Copyright (C) 2018 Free Software Foundation, Inc.
This is free software; see the source for copying conditions.  There is NO
warranty; not even for MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.


C:\>
```

# 什么是C++？

C++ = C + 面向对象 + ……

# 第一个C++程序

```c++
int main(){
    return 0; /* 1、2、3…… */
}
```

编译：g++ test.cpp -o test.exe

执行：test.exe

查看结果（返回值）：`echo %errorlevel%`

# 第二个C++程序

```c++
#include <stdio.h>
int main(){
    printf("Hello, I'm a C++ program.\n");
    return 0;
}
```

库函数：printf

头文件：stdio.h

文件路径：C:\tools\mingw\i686-w64-mingw32\include

如何找到：搜索……

# 第三个C++程序

```c++
#include <iostream>
using namespace std;
int main(){
    cout << "Hello, I'm really a C++ program." << endl;
    return 0;
}
```

输出流对象：cout

头文件：iostream

文件路径：C:\tools\mingw\lib\gcc\i686-w64-mingw32\8.1.0\include\c++

如何找到：搜索……

# C++语言中的I/O

输出流对象：cout

左移运算符：<<

```
int a = 5, b = 4;
printf("a + b = %d \n", a + b);
cout << "a + b = " << a + b << endl; 
```

输入流对象：cin

右移运算符：>>

```
int a, b ;
scanf("%d %d", &a, &b);
cin >> a >> b;
```

【格式化输出参考】

http://c.biancheng.net/view/7578.html

【格式化输入参考】

https://blog.csdn.net/lewsn2008/article/details/2295790

https://blog.csdn.net/dongtingzhizi/article/details/2299358

https://blog.csdn.net/lewsn2008/article/details/2299365

# C++语言中的数据类型

## 逻辑类型

```
    int i,j;
    bool a,b;
    
    a = true;
    b = false;
    cout << "a = " << a << ", " << "b = " << b << endl;
    
    i=a;
    j=b;
    cout << "i = " << i << ", " << "j = " << j << endl;

    a = !i;
    b = !j;
    cout << "a = " << a << ", " << "b = " << b << endl;
```

## 引用类型

https://uestc.feishu.cn/file/boxcnqKAYU21cg0cmeY2FzHQcGb

https://uestc.feishu.cn/file/boxcn2kHXLAUSE6qAK54RtzVLzh

