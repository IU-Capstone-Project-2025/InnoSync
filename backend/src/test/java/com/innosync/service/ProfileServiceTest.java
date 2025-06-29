package com.innosync.service;

import com.innosync.dto.profile.ProfileRequest;
import com.innosync.dto.profile.ProfileResponse;
import com.innosync.dto.profile.WorkExperienceRequest;
import com.innosync.dto.profile.WorkExperienceResponse;
import com.innosync.model.Profile;
import com.innosync.model.Technology;
import com.innosync.model.User;
import com.innosync.model.WorkExperience;
import com.innosync.model.Education;
import com.innosync.model.ExpertiseLevel;
import com.innosync.repository.ProfileRepository;
import com.innosync.repository.TechnologyRepository;
import com.innosync.repository.UserRepository;
import com.innosync.repository.WorkExperienceRepository;
import org.junit.jupiter.api.BeforeEach;
import org.junit.jupiter.api.Test;
import org.mockito.ArgumentCaptor;
import org.mockito.InjectMocks;
import org.mockito.Mock;
import org.mockito.MockitoAnnotations;

import java.time.LocalDate;
import java.util.*;

import static org.junit.jupiter.api.Assertions.*;
import static org.mockito.Mockito.*;

class ProfileServiceTest {

    @Mock
    private ProfileRepository profileRepository;

    @Mock
    private UserRepository userRepository;

    @Mock
    private WorkExperienceRepository workExperienceRepository;

    @Mock
    private TechnologyRepository technologyRepository;

    @InjectMocks
    private ProfileService profileService;

    private User user;

    @BeforeEach
    void setUp() {
        MockitoAnnotations.openMocks(this);
        user = new User();
        user.setEmail("user@example.com");
        user.setFullName("Test User");
    }

    @Test
    void createOrUpdateProfile_CreatesNewProfile_WithTechnologiesAndWorkExperience() {
        // Arrange
        WorkExperienceRequest expReq = new WorkExperienceRequest(
                LocalDate.of(2020,1,1),
                LocalDate.of(2021,1,1),
                "Dev",
                "Company",
                "Desc"
        );
        List<WorkExperienceRequest> expReqs = Collections.singletonList(expReq);
        List<String> techNames = Arrays.asList("Java", "Python");

        ProfileRequest request = new ProfileRequest(
                "@test",
                "github/test",
                "bio",
                "Developer",
                Education.BACHELOR,
                "Backend",
                ExpertiseLevel.SENIOR,
                "resume.pdf",
                expReqs,
                techNames
        );

        when(userRepository.findByEmail(user.getEmail())).thenReturn(Optional.of(user));
        when(profileRepository.findByUser(user)).thenReturn(Optional.empty());

        when(technologyRepository.findByName("Java")).thenReturn(Optional.of(new Technology("Java")));
        when(technologyRepository.findByName("Python")).thenReturn(Optional.empty());
        when(technologyRepository.save(any(Technology.class))).thenAnswer(i -> i.getArgument(0));

        // Stub save to return the passed profile
        when(profileRepository.save(any(Profile.class))).thenAnswer(i -> i.getArgument(0));
        doNothing().when(workExperienceRepository).deleteByProfile(any(Profile.class));
        when(workExperienceRepository.saveAll(anyList())).thenAnswer(i -> i.getArgument(0));

        // Act
        ProfileResponse response = profileService.createOrUpdateProfile(user.getEmail(), request);

        // Assert
        assertEquals(user.getEmail(), response.getEmail());
        assertEquals(user.getFullName(), response.getFullName());
        assertEquals(request.getTelegram(), response.getTelegram());
        assertEquals(techNames, response.getTechnologies());

        verify(technologyRepository, times(2)).findByName(anyString());
        verify(technologyRepository).save(any(Technology.class));

        ArgumentCaptor<Profile> profileCaptor = ArgumentCaptor.forClass(Profile.class);
        verify(profileRepository).save(profileCaptor.capture());
        Profile savedProfile = profileCaptor.getValue();

        verify(workExperienceRepository).deleteByProfile(savedProfile);

        ArgumentCaptor<List<WorkExperience>> expCaptor = ArgumentCaptor.forClass(List.class);
        verify(workExperienceRepository).saveAll(expCaptor.capture());
        List<WorkExperience> savedExps = expCaptor.getValue();
        assertEquals(1, savedExps.size());
        WorkExperience savedExp = savedExps.get(0);
        assertEquals(expReq.getCompany(), savedExp.getCompany());
        assertEquals(expReq.getPosition(), savedExp.getPosition());
    }

    @Test
    void getMyProfile_ReturnsProfileResponse() {
        // Arrange
        Profile profile = new Profile();
        profile.setTelegram("@test");
        profile.setGithub("git");
        profile.setBio("bio");
        profile.setPosition("Dev");
        profile.setEducation(Education.MASTER);
        profile.setExpertise("Backend");
        profile.setExpertiseLevel(ExpertiseLevel.MID);
        profile.setResume("res.pdf");
        profile.setTechnologies(
                Arrays.asList(new Technology("Java"), new Technology("Python"))
        );
        WorkExperience we = new WorkExperience();
        we.setProfile(profile);
        we.setCompany("C");
        we.setPosition("P");
        we.setStartDate(LocalDate.of(2019,1,1));
        we.setEndDate(LocalDate.of(2020,1,1));
        we.setDescription("D");

        when(userRepository.findByEmail(user.getEmail())).thenReturn(Optional.of(user));
        when(profileRepository.findByUser(user)).thenReturn(Optional.of(profile));
        when(workExperienceRepository.findByProfile(profile)).thenReturn(Collections.singletonList(we));

        // Act
        ProfileResponse resp = profileService.getMyProfile(user.getEmail());

        // Assert
        assertEquals(user.getEmail(), resp.getEmail());
        assertEquals(1, resp.getWorkExperience().size());
        WorkExperienceResponse wer = resp.getWorkExperience().get(0);
        assertEquals(we.getCompany(), wer.getCompany());
        assertTrue(resp.getTechnologies().containsAll(Arrays.asList("Java","Python")));
    }
}
