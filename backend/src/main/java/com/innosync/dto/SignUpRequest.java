package com.innosync.dto;

import jakarta.validation.constraints.Email;
import jakarta.validation.constraints.NotBlank;
import lombok.Data;

@Data
public class SignUpRequest {
   @NotBlank(message = "Email is required")
   @Email(message = "Email must be valid")
   String email;

   @NotBlank(message = "Password is required")
   String password;
}
