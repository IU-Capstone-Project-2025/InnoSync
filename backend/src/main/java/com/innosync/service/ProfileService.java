package com.innosync.service;

import com.innosync.dto.profile.WorkExperienceResponse;
import com.innosync.dto.profile.ProfileRequest;
import com.innosync.dto.profile.ProfileResponse;
import com.innosync.model.Profile;
import com.innosync.model.Technology;
import com.innosync.model.User;
import com.innosync.model.WorkExperience;
import com.innosync.repository.ProfileRepository;
import com.innosync.repository.TechnologyRepository;
import com.innosync.repository.UserRepository;
import com.innosync.repository.WorkExperienceRepository;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;
import org.springframework.transaction.annotation.Transactional;

import java.util.*;
import java.util.stream.Collectors;

import java.util.List;


@Service
public class ProfileService {

    private static final Logger logger = LoggerFactory.getLogger(ProfileService.class);

    private final ProfileRepository profileRepository;
    private final UserRepository userRepository;
    private final WorkExperienceRepository workExperienceRepository;
    private final TechnologyRepository technologyRepository;

    public ProfileService(ProfileRepository profileRepository, UserRepository userRepository, 
                         WorkExperienceRepository workExperienceRepository, TechnologyRepository technologyRepository) {
        this.profileRepository = profileRepository;
        this.userRepository = userRepository;
        this.workExperienceRepository = workExperienceRepository;
        this.technologyRepository = technologyRepository;
    }

    @Transactional
    public ProfileResponse createOrUpdateProfile(String email, ProfileRequest request) {
        logger.info("Creating or updating profile for email: {}", email);
        try {
            final User user = userRepository.findByEmail(email)
                    .orElseThrow(() -> new RuntimeException("User not found"));
            logger.debug("User found for profile update: {}", user.getEmail());
            Profile profile = profileRepository.findByUser(user).orElse(new Profile());
            profile.setUser(user);
            profile.setTelegram(request.getTelegram());
            profile.setGithub(request.getGithub());
            profile.setBio(request.getBio());
            profile.setPosition(request.getPosition());
            profile.setEducation(request.getEducation());
            profile.setExpertise(request.getExpertise());
            profile.setExpertiseLevel(request.getExpertiseLevel());
            profile.setExperienceYears(request.getExperienceYears());
            if (request.getTechnologies() != null && !request.getTechnologies().isEmpty()) {
                List<Technology> technologies = request.getTechnologies().stream()
                        .map(name -> technologyRepository.findByName(name)
                                .orElseGet(() -> technologyRepository.save(Technology.builder().name(name).build())))
                        .collect(Collectors.toList());
                profile.setTechnologies(technologies);
                logger.debug("Technologies set for profile: {}", technologies);
            } else {
                profile.setTechnologies(Collections.emptyList());
            }
            final Profile savedProfile = profileRepository.save(profile);
            logger.info("Profile saved for user: {}", user.getEmail());
            workExperienceRepository.deleteByProfile(savedProfile);
            logger.debug("Deleted old work experiences for profile: {}", savedProfile.getId());
            if (request.getWorkExperience() != null && !request.getWorkExperience().isEmpty()) {
                List<WorkExperience> experiences = request.getWorkExperience().stream()
                        .map(req -> WorkExperience.builder()
                                .profile(savedProfile)
                                .startDate(req.getStartDate())
                                .endDate(req.getEndDate())
                                .position(req.getPosition())
                                .company(req.getCompany())
                                .description(req.getDescription())
                                .build())
                        .collect(Collectors.toList());

                workExperienceRepository.saveAll(experiences);
                logger.debug("Saved new work experiences for profile: {}", savedProfile.getId());
            }
            return mapToResponse(savedProfile, user);
        } catch (Exception e) {
            logger.error("Failed to create or update profile for email: {}", email, e);
            throw e;
        }
    }

    public ProfileResponse getMyProfile(String email) {
        final User user = userRepository.findByEmail(email)
                .orElseThrow(() -> new RuntimeException("User not found"));

        final Profile profile = profileRepository.findByUser(user)
                .orElseThrow(() -> new RuntimeException("Profile not found"));

        return mapToResponse(profile, user);
    }

    private ProfileResponse mapToResponse(Profile profile, User user) {
        final ProfileResponse response = new ProfileResponse();
        response.setEmail(user.getEmail());
        response.setFullName(user.getFullName());
        response.setTelegram(profile.getTelegram());
        response.setGithub(profile.getGithub());
        response.setBio(profile.getBio());
        response.setPosition(profile.getPosition());
        response.setEducation(profile.getEducation());
        response.setExpertise(profile.getExpertise());
        response.setExpertiseLevel(profile.getExpertiseLevel());
        response.setExperienceYears(profile.getExperienceYears());
        response.setId(profile.getId());

        // Опыт работы
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

        // Технологии
        List<String> technologyNames = profile.getTechnologies().stream()
                .map(Technology::getName)
                .collect(Collectors.toList());
        response.setTechnologies(technologyNames);

        return response;
    }


    public Profile getProfileByEmail(String email) {
        logger.debug("Getting profile by email: {}", email);
        return profileRepository.findByUser(userRepository.findByEmail(email)
                .orElseThrow(() -> new RuntimeException("User not found")))
                .orElseThrow(() -> new RuntimeException("Profile not found"));
    }

    public User getUserByEmail(String email) {
        logger.debug("Getting user by email: {}", email);
        return userRepository.findByEmail(email).orElseThrow(() -> new RuntimeException("User not found"));
    }

    public void save(Profile profile) {
        logger.debug("Saving profile for user id: {}", profile.getUser().getId());
        profileRepository.save(profile);
    }

    public Profile getProfileById(Long id) {
        logger.debug("Getting profile by id: {}", id);
        return profileRepository.findById(id)
                .orElseThrow(() -> new RuntimeException("Profile not found"));
    }

    public List<ProfileResponse> getAllProfiles() {
        logger.info("Getting all profiles");
        List<Profile> profiles = profileRepository.findAll();
        logger.debug("Found {} profiles", profiles.size());
        return profiles.stream()
                .map(profile -> mapToResponse(profile, profile.getUser()))
                .collect(Collectors.toList());
    }
}
