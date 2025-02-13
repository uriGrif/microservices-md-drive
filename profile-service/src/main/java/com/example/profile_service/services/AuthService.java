package com.example.profile_service.services;

import jakarta.servlet.http.HttpServletRequest;
import org.springframework.stereotype.Service;

@Service
public class AuthService {

  public String getUserId(HttpServletRequest request) {
    return request.getHeader("X-Authenticated-User");
  }

}
