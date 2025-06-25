package com.innosync.dto;

import com.innosync.model.Education;
import com.innosync.model.ExpertiseLevel;
import lombok.Data;

import java.util.List;

@Data
public class ProfileResponse {
    private String email;
    private String fullName;
    private String telegram;
    private String github;
    private String bio;
    private String position;
    private Education education;
    private String expertise;
    private ExpertiseLevel expertiseLevel;
    private String resume;
    private List<WorkExperienceResponse> workExperience;
    private List<String> technologies;
}
