package com.innosync.dto.auth;

import lombok.Data;

@Data
public class LogoutRequest {
    private String refreshToken;
}
