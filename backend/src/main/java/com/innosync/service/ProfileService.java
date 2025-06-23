package com.innosync.service;

import com.innosync.dto.ProfileRequest;
import com.innosync.dto.ProfileResponse;
import com.innosync.dto.WorkExperienceResponse;
import com.innosync.model.Profile;
import com.innosync.model.User;
import com.innosync.model.WorkExperience;
import com.innosync.model.Technology;
import com.innosync.repository.ProfileRepository;
import com.innosync.repository.UserRepository;
import com.innosync.repository.WorkExperienceRepository;
import com.innosync.repository.TechnologyRepository;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;


import java.util.List;
import java.util.Set;
import java.util.stream.Collectors;

@Service
public class ProfileService {

    @Autowired
    private ProfileRepository profileRepository;

    @Autowired
    private UserRepository userRepository;

    @Autowired
    private WorkExperienceRepository workExperienceRepository;

    @Autowired
    private TechnologyRepository technologyRepository;


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

        workExperienceRepository.deleteByProfile(profile);

        // add new entries
        if (request.getWorkExperience() != null) {
            List<WorkExperience> experiences = request.getWorkExperience().stream()
                    .map(req -> WorkExperience.builder()
                            .profile(profile)
                            .startDate(req.getStartDate())
                            .endDate(req.getEndDate())
                            .position(req.getPosition())
                            .company(req.getCompany())
                            .description(req.getDescription())
                            .build())
                    .collect(Collectors.toList());

            workExperienceRepository.saveAll(experiences);
        }

        profile.getTechnologies().clear();

        if (request.getTechnologies() != null) {
            Set<Technology> techSet = request.getTechnologies().stream()
                    .map(name -> technologyRepository.findByNameIgnoreCase(name)
                            .orElseGet(() -> technologyRepository.save(
                                    Technology.builder().name(name).build())))
                    .collect(Collectors.toSet());

            profile.setTechnologies(techSet);
        }


        ProfileResponse response = mapToResponse(profile, user);
        return response;
    }

    public ProfileResponse getMyProfile(String email) {
        User user = userRepository.findByEmail(email)
                .orElseThrow(() -> new RuntimeException("User not found"));

        Profile profile = profileRepository.findByUser(user)
                .orElseThrow(() -> new RuntimeException("Profile not found"));

        return mapToResponse(profile, user);
    }

    private ProfileResponse mapToResponse(Profile profile, User user) {
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

        // Load work experience and mape it into DTO
        List<WorkExperienceResponse> experienceResponses = workExperienceRepository.findByProfile(profile).stream()
                .map(exp -> {
                    WorkExperienceResponse res = new WorkExperienceResponse();
                    res.setStartDate(exp.getStartDate());
                    res.setEndDate(exp.getEndDate());
                    res.setPosition(exp.getPosition());
                    res.setCompany(exp.getCompany());
                    res.setDescription(exp.getDescription());
                    return res;
                })
                .collect(Collectors.toList());

        response.setWorkExperience(experienceResponses);

        response.setTechnologies(
                profile.getTechnologies().stream()
                        .map(Technology::getName)
                        .toList()
        );

        return response;
    }
}
