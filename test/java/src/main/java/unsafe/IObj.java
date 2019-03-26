package unsafe;

import jdk.internal.misc.Unsafe;

import java.lang.reflect.Field;

/**
 * @author ：lbyi
 * @date ：Created in 2019/3/26 14:43
 * @description：
 * @modified By：
 */
public class IObj {
    int objField = 10;
    static Unsafe U;

    static {
        U = Unsafe.getUnsafe();
    }

    //java.exe --add-exports=java.base/jdk.internal.misc=ALL-UNNAMED IObj.java
    //IDEA中VM options：--add-exports=java.base/jdk.internal.misc=test
    public static void main(String[] args) throws NoSuchFieldException {
        Field field = IObj.class.getDeclaredField("objField");
        long offset = U.objectFieldOffset(field);
        IObj obj = new IObj();

        int val = U.getInt(obj, offset);
        System.out.println("1.\t" + val + "\t" + (val == 10));
    }
}
