package com.innosync.service;

import com.innosync.dto.ProfileRequest;
import com.innosync.dto.ProfileResponse;
import com.innosync.model.Profile;
import com.innosync.model.User;
import com.innosync.repository.ProfileRepository;
import com.innosync.repository.UserRepository;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;

@Service
public class ProfileService {

    @Autowired
    private ProfileRepository profileRepository;

    @Autowired
    private UserRepository userRepository;

    public ProfileResponse createOrUpdateProfile(String email, ProfileRequest request) {
        User user = userRepository.findByEmail(email)
                .orElseThrow(() -> new RuntimeException("User not found"));

        Profile profile = profileRepository.findByUser(user).orElse(new Profile());
        profile.setUser(user);
        profile.setTelegram(request.getTelegram());
        profile.setGithub(request.getGithub());
        profile.setBio(request.getBio());
        profile.setPosition(request.getPosition());
        profile.setEducation(request.getEducation());
        profile.setExpertise(request.getExpertise());
        profile.setExpertiseLevel(request.getExpertiseLevel());
        profile.setResume(request.getResume());
        profileRepository.save(profile);

        ProfileResponse response = new ProfileResponse();
        response.setEmail(user.getEmail());
        response.setFullName(user.getFullName());
        response.setTelegram(profile.getTelegram());
        response.setGithub(profile.getGithub());
        response.setBio(profile.getBio());
        response.setPosition(profile.getPosition());
        response.setEducation(profile.getEducation());
        response.setExpertise(profile.getExpertise());
        response.setExpertiseLevel(profile.getExpertiseLevel());
        response.setResume(profile.getResume());


        return response;
    }

    public ProfileResponse getMyProfile(String email) {
        User user = userRepository.findByEmail(email)
                .orElseThrow(() -> new RuntimeException("User not found"));

        Profile profile = profileRepository.findByUser(user)
                .orElseThrow(() -> new RuntimeException("Profile not found"));

        ProfileResponse response = new ProfileResponse();
        response.setEmail(user.getEmail());
        response.setFullName(user.getFullName());
        response.setTelegram(profile.getTelegram());
        response.setGithub(profile.getGithub());
        response.setBio(profile.getBio());
        response.setPosition(profile.getPosition());
        response.setEducation(profile.getEducation());
        response.setExpertise(profile.getExpertise());
        response.setExpertiseLevel(profile.getExpertiseLevel());
        response.setResume(profile.getResume());

        return response;
    }
}
