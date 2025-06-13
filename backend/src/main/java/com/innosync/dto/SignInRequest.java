package com.innosync.dto;

import jakarta.validation.constraints.Email;
import jakarta.validation.constraints.NotBlank;
import jakarta.validation.constraints.Size;
import lombok.Data;

@Data
public class SignInRequest {
    @NotBlank(message = "Email is required")
    @Email(message = "Email must be valid")
    String email;

    @Size(min = 8, max = 20, message = "Password's length must be between 8 to 20")
    @NotBlank(message = "Password is required")
    String password;
}
