package com.innosync.service;

import com.innosync.dto.auth.*;
import com.innosync.model.RefreshToken;
import com.innosync.model.User;
import com.innosync.repository.UserRepository;
import com.innosync.security.JwtUtil;
import org.junit.jupiter.api.BeforeEach;
import org.junit.jupiter.api.Test;
import org.mockito.InjectMocks;
import org.mockito.Mock;
import org.mockito.MockitoAnnotations;
import org.springframework.http.ResponseEntity;
import org.springframework.security.crypto.bcrypt.BCryptPasswordEncoder;

import java.time.Instant;
import java.util.Optional;

import static org.junit.jupiter.api.Assertions.*;
import static org.mockito.Mockito.*;

class AuthServiceTest {

    @Mock
    private UserRepository userRepository;

    @Mock
    private JwtUtil jwtUtil;

    @Mock
    private RefreshTokenService refreshTokenService;

    @Mock
    private BCryptPasswordEncoder passwordEncoder;

    @InjectMocks
    private AuthService authService;

    @BeforeEach
    void setUp() {
        MockitoAnnotations.openMocks(this);
    }

    @Test
    void signUp_Success() {
        SignUpRequest request = new SignUpRequest("user@example.com", "User Name", "password");

        when(userRepository.findByEmail("user@example.com")).thenReturn(Optional.empty());
        when(passwordEncoder.encode("password")).thenReturn("hashedPassword");
        when(jwtUtil.generateToken("user@example.com")).thenReturn("access-token");

        RefreshToken mockRefreshToken = new RefreshToken();
        mockRefreshToken.setToken("refresh-token");
        when(refreshTokenService.createRefreshToken(any(User.class))).thenReturn(mockRefreshToken);

        ResponseEntity<?> response = authService.signUp(request);

        assertEquals(200, response.getStatusCodeValue());
        AuthResponse body = (AuthResponse) response.getBody();
        assertEquals("access-token", body.getAccessToken());
        assertEquals("refresh-token", body.getRefreshToken());
        verify(userRepository).save(any(User.class));
    }

    @Test
    void signUp_EmailAlreadyExists() {
        SignUpRequest request = new SignUpRequest("user@example.com", "User Name", "password");

        when(userRepository.findByEmail("user@example.com")).thenReturn(Optional.of(new User()));

        ResponseEntity<?> response = authService.signUp(request);

        assertEquals(409, response.getStatusCodeValue());
        assertEquals("Email is already registered!", response.getBody());
    }

    @Test
    void signIn_Success() {
        SignInRequest request = new SignInRequest("user@example.com", "rawPassword");
        User user = new User("user@example.com", "User Name", "hashedPassword");

        when(userRepository.findByEmail("user@example.com")).thenReturn(Optional.of(user));
        when(passwordEncoder.matches("rawPassword", "hashedPassword")).thenReturn(true);
        when(jwtUtil.generateToken("user@example.com")).thenReturn("access-token");

        RefreshToken mockRefreshToken = new RefreshToken();
        mockRefreshToken.setToken("refresh-token");
        mockRefreshToken.setUser(user); // if needed
        mockRefreshToken.setExpiryDate(Instant.now().plusSeconds(3600)); // if needed

        when(refreshTokenService.createRefreshToken(user)).thenReturn(mockRefreshToken);

        ResponseEntity<?> response = authService.signIn(request);

        assertEquals(200, response.getStatusCodeValue());
        AuthResponse body = (AuthResponse) response.getBody();
        assertEquals("access-token", body.getAccessToken());
        assertEquals("refresh-token", body.getRefreshToken());
    }

    @Test
    void signIn_InvalidPassword() {
        SignInRequest request = new SignInRequest("user@example.com", "wrongPassword");
        User user = new User("user@example.com", "User Name", "correctHash");

        when(userRepository.findByEmail("user@example.com")).thenReturn(Optional.of(user));
        when(passwordEncoder.matches("wrongPassword", "correctHash")).thenReturn(false);

        ResponseEntity<?> response = authService.signIn(request);

        assertEquals(401, response.getStatusCodeValue());
        assertEquals("Invalid credentials", response.getBody());
    }

    @Test
    void signIn_UserNotFound() {
        SignInRequest request = new SignInRequest("nonexistent@example.com", "password");

        when(userRepository.findByEmail("nonexistent@example.com")).thenReturn(Optional.empty());

        ResponseEntity<?> response = authService.signIn(request);

        assertEquals(401, response.getStatusCodeValue());
        assertEquals("Invalid credentials", response.getBody());
    }

    @Test
    void refreshToken_Success() {
        RefreshTokenRequest request = new RefreshTokenRequest("valid-refresh-token");
        AuthResponse newTokens = new AuthResponse("new-access", "valid-refresh-token");

        when(refreshTokenService.refreshTokenAccess("valid-refresh-token")).thenReturn(Optional.of(newTokens));

        ResponseEntity<?> response = authService.refreshToken(request);

        assertEquals(200, response.getStatusCodeValue());
        assertEquals(newTokens, response.getBody());
    }

    @Test
    void refreshToken_InvalidOrExpired() {
        RefreshTokenRequest request = new RefreshTokenRequest("invalid-token");

        when(refreshTokenService.refreshTokenAccess("invalid-token")).thenReturn(Optional.empty());

        ResponseEntity<?> response = authService.refreshToken(request);

        assertEquals(401, response.getStatusCodeValue());
        AuthResponse body = (AuthResponse) response.getBody();
        assertNull(body.getAccessToken());
        assertEquals("Invalid or expired refresh token", body.getRefreshToken());
    }

    @Test
    void logout_Success() {
        LogoutRequest request = new LogoutRequest("refresh-token");

        doNothing().when(refreshTokenService).deleteByToken("refresh-token");

        ResponseEntity<?> response = authService.logout(request);

        assertEquals(200, response.getStatusCodeValue());
        assertEquals("Logged out successfully", response.getBody());
    }

    @Test
    void logout_Failure() {
        LogoutRequest request = new LogoutRequest("refresh-token");

        doThrow(new RuntimeException()).when(refreshTokenService).deleteByToken("refresh-token");

        ResponseEntity<?> response = authService.logout(request);

        assertEquals(500, response.getStatusCodeValue());
        assertEquals("Logout failed", response.getBody());
    }
}