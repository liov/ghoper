package test;

public class A {

	public A() {
		System.out.println("A的构造方法,i=" + i);
		method();
	}

	int i = 10;

	public void method() {
		System.out.println("test.A 的 method i = " + i);
	}
}
