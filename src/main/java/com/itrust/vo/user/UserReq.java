package com.itrust.vo.user;

import com.itrust.vo.BaseReq;

public class UserReq extends BaseReq {

	private static final long serialVersionUID = 1L;

	private int userId;
	private String userName;
	private int userAge;
	private String userSex;
	private String userAddr;
	
	public int getUserId() {
		return userId;
	}
	public void setUserId(int userId) {
		this.userId = userId;
	}
	public String getUserName() {
		return userName;
	}
	public void setUserName(String userName) {
		this.userName = userName;
	}
	public int getUserAge() {
		return userAge;
	}
	public void setUserAge(int userAge) {
		this.userAge = userAge;
	}
	public String getUserSex() {
		return userSex;
	}
	public void setUserSex(String userSex) {
		this.userSex = userSex;
	}
	public String getUserAddr() {
		return userAddr;
	}
	public void setUserAddr(String userAddr) {
		this.userAddr = userAddr;
	}
	@Override
	public String toString() {
		return "UserReq [userId=" + userId + ", userName=" + userName + ", userAge=" + userAge + ", userSex=" + userSex
				+ ", userAddr=" + userAddr + "]";
	}
	
	
	
	
}
