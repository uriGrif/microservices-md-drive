package com.example.profile_service.exceptions;

public class UserNotAuthorizedException extends Exception {
  public UserNotAuthorizedException() {
    super("user cannot perform this action");
  }
}
