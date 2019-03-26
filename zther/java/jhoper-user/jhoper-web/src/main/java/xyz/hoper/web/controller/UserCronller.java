package xyz.hoper.web.controller;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Controller;
import org.springframework.web.bind.annotation.*;
import xyz.hoper.request.UserReq;
import xyz.hoper.response.UserRes;
import xyz.hoper.web.service.UserService;


@RestController
@RequestMapping("/user")
public class UserCronller {

    @Autowired
    private UserService userService;

    @RequestMapping(value="/get",method = RequestMethod.GET)
    public UserRes toIndex(@RequestParam("id") UserReq req){
        return this.userService.getUser(req);
    }

    @RequestMapping(value="/hello",method = RequestMethod.GET)
    public String hello(){
        return "hello";
    }
}
