package com.example.profile_service.utils.validations;

import jakarta.validation.ConstraintValidator;
import jakarta.validation.ConstraintValidatorContext;
import java.util.Arrays;
import java.util.Set;
import java.util.stream.Collectors;

public class CountryCodeValidator implements ConstraintValidator<ValidCountryCode, String> {
  private static final Set<String> ISO_COUNTRY_CODES = Arrays.stream(java.util.Locale.getISOCountries())
      .collect(Collectors.toSet());

  @Override
  public boolean isValid(String country, ConstraintValidatorContext context) {
    if (country == null) {
      return true; // Let @NotNull handle null checks
    }
    return ISO_COUNTRY_CODES.contains(country);
  }
}
