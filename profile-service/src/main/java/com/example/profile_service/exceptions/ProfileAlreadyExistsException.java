package com.example.profile_service.exceptions;

public class ProfileAlreadyExistsException extends Exception {
  public ProfileAlreadyExistsException() {
    super("Profile already exists");
  }
}
