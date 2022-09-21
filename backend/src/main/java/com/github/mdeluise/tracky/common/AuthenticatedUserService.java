package com.github.mdeluise.tracky.common;

import com.github.mdeluise.tracky.authentication.User;
import com.github.mdeluise.tracky.authentication.UserService;
import org.springframework.security.core.Authentication;
import org.springframework.security.core.context.SecurityContext;
import org.springframework.security.core.context.SecurityContextHolder;
import org.springframework.stereotype.Component;

@Component
public class AuthenticatedUserService {
    private final UserService userService;


    protected AuthenticatedUserService(UserService userService) {
        this.userService = userService;
    }


    public User getAuthenticatedUser() {
        SecurityContext context = SecurityContextHolder.getContext();
        Authentication authentication = context.getAuthentication();
        String username = authentication.getName();
        return userService.get(username);
    }
}
