package com.innosync.controller;

import com.innosync.dto.auth.*;
import com.innosync.model.RefreshToken;
import com.innosync.model.User;
import com.innosync.repository.UserRepository;
import com.innosync.security.JwtUtil;
import com.innosync.service.AuthService;
import com.innosync.service.RefreshTokenService;
import io.swagger.v3.oas.annotations.Operation;
import io.swagger.v3.oas.annotations.tags.Tag;
import jakarta.validation.Valid;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.http.HttpStatus;
import org.springframework.http.ResponseEntity;
import org.springframework.security.crypto.bcrypt.BCryptPasswordEncoder;
import org.springframework.web.bind.annotation.PostMapping;
import org.springframework.web.bind.annotation.RequestBody;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RestController;

@RestController
@RequestMapping("/api/auth")
@Tag(name = "Authentication API", description = "API for user authentication") // Swagger annotation
public class AuthController {

    @Autowired
    AuthService authService;

    @PostMapping("/signup")
    @Operation(summary = "Sign up user")
    public ResponseEntity<?> signUp(@Valid @RequestBody SignUpRequest signUpRequest) {
        return authService.signUp(signUpRequest);
    }

    @PostMapping("/login")
    @Operation(summary = "Log in User")
    public ResponseEntity<?> singIn(@Valid @RequestBody SignInRequest signInRequest) {
        return authService.signIn(signInRequest);
    }

    @PostMapping("/refresh")
    public ResponseEntity<?> refreshToken(@RequestBody RefreshTokenRequest request) {
        return authService.refreshToken(request);
    }



    @PostMapping("logout")
    public ResponseEntity<?> logout(@RequestBody LogoutRequest logoutRequest) {
        return authService.logout(logoutRequest);
    }
 }
