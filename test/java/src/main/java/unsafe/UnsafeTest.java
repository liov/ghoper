package unsafe;

/**
 * @author ：lbyi
 * @date ：Created in 2019/3/26 9:34
 * @description：
 * @modified By：
 */
public class UnsafeTest {

    public static void main(String[] args){
        try {
            MyObj.init();
        } catch (NoSuchFieldException | IllegalAccessException e) {
            e.printStackTrace();
        }
        try {
            MyObj.getObjFieldVal();
        } catch (NoSuchFieldException e) {
            e.printStackTrace();
        }
    }
}
