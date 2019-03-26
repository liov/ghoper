package unsafe;

import sun.misc.Unsafe;

import java.lang.reflect.Field;

/**
 * @author ：lbyi
 * @date ：Created in 2019/3/26 9:27
 * @description：unsafe
 * @modified By：
 */
public class MyObj {
    int objField=10;
    static int clsField=10;
    int[] array={10,20,30,40,50};
    static Unsafe U;
    static {
        try {
            init();
        } catch (NoSuchFieldException | IllegalAccessException e) {
            e.printStackTrace();
        }
    }
    public static void init() throws NoSuchFieldException, IllegalAccessException {
        Field f= Unsafe.class.getDeclaredField("theUnsafe");
        f.setAccessible(true);
        U= (Unsafe) f.get(null);
    }

    static void getObjFieldVal() throws NoSuchFieldException {
        Field field=MyObj.class.getDeclaredField("objField");
        long offset= U.objectFieldOffset(field);
        MyObj obj=new MyObj();

        int val= U.getInt(obj,offset);
        System.out.println("1.\t"+(val==10));

        MyObjChild child=new MyObjChild();
        int corVal1= U.getInt(child,offset);
        System.out.println("2.\t"+(corVal1==10));

        Field fieldChild=MyObj.class.getDeclaredField("objField");
        long offsetChild= U.objectFieldOffset(fieldChild);
        System.out.println("3.\t"+(offset==offsetChild));
        int corVal2= U.getInt(obj,offsetChild);
        System.out.println("4.\t"+(corVal2==10));

        short errVal1=U.getShort(obj,offset);
        System.out.println("5.\t"+(errVal1==10));

        int errVal2=U.getInt("abcd",offset);
        System.out.println("6.\t"+errVal2);

    }
}
