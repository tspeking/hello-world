package com.itrust.service;

import java.util.List;

import com.itrust.vo.user.UserReq;
import com.itrust.vo.user.UserVO;

public interface UserService {
	void createUserTable();
	void addUser(UserReq userReq);
	void updateUser(UserReq userReq);
	void deleteUser(int userId);
	List<UserVO> queryUsers(int userAge);
}
