package com.innosync.dto;

import jakarta.validation.constraints.Email;
import jakarta.validation.constraints.NotBlank;
import jakarta.validation.constraints.Size;
import lombok.Data;

@Data
public class SignUpRequest {
   @NotBlank(message = "Email is required")
   @Email(message = "Email must be valid")
   String email;

   @NotBlank(message = "Password is required")
   @Size(min = 8, max = 20, message = "Password's length must be between 8 to 20")
   String password;
}
