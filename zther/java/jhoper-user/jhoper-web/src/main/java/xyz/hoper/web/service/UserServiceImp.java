package xyz.hoper.web.service;


import org.springframework.beans.BeanUtils;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;
import xyz.hoper.request.UserReq;
import xyz.hoper.response.UserRes;
import xyz.hoper.web.dao.UserMapper;
import xyz.hoper.web.entity.User;


@Service
public class UserServiceImp implements UserService {


    @Autowired
    private UserMapper userMapper;

    @Override
    public UserRes getUser(UserReq req) {
        User user = userMapper.selectByPrimaryKey(req.getId());
        UserRes userRes = new UserRes();
        BeanUtils.copyProperties(user,userRes);
        return userRes;
    }

    @Override
    public boolean addUser(UserReq req) {
        try {
            User user = new User();
            BeanUtils.copyProperties(req,user);
            userMapper.insertSelective(user);
            return true;
        } catch (Exception e) {
            return false;
        }
    }

    @Override
    public boolean deleteUser(UserReq req) {
        try {
            userMapper.deleteByPrimaryKey(req.getId());
            return true;
        } catch (Exception e) {
            return false;
        }
    }

    @Override
    public boolean updateUser(UserReq req) {
        try {
            User user = new User();
            BeanUtils.copyProperties(req,user);
            userMapper.updateByPrimaryKeySelective(user);
            return true;
        } catch (Exception e) {
            return false;
        }
    }

}
