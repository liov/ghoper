package test.hoper;

/**
 * @author ：lbyi
 * @date ：Created in 2019/3/11 13:41
 * @description：匿名内部类
 * @modified By：
 */
public class Outer {

    private String string = "JYB";

    class Inner{}

    private void outerTest(char ch){

        Integer integer = 1;
        new Inner() {
            void innerTest(){
                System.out.println(string);
                System.out.println(ch);
                System.out.println(integer);
            }
        }.innerTest();
    }

    public static void main(String[] args){
        new Outer().outerTest('b');
    }
}
