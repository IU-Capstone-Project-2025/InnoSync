package com.innosync.controller;

import com.innosync.dto.auth.*;
import com.innosync.model.RefreshToken;
import com.innosync.model.User;
import com.innosync.repository.UserRepository;
import com.innosync.security.JwtUtil;
import com.innosync.service.RefreshTokenService;
import io.swagger.v3.oas.annotations.Operation;
import io.swagger.v3.oas.annotations.tags.Tag;
import jakarta.validation.Valid;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.http.HttpStatus;
import org.springframework.http.ResponseEntity;
import org.springframework.security.crypto.password.PasswordEncoder;
import org.springframework.web.bind.annotation.PostMapping;
import org.springframework.web.bind.annotation.RequestBody;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RestController;

@RestController
@RequestMapping("/api/auth")
@Tag(name = "Authentication API", description = "API for user authentication") // Swagger annotation
public class    AuthController {

    private static final Logger logger = LoggerFactory.getLogger(AuthController.class);

    @Autowired
    private UserRepository userRepository;

    @Autowired
    private JwtUtil jwtUtil ;

    @Autowired
    private RefreshTokenService refreshTokenService;

    @Autowired
    private PasswordEncoder passwordEncoder;

    // Add setter for testing
    void setPasswordEncoder(PasswordEncoder passwordEncoder) {
        this.passwordEncoder = passwordEncoder;
    }

    @PostMapping("/signup")
    @Operation(summary = "Sign up user")
    public ResponseEntity<?> signUp(@Valid @RequestBody SignUpRequest signUpRequest) {
        logger.info("Received signup request for email: {}", signUpRequest.getEmail());
        String email = signUpRequest.getEmail();
        if (userRepository.findByEmail(email).isPresent()) {
            logger.warn("Signup attempt with already registered email: {}", email);
            return ResponseEntity.status(HttpStatus.CONFLICT).body("Email is already registered!");
        }

        User user = new User(
                signUpRequest.getEmail(),
                signUpRequest.getFullName(),
                passwordEncoder.encode(signUpRequest.getPassword())
        );
        userRepository.save(user);
        logger.info("User registered successfully: {}", email);

        String accessToken = jwtUtil.generateToken(user.getEmail());
        RefreshToken refreshToken = refreshTokenService.createRefreshToken(user);

        logger.debug("Generated tokens for user: {}", email);
        return ResponseEntity.ok(new AuthResponse(accessToken, refreshToken.getToken()));
    }

    @PostMapping("/login")
    @Operation(summary = "Log in User")
    public ResponseEntity<?> signIn(@Valid @RequestBody SignInRequest signInRequest) {
        logger.info("Login attempt for email: {}", signInRequest.getEmail());
        return userRepository.findByEmail(signInRequest.getEmail())
                .map(user -> {
                    if (passwordEncoder.matches(signInRequest.getPassword(), user.getPasswordHash())) {
                        logger.info("Login successful for email: {}", signInRequest.getEmail());
                        String accessToken = jwtUtil.generateToken(user.getEmail());
                        RefreshToken refreshToken = refreshTokenService.createRefreshToken(user);
                        logger.debug("Generated tokens for user: {}", user.getEmail());
                        return ResponseEntity.ok(new AuthResponse(accessToken, refreshToken.getToken()));
                    } else {
                        logger.warn("Login failed (invalid credentials) for email: {}", signInRequest.getEmail());
                        return ResponseEntity.status(HttpStatus.UNAUTHORIZED).body("Invalid credentials");
                    }
                })
                .orElseGet(() -> {
                    logger.warn("Login failed (user not found) for email: {}", signInRequest.getEmail());
                    return ResponseEntity.status(HttpStatus.UNAUTHORIZED).body("Invalid credentials");
                });
    }

    @PostMapping("/refresh")
    public ResponseEntity<?> refreshToken(@RequestBody RefreshTokenRequest request) {
        logger.info("Refresh token request received");
        return refreshTokenService.refreshTokenAccess(request.getRefreshToken())
                .map(response -> {
                    logger.info("Refresh token valid, new access token issued");
                    return ResponseEntity.ok(response);
                })
                .orElseGet(() -> {
                    logger.warn("Refresh token invalid or expired");
                    return ResponseEntity.status(HttpStatus.UNAUTHORIZED).body(new AuthResponse(null, "Invalid or expired refresh token"));
                });
    }

    @PostMapping("logout")
    public ResponseEntity<?> logout(@RequestBody LogoutRequest logoutRequest) {
        logger.info("Logout request received");
        try {
            refreshTokenService.deleteByToken(logoutRequest.getRefreshToken());
            logger.info("Logout successful");
            return ResponseEntity.ok("Logged out successfully");
        } catch (Exception e) {
            logger.error("Logout failed", e);
            return ResponseEntity.status(HttpStatus.INTERNAL_SERVER_ERROR).body("Logout failed");
        }
    }
 }
