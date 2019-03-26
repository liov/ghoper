package xyz.hoper.web;

import org.mybatis.spring.annotation.MapperScan;
import org.springframework.boot.SpringApplication;
import org.springframework.boot.autoconfigure.SpringBootApplication;
import org.springframework.cloud.client.discovery.EnableDiscoveryClient;
import org.springframework.cloud.netflix.eureka.EnableEurekaClient;
import org.springframework.cloud.openfeign.EnableFeignClients;
import org.springframework.context.annotation.ComponentScan;

@SpringBootApplication(scanBasePackages={"xyz.hoper.web"})
@MapperScan("xyz.hoper.web.dao")
@ComponentScan(basePackages = {"xyz.hoper.web.*"})
@EnableEurekaClient
@EnableDiscoveryClient
@EnableFeignClients
public class JhoperWebApplication {
    /**
     *
     * @param args
     */
    public static void main(String[] args) {
        SpringApplication.run(JhoperWebApplication.class, args);
    }
}
