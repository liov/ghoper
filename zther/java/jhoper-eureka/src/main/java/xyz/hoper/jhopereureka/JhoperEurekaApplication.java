package main.java.xyz.hoper.jhopereureka;

import org.springframework.boot.SpringApplication;
import org.springframework.boot.autoconfigure.SpringBootApplication;
import org.springframework.cloud.netflix.eureka.server.EnableEurekaServer;

@SpringBootApplication
@EnableEurekaServer
public class JhoperEurekaApplication {

    public static void main(String[] args) {
        SpringApplication.run(JhoperEurekaApplication.class, args);
    }

}

