package xyz.hoper.request;

import lombok.Data;

@Data
public class UserReq {
    private Integer id;

    private String name;

    private String password;

    private String email;

    private String phone;
}
