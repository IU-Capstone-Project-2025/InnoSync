package com.innosync.dto.auth;

import jakarta.validation.constraints.Email;
import jakarta.validation.constraints.NotBlank;
import jakarta.validation.constraints.Size;
import lombok.Data;
import lombok.Getter;

@Data
@Getter
public class SignUpRequest {
   @NotBlank(message = "Email is required")
   @Email(message = "Email must be valid")
   String email;

   @NotBlank(message = "Password is required")
   @Size(min = 8, max = 20, message = "Password's length must be between 8 to 20")
   String password;

   @NotBlank(message = "Full Name is required")
   @Size(min = 2, max = 50, message = "Full name's length must be between 2 to 50")
   String fullName;
}
