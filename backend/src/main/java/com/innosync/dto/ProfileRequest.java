package com.innosync.dto;

import com.fasterxml.jackson.annotation.JsonProperty;
import com.innosync.model.Education;
import com.innosync.model.ExpertiseLevel;
import lombok.Data;

import java.util.List;


@Data
public class ProfileRequest {
    private String telegram;
    private String github;
    private String bio;
    private String position;
    private Education education;
    private String expertise;
    @JsonProperty("expertise_level")
    private ExpertiseLevel expertiseLevel;
    private String resume;
    @JsonProperty("work_experience")
    private List<WorkExperienceRequest> workExperience;
    @JsonProperty("technologies")
    private List<String> technologies;

}
