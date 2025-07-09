package com.innosync.service;

import com.innosync.dto.project.RoleApplicationResponse;
import com.innosync.model.ApplicationStatus;
import com.innosync.model.Project;
import com.innosync.model.ProjectRole;
import com.innosync.model.RoleApplication;
import com.innosync.model.Team;
import com.innosync.model.User;
import com.innosync.repository.RoleApplicationRepository;
import com.innosync.repository.ProjectRoleRepository;
import com.innosync.repository.TeamRepository;
import com.innosync.repository.UserRepository;
import jakarta.transaction.Transactional;
import lombok.RequiredArgsConstructor;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;
import org.springframework.security.access.AccessDeniedException;
import org.springframework.stereotype.Service;

import java.time.LocalDateTime;
import java.util.List;
import java.util.stream.Collectors;

@Service
@RequiredArgsConstructor
public class RoleApplicationService {
    private static final Logger logger = LoggerFactory.getLogger(RoleApplicationService.class);
    private final RoleApplicationRepository applicationRepository;
    private final ProjectRoleRepository projectRoleRepository;
    private final UserRepository userRepository;
    private final TeamRepository teamRepository;

    @Transactional
    public RoleApplicationResponse createApplication(Long projectRoleId, String userEmail) {
        User user = userRepository.findByEmail(userEmail)
                .orElseThrow(() -> new RuntimeException("User not found"));

        ProjectRole projectRole = projectRoleRepository.findById(projectRoleId)
                .orElseThrow(() -> new RuntimeException("Project role not found"));

        if (applicationRepository.existsByUserIdAndProjectRoleId(user.getId(), projectRoleId)) {
            throw new IllegalStateException("You have already applied to this role");
        }

        RoleApplication application = new RoleApplication();
        application.setUser(user);
        application.setProjectRole(projectRole);
        application.setStatus(ApplicationStatus.PENDING);
        application.setAppliedAt(LocalDateTime.now());

        return mapToResponse(applicationRepository.save(application));
    }

    @Transactional
    public RoleApplicationResponse updateApplicationStatus(Long applicationId, ApplicationStatus status, String updaterEmail) {
        RoleApplication application = applicationRepository.findById(applicationId)
                .orElseThrow(() -> new RuntimeException("Application not found"));

        // Verify updater is project recruiter
        if (!application.getProjectRole().getProject().getRecruiter().getEmail().equals(updaterEmail)) {
            throw new AccessDeniedException("Only project recruiter can update application status");
        }

        application.setStatus(status);
        application.setUpdatedAt(LocalDateTime.now());

        // If accepted, add user to the project's team
        if (status == ApplicationStatus.ACCEPTED) {
            Project project = application.getProjectRole().getProject();
            Team team = project.getTeam();
            if (team != null) {
                team.getMembers().add(application.getUser());
                teamRepository.save(team);
            }
        }

        return mapToResponse(applicationRepository.save(application));
    }

    public List<RoleApplicationResponse> getApplicationsForRole(Long projectRoleId, String recruiterEmail) {
        ProjectRole role = projectRoleRepository.findById(projectRoleId)
                .orElseThrow(() -> new RuntimeException("Project role not found"));

        // Verify requester is project recruiter
        if (!role.getProject().getRecruiter().getEmail().equals(recruiterEmail)) {
            throw new AccessDeniedException("Only project recruiter can view applications");
        }

        return applicationRepository.findByProjectRoleId(projectRoleId).stream()
                .map(this::mapToResponse)
                .collect(Collectors.toList());
    }


    public List<RoleApplicationResponse> getApplicationsByUser(String email) {
        User user = userRepository.findByEmail(email)
                .orElseThrow(() -> new RuntimeException("User not found"));

        return applicationRepository.findByUserId(user.getId()).stream()
                .map(this::mapToResponse)
                .collect(Collectors.toList());
    }


    private RoleApplicationResponse mapToResponse(RoleApplication application) {
        RoleApplicationResponse response = new RoleApplicationResponse();

        // Basic application info
        response.setId(application.getId());
        response.setStatus(application.getStatus());
        response.setAppliedAt(application.getAppliedAt());
        response.setUpdatedAt(application.getUpdatedAt());

        // User info
        User user = application.getUser();
        response.setUserId(user.getId());
        response.setUserFullName(user.getFullName());

        // Project role info
        ProjectRole role = application.getProjectRole();
        response.setProjectRoleId(role.getId());
        response.setRoleName(role.getRoleName());

        // Project info
        response.setProjectId(role.getProject().getId());
        response.setProjectTitle(role.getProject().getTitle());
        response.setProjectType(role.getProject().getProjectType());

        return response;
    }
}