package com.example.profile_service.exceptions;

import lombok.Getter;

@Getter
public class ProfileNotFoundException extends Exception {
  public ProfileNotFoundException(String userId) {
    super("Profile with id: \"" + userId + "\" not found");
  }
}
