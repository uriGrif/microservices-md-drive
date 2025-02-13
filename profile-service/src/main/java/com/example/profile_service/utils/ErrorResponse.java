package com.example.profile_service.utils;

import lombok.AllArgsConstructor;
import lombok.Data;

@Data
@AllArgsConstructor
public class ErrorResponse {
  private Integer httpStatusCode;
  private String message;
}
