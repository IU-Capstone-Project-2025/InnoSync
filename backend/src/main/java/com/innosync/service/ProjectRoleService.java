package com.innosync.service;

import com.innosync.dto.project.ProjectResponse;
import com.innosync.dto.project.ProjectRoleRequest;
import com.innosync.dto.project.ProjectRoleResponse;
import com.innosync.dto.project.ProjectRoleWithProjectResponse;
import com.innosync.model.Project;
import com.innosync.model.ProjectRole;
import com.innosync.model.Technology;
import com.innosync.repository.ProjectRepository;
import com.innosync.repository.ProjectRoleRepository;
import com.innosync.repository.TechnologyRepository;
import com.innosync.repository.UserRepository;
import lombok.RequiredArgsConstructor;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;
import org.springframework.security.access.AccessDeniedException;
import org.springframework.stereotype.Service;
import org.springframework.web.server.ResponseStatusException;
import org.springframework.http.HttpStatus;

import java.util.ArrayList;
import java.util.Arrays;
import java.util.List;
import java.util.Set;
import java.util.stream.Collectors;

@Service
public class ProjectRoleService {
    private static final Logger logger = LoggerFactory.getLogger(ProjectRoleService.class);
    private final ProjectRoleRepository roleRepository;
    private final ProjectRepository projectRepository;
    private final UserRepository userRepository;
    private final TechnologyRepository technologyRepository;

    public ProjectRoleService(ProjectRoleRepository roleRepository,
                              ProjectRepository projectRepository,
                              UserRepository userRepository,
                              TechnologyRepository technologyRepository) {
        this.roleRepository = roleRepository;
        this.projectRepository = projectRepository;
        this.userRepository = userRepository;
        this.technologyRepository = technologyRepository;
    }
    public ProjectRoleResponse addRoleToProject(Long projectId, ProjectRoleRequest request, String creatorEmail) {
        Project project = projectRepository.findById(projectId).orElseThrow(() -> new ResponseStatusException(HttpStatus.NOT_FOUND, "Project not found"));

        if (!project.getRecruiter().getEmail().equals(creatorEmail)) {
            throw new AccessDeniedException("You are not the creator of this project.");
        }

        ProjectRole role = new ProjectRole();
        role.setProject(project);
        role.setRoleName(request.getRoleName());
        role.setExpertiseLevel(request.getExpertiseLevel());

        Set<Technology> technologies = request.getTechnologies().stream()
                .map(techName -> technologyRepository.findByNameIgnoreCase(techName)
                        .orElseGet(() -> {
                            Technology newTech = new Technology(techName);
                            return technologyRepository.save(newTech);
                        }))
                .collect(Collectors.toSet());

        role.setTechnologies(new ArrayList<>(technologies));

        ProjectRole saved = roleRepository.save(role);
        return toDto(saved);
    }

    public List<ProjectRoleResponse> getRolesByProjectId(Long projectId) {
        return roleRepository.findByProjectId(projectId).stream()
                .map(this::toDto)
                .collect(Collectors.toList());
    }



    public List<ProjectRoleWithProjectResponse> getAllProjectRolesWithProjectInfo() {
        return roleRepository.findAllWithProjectInfo()
                .stream()
                .map(role -> {
                    Project project = role.getProject(); // assuming there's a getProject()
                    return new ProjectRoleWithProjectResponse(
                            role.getId(),
                            role.getRoleName(),
//                            role.getDescription(),
                            project.getId(),
                            project.getTitle(),
                            project.getDescription()
                    );
                })
                .collect(Collectors.toList());
    }

    private ProjectRoleResponse toDto(ProjectRole role) {
        ProjectRoleResponse dto = new ProjectRoleResponse();
        dto.setId(role.getId());
        dto.setRoleName(role.getRoleName());
        dto.setExpertiseLevel(role.getExpertiseLevel());
        dto.setTechnologies(role.getTechnologies().stream()
                .map(Technology::getName)
                .collect(Collectors.toList()));
        return dto;
    }
}
