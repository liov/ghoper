package test;

import java.util.ArrayList;
import java.util.List;
import org.apache.commons.lang3.StringUtils;

/**
 * @author ：lbyi
 * @date ：Created in 2019/4/4
 * @description：test
 */
public class Hoper {

    public static void main(String[] args){

            List<String> stringList = new ArrayList<>();

            stringList.add("test");

            stringList.add("aa");

            String test = StringUtils.join(stringList,",");

            System.out.println(test);

            List<String> newImgUrl = new ArrayList<>();
            for(String s : stringList){
                if (s != null && s.trim().length() > 0) {
                    if (!s.startsWith("http") && !s.startsWith("https")) {
                        s = "xiugai" + s;
                        newImgUrl.add(s);
                    }
                }
            }

            stringList = newImgUrl;
            System.out.println(stringList);
    }
}
