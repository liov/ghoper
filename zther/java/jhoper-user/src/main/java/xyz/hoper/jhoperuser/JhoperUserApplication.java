package main.java.xyz.hoper.jhoperuser;


import org.springframework.boot.SpringApplication;
import org.springframework.boot.autoconfigure.SpringBootApplication;
import org.springframework.cloud.client.discovery.EnableDiscoveryClient;
import org.springframework.cloud.netflix.eureka.EnableEurekaClient;
import org.springframework.cloud.openfeign.EnableFeignClients;


@SpringBootApplication(scanBasePackages={"xyz.hoper.jhoperuser"})
@EnableEurekaClient
@EnableDiscoveryClient
@EnableFeignClients
public class JhoperUserApplication {

    public static void main(String[] args) { SpringApplication.run(JhoperUserApplication.class, args); }

}

