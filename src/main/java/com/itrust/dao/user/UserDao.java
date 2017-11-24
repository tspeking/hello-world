package com.itrust.dao.user;

import java.util.List;

import org.apache.ibatis.annotations.Param;

import com.itrust.vo.user.UserReq;
import com.itrust.vo.user.UserVO;

public interface UserDao {
	int createIndustryListTable(@Param("tableName") String tableName);
	int createUser(@Param("tableName") String tableName);
	int addUser(UserReq userReq);
	int updateUser(UserReq userReq);
	int deleteUser(int userId);
	List<UserVO> queryUsers(int userAge);
}
