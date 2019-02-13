package com.hoper;

import com.sun.deploy.util.StringUtils;

import java.util.ArrayList;
import java.util.List;

public class Hoper {

    public static void main(String[] args){
        List<String> stringList = new ArrayList<>();

        stringList.add("test");

        stringList.add("aa");

        String test = StringUtils.join(stringList,",");

        System.out.println(test);
    }
}
