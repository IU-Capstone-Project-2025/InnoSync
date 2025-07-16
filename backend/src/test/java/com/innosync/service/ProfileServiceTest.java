package com.innosync.service;

import com.innosync.dto.profile.ProfileRequest;
import com.innosync.dto.profile.ProfileResponse;
import com.innosync.dto.profile.WorkExperienceRequest;
import com.innosync.model.*;
import com.innosync.repository.*;
import org.junit.jupiter.api.BeforeEach;
import org.junit.jupiter.api.Test;
import org.junit.jupiter.api.extension.ExtendWith;
import org.mockito.InjectMocks;
import org.mockito.Mock;
import org.mockito.junit.jupiter.MockitoExtension;

import java.time.LocalDate;
import java.util.*;

import static org.assertj.core.api.Assertions.assertThat;
import static org.assertj.core.api.Assertions.assertThatThrownBy;
import static org.mockito.ArgumentMatchers.any;
import static org.mockito.Mockito.*;
import static org.mockito.Mockito.lenient;
import static org.mockito.Mockito.mock;

@ExtendWith(MockitoExtension.class)
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

    private User testUser;
    private Profile testProfile;
    private Technology javaTechnology;
    private Technology springTechnology;
    private ProfileRequest profileRequest;
    private WorkExperienceRequest workExperienceRequest;

    @BeforeEach
    void setUp() {
        testUser = new User("developer@example.com", "John Developer", "password");
        testUser.setId(1L);

        testProfile = new Profile();
        testProfile.setId(1L);
        testProfile.setUser(testUser);
        testProfile.setTelegram("@johndeveloper");
        testProfile.setGithub("johndeveloper");
        testProfile.setBio("Experienced developer");
        testProfile.setPosition("Senior Developer");
        testProfile.setEducation(Education.BACHELOR);
        testProfile.setExpertise("Backend Development");
        testProfile.setExpertiseLevel(ExpertiseLevel.SENIOR);

        javaTechnology = Technology.builder().id(1L).name("Java").build();
        springTechnology = Technology.builder().id(2L).name("Spring").build();
        testProfile.setTechnologies(Arrays.asList(javaTechnology, springTechnology));

        workExperienceRequest = new WorkExperienceRequest();
        workExperienceRequest.setStartDate(LocalDate.of(2020, 1, 1));
        workExperienceRequest.setEndDate(LocalDate.of(2023, 12, 31));
        workExperienceRequest.setPosition("Software Developer");
        workExperienceRequest.setCompany("Tech Company");
        workExperienceRequest.setDescription("Developed web applications");

        profileRequest = new ProfileRequest();
        profileRequest.setTelegram("@johndeveloper");
        profileRequest.setGithub("johndeveloper");
        profileRequest.setBio("Experienced developer");
        profileRequest.setPosition("Senior Developer");
        profileRequest.setEducation(Education.BACHELOR);
        profileRequest.setExpertise("Backend Development");
        profileRequest.setExpertiseLevel(ExpertiseLevel.SENIOR);
        profileRequest.setTechnologies(Arrays.asList("Java", "Spring"));
        profileRequest.setWorkExperience(Arrays.asList(workExperienceRequest));
    }

    @Test
    void createOrUpdateProfile_WithNewProfile_ShouldCreateProfile() {
        // Given
        String userEmail = "developer@example.com";
        
        when(userRepository.findByEmail(userEmail)).thenReturn(Optional.of(testUser));
        when(profileRepository.findByUser(testUser)).thenReturn(Optional.empty());
        when(technologyRepository.findByName("Java")).thenReturn(Optional.of(javaTechnology));
        when(technologyRepository.findByName("Spring")).thenReturn(Optional.of(springTechnology));
        when(profileRepository.save(any(Profile.class))).thenReturn(testProfile);
        when(workExperienceRepository.findByProfile(testProfile)).thenReturn(Collections.emptyList());

        // When
        ProfileResponse result = profileService.createOrUpdateProfile(userEmail, profileRequest);

        // Then
        assertThat(result).isNotNull();
        assertThat(result.getEmail()).isEqualTo("developer@example.com");
        assertThat(result.getFullName()).isEqualTo("John Developer");
        assertThat(result.getTelegram()).isEqualTo("@johndeveloper");
        assertThat(result.getGithub()).isEqualTo("johndeveloper");
        assertThat(result.getBio()).isEqualTo("Experienced developer");
        assertThat(result.getPosition()).isEqualTo("Senior Developer");
        assertThat(result.getEducation()).isEqualTo(Education.BACHELOR);
        assertThat(result.getExpertise()).isEqualTo("Backend Development");
        assertThat(result.getExpertiseLevel()).isEqualTo(ExpertiseLevel.SENIOR);
        assertThat(result.getTechnologies()).containsExactly("Java", "Spring");

        verify(userRepository).findByEmail(userEmail);
        verify(profileRepository).findByUser(testUser);
        verify(profileRepository).save(any(Profile.class));
        verify(workExperienceRepository).deleteByProfile(any(Profile.class));
        verify(workExperienceRepository).saveAll(any());
    }

    @Test
    void createOrUpdateProfile_WithExistingProfile_ShouldUpdateProfile() {
        // Given
        String userEmail = "developer@example.com";
        
        when(userRepository.findByEmail(userEmail)).thenReturn(Optional.of(testUser));
        when(profileRepository.findByUser(testUser)).thenReturn(Optional.of(testProfile));
        when(technologyRepository.findByName("Java")).thenReturn(Optional.of(javaTechnology));
        when(technologyRepository.findByName("Spring")).thenReturn(Optional.of(springTechnology));
        when(profileRepository.save(any(Profile.class))).thenReturn(testProfile);
        when(workExperienceRepository.findByProfile(testProfile)).thenReturn(Collections.emptyList());

        // When
        ProfileResponse result = profileService.createOrUpdateProfile(userEmail, profileRequest);

        // Then
        assertThat(result).isNotNull();
        assertThat(result.getId()).isEqualTo(1L);

        verify(userRepository).findByEmail(userEmail);
        verify(profileRepository).findByUser(testUser);
        verify(profileRepository).save(testProfile);
    }

    @Test
    void createOrUpdateProfile_WithNonExistentUser_ShouldThrowException() {
        // Given
        String userEmail = "nonexistent@example.com";
        
        when(userRepository.findByEmail(userEmail)).thenReturn(Optional.empty());

        // When & Then
        assertThatThrownBy(() -> profileService.createOrUpdateProfile(userEmail, profileRequest))
                .isInstanceOf(RuntimeException.class)
                .hasMessage("User not found");

        verify(userRepository).findByEmail(userEmail);
        verifyNoInteractions(profileRepository);
    }

    @Test
    void createOrUpdateProfile_WithNewTechnology_ShouldCreateTechnology() {
        // Given
        String userEmail = "developer@example.com";
        profileRequest.setTechnologies(Arrays.asList("Java", "NewTech"));
        
        Technology newTechnology = Technology.builder().id(3L).name("NewTech").build();
        
        when(userRepository.findByEmail(userEmail)).thenReturn(Optional.of(testUser));
        when(profileRepository.findByUser(testUser)).thenReturn(Optional.empty());
        when(technologyRepository.findByName("Java")).thenReturn(Optional.of(javaTechnology));
        when(technologyRepository.findByName("NewTech")).thenReturn(Optional.empty());
        when(technologyRepository.save(any(Technology.class))).thenReturn(newTechnology);
        when(profileRepository.save(any(Profile.class))).thenReturn(testProfile);
        when(workExperienceRepository.findByProfile(testProfile)).thenReturn(Collections.emptyList());

        // When
        ProfileResponse result = profileService.createOrUpdateProfile(userEmail, profileRequest);

        // Then
        assertThat(result).isNotNull();

        verify(technologyRepository).findByName("Java");
        verify(technologyRepository).findByName("NewTech");
        verify(technologyRepository).save(any(Technology.class));
    }

    @Test
    void createOrUpdateProfile_WithEmptyTechnologies_ShouldSetEmptyList() {
        // Given
        String userEmail = "developer@example.com";
        profileRequest.setTechnologies(Collections.emptyList());
        
        when(userRepository.findByEmail(userEmail)).thenReturn(Optional.of(testUser));
        when(profileRepository.findByUser(testUser)).thenReturn(Optional.empty());
        when(profileRepository.save(any(Profile.class))).thenReturn(testProfile);
        when(workExperienceRepository.findByProfile(testProfile)).thenReturn(Collections.emptyList());

        // When
        ProfileResponse result = profileService.createOrUpdateProfile(userEmail, profileRequest);

        // Then
        assertThat(result).isNotNull();

        verify(profileRepository).save(argThat(profile -> 
            profile.getTechnologies().isEmpty()
        ));
    }

    @Test
    void createOrUpdateProfile_WithNullTechnologies_ShouldSetEmptyList() {
        // Given
        String userEmail = "developer@example.com";
        profileRequest.setTechnologies(null);
        
        when(userRepository.findByEmail(userEmail)).thenReturn(Optional.of(testUser));
        when(profileRepository.findByUser(testUser)).thenReturn(Optional.empty());
        when(profileRepository.save(any(Profile.class))).thenReturn(testProfile);
        when(workExperienceRepository.findByProfile(testProfile)).thenReturn(Collections.emptyList());

        // When
        ProfileResponse result = profileService.createOrUpdateProfile(userEmail, profileRequest);

        // Then
        assertThat(result).isNotNull();

        verify(profileRepository).save(argThat(profile -> 
            profile.getTechnologies().isEmpty()
        ));
    }

    @Test
    void createOrUpdateProfile_WithWorkExperience_ShouldSaveWorkExperience() {
        // Given
        String userEmail = "developer@example.com";
        
        when(userRepository.findByEmail(userEmail)).thenReturn(Optional.of(testUser));
        when(profileRepository.findByUser(testUser)).thenReturn(Optional.empty());
        when(technologyRepository.findByName("Java")).thenReturn(Optional.of(javaTechnology));
        when(technologyRepository.findByName("Spring")).thenReturn(Optional.of(springTechnology));
        when(profileRepository.save(any(Profile.class))).thenReturn(testProfile);
        when(workExperienceRepository.findByProfile(testProfile)).thenReturn(Collections.emptyList());

        // When
        ProfileResponse result = profileService.createOrUpdateProfile(userEmail, profileRequest);

        // Then
        assertThat(result).isNotNull();

        verify(workExperienceRepository).deleteByProfile(testProfile);
        verify(workExperienceRepository).saveAll(argThat(experiences -> {
            List<WorkExperience> expList = new ArrayList<>();
            experiences.forEach(expList::add);
            return expList.size() == 1 && 
                   expList.get(0).getPosition().equals("Software Developer");
        }));
    }

    @Test
    void createOrUpdateProfile_WithEmptyWorkExperience_ShouldOnlyDeleteExisting() {
        // Given
        String userEmail = "developer@example.com";
        profileRequest.setWorkExperience(Collections.emptyList());
        
        when(userRepository.findByEmail(userEmail)).thenReturn(Optional.of(testUser));
        when(profileRepository.findByUser(testUser)).thenReturn(Optional.empty());
        when(technologyRepository.findByName("Java")).thenReturn(Optional.of(javaTechnology));
        when(technologyRepository.findByName("Spring")).thenReturn(Optional.of(springTechnology));
        when(profileRepository.save(any(Profile.class))).thenReturn(testProfile);
        when(workExperienceRepository.findByProfile(testProfile)).thenReturn(Collections.emptyList());

        // When
        ProfileResponse result = profileService.createOrUpdateProfile(userEmail, profileRequest);

        // Then
        assertThat(result).isNotNull();

        verify(workExperienceRepository).deleteByProfile(testProfile);
        verify(workExperienceRepository, never()).saveAll(any());
    }

    @Test
    void getMyProfile_WithValidUser_ShouldReturnProfile() {
        // Given
        String userEmail = "developer@example.com";
        
        when(userRepository.findByEmail(userEmail)).thenReturn(Optional.of(testUser));
        when(profileRepository.findByUser(testUser)).thenReturn(Optional.of(testProfile));
        when(workExperienceRepository.findByProfile(testProfile)).thenReturn(Collections.emptyList());

        // When
        ProfileResponse result = profileService.getMyProfile(userEmail);

        // Then
        assertThat(result).isNotNull();
        assertThat(result.getEmail()).isEqualTo("developer@example.com");
        assertThat(result.getFullName()).isEqualTo("John Developer");
        assertThat(result.getId()).isEqualTo(1L);

        verify(userRepository).findByEmail(userEmail);
        verify(profileRepository).findByUser(testUser);
        verify(workExperienceRepository).findByProfile(testProfile);
    }

    @Test
    void getMyProfile_WithNonExistentUser_ShouldThrowException() {
        // Given
        String userEmail = "nonexistent@example.com";
        
        when(userRepository.findByEmail(userEmail)).thenReturn(Optional.empty());

        // When & Then
        assertThatThrownBy(() -> profileService.getMyProfile(userEmail))
                .isInstanceOf(RuntimeException.class)
                .hasMessage("User not found");

        verify(userRepository).findByEmail(userEmail);
        verifyNoInteractions(profileRepository);
    }

    @Test
    void getMyProfile_WithNonExistentProfile_ShouldThrowException() {
        // Given
        String userEmail = "nonexistent@example.com";
        User userWithoutProfile = new User(userEmail, "User Without Profile", "password");
        userWithoutProfile.setId(999L);
        
        // Create a new ProfileService instance with fresh mocks for this test
        ProfileRepository freshProfileRepository = mock(ProfileRepository.class);
        UserRepository freshUserRepository = mock(UserRepository.class);
        WorkExperienceRepository freshWorkExperienceRepository = mock(WorkExperienceRepository.class);
        TechnologyRepository freshTechnologyRepository = mock(TechnologyRepository.class);
        
        ProfileService freshProfileService = new ProfileService(
                freshProfileRepository, freshUserRepository, 
                freshWorkExperienceRepository, freshTechnologyRepository
        );
        
        when(freshUserRepository.findByEmail(userEmail)).thenReturn(Optional.of(userWithoutProfile));
        when(freshProfileRepository.findByUser(userWithoutProfile)).thenReturn(Optional.empty());

        // When & Then
        assertThatThrownBy(() -> freshProfileService.getMyProfile(userEmail))
                .isInstanceOf(RuntimeException.class)
                .hasMessage("Profile not found");

        verify(freshUserRepository).findByEmail(userEmail);
        verify(freshProfileRepository).findByUser(userWithoutProfile);
    }

    @Test
    void getProfileByEmail_WithValidEmail_ShouldReturnProfile() {
        // Given
        String userEmail = "developer@example.com";
        
        when(userRepository.findByEmail(userEmail)).thenReturn(Optional.of(testUser));
        when(profileRepository.findByUser(testUser)).thenReturn(Optional.of(testProfile));

        // When
        Profile result = profileService.getProfileByEmail(userEmail);

        // Then
        assertThat(result).isNotNull();
        assertThat(result.getId()).isEqualTo(1L);
        assertThat(result.getUser()).isEqualTo(testUser);

        verify(userRepository).findByEmail(userEmail);
        verify(profileRepository).findByUser(testUser);
    }

    @Test
    void getUserByEmail_WithValidEmail_ShouldReturnUser() {
        // Given
        String userEmail = "developer@example.com";
        
        when(userRepository.findByEmail(userEmail)).thenReturn(Optional.of(testUser));

        // When
        User result = profileService.getUserByEmail(userEmail);

        // Then
        assertThat(result).isNotNull();
        assertThat(result.getId()).isEqualTo(1L);
        assertThat(result.getEmail()).isEqualTo("developer@example.com");

        verify(userRepository).findByEmail(userEmail);
    }

    @Test
    void save_WithValidProfile_ShouldSaveProfile() {
        // Given
        when(profileRepository.save(testProfile)).thenReturn(testProfile);

        // When
        profileService.save(testProfile);

        // Then
        verify(profileRepository).save(testProfile);
    }

    @Test
    void getProfileById_WithValidId_ShouldReturnProfile() {
        // Given
        Long profileId = 1L;
        
        when(profileRepository.findById(profileId)).thenReturn(Optional.of(testProfile));

        // When
        Profile result = profileService.getProfileById(profileId);

        // Then
        assertThat(result).isNotNull();
        assertThat(result.getId()).isEqualTo(1L);

        verify(profileRepository).findById(profileId);
    }

    @Test
    void getProfileById_WithNonExistentId_ShouldThrowException() {
        // Given
        Long profileId = 999L;
        
        when(profileRepository.findById(profileId)).thenReturn(Optional.empty());

        // When & Then
        assertThatThrownBy(() -> profileService.getProfileById(profileId))
                .isInstanceOf(RuntimeException.class)
                .hasMessage("Profile not found");

        verify(profileRepository).findById(profileId);
    }
} 