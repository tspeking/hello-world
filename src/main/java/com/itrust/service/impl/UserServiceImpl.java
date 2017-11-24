package com.itrust.service.impl;

import java.util.List;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;

import com.itrust.dao.user.UserDao;
import com.itrust.service.UserService;
import com.itrust.vo.user.UserReq;
import com.itrust.vo.user.UserVO;

@Service("userService")
public class UserServiceImpl implements UserService {

	@Autowired
	private UserDao userDao;
	
	@Override
	public void createUserTable() {
		userDao.createUser("user_info");
	}

	@Override
	public void addUser(UserReq userReq) {
		userDao.addUser(userReq);
	}

	@Override
	public void updateUser(UserReq userReq) {
		userDao.updateUser(userReq);
	}

	@Override
	public void deleteUser(int userId) {
		userDao.deleteUser(userId);
	}

	@Override
	public List<UserVO> queryUsers(int userAge) {
		return userDao.queryUsers(userAge);
	}

}
