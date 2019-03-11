package main.java.xyz.hoper.jhoperuser.controller;


import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Controller;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RequestMethod;
import org.springframework.web.bind.annotation.ResponseBody;
import xyz.hoper.jhoperuser.remote.UserService;


@Controller
public class UserController {

    @Autowired
    private UserService userService;

    @RequestMapping(value="/hello",method = RequestMethod.GET)
    @ResponseBody
    public String hello(){
        return userService.hello();
    }
}
