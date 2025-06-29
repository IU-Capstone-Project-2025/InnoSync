package com.innosync.dto.auth;

import lombok.AllArgsConstructor;
import lombok.Data;

@Data
@AllArgsConstructor
public class LogoutRequest {
    private String refreshToken;
}
