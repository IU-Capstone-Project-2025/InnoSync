package com.innosync.service;

import com.innosync.dto.auth.*;
import com.innosync.model.RefreshToken;
import com.innosync.model.User;
import com.innosync.repository.UserRepository;
import com.innosync.security.JwtUtil;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.http.HttpStatus;
import org.springframework.http.ResponseEntity;
import org.springframework.security.crypto.bcrypt.BCryptPasswordEncoder;
import org.springframework.stereotype.Service;
import org.springframework.web.bind.annotation.PostMapping;
import org.springframework.web.bind.annotation.RequestBody;

@Service
public class AuthService {

    @Autowired
    private UserRepository userRepository;

    @Autowired
    private JwtUtil jwtUtil ;

    @Autowired
    private RefreshTokenService refreshTokenService;

    @Autowired
    private BCryptPasswordEncoder passwordEncoder;

    public ResponseEntity<?> signUp(SignUpRequest signUpRequest) {
        String email = signUpRequest.getEmail();
        if (userRepository.findByEmail(email).isPresent()) {
            return ResponseEntity.status(HttpStatus.CONFLICT).body("Email is already registered!");
        }

        User user = new User(
                signUpRequest.getEmail(),
                signUpRequest.getFullName(),
                passwordEncoder.encode(signUpRequest.getPassword())

        );
        userRepository.save(user);

        String accessToken = jwtUtil.generateToken(user.getEmail());
        RefreshToken refreshToken = refreshTokenService.createRefreshToken(user);

        return ResponseEntity.ok(new AuthResponse(accessToken, refreshToken.getToken()));
    }

    public ResponseEntity<?> signIn(SignInRequest signInRequest) {
        return userRepository.findByEmail(signInRequest.getEmail())
                .map(user -> {
                    if (passwordEncoder.matches(signInRequest.getPassword(), user.getPasswordHash())) {
                        String accessToken = jwtUtil.generateToken(user.getEmail());
                        RefreshToken refreshToken = refreshTokenService.createRefreshToken(user);
                        return ResponseEntity.ok(new AuthResponse(accessToken, refreshToken.getToken()));
                    } else {
                        return ResponseEntity.status(HttpStatus.UNAUTHORIZED).body("Invalid credentials");
                    }
                })
                .orElseGet(() -> ResponseEntity.status(HttpStatus.UNAUTHORIZED).body("Invalid credentials"));
    }

    public ResponseEntity<?> refreshToken(RefreshTokenRequest request) {
        return refreshTokenService.refreshTokenAccess(request.getRefreshToken())
                .map(ResponseEntity::ok)
                .orElseGet(() -> ResponseEntity.status(HttpStatus.UNAUTHORIZED).body(new AuthResponse(null, "Invalid or expired refresh token")));
    }

    public ResponseEntity<?> logout(LogoutRequest logoutRequest) {
        try {
            refreshTokenService.deleteByToken(logoutRequest.getRefreshToken());
            return ResponseEntity.ok("Logged out successfully");
        } catch (Exception e) {
            e.printStackTrace();  // or use a logger`
            return ResponseEntity.status(HttpStatus.INTERNAL_SERVER_ERROR).body("Logout failed");
        }
    }
}

