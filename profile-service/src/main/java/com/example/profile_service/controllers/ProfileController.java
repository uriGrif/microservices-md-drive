package com.example.profile_service.controllers;

import com.example.profile_service.dtos.CreateProfileDto;
import com.example.profile_service.dtos.SearchProfileDto;
import com.example.profile_service.dtos.UpdateProfileDto;
import com.example.profile_service.exceptions.ProfileAlreadyExistsException;
import com.example.profile_service.exceptions.ProfileNotFoundException;
import com.example.profile_service.exceptions.UserNotAuthenticatedException;
import com.example.profile_service.exceptions.UserNotAuthorizedException;
import com.example.profile_service.models.Profile;
import com.example.profile_service.services.AuthService;
import com.example.profile_service.services.ProfileService;
import com.example.profile_service.utils.ErrorResponse;
import jakarta.servlet.http.HttpServletRequest;
import jakarta.validation.Valid;
import org.springframework.http.HttpStatus;
import org.springframework.validation.FieldError;
import org.springframework.web.bind.MethodArgumentNotValidException;
import org.springframework.web.bind.annotation.DeleteMapping;
import org.springframework.web.bind.annotation.ExceptionHandler;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.PathVariable;
import org.springframework.web.bind.annotation.PostMapping;
import org.springframework.web.bind.annotation.PutMapping;
import org.springframework.web.bind.annotation.RequestBody;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RequestParam;
import org.springframework.web.bind.annotation.ResponseStatus;
import org.springframework.web.bind.annotation.RestController;
import java.util.HashMap;
import java.util.List;
import java.util.Map;

@RestController
@RequestMapping("/profiles")
public class ProfileController {

  private final ProfileService profileService;
  private final AuthService authService;

  public ProfileController(ProfileService profileService, AuthService authService) {
    this.profileService = profileService;
    this.authService = authService;
  }

  @GetMapping("/{id}") // TODO que el id lo saque del header con un middleware
  public Profile getProfile(@PathVariable String id) throws Exception {
    return this.profileService.getProfile(id);
  }

  @PostMapping
  @ResponseStatus(HttpStatus.CREATED)
  public Profile createProfile(HttpServletRequest request, @Valid @RequestBody CreateProfileDto profileDto) throws Exception {
    String userId = this.authService.getUserId(request);
    if (userId == null || userId.isBlank()) {
      throw new UserNotAuthenticatedException();
    }
    return this.profileService.createProfile(profileDto, userId);
  }

  @PutMapping("/{id}")
  public Profile updateProfile(HttpServletRequest request, @PathVariable String id, @Valid @RequestBody UpdateProfileDto profile) throws Exception {
    String userId = this.authService.getUserId(request);
    if (userId == null || userId.isBlank()) {
      throw new UserNotAuthenticatedException();
    }
    if (!userId.equals(id)) {
      throw new UserNotAuthorizedException();
    }
    return this.profileService.updateProfile(id, profile);
  }

  @DeleteMapping("/{id}")
  @ResponseStatus(HttpStatus.NO_CONTENT)
  public void deleteProfile(HttpServletRequest request, @PathVariable String id) throws Exception {
    String userId = this.authService.getUserId(request);
    if (userId == null || userId.isBlank()) {
      throw new UserNotAuthenticatedException();
    }
    if (!userId.equals(id)) {
      throw new UserNotAuthorizedException();
    }
    this.profileService.deleteProfile(id);
  }

  @GetMapping("/search")
  public List<SearchProfileDto> searchProfile(@RequestParam String q) {
    return this.profileService.searchProfileByNickname(q);
  }

  @ExceptionHandler(value = ProfileNotFoundException.class)
  @ResponseStatus(HttpStatus.NOT_FOUND)
  public ErrorResponse handleProfileNotFoundException(ProfileNotFoundException ex) {
    return new ErrorResponse(HttpStatus.NOT_FOUND.value(), ex.getMessage());
  }

  @ExceptionHandler(value = ProfileAlreadyExistsException.class)
  @ResponseStatus(HttpStatus.BAD_REQUEST)
  public ErrorResponse handleProfileAlreadyExistsException(ProfileAlreadyExistsException ex) {
    return new ErrorResponse(HttpStatus.BAD_REQUEST.value(), ex.getMessage());
  }

  @ExceptionHandler(value = UserNotAuthenticatedException.class)
  @ResponseStatus(HttpStatus.UNAUTHORIZED)
  public ErrorResponse handleUserNotAuthenticatedException(UserNotAuthenticatedException ex) {
    return new ErrorResponse(HttpStatus.UNAUTHORIZED.value(), ex.getMessage());
  }

  @ExceptionHandler(value = UserNotAuthorizedException.class)
  @ResponseStatus(HttpStatus.FORBIDDEN)
  public ErrorResponse handleUserNotAuthorizedException(UserNotAuthorizedException ex) {
    return new ErrorResponse(HttpStatus.FORBIDDEN.value(), ex.getMessage());
  }

  @ResponseStatus(HttpStatus.BAD_REQUEST)
  @ExceptionHandler(MethodArgumentNotValidException.class)
  public Map<String, String> handleValidationExceptions(MethodArgumentNotValidException ex) {
    Map<String, String> errors = new HashMap<>();
    ex.getBindingResult().getAllErrors().forEach((error) -> {
      String fieldName = ((FieldError) error).getField();
      String errorMessage = error.getDefaultMessage();
      errors.put(fieldName, errorMessage);
    });
    return errors;
  }
}
