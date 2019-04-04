package xyz.hoper.service;

import xyz.hoper.request.UserReq;
import xyz.hoper.response.UserRes;


public interface UserService {

    UserRes getUser(UserReq req);

    boolean addUser(UserReq req);

    boolean deleteUser(UserReq req);

    boolean updateUser(UserReq req);
}
