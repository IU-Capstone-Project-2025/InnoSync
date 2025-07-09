package com.innosync.service;

import com.innosync.dto.project.ProjectRequest;
import com.innosync.dto.project.ProjectResponse;
import com.innosync.dto.project.ProjectRoleResponse;
import com.innosync.dto.project.ProjectRoleWithProjectResponse;
import com.innosync.model.Project;
import com.innosync.model.User;
import com.innosync.model.Team;
import com.innosync.model.ChatRoom;
import com.innosync.repository.ProjectRepository;
import com.innosync.repository.ProjectRoleRepository;
import com.innosync.repository.UserRepository;
import com.innosync.repository.TeamRepository;
import com.innosync.repository.ChatRoomRepository;
import lombok.RequiredArgsConstructor;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;
import org.springframework.stereotype.Service;
import org.springframework.http.HttpStatus;
import org.springframework.web.server.ResponseStatusException;

import java.util.List;
import java.util.stream.Collectors;
import java.util.HashSet;

@Service
@RequiredArgsConstructor
public class ProjectService {

    private static final Logger logger = LoggerFactory.getLogger(ProjectService.class);

    private final ProjectRepository projectRepository;
    private final UserRepository userRepository;
    private final ProjectRoleRepository projectRoleRepository;
    private final TeamRepository teamRepository;
    private final ChatRoomRepository chatRoomRepository;

    public ProjectResponse createProject(ProjectRequest request, String email) {
        logger.info("Creating project for recruiter: {}", email);
        try {
            User recruiter =  userRepository.findByEmail(email)
                    .orElseThrow(() -> new RuntimeException("User not found"));
            logger.debug("Recruiter found: {}", recruiter.getEmail());
            Project project = Project.builder()
                    .title(request.getTitle())
                    .description(request.getDescription())
                    .projectType(request.getProjectType())
                    .teamSize(request.getTeamSize())
                    .recruiter(recruiter)
                    .build();
            Project saved = projectRepository.save(project);
            // Create and link a team to the project, and add recruiter as member
            Team team = Team.builder().project(saved).name(saved.getTitle() + " Team").build();
            team.getMembers().add(recruiter);
            Team savedTeam = teamRepository.save(team);
            // Automatically create a team chat for the new team
            ChatRoom teamChat = ChatRoom.builder()
                .isTeamChat(true)
                .team(savedTeam)
                .name(savedTeam.getName() + " Chat")
                .participants(new HashSet<>(savedTeam.getMembers()))
                .build();
            chatRoomRepository.save(teamChat);
            logger.info("Project saved with id: {} for recruiter: {}", saved.getId(), recruiter.getEmail());
            return mapToDTO(saved);
        } catch (Exception e) {
            logger.error("Failed to create project for recruiter: {}", email, e);
            throw e;
        }
    }

    public List<ProjectResponse> getMyProjects(String email) {
        logger.debug("Getting projects for recruiter: {}", email);
        try {
            User recruiter =  userRepository.findByEmail(email)
                    .orElseThrow(() -> new RuntimeException("User not found"));
            List<Project> projects = projectRepository.findByRecruiter(recruiter);
            logger.info("Found {} projects for recruiter: {}", projects.size(), email);
            return projects.stream()
                    .map(this::mapToDTO)
                    .collect(Collectors.toList());
        } catch (Exception e) {
            logger.error("Failed to get projects for recruiter: {}", email, e);
            throw e;
        }
    }

    public ProjectResponse getProject(Long id) {
        Project project = projectRepository.findById(id)
                .orElseThrow(() -> new ResponseStatusException(HttpStatus.NOT_FOUND, "Project not found"));

        return mapToDTO(project);
    }

    private ProjectResponse mapToDTO(Project project) {
        ProjectResponse dto = new ProjectResponse();
        dto.setId(project.getId());
        dto.setTitle(project.getTitle());
        dto.setDescription(project.getDescription());
        dto.setCreatedAt(project.getCreatedAt());
        dto.setUpdatedAt(project.getUpdatedAt());
        dto.setProjectType(project.getProjectType());
        dto.setTeamSize(project.getTeamSize());
        return dto;
    }
}
