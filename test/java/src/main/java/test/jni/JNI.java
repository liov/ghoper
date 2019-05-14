package test.jni;

/**
 * @author ：lbyi
 * @date ：Created in 2019/5/13
 * @description：jni
 */

//javac -h . JNI.java or javac JNI.java -h JniH
//java jni的开销只会越来越大,看来java优化足够好了，不需要c了
public class JNI {

    //链动态库
    static {
        System.loadLibrary("hello");
    }

    //方法定义
    public native void testHelloVoid();

    public native String testHello();

    public native long fib(int n);

    public static void main(String[] args){
        //执行
        jnifib();
        javafib();
    }

    private static void jnifib(){
        JNI jni = new JNI();
        long starTime=System.currentTimeMillis();
        System.out.println(starTime);
        System.out.println(jni.fib(43));
        System.out.println(System.currentTimeMillis()-starTime);
    }

    private static void javafib(){
        long starTime=System.currentTimeMillis();
        System.out.println(starTime);
        System.out.println(jfib(43));
        System.out.println(System.currentTimeMillis()-starTime);
    }

    private static long jfib(int n){
        if(n<2) return 1;
        return jfib(n-1)+jfib(n-2);
    }

}
