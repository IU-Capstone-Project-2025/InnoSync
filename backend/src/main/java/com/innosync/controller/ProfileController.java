package com.innosync.controller;

import com.innosync.dto.ProfileRequest;
import com.innosync.dto.ProfileResponse;
import com.innosync.service.ProfileService;
import io.swagger.v3.oas.annotations.Operation;
import io.swagger.v3.oas.annotations.tags.Tag;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.security.core.Authentication;
import org.springframework.web.bind.annotation.*;

@RestController
@RequestMapping("/api/profile")
@Tag(name = "Profile API", description = "API for user profile") // Swagger annotation
public class ProfileController {

    @Autowired
    private ProfileService profileService;

    @PostMapping
    @Operation(summary = "Create user profile")
    public ProfileResponse createOrUpdateProfile(@RequestBody ProfileRequest request, Authentication auth) {
        String email = auth.getName();
        return profileService.createOrUpdateProfile(email, request);
    }

    @PutMapping
    @Operation(summary = "Update user profile")
    public ProfileResponse updateProfile(@RequestBody ProfileRequest request, Authentication auth) {
        String email = auth.getName();
        return profileService.createOrUpdateProfile(email, request);
    }

    @GetMapping("/me")
    @Operation(summary = "Show personal profile")
    public ProfileResponse getMyProfile(Authentication auth) {
        String email = auth.getName();
        return profileService.getMyProfile(email);
    }
}
