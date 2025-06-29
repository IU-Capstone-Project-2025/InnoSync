package com.innosync.service;

import com.innosync.dto.auth.AuthResponse;
import com.innosync.model.RefreshToken;
import com.innosync.model.User;
import com.innosync.repository.RefreshTokenRepository;
import com.innosync.repository.UserRepository;
import com.innosync.security.JwtUtil;
import org.junit.jupiter.api.BeforeEach;
import org.junit.jupiter.api.Test;
import org.mockito.ArgumentCaptor;
import org.mockito.InjectMocks;
import org.mockito.Mock;
import org.mockito.MockitoAnnotations;
import org.springframework.transaction.annotation.Transactional;

import java.time.Instant;
import java.time.temporal.ChronoUnit;
import java.util.Optional;
import java.util.UUID;

import static org.junit.jupiter.api.Assertions.*;
import static org.mockito.Mockito.*;

class RefreshTokenServiceTest {

    @Mock
    private UserRepository userRepository;
    @Mock
    private RefreshTokenRepository refreshTokenRepository;
    @Mock
    private JwtUtil jwtUtil;

    @InjectMocks
    private RefreshTokenService refreshTokenService;

    @BeforeEach
    void setUp() {
        MockitoAnnotations.openMocks(this);
    }

    @Test
    void createRefreshToken_SetsPropertiesAndSaves() {
        User user = new User();
        user.setEmail("test@example.com");

        ArgumentCaptor<RefreshToken> captor = ArgumentCaptor.forClass(RefreshToken.class);
        when(refreshTokenRepository.save(any(RefreshToken.class))).thenAnswer(i -> i.getArgument(0));

        RefreshToken token = refreshTokenService.createRefreshToken(user);

        verify(refreshTokenRepository).save(captor.capture());
        RefreshToken saved = captor.getValue();
        assertEquals(user, saved.getUser());
        assertNotNull(saved.getToken());
        assertDoesNotThrow(() -> UUID.fromString(saved.getToken()));
        Instant now = Instant.now();
        assertTrue(saved.getExpiryDate().isAfter(now));
        assertTrue(saved.getExpiryDate().isBefore(now.plus(7, ChronoUnit.DAYS)));
    }

    @Test
    void verifyExpiration_NotExpired_ReturnsToken() {
        RefreshToken token = new RefreshToken();
        token.setExpiryDate(Instant.now().plusSeconds(60));
        token.setToken("token");

        Optional<RefreshToken> result = refreshTokenService.verifyExpiration(token);

        assertTrue(result.isPresent());
        assertEquals(token, result.get());
    }

    @Test
    void verifyExpiration_Expired_DeletesAndReturnsEmpty() {
        RefreshToken token = new RefreshToken();
        token.setExpiryDate(Instant.now().minusSeconds(60));
        token.setToken("expired-token");

        Optional<RefreshToken> result = refreshTokenService.verifyExpiration(token);

        assertFalse(result.isPresent());
        verify(refreshTokenRepository).deleteByToken("expired-token");
    }

    @Test
    void deleteToken_InvokesRepository() {
        User user = new User();
        refreshTokenService.deleteToken(user);

        verify(refreshTokenRepository).deleteByUser(user);
    }

    @Test
    void deleteByToken_InvokesRepository() {
        String tokenValue = "some-token";
        refreshTokenService.deleteByToken(tokenValue);

        verify(refreshTokenRepository).deleteByToken(tokenValue);
    }

    @Test
    void refreshTokenAccess_ValidToken_ReturnsAuthResponse() {
        String tokenValue = "valid-token";
        User user = new User();
        user.setEmail("user@example.com");
        RefreshToken refreshToken = new RefreshToken();
        refreshToken.setUser(user);
        refreshToken.setToken(tokenValue);
        refreshToken.setExpiryDate(Instant.now().plusSeconds(60));

        when(refreshTokenRepository.findByToken(tokenValue)).thenReturn(Optional.of(refreshToken));
        when(jwtUtil.generateToken("user@example.com")).thenReturn("new-access-token");

        Optional<AuthResponse> response = refreshTokenService.refreshTokenAccess(tokenValue);

        assertTrue(response.isPresent());
        AuthResponse auth = response.get();
        assertEquals("new-access-token", auth.getAccessToken());
        assertEquals(tokenValue, auth.getRefreshToken());
    }

    @Test
    void refreshTokenAccess_TokenNotFound_ReturnsEmpty() {
        when(refreshTokenRepository.findByToken("missing")).thenReturn(Optional.empty());

        Optional<AuthResponse> response = refreshTokenService.refreshTokenAccess("missing");

        assertTrue(response.isEmpty());
    }

    @Test
    void refreshTokenAccess_ExpiredToken_ReturnsEmpty() {
        String tokenValue = "expired-token";
        User user = new User();
        user.setEmail("user@example.com");
        RefreshToken refreshToken = new RefreshToken();
        refreshToken.setUser(user);
        refreshToken.setToken(tokenValue);
        refreshToken.setExpiryDate(Instant.now().minusSeconds(60));

        when(refreshTokenRepository.findByToken(tokenValue)).thenReturn(Optional.of(refreshToken));

        Optional<AuthResponse> response = refreshTokenService.refreshTokenAccess(tokenValue);

        assertTrue(response.isEmpty());
        verify(refreshTokenRepository).deleteByToken(tokenValue);
    }
}
