package com.innosync.service;

import com.innosync.dto.project.ProjectRequest;
import com.innosync.dto.project.ProjectResponse;
import com.innosync.model.Project;
import com.innosync.model.ProjectType;
import com.innosync.model.TeamSize;
import com.innosync.model.User;
import com.innosync.model.Team;
import com.innosync.model.ChatRoom;
import com.innosync.repository.ProjectRepository;
import com.innosync.repository.ProjectRoleRepository;
import com.innosync.repository.UserRepository;
import com.innosync.repository.TeamRepository;
import com.innosync.repository.ChatRoomRepository;
import org.junit.jupiter.api.BeforeEach;
import org.junit.jupiter.api.Test;
import org.junit.jupiter.api.extension.ExtendWith;
import org.mockito.InjectMocks;
import org.mockito.Mock;
import org.mockito.junit.jupiter.MockitoExtension;
import org.springframework.http.HttpStatus;
import org.springframework.web.server.ResponseStatusException;

import java.time.LocalDateTime;
import java.util.Arrays;
import java.util.Collections;
import java.util.List;
import java.util.Optional;

import static org.assertj.core.api.Assertions.assertThat;
import static org.assertj.core.api.Assertions.assertThatThrownBy;
import static org.mockito.ArgumentMatchers.any;
import static org.mockito.Mockito.verify;
import static org.mockito.Mockito.when;

@ExtendWith(MockitoExtension.class)
class ProjectServiceTest {

    @Mock
    private ProjectRepository projectRepository;
    
    @Mock
    private UserRepository userRepository;
    
    @Mock
    private ProjectRoleRepository projectRoleRepository;

    @Mock
    private TeamRepository teamRepository;

    @Mock
    private ChatRoomRepository chatRoomRepository;

    @InjectMocks
    private ProjectService projectService;

    private User testUser;
    private Project testProject;
    private ProjectRequest projectRequest;
    private Team testTeam;

    @BeforeEach
    void setUp() {
        testUser = new User("recruiter@example.com", "Test Recruiter", "password");
        testUser.setId(1L);

        testProject = Project.builder()
                .id(1L)
                .title("Test Project")
                .description("Test Description")
                .projectType(ProjectType.FREELANCE)
                .teamSize(TeamSize.OneThree)
                .recruiter(testUser)
                .createdAt(LocalDateTime.now())
                .updatedAt(LocalDateTime.now())
                .build();

        testTeam = Team.builder()
                .id(1L)
                .name("Test Project Team")
                .project(testProject)
                .build();

        projectRequest = new ProjectRequest();
        projectRequest.setTitle("Test Project");
        projectRequest.setDescription("Test Description");
        projectRequest.setProjectType(ProjectType.FREELANCE);
        projectRequest.setTeamSize(TeamSize.OneThree);
    }

    @Test
    void createProject_WithValidData_ShouldReturnProjectResponse() {
        // Given
        String recruiterEmail = "recruiter@example.com";
        when(userRepository.findByEmail(recruiterEmail)).thenReturn(Optional.of(testUser));
        when(projectRepository.save(any(Project.class))).thenReturn(testProject);
        when(teamRepository.save(any(Team.class))).thenReturn(testTeam);
        when(chatRoomRepository.save(any(ChatRoom.class))).thenReturn(any(ChatRoom.class));

        // When
        ProjectResponse result = projectService.createProject(projectRequest, recruiterEmail);

        // Then
        assertThat(result).isNotNull();
        assertThat(result.getId()).isEqualTo(1L);
        assertThat(result.getTitle()).isEqualTo("Test Project");
        assertThat(result.getDescription()).isEqualTo("Test Description");
        assertThat(result.getProjectType()).isEqualTo(ProjectType.FREELANCE);
        assertThat(result.getTeamSize()).isEqualTo(TeamSize.OneThree);
        
        verify(userRepository).findByEmail(recruiterEmail);
        verify(projectRepository).save(any(Project.class));
    }

    @Test
    void createProject_WithInvalidUser_ShouldThrowException() {
        // Given
        String invalidEmail = "invalid@example.com";
        when(userRepository.findByEmail(invalidEmail)).thenReturn(Optional.empty());

        // When & Then
        assertThatThrownBy(() -> projectService.createProject(projectRequest, invalidEmail))
                .isInstanceOf(RuntimeException.class)
                .hasMessage("User not found");
        
        verify(userRepository).findByEmail(invalidEmail);
    }

    @Test
    void createProject_WithNullRequest_ShouldThrowException() {
        // Given
        String recruiterEmail = "recruiter@example.com";
        when(userRepository.findByEmail(recruiterEmail)).thenReturn(Optional.of(testUser));

        // When & Then
        assertThatThrownBy(() -> projectService.createProject(null, recruiterEmail))
                .isInstanceOf(NullPointerException.class);
    }

    @Test
    void getMyProjects_WithValidUser_ShouldReturnProjectList() {
        // Given
        String recruiterEmail = "recruiter@example.com";
        List<Project> projects = Arrays.asList(testProject);
        
        when(userRepository.findByEmail(recruiterEmail)).thenReturn(Optional.of(testUser));
        when(projectRepository.findByRecruiter(testUser)).thenReturn(projects);

        // When
        List<ProjectResponse> result = projectService.getMyProjects(recruiterEmail);

        // Then
        assertThat(result).hasSize(1);
        assertThat(result.get(0).getTitle()).isEqualTo("Test Project");
        assertThat(result.get(0).getId()).isEqualTo(1L);
        
        verify(userRepository).findByEmail(recruiterEmail);
        verify(projectRepository).findByRecruiter(testUser);
    }

    @Test
    void getMyProjects_WithInvalidUser_ShouldThrowException() {
        // Given
        String invalidEmail = "invalid@example.com";
        when(userRepository.findByEmail(invalidEmail)).thenReturn(Optional.empty());

        // When & Then
        assertThatThrownBy(() -> projectService.getMyProjects(invalidEmail))
                .isInstanceOf(RuntimeException.class)
                .hasMessage("User not found");
        
        verify(userRepository).findByEmail(invalidEmail);
    }

    @Test
    void getMyProjects_WithNoProjects_ShouldReturnEmptyList() {
        // Given
        String recruiterEmail = "recruiter@example.com";
        
        when(userRepository.findByEmail(recruiterEmail)).thenReturn(Optional.of(testUser));
        when(projectRepository.findByRecruiter(testUser)).thenReturn(Collections.emptyList());

        // When
        List<ProjectResponse> result = projectService.getMyProjects(recruiterEmail);

        // Then
        assertThat(result).isEmpty();
        
        verify(userRepository).findByEmail(recruiterEmail);
        verify(projectRepository).findByRecruiter(testUser);
    }

    @Test
    void getProject_WithValidId_ShouldReturnProject() {
        // Given
        Long projectId = 1L;
        when(projectRepository.findById(projectId)).thenReturn(Optional.of(testProject));

        // When
        ProjectResponse result = projectService.getProject(projectId);

        // Then
        assertThat(result).isNotNull();
        assertThat(result.getId()).isEqualTo(projectId);
        assertThat(result.getTitle()).isEqualTo("Test Project");
        assertThat(result.getDescription()).isEqualTo("Test Description");
        
        verify(projectRepository).findById(projectId);
    }

    @Test
    void getProject_WithInvalidId_ShouldThrowException() {
        // Given
        Long invalidId = 999L;
        when(projectRepository.findById(invalidId)).thenReturn(Optional.empty());

        // When & Then
        assertThatThrownBy(() -> projectService.getProject(invalidId))
                .isInstanceOf(ResponseStatusException.class)
                .hasMessageContaining("Project not found")
                .hasFieldOrPropertyWithValue("status", HttpStatus.NOT_FOUND);
        
        verify(projectRepository).findById(invalidId);
    }

    @Test
    void getProject_WithNullId_ShouldCallRepository() {
        // Given
        when(projectRepository.findById(null)).thenReturn(Optional.empty());

        // When & Then
        assertThatThrownBy(() -> projectService.getProject(null))
                .isInstanceOf(ResponseStatusException.class)
                .hasMessageContaining("Project not found")
                .hasFieldOrPropertyWithValue("status", HttpStatus.NOT_FOUND);
        
        verify(projectRepository).findById(null);
    }

    @Test
    void createProject_WithMobileAppType_ShouldCreateSuccessfully() {
        // Given
        projectRequest.setProjectType(ProjectType.FREELANCE);
        projectRequest.setTeamSize(TeamSize.SevenPlus);
        
        Project mobileProject = Project.builder()
                .id(2L)
                .title("Mobile App")
                .description("Test Description")
                .projectType(ProjectType.FREELANCE)
                .teamSize(TeamSize.SevenPlus)
                .recruiter(testUser)
                .build();

        Team mobileTeam = Team.builder()
                .id(2L)
                .name("Mobile App Team")
                .project(mobileProject)
                .build();

        String recruiterEmail = "recruiter@example.com";
        when(userRepository.findByEmail(recruiterEmail)).thenReturn(Optional.of(testUser));
        when(projectRepository.save(any(Project.class))).thenReturn(mobileProject);
        when(teamRepository.save(any(Team.class))).thenReturn(mobileTeam);
        when(chatRoomRepository.save(any(ChatRoom.class))).thenReturn(any(ChatRoom.class));

        // When
        ProjectResponse result = projectService.createProject(projectRequest, recruiterEmail);

        // Then
        assertThat(result).isNotNull();
        assertThat(result.getProjectType()).isEqualTo(ProjectType.FREELANCE);
        assertThat(result.getTeamSize()).isEqualTo(TeamSize.SevenPlus);
    }
}