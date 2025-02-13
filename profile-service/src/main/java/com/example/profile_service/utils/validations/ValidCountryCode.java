package com.example.profile_service.utils.validations;

import jakarta.validation.Constraint;
import jakarta.validation.Payload;
import java.lang.annotation.ElementType;
import java.lang.annotation.Retention;
import java.lang.annotation.RetentionPolicy;
import java.lang.annotation.Target;

@Constraint(validatedBy = CountryCodeValidator.class)
@Target({ElementType.FIELD, ElementType.PARAMETER})
@Retention(RetentionPolicy.RUNTIME)
public @interface ValidCountryCode {
  String message() default "Invalid country code. It must be a valid 2-letter ISO country code.";

  Class<?>[] groups() default {};

  Class<? extends Payload>[] payload() default {};
}