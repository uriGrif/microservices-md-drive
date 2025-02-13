package com.example.profile_service.dtos;

import com.example.profile_service.utils.validations.ValidCountryCode;
import jakarta.validation.constraints.NotBlank;
import jakarta.validation.constraints.NotNull;
import jakarta.validation.constraints.Pattern;
import lombok.Data;

@Data
public class UpdateProfileDto {
  @NotBlank(message = "nickname cannot be blank")
  @Pattern(regexp = "[A-Za-z0-9]+")
  private String nickname;

  @NotNull(message = "a country must be specified")
  @ValidCountryCode
  private String country;

  private String description;
}
