package com.itrust.controller.user;

import java.util.List;

import javax.annotation.Resource;
import javax.servlet.http.HttpServletRequest;
import javax.servlet.http.HttpServletResponse;

import org.slf4j.Logger;
import org.slf4j.LoggerFactory;
import org.springframework.stereotype.Controller;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RequestMethod;
import org.springframework.web.bind.annotation.ResponseBody;

import com.itrust.controller.BaseController;
import com.itrust.service.UserService;
import com.itrust.vo.user.UserReq;
import com.itrust.vo.user.UserVO;

@Controller
public class UserController extends BaseController{

	protected Logger logger = LoggerFactory.getLogger(getClass());
	
	@Resource(name = "userService")
	private UserService userService;
	
	@RequestMapping(value = "/createUser",method=RequestMethod.POST)
	@ResponseBody
	public String createUser(HttpServletRequest request, HttpServletResponse response) {
		logger.info(">>>>>>>>>>>>>>>>>>>>>>>>> 创建user 表");
		userService.createUserTable();
		
		return "创建用户成功";
	}
	
	@RequestMapping(value = "/addUser",method=RequestMethod.POST)
	@ResponseBody
	public String addUser(UserReq userReq,HttpServletRequest request, HttpServletResponse response) {
		userService.addUser(userReq);
		logger.info(">>>>>>>>>>>>>>>>>>>>>>>>> 添加用户成功");
		return "添加用户成功";
	}
	
	@RequestMapping(value = "/updateUser",method=RequestMethod.POST)
	@ResponseBody
	public String updateUser(UserReq userReq,HttpServletRequest request, HttpServletResponse response) {
		userService.updateUser(userReq);
		logger.info(">>>>>>>>>>>>>>>>>>>>>>>>> 更新用户成功");
		return "更新用户成功";
	}
	
	@RequestMapping(value = "/deleteUser",method=RequestMethod.POST)
	@ResponseBody
	public String deleteUser(int userId,HttpServletRequest request, HttpServletResponse response) {
		userService.deleteUser(userId);
		logger.info(">>>>>>>>>>>>>>>>>>>>>>>>> 删除用户成功");
		return "删除用户成功";
	}
	
	@RequestMapping(value = "/queryUsers",method=RequestMethod.POST)
	@ResponseBody
	public String queryUsers(int userAge,HttpServletRequest request, HttpServletResponse response) {
		List<UserVO> users = userService.queryUsers(userAge);
		logger.info(">>>>>>>>>>>>>>>>>>>>>>>>> 查询用户成功");
		return users.toString();
	}
}
