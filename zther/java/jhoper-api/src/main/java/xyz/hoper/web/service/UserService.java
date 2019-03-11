package main.java.xyz.hoper.web.service;

import main.java.xyz.hoper.request.UserReq;
import main.java.xyz.hoper.response.UserRes;


public interface UserService {

    UserRes getUser(UserReq req);

    boolean addUser(UserReq req);

    boolean deleteUser(UserReq req);

    boolean updateUser(UserReq req);
}
