package main.java.xyz.hoper.jhoperuser.remote;


import org.springframework.cloud.openfeign.FeignClient;
import org.springframework.stereotype.Service;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RequestMethod;


@FeignClient(value = "jhoper-web")
public interface UserService {

    @RequestMapping(value = "/user/hello",method = RequestMethod.GET)
     String hello();
}
