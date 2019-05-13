package test.jni;

/**
 * @author ：lbyi
 * @date ：Created in 2019/5/13
 * @description：jni
 */

//javac -h . JNI.java or javac JNI.java -h JniH

public class JNI {

    //链动态库
    static {
        System.loadLibrary("hello");
    }

    //方法定义
    public native void testHelloVoid();

    public native String testHello();

    public static void main(String[] args){
        //执行
        JNI jni = new JNI();
        jni.testHelloVoid();
        System.out.println(jni.testHello());
    }
}
