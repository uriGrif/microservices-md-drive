package com.example.profile_service.exceptions;

public class UserNotAuthenticatedException extends Exception {
  public UserNotAuthenticatedException() {
    super("user is not authenticated");
  }
}
