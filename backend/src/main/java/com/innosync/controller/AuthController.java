package com.innosync.controller;

import com.innosync.dto.JwtResponse;
import com.innosync.dto.SignInRequest;
import com.innosync.dto.SignUpRequest;
import com.innosync.model.User;
import com.innosync.repository.UserRepository;
import com.innosync.security.JwtUtil;
import io.swagger.v3.oas.annotations.Operation;
import io.swagger.v3.oas.annotations.tags.Tag;
import jakarta.validation.Valid;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.http.HttpStatus;
import org.springframework.http.ResponseEntity;
import org.springframework.security.crypto.bcrypt.BCryptPasswordEncoder;
import org.springframework.security.crypto.password.PasswordEncoder;
import org.springframework.web.bind.annotation.PostMapping;
import org.springframework.web.bind.annotation.RequestBody;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RestController;

@RestController
@RequestMapping("/api/auth")
@Tag(name = "Authentication API", description = "API for user authentication") // Swagger annotation
public class AuthController {

    @Autowired
    private UserRepository userRepository;

    @Autowired
    private JwtUtil jwtUtil ;

    private BCryptPasswordEncoder passwordEncoder = new BCryptPasswordEncoder();

    @PostMapping("/signup")
    @Operation(summary = "Sign in user")
    public ResponseEntity<?> signUp(@Valid @RequestBody SignInRequest signUpRequest) {
        String email = signUpRequest.getEmail();
        if (userRepository.findByEmail(email).isPresent()) {
            return ResponseEntity.status(HttpStatus.CONFLICT).body("Email is already registered!");
        }

        User user = new User(
                signUpRequest.getEmail(),
                passwordEncoder.encode(signUpRequest.getPassword())
        );
        userRepository.save(user);
        return ResponseEntity.ok("User registered successfully");
    }

    @PostMapping("/login")
    @Operation(summary = "Log in User")
    public ResponseEntity<?> singIn(@Valid @RequestBody SignUpRequest signInRequest) {
        return userRepository.findByEmail(signInRequest.getEmail())
                .map(user -> {
                    if (passwordEncoder.matches(signInRequest.getPassword(), user.getPasswordHash())) {
                        String token = jwtUtil.generateToken(user.getEmail());
                        return ResponseEntity.ok(new JwtResponse(token));
                    } else {
                        return ResponseEntity.status(HttpStatus.UNAUTHORIZED).body("Invalid credentials");
                    }
                })
                .orElseGet(() -> ResponseEntity.status(HttpStatus.UNAUTHORIZED).body("Invalid credentials"));
    }
 }
