/*************************************************************************
	> File Name : main.cpp
	> Author : qi.li 
	> Mail : 2273172321@qq.com 
	> Created Time: Tue 24 Dec 2019 09:31:11 PM CST
 ************************************************************************/

#include<iostream>
#include <stdio.h>
using namespace std;

struct T {
    char c1;
    char c2;
    int a:8;
    int b:8;
    int c:8;
    int d:8;
    char c3;
    char c4;
};

int main()
{
    char s1[1024];
    s1[0] = 'a';
    s1[1] = 'b';
    s1[2] = 'c';
    s1[1023] = 'd';
    char w1 = 'q';
    char w2 = 'w';
    char w3 = 'e';
    std::cout<<w1<<std::endl;
    T t;
    t.a = 0;
    t.b = 65;
    t.c = 0;
    t.d = 0;
    std::cout<<w1<<std::endl;
    std::cout<<w1<<std::endl;
    std::cout<<w1<<std::endl;
    t.b = 855;
    std::cout<<w1<<std::endl;
    std::cout<<w1<<std::endl;
}
