package com.innosync.dto;

import lombok.Data;

import java.time.LocalDate;

@Data
public class WorkExperienceResponse {
    private LocalDate startDate;
    private LocalDate endDate;
    private String position;
    private String company;
    private String description;
}
